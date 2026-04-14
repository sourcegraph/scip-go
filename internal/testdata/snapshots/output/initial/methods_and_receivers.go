  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type MyStruct struct{ f, y int }
//     ^^^^^^^^ definition 0.1.test `sg/initial`/MyStruct#
//              signature_documentation
//              > type MyStruct struct {
//              >     f int
//              >     y int
//              > }
//                      ^ definition 0.1.test `sg/initial`/MyStruct#f.
//                        signature_documentation
//                        > struct field f int
//                         ^ definition 0.1.test `sg/initial`/MyStruct#y.
//                           signature_documentation
//                           > struct field y int
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/MyStruct#RecvFunction().
  func (m MyStruct) RecvFunction(b int) int { return m.f + b }
//      ^ definition local 0
//        display_name m
//        signature_documentation
//        > var m MyStruct
//        ^^^^^^^^ reference 0.1.test `sg/initial`/MyStruct#
//                  ^^^^^^^^^^^^ definition 0.1.test `sg/initial`/MyStruct#RecvFunction().
//                               signature_documentation
//                               > func (MyStruct).RecvFunction(b int) int
//                               ^ definition local 1
//                                 display_name b
//                                 signature_documentation
//                                 > var b int
//                                                   ^ reference local 0
//                                                     ^ reference 0.1.test `sg/initial`/MyStruct#f.
//                                                         ^ reference local 1
//                                                           ⌃ enclosing_range_end 0.1.test `sg/initial`/MyStruct#RecvFunction().
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/SomethingElse().
  func SomethingElse() {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/SomethingElse().
//                   signature_documentation
//                   > func SomethingElse()
   s := MyStruct{f: 0}
// ^ definition local 2
//   display_name s
//   signature_documentation
//   > var s MyStruct
//      ^^^^^^^^ reference 0.1.test `sg/initial`/MyStruct#
//               ^ reference 0.1.test `sg/initial`/MyStruct#f.
   fmt.Println(s)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^ reference local 2
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/SomethingElse().
  
