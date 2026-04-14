  package main
//        ^^^^ definition github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/
//             display_name main
//             signature_documentation
//             > package main
  
  // Check that we don't panic
  // Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
  type (
   T struct{}
// ^ definition github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/T#
//   signature_documentation
//   > type T struct{}
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
   U = T
// ^ definition github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
//   signature_documentation
//   > type U = T
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
//     ^ reference github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/T#
   V = U
// ^ definition github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/V#
//   signature_documentation
//   > type V = U
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
//     ^ reference github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
   S U
// ^ definition github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/S#
//   signature_documentation
//   > type S struct{}
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
//   ^ reference github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
   Z int32
// ^ definition github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/Z#
//   signature_documentation
//   > type Z int32
//   documentation
//   > Check that we don't panic
//   > Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
  )
  
//⌄ enclosing_range_start github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/f().
  func f(u U) {}
//     ^ definition github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/f().
//       signature_documentation
//       > func f(u U)
//       ^ definition local 0
//         display_name u
//         signature_documentation
//         > var u U
//         ^ reference github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/U#
//             ⌃ enclosing_range_end github.com/sourcegraph/scip-go 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/snapshots/input/alias`/f().
  
