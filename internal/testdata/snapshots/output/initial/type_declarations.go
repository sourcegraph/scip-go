  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  type LiteralType int
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/LiteralType#
//                 signature_documentation
//                 > type LiteralType int
  
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition 0.1.test `sg/initial`/FuncType#
//              signature_documentation
//              > type FuncType func(LiteralType, int) bool
//                   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
  type IfaceType interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/IfaceType#
//               signature_documentation
//               > type IfaceType interface
//               > interface {
//               >     Method() LiteralType
//               > }
   Method() LiteralType
// ^^^^^^ definition 0.1.test `sg/initial`/IfaceType#Method.
//        signature_documentation
//        > func (IfaceType).Method() LiteralType
//          ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  }
  
  type StructType struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#
//                signature_documentation
//                > type StructType struct
//                > struct {
//                >     m IfaceType
//                >     f LiteralType
//                >     anon struct {
//                >         sub int
//                >     }
//                >     i interface {
//                >         AnonMethod() bool
//                >     }
//                > }
   m IfaceType
// ^ definition 0.1.test `sg/initial`/StructType#m.
//   signature_documentation
//   > struct field m sg/initial.IfaceType
//   ^^^^^^^^^ reference 0.1.test `sg/initial`/IfaceType#
   f LiteralType
// ^ definition 0.1.test `sg/initial`/StructType#f.
//   signature_documentation
//   > struct field f sg/initial.LiteralType
//   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
   // anonymous struct
   anon struct {
// ^^^^ definition 0.1.test `sg/initial`/StructType#anon.
//      signature_documentation
//      > struct field anon struct{sub int}
    sub int
//  ^^^ definition 0.1.test `sg/initial`/StructType#$anon_0ba9ace1dcfd6761#sub.
//      signature_documentation
//      > struct field sub int
   }
  
   // interface within struct
   i interface {
// ^ definition 0.1.test `sg/initial`/StructType#i.
//   signature_documentation
//   > struct field i interface{AnonMethod() bool}
    AnonMethod() bool
//  ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#$anon_97e7de633e3ef8e8#AnonMethod.
//             signature_documentation
//             > func (interface).AnonMethod() bool
   }
  }
  
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#
//                    signature_documentation
//                    > type DeclaredBefore struct
//                    > struct {
//                    >     DeclaredAfter
//                    > }
//                            ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#DeclaredAfter.
//                                          signature_documentation
//                                          > struct field DeclaredAfter sg/initial.DeclaredAfter
//                            ^^^^^^^^^^^^^ reference 0.1.test `sg/initial`/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredAfter#
//                   signature_documentation
//                   > type DeclaredAfter struct
//                   > struct{}
  
