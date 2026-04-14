  package generallyeric
//        ^^^^^^^^^^^^^ definition 0.1.test `sg/generallyeric`/
  
  type Number interface {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Number#
//            signature_documentation
//            > type Number interface {
//            >     ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
//            > }
   ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~float32 | ~float64
  }
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/Double().
  func Double[T Number](value T) T {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Double().
//            signature_documentation
//            > func Double[T Number](value T) T
//            ^ definition local 0
//              display_name T
//              signature_documentation
//              > T T
//              ^^^^^^ reference 0.1.test `sg/generallyeric`/Number#
//                      ^^^^^ definition local 1
//                            display_name value
//                            signature_documentation
//                            > var value T
//                            ^ reference local 0
//                               ^ reference local 0
   return value * 2
//        ^^^^^ reference local 1
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/Double().
  
  type Box[T any] struct {
//     ^^^ definition 0.1.test `sg/generallyeric`/Box#
//         signature_documentation
//         > type Box struct{ Something T }
//         ^ definition local 2
//           display_name T
//           signature_documentation
//           > T T
   Something T
// ^^^^^^^^^ definition 0.1.test `sg/generallyeric`/Box#Something.
//           signature_documentation
//           > struct field Something T
//           ^ reference local 2
  }
  
  type handler[T any] struct {
//     ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#
//             signature_documentation
//             > type handler struct {
//             >     Box[T]
//             >     Another string
//             > }
//             ^ definition local 3
//               display_name T
//               signature_documentation
//               > T T
   Box[T]
// ^^^ definition 0.1.test `sg/generallyeric`/handler#Box.
//     signature_documentation
//     > struct field Box sg/generallyeric.Box[T]
// ^^^ reference 0.1.test `sg/generallyeric`/Box#
//     ^ reference local 3
   Another string
// ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#Another.
//         signature_documentation
//         > struct field Another string
  }
  
