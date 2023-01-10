  package initial
//        ^^^^^^^ reference 0.1.test sg/initial/
  
  type LiteralType int
//     ^^^^^^^^^^^ definition 0.1.test sg/initial/LiteralType#
//     documentation ```go
  
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition 0.1.test sg/initial/FuncType#
//     documentation ```go
//                   ^^^^^^^^^^^ reference 0.1.test sg/initial/LiteralType#
  
  type IfaceType interface {
//     ^^^^^^^^^ definition 0.1.test sg/initial/IfaceType#
//     documentation ```go
//     documentation ```go
   Method() LiteralType
// ^^^^^^ definition 0.1.test sg/initial/IfaceType#Method.
// documentation ```go
//          ^^^^^^^^^^^ reference 0.1.test sg/initial/LiteralType#
  }
  
  type StructType struct {
//     ^^^^^^^^^^ definition 0.1.test sg/initial/StructType#
//     documentation ```go
//     documentation ```go
   m IfaceType
// ^ definition 0.1.test sg/initial/StructType#m.
// documentation ```go
//   ^^^^^^^^^ reference 0.1.test sg/initial/IfaceType#
   f LiteralType
// ^ definition 0.1.test sg/initial/StructType#f.
// documentation ```go
//   ^^^^^^^^^^^ reference 0.1.test sg/initial/LiteralType#
  
   // anonymous struct
   anon struct {
// ^^^^ definition 0.1.test sg/initial/StructType#anon.
// documentation ```go
    sub int
//  ^^^ definition 0.1.test sg/initial/StructType#anon.sub.
//  documentation ```go
   }
  
   // interface within struct
   i interface {
// ^ definition 0.1.test sg/initial/StructType#i.
// documentation ```go
    AnonMethod() bool
//  ^^^^^^^^^^ definition 0.1.test sg/initial/StructType#i.AnonMethod.
//  documentation ```go
   }
  }
  
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition 0.1.test sg/initial/DeclaredBefore#
//     documentation ```go
//     documentation ```go
//                            ^^^^^^^^^^^^^ definition 0.1.test sg/initial/DeclaredBefore#DeclaredAfter.
//                            documentation ```go
//                            ^^^^^^^^^^^^^ reference 0.1.test sg/initial/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition 0.1.test sg/initial/DeclaredAfter#
//     documentation ```go
//     documentation ```go
  
