  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  type LiteralType int
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/LiteralType#
//                 display_name LiteralType
//                 signature_documentation
//                 > type LiteralType int
  
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition 0.1.test `sg/initial`/FuncType#
//              display_name FuncType
//              signature_documentation
//              > type FuncType func(LiteralType, int) bool
//                   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
  type IfaceType interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/IfaceType#
//               display_name IfaceType
//               signature_documentation
//               > type IfaceType interface{ Method() LiteralType }
   Method() LiteralType
// ^^^^^^ definition 0.1.test `sg/initial`/IfaceType#Method.
//        display_name Method
//        signature_documentation
//        > func (IfaceType).Method() LiteralType
//          ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  }
  
  type StructType struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#
//                display_name StructType
//                signature_documentation
//                > type StructType struct {
//                >     m    IfaceType
//                >     f    LiteralType
//                >     anon struct{ sub int }
//                >     i    interface{ AnonMethod() bool }
//                > }
   m IfaceType
// ^ definition 0.1.test `sg/initial`/StructType#m.
//   display_name m
//   signature_documentation
//   > struct field m IfaceType
//   ^^^^^^^^^ reference 0.1.test `sg/initial`/IfaceType#
   f LiteralType
// ^ definition 0.1.test `sg/initial`/StructType#f.
//   display_name f
//   signature_documentation
//   > struct field f LiteralType
//   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
   // anonymous struct
   anon struct {
// ^^^^ definition 0.1.test `sg/initial`/StructType#anon.
//      display_name anon
//      signature_documentation
//      > struct field anon struct{sub int}
//      documentation
//      > anonymous struct
    sub int
//  ^^^ definition 0.1.test `sg/initial`/StructType#$anon_0ba9ace1dcfd6761#sub.
//      display_name sub
//      signature_documentation
//      > struct field sub int
   }
  
   // interface within struct
   i interface {
// ^ definition 0.1.test `sg/initial`/StructType#i.
//   display_name i
//   signature_documentation
//   > struct field i interface{AnonMethod() bool}
//   documentation
//   > interface within struct
    AnonMethod() bool
//  ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#$anon_97e7de633e3ef8e8#AnonMethod.
//             display_name AnonMethod
//             signature_documentation
//             > func (interface).AnonMethod() bool
   }
  }
  
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#
//                    display_name DeclaredBefore
//                    signature_documentation
//                    > type DeclaredBefore struct{ DeclaredAfter }
//                            ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#DeclaredAfter.
//                                          display_name DeclaredAfter
//                                          signature_documentation
//                                          > struct field DeclaredAfter DeclaredAfter
//                            ^^^^^^^^^^^^^ reference 0.1.test `sg/initial`/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredAfter#
//                   display_name DeclaredAfter
//                   signature_documentation
//                   > type DeclaredAfter struct{}
  
