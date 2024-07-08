.PHONY: clean checks test build install

export CGO_ENABLED=0

default: clean checks test build

clean:
	rm -rf dist/ builds/ cover.out

build: clean
	go build -ldflags "-s -w" -trimpath

install: clean
	go install -ldflags "-s -w" -trimpath

test: clean
	go test -v -cover ./...

checks:
	golangci-lint run
