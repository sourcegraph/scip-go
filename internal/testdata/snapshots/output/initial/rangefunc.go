  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
  import (
    "slices"
//   ^^^^^^ reference github.com/golang/go/src go1.22 slices/
  )
  
  func f(xs []int) int {
//     ^ definition 0.1.test `sg/initial`/f().
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
  
