  package pr95
//        ^^^^ definition 0.1.test `sg/pr95`/
  
  // Phase 2: lookup.Set conflict scenarios.
  // When multiple fields share the same anonymous struct type, their nested
  // fields map to the same token.Pos. The current lookup.Set silently
  // overwrites, masking the conflict.
  
  type ConflictingFields struct {
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/ConflictingFields#
//     documentation
//     > ```go
//     > type ConflictingFields struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     first struct {
//     >         shared int
//     >     }
//     >     second struct {
//     >         shared int
//     >     }
//     > }
//     > ```
   first struct {
// ^^^^^ definition 0.1.test `sg/pr95`/ConflictingFields#first.
// documentation
// > ```go
// > struct field first struct{shared int}
// > ```
    shared int
//  ^^^^^^ definition 0.1.test `sg/pr95`/ConflictingFields#first.shared.
//  documentation
//  > ```go
//  > struct field shared int
//  > ```
   }
   second struct {
// ^^^^^^ definition 0.1.test `sg/pr95`/ConflictingFields#second.
// documentation
// > ```go
// > struct field second struct{shared int}
// > ```
    shared int
//  ^^^^^^ definition 0.1.test `sg/pr95`/ConflictingFields#second.shared.
//  documentation
//  > ```go
//  > struct field shared int
//  > ```
   }
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/useConflictingFields().
  func useConflictingFields() {
//     ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/useConflictingFields().
//     documentation
//     > ```go
//     > func useConflictingFields()
//     > ```
   var c ConflictingFields
//     ^ definition local 0
//       ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr95`/ConflictingFields#
   c.first.shared = 1
// ^ reference local 0
//   ^^^^^ reference 0.1.test `sg/pr95`/ConflictingFields#first.
//         ^^^^^^ reference 0.1.test `sg/pr95`/ConflictingFields#first.shared.
   c.second.shared = 2
// ^ reference local 0
//   ^^^^^^ reference 0.1.test `sg/pr95`/ConflictingFields#second.
//          ^^^^^^ reference 0.1.test `sg/pr95`/ConflictingFields#second.shared.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/useConflictingFields().
  
  // Multi-name fields also trigger the same-pos conflict.
  type MultiNameConflict struct {
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/MultiNameConflict#
//     documentation
//     > ```go
//     > type MultiNameConflict struct
//     > ```
//     documentation
//     > Multi-name fields also trigger the same-pos conflict.
//     documentation
//     > ```go
//     > struct {
//     >     x struct {
//     >         field int
//     >     }
//     >     y struct {
//     >         field int
//     >     }
//     > }
//     > ```
   x, y struct {
// ^ definition 0.1.test `sg/pr95`/MultiNameConflict#x.
// documentation
// > ```go
// > struct field x struct{field int}
// > ```
//    ^ definition 0.1.test `sg/pr95`/MultiNameConflict#x.y.
//    documentation
//    > ```go
//    > struct field y struct{field int}
//    > ```
    field int
//  ^^^^^ definition 0.1.test `sg/pr95`/MultiNameConflict#x.y.field.
//  documentation
//  > ```go
//  > struct field field int
//  > ```
   }
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/useMultiNameConflict().
  func useMultiNameConflict() {
//     ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/useMultiNameConflict().
//     documentation
//     > ```go
//     > func useMultiNameConflict()
//     > ```
   var m MultiNameConflict
//     ^ definition local 1
//       ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr95`/MultiNameConflict#
   m.x.field = 10
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/MultiNameConflict#x.
//     ^^^^^ reference 0.1.test `sg/pr95`/MultiNameConflict#x.y.field.
   m.y.field = 20
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/MultiNameConflict#x.y.
//     ^^^^^ reference 0.1.test `sg/pr95`/MultiNameConflict#x.y.field.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/useMultiNameConflict().
  
