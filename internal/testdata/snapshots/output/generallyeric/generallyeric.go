  // generallyeric -> generic for short
  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test `sg/generallyeric`/
//        documentation
//        > generallyeric -> generic for short
  
  import "fmt"
//        ^^^ definition local 0
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/Print().
  func Print[T any](s []T) {
//     ^^^^^ definition 0.1.test `sg/generallyeric`/Print().
//     documentation
//     > ```go
//     > func Print[T any](s []T)
//     > ```
//           ^ definition local 1
//                  ^ definition local 2
//                      ^ reference local 1
   for _, v := range s {
//        ^ definition local 3
//                   ^ reference local 2
    fmt.Print(v)
//  ^^^ reference local 0
//      ^^^^^ reference github.com/golang/go/src go1.22 fmt/Print().
//            ^ reference local 3
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/Print().
  
