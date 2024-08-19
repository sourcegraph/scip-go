  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/
  
  type InFuncSig struct {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#
//     documentation
//     > ```go
//     > type InFuncSig struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     value bool
//     > }
//     > ```
   value bool
// ^^^^^ definition 0.1.test `sg/inlinestruct`/InFuncSig#value.
// documentation
// > ```go
// > struct field value bool
// > ```
  }
  
  var rowsCloseHook = func() func(InFuncSig, *error) { return nil }
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/rowsCloseHook.
//    documentation
//    > ```go
//    > var rowsCloseHook func() func(InFuncSig, *error)
//    > ```
//                                ^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/InFuncSig#
  
