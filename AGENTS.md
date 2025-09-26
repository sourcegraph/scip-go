# scip-go Development Guide

## Commands

- **Build**: `go build ./cmd/scip-go`
- **Test all**: `go test ./...`
- **Test single package**: `go test ./internal/index`
- **Update snapshots**: `go test ./internal/index -update-snapshots`
- **Install**: `go install ./cmd/scip-go`
- **Version check**: `go version` (uses Go 1.24.3)

## Testing

- **Snapshot testing**: Primary testing method using
  `internal/testdata/snapshots` directory with `internal/index/scip_test.go` as
  the test runner
- **Test data**: Located in `internal/testdata/` with test cases as
  subdirectories
- **Update snapshots**: Run with `-update-snapshots` flag to regenerate expected
  outputs
- **Filter tests**: Use `-filter <name>` to run specific test cases

## Architecture

SCIP (Source Code Intelligence Protocol) indexer for Go. Main entry point is
`cmd/scip-go/main.go`. Key packages: `internal/index` (core indexing),
`internal/loader` (package loading), `internal/visitors` (AST traversal),
`internal/document` (SCIP document generation).

## Code Style

- Package naming: lowercase, single word preferred (e.g., `index`, `loader`,
  `symbols`)
- Imports: grouped by std library, third-party, then internal packages with
  blank lines
- Test files: use `package_test` convention for black-box testing
- Embed files: use `//go:embed` for version.txt and similar resources
- Error handling: return explicit errors, no panics in library code
- Logging: use `charmbracelet/log` for structured logging
