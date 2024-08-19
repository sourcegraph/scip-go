  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
  type LiteralType int
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/LiteralType#
//     documentation
//     > ```go
//     > int
//     > ```
  
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition 0.1.test `sg/initial`/FuncType#
//     documentation
//     > ```go
//     > func(LiteralType, int) bool
//     > ```
//                   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
  type IfaceType interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/IfaceType#
//     documentation
//     > ```go
//     > type IfaceType interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Method() LiteralType
//     > }
//     > ```
   Method() LiteralType
// ^^^^^^ definition 0.1.test `sg/initial`/IfaceType#Method.
// documentation
// > ```go
// > func (IfaceType).Method() LiteralType
// > ```
//          ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  }
  
  type StructType struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#
//     documentation
//     > ```go
//     > type StructType struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     m IfaceType
//     >     f LiteralType
//     >     anon struct {
//     >         sub int
//     >     }
//     >     i interface {
//     >         AnonMethod() bool
//     >     }
//     > }
//     > ```
   m IfaceType
// ^ definition 0.1.test `sg/initial`/StructType#m.
// documentation
// > ```go
// > struct field m sg/initial.IfaceType
// > ```
//   ^^^^^^^^^ reference 0.1.test `sg/initial`/IfaceType#
   f LiteralType
// ^ definition 0.1.test `sg/initial`/StructType#f.
// documentation
// > ```go
// > struct field f sg/initial.LiteralType
// > ```
//   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
   // anonymous struct
   anon struct {
// ^^^^ definition 0.1.test `sg/initial`/StructType#anon.
// documentation
// > ```go
// > struct field anon struct{sub int}
// > ```
    sub int
//  ^^^ definition 0.1.test `sg/initial`/StructType#anon.sub.
//  documentation
//  > ```go
//  > struct field sub int
//  > ```
   }
  
   // interface within struct
   i interface {
// ^ definition 0.1.test `sg/initial`/StructType#i.
// documentation
// > ```go
// > struct field i interface{AnonMethod() bool}
// > ```
    AnonMethod() bool
//  ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#i.AnonMethod.
//  documentation
//  > ```go
//  > func (interface).AnonMethod() bool
//  > ```
   }
  }
  
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#
//     documentation
//     > ```go
//     > type DeclaredBefore struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     DeclaredAfter
//     > }
//     > ```
//                            ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#DeclaredAfter.
//                            documentation
//                            > ```go
//                            > struct field DeclaredAfter sg/initial.DeclaredAfter
//                            > ```
//                            ^^^^^^^^^^^^^ reference 0.1.test `sg/initial`/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredAfter#
//     documentation
//     > ```go
//     > type DeclaredAfter struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
