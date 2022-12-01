  package initial
//        ^^^^^^^ reference sg/initial/
  
  type LiteralType int
//     ^^^^^^^^^^^ definition sg/initial/LiteralType#
//     documentation ```go
  
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition sg/initial/FuncType#
//     documentation ```go
//                   ^^^^^^^^^^^ reference sg/initial/LiteralType#
  
  type IfaceType interface {
//     ^^^^^^^^^ definition sg/initial/IfaceType#
//     documentation ```go
//     documentation ```go
   Method() LiteralType
// ^^^^^^ definition sg/initial/IfaceType#Method.
// documentation ```go
//          ^^^^^^^^^^^ reference sg/initial/LiteralType#
  }
  
  type StructType struct {
//     ^^^^^^^^^^ definition sg/initial/StructType#
//     documentation ```go
//     documentation ```go
   m IfaceType
// ^ definition sg/initial/StructType#m.
// documentation ```go
//   ^^^^^^^^^ reference sg/initial/IfaceType#
   f LiteralType
// ^ definition sg/initial/StructType#f.
// documentation ```go
//   ^^^^^^^^^^^ reference sg/initial/LiteralType#
  
   // anonymous struct
   anon struct {
// ^^^^ definition sg/initial/StructType#anon.
// documentation ```go
    sub int
//  ^^^ definition sg/initial/StructType#anon.sub.
//  documentation ```go
   }
  
   // interface within struct
   i interface {
// ^ definition sg/initial/StructType#i.
// documentation ```go
    AnonMethod() bool
//  ^^^^^^^^^^ definition sg/initial/StructType#i.AnonMethod.
//  documentation ```go
   }
  }
  
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition sg/initial/DeclaredBefore#
//     documentation ```go
//     documentation ```go
//                            ^^^^^^^^^^^^^ definition sg/initial/DeclaredBefore#DeclaredAfter.
//                            documentation ```go
//                            ^^^^^^^^^^^^^ reference sg/initial/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition sg/initial/DeclaredAfter#
//     documentation ```go
//     documentation ```go
  
