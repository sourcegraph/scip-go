  package main
//        ^^^^ definition github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/
//        documentation
//        > package main
  
  // Check that we don't panic
  // Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
  type (
   T struct{}
// ^ definition github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/T#
// documentation
// > ```go
// > type T struct
// > ```
// documentation
// > Check that we don't panic
// > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
// documentation
// > ```go
// > struct{}
// > ```
   U = T
// ^ definition github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
// documentation
// > ```go
// > type U = T
// > ```
// documentation
// > Check that we don't panic
// > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
// documentation
// > ```go
// > struct{}
// > ```
//     ^ reference github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/T#
   S U
// ^ definition github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/S#
// documentation
// > ```go
// > type S struct
// > ```
// documentation
// > Check that we don't panic
// > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
// documentation
// > ```go
// > struct{}
// > ```
//   ^ reference github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
  )
  
