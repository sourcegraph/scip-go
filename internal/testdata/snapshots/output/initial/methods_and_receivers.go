  package initial
//        ^^^^^^^ reference 0.1.test sg/initial/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.19 fmt/
  
  type MyStruct struct{ f, y int }
//     ^^^^^^^^ definition 0.1.test sg/initial/MyStruct#
//     documentation ```go
//     documentation ```go
//                      ^ definition 0.1.test sg/initial/MyStruct#f.
//                      documentation ```go
//                         ^ definition 0.1.test sg/initial/MyStruct#y.
//                         documentation ```go
  
  func (m MyStruct) RecvFunction(b int) int { return m.f + b }
//      ^ definition local 0
//        ^^^^^^^^ reference 0.1.test sg/initial/MyStruct#
//                  ^^^^^^^^^^^^ definition 0.1.test sg/initial/MyStruct#RecvFunction().
//                  documentation ```go
//                               ^ definition local 1
//                                                   ^ reference local 0
//                                                     ^ reference 0.1.test sg/initial/MyStruct#f.
//                                                         ^ reference local 1
  
  func SomethingElse() {
//     ^^^^^^^^^^^^^ definition 0.1.test sg/initial/SomethingElse().
//     documentation ```go
   s := MyStruct{f: 0}
// ^ definition local 2
//      ^^^^^^^^ reference 0.1.test sg/initial/MyStruct#
//               ^ reference 0.1.test sg/initial/MyStruct#f.
   fmt.Println(s)
// ^^^ reference github.com/golang/go/src go1.19 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.19 fmt/Println().
//             ^ reference local 2
  }
  
