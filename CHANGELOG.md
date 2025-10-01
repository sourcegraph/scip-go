# ChangeLog


## v0.1.26

- Add information about local symbols to generated index

## v0.1.25

- Upgrades Dockerfile to use Go 1.25.0 (released Aug 12 2025)

## v0.1.24

- Upgrades Dockerfile to use Go 1.24.3 (released May 6 2025).
- Sets `GOTOOLCHAIN=auto` in Dockerfile by default, to allow
  for transparent toolchain upgrading if the network
  configuration allows it.

## v0.1.23

- Upgrades Dockerfile to use Go 1.24.0 (released Feb 11 2025).
  (https://github.com/sourcegraph/scip-go/pull/146)

## v0.1.22

- Fixes a panic when using a custom GOPACKAGESDRIVER along
  with a build system other than the default Go build system
  (https://github.com/sourcegraph/scip-go/pull/138)
- Optionally allows passing package patterns to scip-go for
  only indexing a subset of packages.
  (https://github.com/sourcegraph/scip-go/pull/139)
- Upgrades Dockerfile to use Go 1.23.3 (released Oct 1 2024).
  (https://github.com/sourcegraph/scip-go/pull/141)

## v0.1.21

- Upgrades Dockerfile to use Go 1.23.2 (released Oct 1 2024).
  (https://github.com/sourcegraph/scip-go/pull/136)

## v0.1.20

- Fixes a bug which caused test files using the same package name
  as the main package to not be indexed.
  (https://github.com/sourcegraph/scip-go/pull/134)

## v0.1.19

- Upgrades Dockerfile to use Go 1.23.1 (released on Sept 05 2024).
- Removes incorrect warning about missing ASTs for packages
  only containing test files.
  (https://github.com/sourcegraph/scip-go/pull/129)

## v0.1.18

- Adds more detailed logging for diagnosing hangs
  during indexing.
  (https://github.com/sourcegraph/scip-go/pull/126)

## v0.1.17

- Fixes panic due to alias types in v0.1.16.
  (https://github.com/sourcegraph/scip-go/pull/121)

## v0.1.16

NOTE: This release panics on alias types due to
upstream Go bug [68877](https://github.com/golang/go/issues/68877);
you can work around that by using `GODEBUG=gotypesalias=0`.

- Updates the indexer and Dockerfile for Go 1.23.0
  (https://github.com/sourcegraph/scip-go/pull/116)


## v0.1.15

- Updates the Dockerfile to use Go 1.22.5.
  (https://github.com/sourcegraph/scip-go/pull/111)

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
