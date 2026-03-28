  package generallyeric
//        ^^^^^^^^^^^^^ reference 0.1.test `sg/generallyeric`/
  
  type Number interface {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Number#
//     documentation
//     > ```go
//     > type Number interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
//     > }
//     > ```
   ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~float32 | ~float64
  }
  
//⌄ enclosing_range_start 0.1.test `sg/generallyeric`/Double().
  func Double[T Number](value T) T {
//     ^^^^^^ definition 0.1.test `sg/generallyeric`/Double().
//     documentation
//     > ```go
//     > func Double[T Number](value T) T
//     > ```
//            ^ definition local 0
//              ^^^^^^ reference 0.1.test `sg/generallyeric`/Number#
//                      ^^^^^ definition local 1
//                            ^ reference local 0
//                               ^ reference local 0
   return value * 2
//        ^^^^^ reference local 1
  }
//⌃ enclosing_range_end 0.1.test `sg/generallyeric`/Double().
  
  type Box[T any] struct {
//     ^^^ definition 0.1.test `sg/generallyeric`/Box#
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
   Something T
// ^^^^^^^^^ definition 0.1.test `sg/generallyeric`/Box#Something.
// documentation
// > ```go
// > struct field Something T
// > ```
//           ^ reference local 2
  }
  
  type handler[T any] struct {
//     ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#
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
   Box[T]
// ^^^ definition 0.1.test `sg/generallyeric`/handler#Box.
// documentation
// > ```go
// > struct field Box sg/generallyeric.Box[T]
// > ```
// ^^^ reference 0.1.test `sg/generallyeric`/Box#
//     ^ reference local 3
   Another string
// ^^^^^^^ definition 0.1.test `sg/generallyeric`/handler#Another.
// documentation
// > ```go
// > struct field Another string
// > ```
  }
  
