  package generallyeric
//        ^^^^^^^^^^^^^ reference 0.1.test sg/generallyeric/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.19 fmt/
  
  type Person interface {
//     ^^^^^^ definition 0.1.test sg/generallyeric/Person#
//     documentation ```go
//     documentation ```go
   Work()
// ^^^^ definition 0.1.test sg/generallyeric/Person#Work.
// documentation ```go
  }
  
  type worker string
//     ^^^^^^ definition 0.1.test sg/generallyeric/worker#
//     documentation ```go
//     relationship 0.1.test sg/generallyeric/Person# implementation
  
  func (w worker) Work() {
//      ^ definition local 0
//        ^^^^^^ reference 0.1.test sg/generallyeric/worker#
//                ^^^^ definition 0.1.test sg/generallyeric/worker#Work().
//                documentation ```go
//                relationship 0.1.test sg/generallyeric/Person#Work. implementation
   fmt.Printf("%s is working\n", w)
// ^^^ reference github.com/golang/go/src go1.19 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.19 fmt/Printf().
//                               ^ reference local 0
  }
  
  func DoWork[T Person](things []T) {
//     ^^^^^^ definition 0.1.test sg/generallyeric/DoWork().
//     documentation ```go
//            ^ definition local 1
//              ^^^^^^ reference 0.1.test sg/generallyeric/Person#
//                      ^^^^^^ definition local 2
//                               ^ reference local 1
   for _, v := range things {
//        ^ definition local 3
//                   ^^^^^^ reference local 2
    v.Work()
//  ^ reference local 3
//    ^^^^ reference 0.1.test sg/generallyeric/Person#Work.
   }
  }
  
  func main() {
//     ^^^^ definition 0.1.test sg/generallyeric/main().
//     documentation ```go
   var a, b, c worker
//     ^ definition local 4
//        ^ definition local 5
//           ^ definition local 6
//             ^^^^^^ reference 0.1.test sg/generallyeric/worker#
   a = "A"
// ^ reference local 4
   b = "B"
// ^ reference local 5
   c = "C"
// ^ reference local 6
   DoWork([]worker{a, b, c})
// ^^^^^^ reference 0.1.test sg/generallyeric/DoWork().
//          ^^^^^^ reference 0.1.test sg/generallyeric/worker#
//                 ^ reference local 4
//                    ^ reference local 5
//                       ^ reference local 6
  }
  
