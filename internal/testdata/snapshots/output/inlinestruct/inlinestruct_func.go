  package inlinestruct
//        ^^^^^^^^^^^^ reference sg/inlinestruct/
  
  type InFuncSig struct {
//     ^^^^^^^^^ definition sg/inlinestruct/InFuncSig#
//     documentation ```go
//     documentation ```go
   value bool
// ^^^^^ definition sg/inlinestruct/InFuncSig#value.
// documentation ```go
  }
  
  var rowsCloseHook = func() func(InFuncSig, *error) { return nil }
//    ^^^^^^^^^^^^^ definition rowsCloseHook.
//    documentation ```go
//                                ^^^^^^^^^ reference sg/inlinestruct/InFuncSig#
  
