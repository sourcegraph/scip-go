  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
  type Params struct{}
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/Params#
//            signature_documentation
//            > type Params struct{}
  type HighlightedCode struct{}
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/HighlightedCode#
//                     signature_documentation
//                     > type HighlightedCode struct{}
  
  var Mocks, emptyMocks struct {
//    ^^^^^ definition 0.1.test `sg/inlinestruct`/Mocks.
//          signature_documentation
//          > var Mocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
//           ^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/emptyMocks.
//                      signature_documentation
//                      > var emptyMocks struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/inline-6-5:Code.
//      signature_documentation
//      > struct field Code func(p sg/inlinestruct.Params) (response *sg/inlinestruct.HighlightedCode, aborted bool, err error)
//           ^ definition local 0
//             display_name p
//             signature_documentation
//             > var p sg/inlinestruct.Params
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 1
//                               display_name response
//                               signature_documentation
//                               > var response *sg/inlinestruct.HighlightedCode
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 2
//                                                         display_name aborted
//                                                         signature_documentation
//                                                         > var aborted bool
//                                                               ^^^ definition local 3
//                                                                   display_name err
//                                                                   signature_documentation
//                                                                   > var err error
  }
  
  var MocksSingle struct {
//    ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle.
//                signature_documentation
//                > var MocksSingle struct{Code func(p Params) (response *HighlightedCode, aborted bool, err error)}
   Code func(p Params) (response *HighlightedCode, aborted bool, err error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/MocksSingle:Code.
//      signature_documentation
//      > struct field Code func(p sg/inlinestruct.Params) (response *sg/inlinestruct.HighlightedCode, aborted bool, err error)
//           ^ definition local 4
//             display_name p
//             signature_documentation
//             > var p sg/inlinestruct.Params
//             ^^^^^^ reference 0.1.test `sg/inlinestruct`/Params#
//                      ^^^^^^^^ definition local 5
//                               display_name response
//                               signature_documentation
//                               > var response *sg/inlinestruct.HighlightedCode
//                                ^^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/HighlightedCode#
//                                                 ^^^^^^^ definition local 6
//                                                         display_name aborted
//                                                         signature_documentation
//                                                         > var aborted bool
//                                                               ^^^ definition local 7
//                                                                   display_name err
//                                                                   signature_documentation
//                                                                   > var err error
  }
  
  var (
   okReply   interface{} = "OK"
// ^^^^^^^ definition 0.1.test `sg/inlinestruct`/okReply.
//         signature_documentation
//         > var okReply interface{}
   pongReply interface{} = "PONG"
// ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/pongReply.
//           signature_documentation
//           > var pongReply interface{}
  )
  
