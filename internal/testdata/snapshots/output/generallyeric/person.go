  package generallyeric
//        ^^^^^^^^^^^^^ reference 0.1.test `sg/generallyeric`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type Person interface {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Person#
//     kind Interface
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
// ^^^^ definition 0.1.test `sg/generallyeric`/Person#Work.
// kind Method
// documentation
// > ```go
// > func (Person).Work()
// > ```
  }
  
  type worker string
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/worker#
//     kind Type
//     documentation
//     > ```go
//     > string
//     > ```
//     relationship 0.1.test `sg/generallyeric`/Person# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/worker#Work().
  func (w worker) Work() {
//      ^ definition local 0
//      kind Variable
//        ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
//                ^^^^ definition 0.1.test `sg/generallyeric`/worker#Work().
//                kind Method
//                documentation
//                > ```go
//                > func (worker).Work()
//                > ```
//                relationship 0.1.test `sg/generallyeric`/Person#Work. implementation
   fmt.Printf("%s is working\n", w)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                               ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/worker#Work().
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/DoWork().
  func DoWork[T Person](things []T) {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/DoWork().
//     kind Function
//     documentation
//     > ```go
//     > func DoWork[T Person](things []T)
//     > ```
//            ^ definition local 1
//            kind Interface
//              ^^^^^^ reference 0.1.test `sg/generallyeric`/Person#
//                      ^^^^^^ definition local 2
//                      kind Variable
//                               ^ reference local 1
   for _, v := range things {
//        ^ definition local 3
//        kind Variable
//                   ^^^^^^ reference local 2
    v.Work()
//  ^ reference local 3
//    ^^^^ reference 0.1.test `sg/generallyeric`/Person#Work.
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/DoWork().
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/main().
  func main() {
//     ^^^^ definition 0.1.test `sg/generallyeric`/main().
//     kind Function
//     documentation
//     > ```go
//     > func main()
//     > ```
   var a, b, c worker
//     ^ definition local 4
//     kind Variable
//        ^ definition local 5
//        kind Variable
//           ^ definition local 6
//           kind Variable
//             ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
   a = "A"
// ^ reference local 4
   b = "B"
// ^ reference local 5
   c = "C"
// ^ reference local 6
   DoWork([]worker{a, b, c})
// ^^^^^^ reference 0.1.test `sg/generallyeric`/DoWork().
//          ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
//                 ^ reference local 4
//                    ^ reference local 5
//                       ^ reference local 6
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/main().
  
