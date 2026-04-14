  package pr95
//        ^^^^ definition 0.1.test `sg/pr95`/
  
  // Identical anonymous struct types should share symbols for nested fields.
  
  type IdenticalAnonFields struct {
//     ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/IdenticalAnonFields#
//     documentation
//     > ```go
//     > type IdenticalAnonFields struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     x struct {
//     >         t int
//     >     }
//     >     z struct {
//     >         t int
//     >     }
//     > }
//     > ```
   x struct{ t int }
// ^ definition 0.1.test `sg/pr95`/IdenticalAnonFields#x.
// documentation
// > ```go
// > struct field x struct{t int}
// > ```
//           ^ definition 0.1.test `sg/pr95`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
//           documentation
//           > ```go
//           > struct field t int
//           > ```
   z struct{ t int }
// ^ definition 0.1.test `sg/pr95`/IdenticalAnonFields#z.
// documentation
// > ```go
// > struct field z struct{t int}
// > ```
//           ^ definition 0.1.test `sg/pr95`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
//           documentation
//           > ```go
//           > struct field t int
//           > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/useIdenticalAnonFields().
  func useIdenticalAnonFields() {
//     ^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/useIdenticalAnonFields().
//     documentation
//     > ```go
//     > func useIdenticalAnonFields()
//     > ```
   var y IdenticalAnonFields
//     ^ definition local 0
//       ^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#
   y.x = y.z
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#x.
//       ^ reference local 0
//         ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#z.
   y.x.t = 1
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#x.
//     ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
   y.z.t = 2
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#z.
//     ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#$anon_44af0565eb406c15#t.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/useIdenticalAnonFields().
  
  // Different field order means different type — symbols must NOT unify.
  type FieldOrderMatters struct {
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/FieldOrderMatters#
//     documentation
//     > ```go
//     > type FieldOrderMatters struct
//     > ```
//     documentation
//     > Different field order means different type — symbols must NOT unify.
//     documentation
//     > ```go
//     > struct {
//     >     a struct {
//     >         x int
//     >         y string
//     >     }
//     >     b struct {
//     >         y string
//     >         x int
//     >     }
//     > }
//     > ```
   a struct{ x int; y string }
// ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#a.
// documentation
// > ```go
// > struct field a struct{x int; y string}
// > ```
//           ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#$anon_c0a8952b3a214f68#x.
//           documentation
//           > ```go
//           > struct field x int
//           > ```
//                  ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#$anon_c0a8952b3a214f68#y.
//                  documentation
//                  > ```go
//                  > struct field y string
//                  > ```
   b struct{ y string; x int }
// ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#b.
// documentation
// > ```go
// > struct field b struct{y string; x int}
// > ```
//           ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#$anon_b8d88f3211c0d7a4#y.
//           documentation
//           > ```go
//           > struct field y string
//           > ```
//                     ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#$anon_b8d88f3211c0d7a4#x.
//                     documentation
//                     > ```go
//                     > struct field x int
//                     > ```
  }
  
  // Different struct tags — symbols must NOT unify.
  type DifferentTags struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/DifferentTags#
//     documentation
//     > ```go
//     > type DifferentTags struct
//     > ```
//     documentation
//     > Different struct tags — symbols must NOT unify.
//     documentation
//     > ```go
//     > struct {
//     >     a struct {
//     >         Name string "json:\"name\""
//     >     }
//     >     b struct {
//     >         Name string "json:\"full_name\""
//     >     }
//     > }
//     > ```
   a struct{ Name string `json:"name"` }
// ^ definition 0.1.test `sg/pr95`/DifferentTags#a.
// documentation
// > ```go
// > struct field a struct{Name string "json:\"name\""}
// > ```
//           ^^^^ definition 0.1.test `sg/pr95`/DifferentTags#$anon_ed545d904f2246eb#Name.
//           documentation
//           > ```go
//           > struct field Name string
//           > ```
   b struct{ Name string `json:"full_name"` }
// ^ definition 0.1.test `sg/pr95`/DifferentTags#b.
// documentation
// > ```go
// > struct field b struct{Name string "json:\"full_name\""}
// > ```
//           ^^^^ definition 0.1.test `sg/pr95`/DifferentTags#$anon_29f1ad2683b11ed0#Name.
//           documentation
//           > ```go
//           > struct field Name string
//           > ```
  }
  
