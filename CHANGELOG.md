# ChangeLog

## v0.1.14

- Fixes a bug with cross-repo navigation when depending
  on unpublished versions of libraries.
  (https://github.com/sourcegraph/scip-go/pull/99)

## v0.1.13

- Adds workaround for a panic triggerd by the presence of multiple
  field definitions with the same anonymous type.
  (https://github.com/sourcegraph/scip-go/pull/96)

  ```go
  type T struct {
      A, B struct { B int }
  }
  ```

## v0.1.12

- Fixes the Dockerfile for the indexer. Due to a bug in the Dockerfile,
  the v0.1.11 release does not have any accompanying Docker image.
  (https://github.com/sourcegraph/scip-go/pull/86)

## v0.1.11

- Updates the indexer to build using Go 1.22.1.
  (https://github.com/sourcegraph/scip-go/pull/81)

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
