# ChangeLog

## v0.1.10

- Updates the indexer to build using Go 1.21.5.
  (https://github.com/sourcegraph/scip-go/pull/50)

## v0.1.9

- Fixes a bug where the indexer would emit an empty SCIP index
  when hitting a panic.
  (https://github.com/sourcegraph/scip-go/pull/62)

## v0.1.8

- Fixed the version number emitted in SCIP indexes
  and printed by `scip-go --version`.
  (https://github.com/sourcegraph/scip-go/pull/60)
