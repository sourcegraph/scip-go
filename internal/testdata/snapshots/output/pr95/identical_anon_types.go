  package pr95
//        ^^^^ definition 0.1.test `sg/pr95`/
  
  // Phase 3: Identical anonymous struct types across different fields.
  // Fields with the same anonymous struct type should share symbol names
  // for their nested fields, enabling cross-field Find References.
  
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
//           ^ definition 0.1.test `sg/pr95`/IdenticalAnonFields#x.t.
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
//           ^ definition 0.1.test `sg/pr95`/IdenticalAnonFields#z.t.
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
//     ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#x.t.
   y.z.t = 2
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#z.
//     ^ reference 0.1.test `sg/pr95`/IdenticalAnonFields#z.t.
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
   a struct {
// ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#a.
// documentation
// > ```go
// > struct field a struct{x int; y string}
// > ```
    x int
//  ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#a.x.
//  documentation
//  > ```go
//  > struct field x int
//  > ```
    y string
//  ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#a.y.
//  documentation
//  > ```go
//  > struct field y string
//  > ```
   }
   b struct {
// ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#b.
// documentation
// > ```go
// > struct field b struct{y string; x int}
// > ```
    y string
//  ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#b.y.
//  documentation
//  > ```go
//  > struct field y string
//  > ```
    x int
//  ^ definition 0.1.test `sg/pr95`/FieldOrderMatters#b.x.
//  documentation
//  > ```go
//  > struct field x int
//  > ```
   }
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/useFieldOrderMatters().
  func useFieldOrderMatters() {
//     ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/useFieldOrderMatters().
//     documentation
//     > ```go
//     > func useFieldOrderMatters()
//     > ```
   var f FieldOrderMatters
//     ^ definition local 1
//       ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr95`/FieldOrderMatters#
   f.a.x = 1
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#a.
//     ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#a.x.
   f.a.y = "hello"
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#a.
//     ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#a.y.
   f.b.y = "world"
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#b.
//     ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#b.y.
   f.b.x = 2
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#b.
//     ^ reference 0.1.test `sg/pr95`/FieldOrderMatters#b.x.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/useFieldOrderMatters().
  
  // Different struct tags mean different types — symbols must NOT unify.
  type DifferentTags struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/DifferentTags#
//     documentation
//     > ```go
//     > type DifferentTags struct
//     > ```
//     documentation
//     > Different struct tags mean different types — symbols must NOT unify.
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
   a struct {
// ^ definition 0.1.test `sg/pr95`/DifferentTags#a.
// documentation
// > ```go
// > struct field a struct{Name string "json:\"name\""}
// > ```
    Name string `json:"name"`
//  ^^^^ definition 0.1.test `sg/pr95`/DifferentTags#a.Name.
//  documentation
//  > ```go
//  > struct field Name string
//  > ```
   }
   b struct {
// ^ definition 0.1.test `sg/pr95`/DifferentTags#b.
// documentation
// > ```go
// > struct field b struct{Name string "json:\"full_name\""}
// > ```
    Name string `json:"full_name"`
//  ^^^^ definition 0.1.test `sg/pr95`/DifferentTags#b.Name.
//  documentation
//  > ```go
//  > struct field Name string
//  > ```
   }
  }
  
  // Exported vs unexported field names — different types.
  type ExportedVsUnexported struct {
//     ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/ExportedVsUnexported#
//     documentation
//     > ```go
//     > type ExportedVsUnexported struct
//     > ```
//     documentation
//     > Exported vs unexported field names — different types.
//     documentation
//     > ```go
//     > struct {
//     >     a struct {
//     >         X int
//     >         Y string
//     >     }
//     >     b struct {
//     >         x int
//     >         y string
//     >     }
//     > }
//     > ```
   a struct {
// ^ definition 0.1.test `sg/pr95`/ExportedVsUnexported#a.
// documentation
// > ```go
// > struct field a struct{X int; Y string}
// > ```
    X int
//  ^ definition 0.1.test `sg/pr95`/ExportedVsUnexported#a.X.
//  documentation
//  > ```go
//  > struct field X int
//  > ```
    Y string
//  ^ definition 0.1.test `sg/pr95`/ExportedVsUnexported#a.Y.
//  documentation
//  > ```go
//  > struct field Y string
//  > ```
   }
   b struct {
// ^ definition 0.1.test `sg/pr95`/ExportedVsUnexported#b.
// documentation
// > ```go
// > struct field b struct{x int; y string}
// > ```
    x int
//  ^ definition 0.1.test `sg/pr95`/ExportedVsUnexported#b.x.
//  documentation
//  > ```go
//  > struct field x int
//  > ```
    y string
//  ^ definition 0.1.test `sg/pr95`/ExportedVsUnexported#b.y.
//  documentation
//  > ```go
//  > struct field y string
//  > ```
   }
  }
  
