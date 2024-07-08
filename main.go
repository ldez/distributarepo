package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/google/go-github/v62/github"
	"github.com/jszwec/csvutil"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

const maxWorkers = 15

var version = "devel"

// Result fork information.
type Result struct {
	ForkURL       string `json:"forkURL" csv:"forkURL"`
	Owner         string `json:"-" csv:"-"`
	Repo          string `json:"-" csv:"-"`
	Ahead         int    `json:"ahead" csv:"ahead"`
	AheadURL      string `json:"aheadURL" csv:"aheadURL"`
	Behind        int    `json:"behind" csv:"behind"`
	Stars         int    `json:"stars" csv:"stars"`
	Forks         int    `json:"forks" csv:"forks"`
	Issues        int    `json:"issues" csv:"issues"`
	DefaultBranch string `json:"-" csv:"-"`
}

func main() {
	app := cli.NewApp()
	app.Name = "distributarepo"
	app.HelpName = "distributarepo"
	app.Usage = "Helper to get an overview of the forks of a GitHub repository."
	app.EnableBashCompletion = true
	app.Version = version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "owner",
			Aliases:  []string{"o"},
			Usage:    "GitHub owner of the repository",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "repo",
			Aliases:  []string{"r"},
			Usage:    "Name of the GitHub repository",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "token",
			EnvVars: []string{"GITHUB_TOKEN"},
			Usage:   "GitHub token",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "format",
			Aliases: []string{"f"},
			Usage:   "format output: csv, json, md, markdown",
			Value:   "markdown",
		},
		&cli.PathFlag{
			Name:  "output",
			Usage: "output file (default: stdout)",
			Value: "",
		},
	}
	app.Action = func(cliCtx *cli.Context) error {
		formatter, err := getFormatter(cliCtx.String("format"))
		if err != nil {
			return err
		}

		writer, err := getWriter(cliCtx.String("output"))
		if err != nil {
			return err
		}

		ctx := context.Background()

		d := &Distributary{
			client:   newGitHubClient(ctx, cliCtx.String("token")),
			owner:    cliCtx.String("owner"),
			repoName: cliCtx.String("repo"),
			writer:   writer,
		}

		return d.run(ctx, formatter)
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("distributarepo version %s %s/%s\n", c.App.Version, runtime.GOOS, runtime.GOARCH)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getWriter(output string) (io.Writer, error) {
	if output == "" {
		return os.Stdout, nil
	}

	file, err := os.Create(filepath.Clean(output))
	if err != nil {
		return nil, fmt.Errorf("create output file: %w", err)
	}

	return file, nil
}

// Formatter output formatter.
type Formatter func(w io.Writer, results []*Result) error

// Distributary finds and displays forks.
type Distributary struct {
	client *github.Client

	owner    string
	repoName string

	writer io.Writer
}

func (d *Distributary) run(ctx context.Context, formatter Formatter) error {
	repo, _, err := d.client.Repositories.Get(ctx, d.owner, d.repoName)
	if err != nil {
		return fmt.Errorf("get repository: %w", err)
	}

	defaultBranch := repo.GetDefaultBranch()

	resultsChan := make(chan *Result)
	forkChan := make(chan *github.Repository, maxWorkers)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(maxWorkers)

		for range maxWorkers {
			go func() {
				d.work(ctx, defaultBranch, forkChan, resultsChan)
				wg.Done()
			}()
		}

		wg.Wait()
		close(resultsChan)
	}()

	options := &github.RepositoryListForksOptions{}

	go func() {
		for {
			forks, resp, errG := d.client.Repositories.ListForks(ctx, d.owner, d.repoName, options)
			if errG != nil {
				log.Fatal(errG)
			}

			for _, fork := range forks {
				if fork.GetArchived() {
					continue
				}

				forkChan <- fork
			}

			if resp.NextPage == 0 {
				break
			}

			options.Page = resp.NextPage
		}

		close(forkChan)
	}()

	var results []*Result

	for result := range resultsChan {
		if result == nil {
			continue
		}

		results = append(results, result)
	}

	err = formatter(d.writer, results)
	if err != nil {
		return err
	}

	return nil
}

func (d *Distributary) work(ctx context.Context, defaultBranch string, forkChan <-chan *github.Repository, resultsChan chan<- *Result) {
	for fork := range forkChan {
		head := fmt.Sprintf("%s:%s", fork.GetOwner().GetLogin(), fork.GetDefaultBranch())

		commits, cResp, err := d.client.Repositories.CompareCommits(ctx, d.owner, d.repoName, defaultBranch, head, nil)
		if cResp != nil && cResp.StatusCode == http.StatusNotFound {
			// some forks are still returned even if the user/orga has been deleted.
			continue
		}
		if err != nil { //nolint:wsl // the response should be handled before the error.
			log.Fatal(err)
		}

		if commits.GetAheadBy() == 0 && commits.GetBehindBy() >= 0 {
			continue
		}

		result := &Result{
			Owner:         fork.GetOwner().GetLogin(),
			Repo:          fork.GetName(),
			Ahead:         commits.GetAheadBy(),
			Behind:        commits.GetBehindBy(),
			Stars:         fork.GetStargazersCount(),
			Forks:         fork.GetForksCount(),
			Issues:        fork.GetOpenIssuesCount(),
			DefaultBranch: fork.GetDefaultBranch(),
		}

		result.ForkURL = fmt.Sprintf("https://github.com/%s/%s", result.Owner, result.Repo)

		if result.Ahead > 0 {
			// https://github.com/<src-owner>/<src-repo>/compare/<src-default-branch>...<fork-owner>:<fork-repo>:<fork-default-branch>
			result.AheadURL = fmt.Sprintf(
				"https://github.com/%s/%s/compare/%s...%s:%s:%s",
				d.owner, d.repoName,
				defaultBranch, result.Owner, result.Repo, result.DefaultBranch,
			)
		}

		resultsChan <- result
	}
}

func toMarkdown(w io.Writer, results []*Result) error {
	var data [][]string

	for _, result := range results {
		ahead := "0"
		if result.Ahead > 0 {
			ahead = fmt.Sprintf("[%d](%s)", result.Ahead, result.AheadURL)
		}

		data = append(data, []string{
			fmt.Sprintf("[%s/%s](%s)", result.Owner, result.Repo, result.ForkURL),
			ahead, strconv.Itoa(result.Behind),
			strconv.Itoa(result.Stars), strconv.Itoa(result.Forks), strconv.Itoa(result.Issues),
		})
	}

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Fork", "Ahead", "Behind", "Stars", "Forks", "Issues"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	return nil
}

func toCSV(w io.Writer, results []*Result) error {
	cw := csv.NewWriter(w)

	err := csvutil.NewEncoder(cw).Encode(results)
	if err != nil {
		return fmt.Errorf("encode csv: %w", err)
	}

	cw.Flush()

	return nil
}

func toJSON(w io.Writer, results []*Result) error {
	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

func getFormatter(format string) (Formatter, error) {
	switch strings.ToLower(format) {
	case "csv":
		return toCSV, nil
	case "json":
		return toJSON, nil
	case "md", "markdown":
		return toMarkdown, nil
	default:
		return nil, fmt.Errorf("unknown formatter: %s", format)
	}
}

func newGitHubClient(ctx context.Context, token string) *github.Client {
	if token == "" {
		return github.NewClient(nil)
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	return github.NewClient(oauth2.NewClient(ctx, ts))
}
