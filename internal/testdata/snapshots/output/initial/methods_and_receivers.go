  package initial
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src fmt/
  
  type MyStruct struct{ f, y int }
//     ^^^^^^^^ definition sg/initial/MyStruct#
//     documentation ```go
//     documentation ```go
//                      ^ definition sg/initial/MyStruct#f.
//                      documentation ```go
//                         ^ definition sg/initial/MyStruct#y.
//                         documentation ```go
  
  func (m MyStruct) RecvFunction(b int) int { return m.f + b }
//      ^ definition local 0
//        ^^^^^^^^ reference sg/initial/MyStruct#
//                  ^^^^^^^^^^^^ definition sg/initial/MyStruct#RecvFunction().
//                  documentation ```go
//                               ^ definition local 1
//                                                   ^ reference local 0
//                                                     ^ reference sg/initial/MyStruct#f.
//                                                         ^ reference local 1
  
  func SomethingElse() {
//     ^^^^^^^^^^^^^ definition sg/initial/SomethingElse().
//     documentation ```go
   s := MyStruct{f: 0}
// ^ definition local 2
//      ^^^^^^^^ reference sg/initial/MyStruct#
//               ^ reference sg/initial/MyStruct#f.
   fmt.Println(s)
//     ^^^^^^^ reference github.com/golang/go/src fmt/Println().
//             ^ reference local 2
  }
  
