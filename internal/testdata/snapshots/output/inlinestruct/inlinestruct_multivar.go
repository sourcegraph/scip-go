  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/
  
  type Params struct{}
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/Params#
//     documentation
//     > ```go
//     > type Params struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  type HighlightedCode struct{}
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/HighlightedCode#
//     documentation
//     > ```go
//     > type HighlightedCode struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
  var Mocks, emptyMocks struct {
//    ^^^^^ definition 0.1.test `sg/inlinestruct`/Mocks.
//    documentation
//    > ```go
//    > var Mocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//    > ```
//           ^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/emptyMocks.
//           documentation
//           > ```go
//           > var emptyMocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//           > ```
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/inline-6-5:Code.
// documentation
// > ```go
// > struct field Code func(p sg/inlinestruct.Params) (response *sg/inlinestruct.HighlightedCode, aborted bool, err error)
// > ```
//           ^ definition local 0
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 1
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 2
//                                                               ^^^ definition local 3
  }
  
  var MocksSingle struct {
//    ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle.
//    documentation
//    > ```go
//    > var MocksSingle struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//    > ```
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle:Code.
// documentation
// > ```go
// > struct field Code func(p sg/inlinestruct.Params) (response *sg/inlinestruct.HighlightedCode, aborted bool, err error)
// > ```
//           ^ definition local 4
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 5
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 6
//                                                               ^^^ definition local 7
  }
  
  var (
   okReply   interface{} = "OK"
// ^^^^^^^ definition 0.1.test `sg/inlinestruct`/okReply.
// documentation
// > ```go
// > var okReply interface{}
// > ```
   pongReply interface{} = "PONG"
// ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/pongReply.
// documentation
// > ```go
// > var pongReply interface{}
// > ```
  )
  
