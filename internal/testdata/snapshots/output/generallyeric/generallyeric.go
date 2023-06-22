  // generallyeric -> generic for short
  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test sg/generallyeric/
//        documentation generallyeric -> generic for short
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.19 fmt/
  
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
//  ^^^ reference github.com/golang/go/src go1.19 fmt/
//      ^^^^^ reference github.com/golang/go/src go1.19 fmt/Print().
//            ^ reference local 2
   }
  }
  
