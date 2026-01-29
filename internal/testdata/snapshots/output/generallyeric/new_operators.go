  package generallyeric
//        ^^^^^^^^^^^^^ reference 0.1.test `sg/generallyeric`/
  
  import "golang.org/x/exp/constraints"
//        ^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference golang.org/x/exp 47842c84f3db `golang.org/x/exp/constraints`/
  
  type Number interface {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Number#
//     kind Interface
//     documentation
//     > ```go
//     > type Number interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Float | Integer | Complex
//     > }
//     > ```
   constraints.Float | constraints.Integer | constraints.Complex
// ^^^^^^^^^^^ reference golang.org/x/exp 47842c84f3db `golang.org/x/exp/constraints`/
//             ^^^^^ reference golang.org/x/exp 47842c84f3db `golang.org/x/exp/constraints`/Float#
//                     ^^^^^^^^^^^ reference golang.org/x/exp 47842c84f3db `golang.org/x/exp/constraints`/
//                                 ^^^^^^^ reference golang.org/x/exp 47842c84f3db `golang.org/x/exp/constraints`/Integer#
//                                           ^^^^^^^^^^^ reference golang.org/x/exp 47842c84f3db `golang.org/x/exp/constraints`/
//                                                       ^^^^^^^ reference golang.org/x/exp 47842c84f3db `golang.org/x/exp/constraints`/Complex#
  }
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/Double().
  func Double[T Number](value T) T {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Double().
//     kind Function
//     documentation
//     > ```go
//     > func Double[T Number](value T) T
//     > ```
//            ^ definition local 0
//            kind Interface
//              ^^^^^^ reference 0.1.test `sg/generallyeric`/Number#
//                      ^^^^^ definition local 1
//                      kind Variable
//                            ^ reference local 0
//                               ^ reference local 0
   return value * 2
//        ^^^^^ reference local 1
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/Double().
  
  type Box[T any] struct {
//     ^^^ definition 0.1.test `sg/generallyeric`/Box#
//     kind Class
//     documentation
//     > ```go
//     > type Box struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Something T
//     > }
//     > ```
//         ^ definition local 2
//         kind Interface
   Something T
// ^^^^^^^^^ definition 0.1.test `sg/generallyeric`/Box#Something.
// kind Field
// documentation
// > ```go
// > struct field Something T
// > ```
//           ^ reference local 2
  }
  
  type handler[T any] struct {
//     ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#
//     kind Class
//     documentation
//     > ```go
//     > type handler struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Box[T]
//     >     Another string
//     > }
//     > ```
//             ^ definition local 3
//             kind Interface
   Box[T]
// ^^^ definition 0.1.test `sg/generallyeric`/handler#Box.
// kind Field
// documentation
// > ```go
// > struct field Box sg/generallyeric.Box[T]
// > ```
// ^^^ reference 0.1.test `sg/generallyeric`/Box#
//     ^ reference local 3
   Another string
// ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#Another.
// kind Field
// documentation
// > ```go
// > struct field Another string
// > ```
  }
  
