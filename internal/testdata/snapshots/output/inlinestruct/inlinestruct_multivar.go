  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/
  
  type Params struct{}
//     ^^^^^^ definition 0.1.test sg/inlinestruct/Params#
//     documentation ```go
//     documentation ```go
  type HighlightedCode struct{}
//     ^^^^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/HighlightedCode#
//     documentation ```go
//     documentation ```go
  
  var Mocks, emptyMocks struct {
//    ^^^^^ definition 0.1.test sg/inlinestruct/Mocks.
//    documentation ```go
//           ^^^^^^^^^^ definition 0.1.test sg/inlinestruct/emptyMocks.
//           documentation ```go
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test sg/inlinestruct/inline-30:Code.
// documentation ```go
//           ^ definition local 0
//             ^^^^^^ reference 0.1.test sg/inlinestruct/Params#
//                      ^^^^^^^^ definition local 1
//                                ^^^^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/HighlightedCode#
//                                                 ^^^^^^^ definition local 2
//                                                               ^^^ definition local 3
  }
  
  var MocksSingle struct {
//    ^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/MocksSingle.
//    documentation ```go
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test sg/inlinestruct/MocksSingle:Code.
// documentation ```go
//           ^ definition local 4
//             ^^^^^^ reference 0.1.test sg/inlinestruct/Params#
//                      ^^^^^^^^ definition local 5
//                                ^^^^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/HighlightedCode#
//                                                 ^^^^^^^ definition local 6
//                                                               ^^^ definition local 7
  }
  
  var (
   okReply   interface{} = "OK"
// ^^^^^^^ definition 0.1.test sg/inlinestruct/okReply.
// documentation ```go
   pongReply interface{} = "PONG"
// ^^^^^^^^^ definition 0.1.test sg/inlinestruct/pongReply.
// documentation ```go
  )
  
