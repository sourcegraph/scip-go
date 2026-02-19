  package generallyeric
//        ^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type Person interface {
//     ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/Person#
//     documentation
//     > ```go
//     > type Person interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Work()
//     > }
//     > ```
   Work()
// ^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/Person#Work.
// documentation
// > ```go
// > func (Person).Work()
// > ```
  }
  
  type worker string
//     ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/worker#
//     documentation
//     > ```go
//     > string
//     > ```
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/Person# implementation
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/worker#Work().
  func (w worker) Work() {
//      ^ definition local 0
//        ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/worker#
//                ^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/worker#Work().
//                documentation
//                > ```go
//                > func (worker).Work()
//                > ```
//                relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/Person#Work. implementation
   fmt.Printf("%s is working\n", w)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                               ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/worker#Work().
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/DoWork().
  func DoWork[T Person](things []T) {
//     ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/DoWork().
//     documentation
//     > ```go
//     > func DoWork[T Person](things []T)
//     > ```
//            ^ definition local 1
//              ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/Person#
//                      ^^^^^^ definition local 2
//                               ^ reference local 1
   for _, v := range things {
//        ^ definition local 3
//                   ^^^^^^ reference local 2
    v.Work()
//  ^ reference local 3
//    ^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/Person#Work.
   }
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/DoWork().
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/main().
  func main() {
//     ^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   var a, b, c worker
//     ^ definition local 4
//        ^ definition local 5
//           ^ definition local 6
//             ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/worker#
   a = "A"
// ^ reference local 4
   b = "B"
// ^ reference local 5
   c = "C"
// ^ reference local 6
   DoWork([]worker{a, b, c})
// ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/DoWork().
//          ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/worker#
//                 ^ reference local 4
//                    ^ reference local 5
//                       ^ reference local 6
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/generallyeric`/main().
  
