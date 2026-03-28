# Development docs

## Run tests

Run all tests:

```bash
go test ./...
```

Update snapshot tests:

```bash
go test ./internal/index -update-snapshots
```

## Cutting releases

1. Land a PR with the following changes:

   - A ChangeLog entry with `## vM.N.P`
   - Updated version in `internal/index/version.txt`

2. From the `main` branch, trigger the [release workflow](https://github.com/sourcegraph/scip-go/actions/workflows/release.yml)
   via **Actions → release → Run workflow**, entering `M.N.P` as the version.
