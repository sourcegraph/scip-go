  package generallyeric
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src fmt/
  
  type Person interface {
//     ^^^^^^ definition sg/generallyeric/Person#
//     documentation ```go
//     documentation ```go
   Work()
// ^^^^ definition sg/generallyeric/Person#Work.
// documentation ```go
  }
  
  type worker string
//     ^^^^^^ definition sg/generallyeric/worker#
//     documentation ```go
//     relationship sg/generallyeric/Person# implementation
  
  func (w worker) Work() {
//      ^ definition local 0
//        ^^^^^^ reference sg/generallyeric/worker#
//                ^^^^ definition sg/generallyeric/worker#Work().
//                documentation ```go
//                relationship sg/generallyeric/Person#Work. implementation
   fmt.Printf("%s is working\n", w)
// ^^^ reference github.com/golang/go/src fmt/
//     ^^^^^^ reference github.com/golang/go/src fmt/Printf().
//                               ^ reference local 0
  }
  
  func DoWork[T Person](things []T) {
//     ^^^^^^ definition sg/generallyeric/DoWork().
//     documentation ```go
//            ^ definition local 1
//              ^^^^^^ reference sg/generallyeric/Person#
//                      ^^^^^^ definition local 2
//                               ^ reference local 1
   for _, v := range things {
//        ^ definition local 3
//                   ^^^^^^ reference local 2
    v.Work()
//  ^ reference local 3
//    ^^^^ reference sg/generallyeric/Person#Work.
   }
  }
  
  func main() {
//     ^^^^ definition sg/generallyeric/main().
//     documentation ```go
   var a, b, c worker
//     ^ definition local 4
//        ^ definition local 5
//           ^ definition local 6
//             ^^^^^^ reference sg/generallyeric/worker#
   a = "A"
// ^ reference local 4
   b = "B"
// ^ reference local 5
   c = "C"
// ^ reference local 6
   DoWork([]worker{a, b, c})
// ^^^^^^ reference sg/generallyeric/DoWork().
//          ^^^^^^ reference sg/generallyeric/worker#
//                 ^ reference local 4
//                    ^ reference local 5
//                       ^ reference local 6
  }
  
