  // generallyeric -> generic for short
  package generallyeric
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src fmt/
  
  func Print[T any](s []T) {
//     ^^^^^ definition sg/generallyeric/Print().
//           ^ definition local 0
//                  ^ definition local 1
//                      ^ reference local 0
   for _, v := range s {
//        ^ definition local 2
//                   ^ reference local 1
    fmt.Print(v)
//      ^^^^^ reference github.com/golang/go/src fmt/Print().
//            ^ reference local 2
   }
  }
  
