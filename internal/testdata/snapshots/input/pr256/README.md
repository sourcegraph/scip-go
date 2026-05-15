# Directories with only `_test.go` files

Reproduction for issue [#255].

A directory that only contains `*_test.go` files belonging to an external
`*_test` package would cause `scip-go` to panic with
`index out of range [0] with length 0` because the synthetic regular package
returned by `packages.Load` had an empty `Syntax` slice.

The fix skips package-symbol attachment for packages with no parsed source
files; the external `*_test` package still has its own non-empty entry and is
indexed normally.

  [#255]: https://github.com/scip-code/scip-go/issues/255
