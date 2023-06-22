  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/
  
  type InFuncSig struct {
//     ^^^^^^^^^ definition 0.1.test sg/inlinestruct/InFuncSig#
//     documentation ```go
//     documentation ```go
   value bool
// ^^^^^ definition 0.1.test sg/inlinestruct/InFuncSig#value.
// documentation ```go
  }
  
  var rowsCloseHook = func() func(InFuncSig, *error) { return nil }
//    ^^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/rowsCloseHook.
//    documentation ```go
//                                ^^^^^^^^^ reference 0.1.test sg/inlinestruct/InFuncSig#
  
