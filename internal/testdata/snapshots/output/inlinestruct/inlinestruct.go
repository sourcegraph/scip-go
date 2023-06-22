  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/
//        documentation package inlinestruct
  
  type FieldInterface interface {
//     ^^^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/FieldInterface#
//     documentation ```go
//     documentation ```go
   SomeMethod() string
// ^^^^^^^^^^ definition 0.1.test sg/inlinestruct/FieldInterface#SomeMethod.
// documentation ```go
  }
  
  var MyInline = struct {
//    ^^^^^^^^ definition 0.1.test sg/inlinestruct/MyInline.
//    documentation ```go
   privateField FieldInterface
// ^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/MyInline:privateField.
// documentation ```go
//              ^^^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/FieldInterface#
   PublicField  FieldInterface
// ^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/MyInline:PublicField.
// documentation ```go
//              ^^^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/FieldInterface#
  }{}
  
  func MyFunc() {
//     ^^^^^^ definition 0.1.test sg/inlinestruct/MyFunc().
//     documentation ```go
   _ = MyInline.privateField
//     ^^^^^^^^ reference 0.1.test sg/inlinestruct/MyInline.
//              ^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/MyInline:privateField.
   _ = MyInline.PublicField
//     ^^^^^^^^ reference 0.1.test sg/inlinestruct/MyInline.
//              ^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/MyInline:PublicField.
  }
  
