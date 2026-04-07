  package generallyeric
//        ^^^^^^^^^^^^^ reference 0.1.test `sg/generallyeric`/
  
  import "fmt"
//        ^^^ definition local 0
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type Person interface {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Person#
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
// documentation
// > ```go
// > func (Person).Work()
// > ```
  }
  
  type worker string
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/worker#
//     documentation
//     > ```go
//     > string
//     > ```
//     relationship 0.1.test `sg/generallyeric`/Person# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/worker#Work().
  func (w worker) Work() {
//      ^ definition local 1
//        ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
//                ^^^^ definition 0.1.test `sg/generallyeric`/worker#Work().
//                documentation
//                > ```go
//                > func (worker).Work()
//                > ```
//                relationship 0.1.test `sg/generallyeric`/Person#Work. implementation
   fmt.Printf("%s is working\n", w)
// ^^^ reference local 0
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                               ^ reference local 1
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/worker#Work().
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/DoWork().
  func DoWork[T Person](things []T) {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/DoWork().
//     documentation
//     > ```go
//     > func DoWork[T Person](things []T)
//     > ```
//            ^ definition local 2
//              ^^^^^^ reference 0.1.test `sg/generallyeric`/Person#
//                      ^^^^^^ definition local 3
//                               ^ reference local 2
   for _, v := range things {
//        ^ definition local 4
//                   ^^^^^^ reference local 3
    v.Work()
//  ^ reference local 4
//    ^^^^ reference 0.1.test `sg/generallyeric`/Person#Work.
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/DoWork().
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/main().
  func main() {
//     ^^^^ definition 0.1.test `sg/generallyeric`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
   var a, b, c worker
//     ^ definition local 5
//        ^ definition local 6
//           ^ definition local 7
//             ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
   a = "A"
// ^ reference local 5
   b = "B"
// ^ reference local 6
   c = "C"
// ^ reference local 7
   DoWork([]worker{a, b, c})
// ^^^^^^ reference 0.1.test `sg/generallyeric`/DoWork().
//          ^^^^^^ reference 0.1.test `sg/generallyeric`/worker#
//                 ^ reference local 5
//                    ^ reference local 6
//                       ^ reference local 7
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/main().
  
