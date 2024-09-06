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
   V = U
// ^ definition github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/V#
// documentation
// > ```go
// > type V = U
// > ```
// documentation
// > Check that we don't panic
// > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
// documentation
// > ```go
// > struct{}
// > ```
//     ^ reference github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
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
   Z int32
// ^ definition github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/Z#
// documentation
// > Check that we don't panic
// > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
// documentation
// > ```go
// > int32
// > ```
  )
  
  func f(u U) {}
//     ^ definition github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/f().
//     documentation
//     > ```go
//     > func f(u U)
//     > ```
//       ^ definition local 0
//         ^ reference github.com/sourcegraph/scip-go . `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
  
