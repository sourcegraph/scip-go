---
name: using-scip-go
description: >
  Generate a SCIP index for a Go module with scip-go. Use when asked
  to index Go code, produce index.scip / SCIP output, or handle Go
  monorepos and GOPACKAGESDRIVER-based builds.
---

# scip-go

Primary command:

``` bash
# Run from the directory containing the target go.mod
scip-go index -o index.scip [package-pattern...]
```

## Standard module

``` bash
cd path/to/module
scip-go index -o index.scip
```

## Specific packages only

``` bash
scip-go index -o index.scip ./cmd/... ./pkg/...
```

See `go help packages` for the full pattern syntax.

## Submodule without changing directories

``` bash
scip-go index --module-root=path/to/module -o module.scip
```

## Multi-module / go.work repo

Index each module separately (run once per `go.mod`):

``` bash
scip-go index -o root.scip
cd staging/src/k8s.io/api
scip-go index -o api.scip
```

## Custom build system (Bazel, Buck2, Please)

``` bash
GOPACKAGESDRIVER=./driver.sh \
  scip-go index \
    --module-path=example.com/repo \
    --module-version=abc1234 \
    -o index.scip
```

With `GOPACKAGESDRIVER`, module metadata may be incomplete. Pass
`--module-path` and, if known, `--module-version` to improve symbol
identity. Dependency cross-repo navigation may still be limited.

## Useful flags

| Flag                     | Short | Description                              |
|--------------------------|-------|------------------------------------------|
| `--output`               | `-o`  | Output file (default: `index.scip`)      |
| `--module-root`          |       | Directory containing `go.mod`            |
| `--module-path`          |       | Override module path from `go.mod`       |
| `--module-version`       |       | Module version (default: from git)       |
| `--go-version`           |       | Go stdlib version, e.g. `go1.22`        |
| `--skip-tests`           |       | Exclude test files                       |
| `--skip-implementations` |       | Skip implementation relationship edges   |
| `--verbose`              | `-V`  | Info logs; use `-VV` for debug           |

## Diagnostics

``` bash
scip-go missing   # list files not covered by the index
scip-go packages  # list current and dependency packages
```

These accept the same flags and package patterns as `index`.

## Guardrails

- Always run from the target module root or pass `--module-root`
  explicitly. Do not rely on automatic detection.
- `--module-path` sets the module identity in the index but does not
  replace a working `go.mod` or package driver. If `go/packages`
  cannot load the code, indexing will fail.
