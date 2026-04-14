  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
//                     display_name inlinestruct
//                     signature_documentation
//                     > package inlinestruct
  
  type FieldInterface interface {
//     ^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/FieldInterface#
//                    signature_documentation
//                    > type FieldInterface interface
//                    > interface {
//                    >     SomeMethod() string
//                    > }
   SomeMethod() string
// ^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/FieldInterface#SomeMethod.
//            signature_documentation
//            > func (FieldInterface).SomeMethod() string
  }
  
  var MyInline = struct {
//    ^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline.
//             signature_documentation
//             > var MyInline struct{privateField FieldInterface; PublicField FieldInterface}
   privateField FieldInterface
// ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline:privateField.
//              signature_documentation
//              > struct field privateField sg/inlinestruct.FieldInterface
//              ^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/FieldInterface#
   PublicField  FieldInterface
// ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/MyInline:PublicField.
//             signature_documentation
//             > struct field PublicField sg/inlinestruct.FieldInterface
//              ^^^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/FieldInterface#
  }{}
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/MyFunc().
  func MyFunc() {
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/MyFunc().
//            signature_documentation
//            > func MyFunc()
   _ = MyInline.privateField
//     ^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline.
//              ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline:privateField.
   _ = MyInline.PublicField
//     ^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline.
//              ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/MyInline:PublicField.
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/MyFunc().
  
