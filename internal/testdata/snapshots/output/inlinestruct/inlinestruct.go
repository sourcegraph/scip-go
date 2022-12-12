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
// ^^^^^^^^^^^^ definition local 0
//              ^^^^^^^^^^^^^^ reference sg/inlinestruct/FieldInterface#
   PublicField  FieldInterface
// ^^^^^^^^^^^ definition local 1
//              ^^^^^^^^^^^^^^ reference sg/inlinestruct/FieldInterface#
  }{}
  
  func MyFunc() {
//     ^^^^^^ definition sg/inlinestruct/MyFunc().
//     documentation ```go
   _ = MyInline.privateField
//     ^^^^^^^^ reference MyInline.
//              ^^^^^^^^^^^^ reference local 0
   _ = MyInline.PublicField
//     ^^^^^^^^ reference MyInline.
//              ^^^^^^^^^^^ reference local 1
  }
  
