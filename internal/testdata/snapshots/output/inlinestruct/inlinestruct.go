  package inlinestruct
//        ^^^^^^^^^^^^ definition sg/inlinestruct/
  
  type FieldInterface interface {
//     ^^^^^^^^^^^^^^ definition sg/inlinestruct/FieldInterface#
//     documentation ```go
//     documentation ```go
   SomeMethod() string
// ^^^^^^^^^^ definition sg/inlinestruct/FieldInterface#SomeMethod.
// documentation ```go
  }
  
  var MyInline = struct {
//    ^^^^^^^^ definition MyInline.
//    documentation ```go
   privateField FieldInterface
// ^^^^^^^^^^^^ definition sg/inlinestruct/MyInline:privateField.
// documentation ```go
//              ^^^^^^^^^^^^^^ reference sg/inlinestruct/FieldInterface#
   PublicField  FieldInterface
// ^^^^^^^^^^^ definition sg/inlinestruct/MyInline:PublicField.
// documentation ```go
//              ^^^^^^^^^^^^^^ reference sg/inlinestruct/FieldInterface#
  }{}
  
  func MyFunc() {
//     ^^^^^^ definition sg/inlinestruct/MyFunc().
//     documentation ```go
   _ = MyInline.privateField
//     ^^^^^^^^ reference MyInline.
//              ^^^^^^^^^^^^ reference sg/inlinestruct/MyInline:privateField.
   _ = MyInline.PublicField
//     ^^^^^^^^ reference MyInline.
//              ^^^^^^^^^^^ reference sg/inlinestruct/MyInline:PublicField.
  }
  
