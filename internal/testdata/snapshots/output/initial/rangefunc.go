  package initial
//        ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/
  
  import (
    "slices"
//   ^^^^^^ reference github.com/golang/go/src go1.22 slices/
  )
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/f().
  func f(xs []int) int {
//     ^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/f().
//     documentation
//     > ```go
//     > func f(xs []int) int
//     > ```
//       ^^ definition local 0
    for _, x := range slices.All(xs) {
//         ^ definition local 1
//                    ^^^^^^ reference github.com/golang/go/src go1.22 slices/
//                           ^^^ reference github.com/golang/go/src go1.22 slices/All().
//                               ^^ reference local 0
      return x
//           ^ reference local 1
    }
    return -1
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/f().
  
