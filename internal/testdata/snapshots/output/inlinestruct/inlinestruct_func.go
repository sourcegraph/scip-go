  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
  type InFuncSig struct {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#
//               signature_documentation
//               > type InFuncSig struct{ value bool }
   value bool
// ^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#value.
//       signature_documentation
//       > struct field value bool
  }
  
  var rowsCloseHook = func() func(InFuncSig, *error) { return nil }
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/rowsCloseHook.
//                  signature_documentation
//                  > var rowsCloseHook func() func(InFuncSig, *error)
//                                ^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/InFuncSig#
  
