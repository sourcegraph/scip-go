  package inlinestruct
//        ^^^^^^^^^^^^ reference sg/inlinestruct/
  
  type Params struct{}
//     ^^^^^^ definition sg/inlinestruct/Params#
//     documentation ```go
//     documentation ```go
  type HighlightedCode struct{}
//     ^^^^^^^^^^^^^^^ definition sg/inlinestruct/HighlightedCode#
//     documentation ```go
//     documentation ```go
  
  var Mocks, emptyMocks struct {
//    ^^^^^ definition Mocks.
//    documentation ```go
//           ^^^^^^^^^^ definition emptyMocks.
//           documentation ```go
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition sg/inlinestruct/inline-18:Code.
// documentation ```go
//           ^ definition local 0
//             ^^^^^^ reference sg/inlinestruct/Params#
//                      ^^^^^^^^ definition local 1
//                                ^^^^^^^^^^^^^^^ reference sg/inlinestruct/HighlightedCode#
//                                                 ^^^^^^^ definition local 2
//                                                               ^^^ definition local 3
  }
  
  var MocksSingle struct {
//    ^^^^^^^^^^^ definition MocksSingle.
//    documentation ```go
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition sg/inlinestruct/MocksSingle:Code.
// documentation ```go
//           ^ definition local 4
//             ^^^^^^ reference sg/inlinestruct/Params#
//                      ^^^^^^^^ definition local 5
//                                ^^^^^^^^^^^^^^^ reference sg/inlinestruct/HighlightedCode#
//                                                 ^^^^^^^ definition local 6
//                                                               ^^^ definition local 7
  }
  
  var (
   okReply   interface{} = "OK"
// ^^^^^^^ definition okReply.
// documentation ```go
   pongReply interface{} = "PONG"
// ^^^^^^^^^ definition pongReply.
// documentation ```go
  )
  
