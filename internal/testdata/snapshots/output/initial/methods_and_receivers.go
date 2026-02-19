  package initial
//        ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type MyStruct struct{ f, y int }
//     ^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#
//     documentation
//     > ```go
//     > type MyStruct struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     f int
//     >     y int
//     > }
//     > ```
//                      ^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#f.
//                      documentation
//                      > ```go
//                      > struct field f int
//                      > ```
//                         ^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#y.
//                         documentation
//                         > ```go
//                         > struct field y int
//                         > ```
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#RecvFunction().
  func (m MyStruct) RecvFunction(b int) int { return m.f + b }
//      ^ definition local 0
//        ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#
//                  ^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#RecvFunction().
//                  documentation
//                  > ```go
//                  > func (MyStruct).RecvFunction(b int) int
//                  > ```
//                               ^ definition local 1
//                                                   ^ reference local 0
//                                                     ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#f.
//                                                         ^ reference local 1
//                                                           ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#RecvFunction().
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/SomethingElse().
  func SomethingElse() {
//     ^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/SomethingElse().
//     documentation
//     > ```go
//     > func SomethingElse()
//     > ```
   s := MyStruct{f: 0}
// ^ definition local 2
//      ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#
//               ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/MyStruct#f.
   fmt.Println(s)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^ reference local 2
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/SomethingElse().
  
