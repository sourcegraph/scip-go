  package generallyeric
//        ^^^^^^^^^^^^^ reference 0.1.test sg/generallyeric/
  
  import "golang.org/x/exp/constraints"
//        ^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference golang.org/x/exp v0.0.0-20221205204356-47842c84f3db golang.org/x/exp/constraints/
  
  type Number interface {
//     ^^^^^^ definition 0.1.test sg/generallyeric/Number#
//     documentation ```go
//     documentation ```go
   constraints.Float | constraints.Integer | constraints.Complex
// ^^^^^^^^^^^ reference golang.org/x/exp v0.0.0-20221205204356-47842c84f3db golang.org/x/exp/constraints/
//             ^^^^^ reference golang.org/x/exp v0.0.0-20221205204356-47842c84f3db golang.org/x/exp/constraints/Float#
//                     ^^^^^^^^^^^ reference golang.org/x/exp v0.0.0-20221205204356-47842c84f3db golang.org/x/exp/constraints/
//                                 ^^^^^^^ reference golang.org/x/exp v0.0.0-20221205204356-47842c84f3db golang.org/x/exp/constraints/Integer#
//                                           ^^^^^^^^^^^ reference golang.org/x/exp v0.0.0-20221205204356-47842c84f3db golang.org/x/exp/constraints/
//                                                       ^^^^^^^ reference golang.org/x/exp v0.0.0-20221205204356-47842c84f3db golang.org/x/exp/constraints/Complex#
  }
  
  func Double[T Number](value T) T {
//     ^^^^^^ definition 0.1.test sg/generallyeric/Double().
//     documentation ```go
//            ^ definition local 0
//              ^^^^^^ reference 0.1.test sg/generallyeric/Number#
//                      ^^^^^ definition local 1
//                            ^ reference local 0
//                               ^ reference local 0
   return value * 2
//        ^^^^^ reference local 1
  }
  
  type Box[T any] struct {
//     ^^^ definition 0.1.test sg/generallyeric/Box#
//     documentation ```go
//     documentation ```go
//         ^ definition local 2
   Something T
// ^^^^^^^^^ definition 0.1.test sg/generallyeric/Box#Something.
// documentation ```go
//           ^ reference local 2
  }
  
  type handler[T any] struct {
//     ^^^^^^^ definition 0.1.test sg/generallyeric/handler#
//     documentation ```go
//     documentation ```go
//             ^ definition local 3
   Box[T]
// ^^^ definition 0.1.test sg/generallyeric/handler#Box.
// documentation ```go
// ^^^ reference 0.1.test sg/generallyeric/Box#
//     ^ reference local 3
   Another string
// ^^^^^^^ definition 0.1.test sg/generallyeric/handler#Another.
// documentation ```go
  }
  
