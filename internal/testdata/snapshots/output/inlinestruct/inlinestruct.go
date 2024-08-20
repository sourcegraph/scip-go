  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
//        documentation
//        > package inlinestruct
  
  type FieldInterface interface {
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/FieldInterface#
//     documentation
//     > ```go
//     > type FieldInterface interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     SomeMethod() string
//     > }
//     > ```
   SomeMethod() string
// ^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/FieldInterface#SomeMethod.
// documentation
// > ```go
// > func (FieldInterface).SomeMethod() string
// > ```
  }
  
  var MyInline = struct {
//    ^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline.
//    documentation
//    > ```go
//    > var MyInline struct{privateField FieldInterface; PublicField FieldInterface}
//    > ```
   privateField FieldInterface
// ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline:privateField.
// documentation
// > ```go
// > struct field privateField sg/inlinestruct.FieldInterface
// > ```
//              ^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/FieldInterface#
   PublicField  FieldInterface
// ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline:PublicField.
// documentation
// > ```go
// > struct field PublicField sg/inlinestruct.FieldInterface
// > ```
//              ^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/FieldInterface#
  }{}
  
  func MyFunc() {
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/MyFunc().
//     documentation
//     > ```go
//     > func MyFunc()
//     > ```
   _ = MyInline.privateField
//     ^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline.
//              ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline:privateField.
   _ = MyInline.PublicField
//     ^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline.
//              ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline:PublicField.
  }
  
