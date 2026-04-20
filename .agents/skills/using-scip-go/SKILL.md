---
name: using-scip-go
description: >
  Indexes Go projects with scip-go to produce SCIP indexes. Use when
  asked to index, run scip-go, or generate code navigation data for
  Go code.
---

# Using scip-go

scip-go is an SCIP indexer for Go. It produces `.scip` index files that
power code navigation (go-to-definition, find-references,
find-implementations).

## Basic Usage

Run from the root of a Go module (where `go.mod` lives):

``` bash
scip-go index
```

This indexes all packages (`./...`) and writes `index.scip` in the
current directory.

## Commands

| Command    | Description                                        |
|------------|----------------------------------------------------|
| `index`    | Index Go source code and emit SCIP index (default) |
| `packages` | List current and dependency packages               |
| `missing`  | List missing documents                             |

The `index` command is the default and runs when no command is
specified.

## Common Flags

| Flag                      | Short | Description                                                          |
|---------------------------|-------|----------------------------------------------------------------------|
| `--output`                | `-o`  | Output file path (default: `index.scip`)                             |
| `--module-root`           |       | Directory containing `go.mod` (default: auto-detected)               |
| `--module-path`           |       | Override module path inferred from `go.mod`                          |
| `--module-version`        |       | Module version (default: inferred from git)                          |
| `--go-version`            |       | Go stdlib version to link to, e.g. `go1.22` (default: from `go.mod`) |
| `--repository-remote`     |       | Canonical repository remote name (default: from git)                 |
| `--skip-implementations`  |       | Skip generating implementation relationships                        |
| `--skip-tests`            |       | Skip indexing test files                                             |
| `--verbose`               | `-V`  | Enable info logs. Use `-VV` for debug logs                           |
| `--quiet`                 | `-q`  | Suppress all output                                                  |
| `--version`               | `-v`  | Show version                                                         |

All flags can also be set via environment variables with the flag name
uppercased and hyphens replaced by underscores (e.g. `MODULE_ROOT`,
`SKIP_TESTS`, `OUTPUT`).

## Package Patterns

By default scip-go indexes `./...` (all packages recursively). You can
specify explicit patterns:

``` bash
# Index only specific packages
scip-go index ./pkg/... ./cmd/...

# Index a single package
scip-go index ./internal/loader
```

See `go help packages` for the full pattern syntax.

## Standard Go Module Projects

For most projects, just run `scip-go` from the module root:

``` bash
cd my-project
scip-go index -o index.scip
```

## Multi-Module Repositories (Monorepos)

For repositories with multiple `go.mod` files (e.g. Kubernetes), index
each module separately:

``` bash
# Index root module
scip-go index -o root-index.scip

# Index sub-modules
cd staging/src/k8s.io/api
scip-go index -o api-index.scip

cd staging/src/k8s.io/client-go
scip-go index -o client-go-index.scip
```

Alternatively, use `--module-root` to point at a sub-module without
changing directories:

``` bash
scip-go index --module-root=staging/src/k8s.io/api -o api-index.scip
```

## Alternative Build Systems (Bazel, Buck2, Please)

For projects using non-Go build systems, set the `GOPACKAGESDRIVER`
environment variable:

``` bash
GOPACKAGESDRIVER=./gopackagesdriver.sh scip-go index
```

Note: cross-repo navigation does not work with custom package drivers
because the driver protocol does not expose module metadata. Use
`--module-path` to set the module path manually.

## Troubleshooting

### Go standard library navigation not working

Specify the Go version explicitly:

``` bash
scip-go index --go-version=go1.22
```

To enable navigation *into* stdlib source, index the Go source tree
separately:

``` bash
cd $(go env GOROOT)/src
scip-go index --go-version=go1.22
```

### `unsafe.Pointer` warnings

Warnings about `Unable to find symbol` for `unsafe.Pointer` or
`unsafe.Sizeof` are expected and harmless — the `unsafe` package is a
compiler builtin without real source.

### `*types.Label` warnings

Warnings about `*types.Label` (for `goto`/`break`/`continue` labels)
are expected — Go labels are not symbolized.

### Projects without `go.mod`

Projects without a `go.mod` may have issues. Supply the module name
explicitly:

``` bash
scip-go index --module-path="github.com/owner/repo"
```
