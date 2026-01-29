  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
  type LiteralType int
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/LiteralType#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
  
  type FuncType func(LiteralType, int) bool
//     ^^^^^^^^ definition 0.1.test `sg/initial`/FuncType#
//     kind Type
//     documentation
//     > ```go
//     > func(LiteralType, int) bool
//     > ```
//                   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
  type IfaceType interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/IfaceType#
//     kind Interface
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
// kind Method
// documentation
// > ```go
// > func (IfaceType).Method() LiteralType
// > ```
//          ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  }
  
  type StructType struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#
//     kind Class
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
// kind Field
// documentation
// > ```go
// > struct field m sg/initial.IfaceType
// > ```
//   ^^^^^^^^^ reference 0.1.test `sg/initial`/IfaceType#
   f LiteralType
// ^ definition 0.1.test `sg/initial`/StructType#f.
// kind Field
// documentation
// > ```go
// > struct field f sg/initial.LiteralType
// > ```
//   ^^^^^^^^^^^ reference 0.1.test `sg/initial`/LiteralType#
  
   // anonymous struct
   anon struct {
// ^^^^ definition 0.1.test `sg/initial`/StructType#anon.
// kind Field
// documentation
// > ```go
// > struct field anon struct{sub int}
// > ```
    sub int
//  ^^^ definition 0.1.test `sg/initial`/StructType#anon.sub.
//  kind Field
//  documentation
//  > ```go
//  > struct field sub int
//  > ```
   }
  
   // interface within struct
   i interface {
// ^ definition 0.1.test `sg/initial`/StructType#i.
// kind Field
// documentation
// > ```go
// > struct field i interface{AnonMethod() bool}
// > ```
    AnonMethod() bool
//  ^^^^^^^^^^ definition 0.1.test `sg/initial`/StructType#i.AnonMethod.
//  kind Method
//  documentation
//  > ```go
//  > func (interface).AnonMethod() bool
//  > ```
   }
  }
  
  type DeclaredBefore struct{ DeclaredAfter }
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredBefore#
//     kind Class
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
//                            kind Field
//                            documentation
//                            > ```go
//                            > struct field DeclaredAfter sg/initial.DeclaredAfter
//                            > ```
//                            ^^^^^^^^^^^^^ reference 0.1.test `sg/initial`/DeclaredAfter#
  type DeclaredAfter struct{}
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/DeclaredAfter#
//     kind Class
//     documentation
//     > ```go
//     > type DeclaredAfter struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
