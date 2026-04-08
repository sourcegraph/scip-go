  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  import (
   "slices"
//  ^^^^^^ definition local 0
//  ^^^^^^ reference github.com/golang/go/src go1.22 slices/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/f().
  func f(xs []int) int {
//     ^ definition 0.1.test `sg/initial`/f().
//     documentation
//     > ```go
//     > func f(xs []int) int
//     > ```
//       ^^ definition local 1
   for _, x := range slices.All(xs) {
//        ^ definition local 2
//                   ^^^^^^ reference local 0
//                          ^^^ reference github.com/golang/go/src go1.22 slices/All().
//                              ^^ reference local 1
    return x
//         ^ reference local 2
   }
   return -1
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/f().
  
