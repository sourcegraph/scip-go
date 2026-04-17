# GOPACKAGESDRIVER support

Tests indexing through a custom `GOPACKAGESDRIVER` instead of `go list`.

The `driver/` subdirectory contains a proxy driver that loads packages via
`go list` (with `GOPACKAGESDRIVER=off`) and re-emits them using the driver wire
format --- which has no `Module` field. This exercises the `normalizePackage`
code path that fills in module info from CLI flags.

The `driver/` has its own `go.mod` to create a module boundary so the indexer's
`./...` pattern skips it. Any snapshot input directory can opt into driver
testing by adding a `driver/` subdirectory with `main.go` and `go.mod`.
