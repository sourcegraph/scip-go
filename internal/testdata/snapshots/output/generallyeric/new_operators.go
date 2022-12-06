  package generallyeric
//        ^^^^^^^^^^^^^ reference sg/generallyeric/
  
  import "golang.org/x/exp/constraints"
//        ^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference golang.org/x/exp golang.org/x/exp/constraints/
  
  type Number interface {
//     ^^^^^^ definition sg/generallyeric/Number#
//     documentation ```go
//     documentation ```go
   constraints.Float | constraints.Integer | constraints.Complex
// ^^^^^^^^^^^ reference golang.org/x/exp golang.org/x/exp/constraints/
//             ^^^^^ reference golang.org/x/exp golang.org/x/exp/constraints/Float#
//                     ^^^^^^^^^^^ reference golang.org/x/exp golang.org/x/exp/constraints/
//                                 ^^^^^^^ reference golang.org/x/exp golang.org/x/exp/constraints/Integer#
//                                           ^^^^^^^^^^^ reference golang.org/x/exp golang.org/x/exp/constraints/
//                                                       ^^^^^^^ reference golang.org/x/exp golang.org/x/exp/constraints/Complex#
  }
  
  func Double[T Number](value T) T {
//     ^^^^^^ definition sg/generallyeric/Double().
//     documentation ```go
//            ^ definition local 0
//              ^^^^^^ reference sg/generallyeric/Number#
//                      ^^^^^ definition local 1
//                            ^ reference local 0
//                               ^ reference local 0
   return value * 2
//        ^^^^^ reference local 1
  }
  
  type Box[T any] struct {
//     ^^^ definition sg/generallyeric/Box#
//     documentation ```go
//     documentation ```go
//         ^ definition local 2
   Something T
// ^^^^^^^^^ definition sg/generallyeric/Box#Something.
// documentation ```go
//           ^ reference local 2
  }
  
  type handler[T any] struct {
//     ^^^^^^^ definition sg/generallyeric/handler#
//     documentation ```go
//     documentation ```go
//             ^ definition local 3
   Box[T]
// ^^^ definition sg/generallyeric/handler#Box.
// documentation ```go
// ^^^ reference sg/generallyeric/Box#
//     ^ reference local 3
   Another string
// ^^^^^^^ definition sg/generallyeric/handler#Another.
// documentation ```go
  }
  
