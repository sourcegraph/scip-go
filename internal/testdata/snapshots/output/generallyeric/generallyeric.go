  // generallyeric -> generic for short
  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test sg/generallyeric/
  
  import "fmt"
//        ^^^ reference v1.19 fmt/
  
  func Print[T any](s []T) {
//     ^^^^^ definition 0.1.test sg/generallyeric/Print().
//     documentation ```go
//           ^ definition local 0
//                  ^ definition local 1
//                      ^ reference local 0
   for _, v := range s {
//        ^ definition local 2
//                   ^ reference local 1
    fmt.Print(v)
//  ^^^ reference v1.19 fmt/
//      ^^^^^ reference v1.19 fmt/Print().
//            ^ reference local 2
   }
  }
  
