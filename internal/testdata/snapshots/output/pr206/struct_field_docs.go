  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  // MyStruct is a struct with documented fields.
  type MyStruct struct {
//     ^^^^^^^^ definition 0.1.test `sg/pr206`/MyStruct#
//              signature_documentation
//              > type MyStruct struct {
//              >     Name  string
//              >     Score float64
//              > }
//              documentation
//              > MyStruct is a struct with documented fields.
   // Name is the user's display name.
   Name string
// ^^^^ definition 0.1.test `sg/pr206`/MyStruct#Name.
//      signature_documentation
//      > struct field Name string
  
   Score float64 // inline comment on Score
// ^^^^^ definition 0.1.test `sg/pr206`/MyStruct#Score.
//       signature_documentation
//       > struct field Score float64
  }
  
