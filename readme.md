# Distributarepo

Helper to get an overview of the forks of a GitHub repository.

```
NAME:
   distributarepo - Helper to get an overview of the forks of a GitHub repository.

USAGE:
   distributarepo [global options] command [command options] 

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --owner value, -o value   GitHub owner of the repository
   --repo value, -r value    Name of the GitHub repository
   --token value             GitHub token [$GITHUB_TOKEN]
   --format value, -f value  format output: csv, json, md, markdown (default: "markdown")
   --output value            output file (default: stdout)
   --help, -h                show help
   --version, -v             print the version
```

_[distributary](https://en.wikipedia.org/wiki/Distributary)_


## Examples

<details><summary>Markdown</summary>

```console
$ distributarepo -o "gofrs" -r "flock"
|                                FORK                                 |                                      AHEAD                                       | BEHIND | STARS | FORKS | ISSUES |
|---------------------------------------------------------------------|----------------------------------------------------------------------------------|--------|-------|-------|--------|
| [JackMordaunt/flock](https://github.com/JackMordaunt/flock)         | [1](https://github.com/gofrs/flock/compare/main...JackMordaunt:flock:master)     |     55 |     0 |     0 |      0 |
| [moskyb/flock](https://github.com/moskyb/flock)                     | [2](https://github.com/gofrs/flock/compare/main...moskyb:flock:master)           |     55 |     0 |     0 |      0 |
| [aaydin-tr/flock](https://github.com/aaydin-tr/flock)               | [2](https://github.com/gofrs/flock/compare/main...aaydin-tr:flock:master)        |     55 |     0 |     0 |      0 |
| [cluetrust/flock](https://github.com/cluetrust/flock)               | [1](https://github.com/gofrs/flock/compare/main...cluetrust:flock:master)        |     22 |     0 |     0 |      0 |
| [mikhail-artemev/flock](https://github.com/mikhail-artemev/flock)   | [1](https://github.com/gofrs/flock/compare/main...mikhail-artemev:flock:master)  |     55 |     0 |     0 |      0 |
| [onflowser/flock](https://github.com/onflowser/flock)               | [3](https://github.com/gofrs/flock/compare/main...onflowser:flock:master)        |     55 |     0 |     0 |      0 |
| [juicedata/flock](https://github.com/juicedata/flock)               | [1](https://github.com/gofrs/flock/compare/main...juicedata:flock:master)        |     55 |     0 |     0 |      0 |
| [pgavlin/flock](https://github.com/pgavlin/flock)                   | [1](https://github.com/gofrs/flock/compare/main...pgavlin:flock:master)          |     55 |     0 |     0 |      0 |
| [trying2016/flock](https://github.com/trying2016/flock)             | [2](https://github.com/gofrs/flock/compare/main...trying2016:flock:master)       |     57 |     0 |     0 |      0 |
| [88250/flock](https://github.com/88250/flock)                       | [2](https://github.com/gofrs/flock/compare/main...88250:flock:master)            |     55 |     0 |     0 |      0 |
| [ujjwalsh/flock](https://github.com/ujjwalsh/flock)                 | [1](https://github.com/gofrs/flock/compare/main...ujjwalsh:flock:master)         |     57 |     0 |     0 |      0 |
| [kakami/flock](https://github.com/kakami/flock)                     | [1](https://github.com/gofrs/flock/compare/main...kakami:flock:master)           |     57 |     0 |     0 |      0 |
| [fearful-symmetry/flock](https://github.com/fearful-symmetry/flock) | [1](https://github.com/gofrs/flock/compare/main...fearful-symmetry:flock:master) |     68 |     0 |     0 |      0 |
| [wataash/flock](https://github.com/wataash/flock)                   | [1](https://github.com/gofrs/flock/compare/main...wataash:flock:master)          |     68 |     0 |     0 |      0 |
| [virtuald/go-flock](https://github.com/virtuald/go-flock)           | [4](https://github.com/gofrs/flock/compare/main...virtuald:go-flock:master)      |     93 |     0 |     0 |      0 |
```

</details>

<details><summary>JSON</summary>

```console
$ distributarepo -o "gofrs" -r "flock" -f json
[{"forkURL":"https://github.com/aaydin-tr/flock","ahead":2,"aheadURL":"https://github.com/gofrs/flock/compare/main...aaydin-tr:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/mikhail-artemev/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...mikhail-artemev:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/cluetrust/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...cluetrust:flock:master","behind":22,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/moskyb/flock","ahead":2,"aheadURL":"https://github.com/gofrs/flock/compare/main...moskyb:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/onflowser/flock","ahead":3,"aheadURL":"https://github.com/gofrs/flock/compare/main...onflowser:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/JackMordaunt/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...JackMordaunt:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/pgavlin/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...pgavlin:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/88250/flock","ahead":2,"aheadURL":"https://github.com/gofrs/flock/compare/main...88250:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/trying2016/flock","ahead":2,"aheadURL":"https://github.com/gofrs/flock/compare/main...trying2016:flock:master","behind":57,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/juicedata/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...juicedata:flock:master","behind":55,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/ujjwalsh/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...ujjwalsh:flock:master","behind":57,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/kakami/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...kakami:flock:master","behind":57,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/wataash/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...wataash:flock:master","behind":68,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/fearful-symmetry/flock","ahead":1,"aheadURL":"https://github.com/gofrs/flock/compare/main...fearful-symmetry:flock:master","behind":68,"stars":0,"forks":0,"issues":0},{"forkURL":"https://github.com/virtuald/go-flock","ahead":4,"aheadURL":"https://github.com/gofrs/flock/compare/main...virtuald:go-flock:master","behind":93,"stars":0,"forks":0,"issues":0}]
```

</details>

<details><summary>CSV</summary>

```console
$ distributarepo -o "gofrs" -r "flock" -f csv
forkURL,ahead,aheadURL,behind,stars,forks,issues
https://github.com/JackMordaunt/flock,1,https://github.com/gofrs/flock/compare/main...JackMordaunt:flock:master,55,0,0,0
https://github.com/cluetrust/flock,1,https://github.com/gofrs/flock/compare/main...cluetrust:flock:master,22,0,0,0
https://github.com/mikhail-artemev/flock,1,https://github.com/gofrs/flock/compare/main...mikhail-artemev:flock:master,55,0,0,0
https://github.com/onflowser/flock,3,https://github.com/gofrs/flock/compare/main...onflowser:flock:master,55,0,0,0
https://github.com/moskyb/flock,2,https://github.com/gofrs/flock/compare/main...moskyb:flock:master,55,0,0,0
https://github.com/aaydin-tr/flock,2,https://github.com/gofrs/flock/compare/main...aaydin-tr:flock:master,55,0,0,0
https://github.com/juicedata/flock,1,https://github.com/gofrs/flock/compare/main...juicedata:flock:master,55,0,0,0
https://github.com/pgavlin/flock,1,https://github.com/gofrs/flock/compare/main...pgavlin:flock:master,55,0,0,0
https://github.com/88250/flock,2,https://github.com/gofrs/flock/compare/main...88250:flock:master,55,0,0,0
https://github.com/trying2016/flock,2,https://github.com/gofrs/flock/compare/main...trying2016:flock:master,57,0,0,0
https://github.com/ujjwalsh/flock,1,https://github.com/gofrs/flock/compare/main...ujjwalsh:flock:master,57,0,0,0
https://github.com/kakami/flock,1,https://github.com/gofrs/flock/compare/main...kakami:flock:master,57,0,0,0
https://github.com/fearful-symmetry/flock,1,https://github.com/gofrs/flock/compare/main...fearful-symmetry:flock:master,68,0,0,0
https://github.com/wataash/flock,1,https://github.com/gofrs/flock/compare/main...wataash:flock:master,68,0,0,0
https://github.com/virtuald/go-flock,4,https://github.com/gofrs/flock/compare/main...virtuald:go-flock:master,93,0,0,0
```

</details>
