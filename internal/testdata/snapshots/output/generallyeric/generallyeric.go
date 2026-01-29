  // generallyeric -> generic for short
  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test `sg/generallyeric`/
//        documentation
//        > generallyeric -> generic for short
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/Print().
  func Print[T any](s []T) {
//     ^^^^^ definition 0.1.test `sg/generallyeric`/Print().
//     kind Function
//     documentation
//     > ```go
//     > func Print[T any](s []T)
//     > ```
//           ^ definition local 0
//           kind Interface
//                  ^ definition local 1
//                  kind Variable
//                      ^ reference local 0
   for _, v := range s {
//        ^ definition local 2
//        kind Variable
//                   ^ reference local 1
    fmt.Print(v)
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
//      ^^^^^ reference github.com/golang/go/src go1.22 fmt/Print().
//            ^ reference local 2
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/Print().
  
