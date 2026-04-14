  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  import (
   "slices"
//  ^^^^^^ reference github.com/golang/go/src go1.22 slices/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/f().
  func f(xs []int) int {
//     ^ definition 0.1.test `sg/initial`/f().
//       signature_documentation
//       > func f(xs []int) int
//       ^^ definition local 0
//          display_name xs
//          signature_documentation
//          > var xs []int
   for _, x := range slices.All(xs) {
//        ^ definition local 1
//          display_name x
//          signature_documentation
//          > var x int
//                   ^^^^^^ reference github.com/golang/go/src go1.22 slices/
//                          ^^^ reference github.com/golang/go/src go1.22 slices/All().
//                              ^^ reference local 0
    return x
//         ^ reference local 1
   }
   return -1
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/f().
  
