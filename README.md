# scip-go

SCIP indexer for Golang.

# Quick Start

## Installation

This will build and install the latest version of `scip-go`

```
go install github.com/sourcegraph/scip-go/cmd/scip-go@latest
```

You can confirm it's been installed by running:

```
scip-go --version
```

## Indexing a Go project

From the root of your project, you can run:

```
scip-go
```

If `scip-go` is unable to determine some project information, you may need to provide some command-line arguments.


```
scip-go --module-name=NAME --module-version=VERSION
```

If this doesn't solve the problem, check the rest of the available flags in:

```
scip-go --help
```

### Common Problems:

- Unable to navigate to Go standard library.
  - To solve this, you may want to use the `--go-version=go1.X.Y` flag when indexing and then also index the go versions manually.
  - To index the Go standard library, you'll want to check out https://github.com/golang/go, checkout the tag you want, navigate to the `src/` directory, and then run:
    - `$ scip-go --go-version=go1.X.Y` and then upload that index to your local sourcegraph instance.
    - After you've done this, you should be able to navigate to the standard library.


(NOTE: Projects without a `go.mod` may experience challenges indexing. See next section for details)

## Indexing without shelling to `go` binary

`scip-go` by default uses a few different `go` commands from the command line to
gain information about the project and module. To avoid running `go` directly
(perhaps you have some other build system), you will need to supply the folling args.

```
scip-go --module-name="<my modules name here>"
```

NOTE: The rest of this isn't properly implemented yet. It's on the todo list for scip-go.

## Indexing in CI

```
# Install scip-go
go install github.com/sourcegraph/scip-go/cmd/scip-go@latest

# Run scip-go
scip-go

# Upload index with any necessary tokens (shown here using GitHub workflow syntax)
src lsif upload -github-token='${{ secrets.GITHUB_TOKEN }}' -no-progress
```


# Contributing

Contributors should follow the [Sourcegraph Community Code of Conduct](https://handbook.sourcegraph.com/company-info-and-process/community/code_of_conduct/).
