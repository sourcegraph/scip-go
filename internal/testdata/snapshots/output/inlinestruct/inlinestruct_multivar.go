  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/
  
  type Params struct{}
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/Params#
//     kind Class
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
//     kind Class
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
//    kind Variable
//    documentation
//    > ```go
//    > var Mocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//    > ```
//           ^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/emptyMocks.
//           kind Variable
//           documentation
//           > ```go
//           > var emptyMocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//           > ```
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/inline-6-5:Code.
// kind Field
// documentation
// > ```go
// > struct field Code func(p sg/inlinestruct.Params) (response *sg/inlinestruct.HighlightedCode, aborted bool, err error)
// > ```
//           ^ definition local 0
//           kind Variable
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 1
//                      kind Variable
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 2
//                                                 kind Variable
//                                                               ^^^ definition local 3
//                                                               kind Variable
  }
  
  var MocksSingle struct {
//    ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle.
//    kind Variable
//    documentation
//    > ```go
//    > var MocksSingle struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//    > ```
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle:Code.
// kind Field
// documentation
// > ```go
// > struct field Code func(p sg/inlinestruct.Params) (response *sg/inlinestruct.HighlightedCode, aborted bool, err error)
// > ```
//           ^ definition local 4
//           kind Variable
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 5
//                      kind Variable
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 6
//                                                 kind Variable
//                                                               ^^^ definition local 7
//                                                               kind Variable
  }
  
  var (
   okReply   interface{} = "OK"
// ^^^^^^^ definition 0.1.test `sg/inlinestruct`/okReply.
// kind Variable
// documentation
// > ```go
// > var okReply interface{}
// > ```
   pongReply interface{} = "PONG"
// ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/pongReply.
// kind Variable
// documentation
// > ```go
// > var pongReply interface{}
// > ```
  )
  
