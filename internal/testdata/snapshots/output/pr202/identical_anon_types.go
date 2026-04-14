  package pr202
//        ^^^^^ definition 0.1.test `sg/pr202`/
//              display_name pr202
//              signature_documentation
//              > package pr202
  
  // Identical anonymous struct types should share symbols for nested fields.
  
  type IdenticalAnonFields struct {
//     ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#
//                         documentation
//                         > ```go
//                         > type IdenticalAnonFields struct
//                         > ```
//                         documentation
//                         > ```go
//                         > struct {
//                         >     x struct {
//                         >         t int
//                         >     }
//                         >     z struct {
//                         >         t int
//                         >     }
//                         > }
//                         > ```
   x struct{ t int }
// ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#x.
//   documentation
//   > ```go
//   > struct field x struct{t int}
//   > ```
//           ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
//             documentation
//             > ```go
//             > struct field t int
//             > ```
   z struct{ t int }
// ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#z.
//   documentation
//   > ```go
//   > struct field z struct{t int}
//   > ```
//           ^ definition 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
//             documentation
//             > ```go
//             > struct field t int
//             > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/useIdenticalAnonFields().
  func useIdenticalAnonFields() {
//     ^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/useIdenticalAnonFields().
//                            documentation
//                            > ```go
//                            > func useIdenticalAnonFields()
//                            > ```
   var y IdenticalAnonFields
//     ^ definition local 0
//       display_name y
//       signature_documentation
//       > var y sg/pr202.IdenticalAnonFields
//       ^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#
   y.x = y.z
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#x.
//       ^ reference local 0
//         ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#z.
   y.x.t = 1
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#x.
//     ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
   y.z.t = 2
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#z.
//     ^ reference 0.1.test `sg/pr202`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/useIdenticalAnonFields().
  
  // Different field order means different type — symbols must NOT unify.
  type FieldOrderMatters struct {
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/FieldOrderMatters#
//                       documentation
//                       > ```go
//                       > type FieldOrderMatters struct
//                       > ```
//                       documentation
//                       > Different field order means different type — symbols must NOT unify.
//                       documentation
//                       > ```go
//                       > struct {
//                       >     a struct {
//                       >         x int
//                       >         y string
//                       >     }
//                       >     b struct {
//                       >         y string
//                       >         x int
//                       >     }
//                       > }
//                       > ```
   a struct {
// ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#a.
//   documentation
//   > ```go
//   > struct field a struct{x int; y string}
//   > ```
    x int
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_c0a8952b3a214f68#x.
//    documentation
//    > ```go
//    > struct field x int
//    > ```
    y string
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_c0a8952b3a214f68#y.
//    documentation
//    > ```go
//    > struct field y string
//    > ```
   }
   b struct {
// ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#b.
//   documentation
//   > ```go
//   > struct field b struct{y string; x int}
//   > ```
    y string
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_b8d88f3211c0d7a4#y.
//    documentation
//    > ```go
//    > struct field y string
//    > ```
    x int
//  ^ definition 0.1.test `sg/pr202`/FieldOrderMatters#$anon_b8d88f3211c0d7a4#x.
//    documentation
//    > ```go
//    > struct field x int
//    > ```
   }
  }
  
  // Different struct tags — symbols must NOT unify.
  type DifferentTags struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/DifferentTags#
//                   documentation
//                   > ```go
//                   > type DifferentTags struct
//                   > ```
//                   documentation
//                   > Different struct tags — symbols must NOT unify.
//                   documentation
//                   > ```go
//                   > struct {
//                   >     a struct {
//                   >         Name string `json:"name"`
//                   >     }
//                   >     b struct {
//                   >         Name string `json:"full_name"`
//                   >     }
//                   > }
//                   > ```
   a struct {
// ^ definition 0.1.test `sg/pr202`/DifferentTags#a.
//   documentation
//   > ```go
//   > struct field a struct{Name string `json:"name"`}
//   > ```
    Name string `json:"name"`
//  ^^^^ definition 0.1.test `sg/pr202`/DifferentTags#$anon_ed545d904f2246eb#Name.
//       documentation
//       > ```go
//       > struct field Name string
//       > ```
   }
   b struct {
// ^ definition 0.1.test `sg/pr202`/DifferentTags#b.
//   documentation
//   > ```go
//   > struct field b struct{Name string `json:"full_name"`}
//   > ```
    Name string `json:"full_name"`
//  ^^^^ definition 0.1.test `sg/pr202`/DifferentTags#$anon_29f1ad2683b11ed0#Name.
//       documentation
//       > ```go
//       > struct field Name string
//       > ```
   }
  }
  
