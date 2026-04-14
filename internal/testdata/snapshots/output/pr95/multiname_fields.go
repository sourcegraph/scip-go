  package pr95
//        ^^^^ definition 0.1.test `sg/pr95`/
  
  // Phase 1: defer-in-loop bug — multi-name field declarations.
  // The `defer v.scope.pop()` inside the for-loop over field names causes
  // scopes to accumulate, producing incorrect nested symbols like #a.b.x
  // instead of independent #a.x and #b.x.
  
  type MultiNameStruct struct {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/MultiNameStruct#
//     documentation
//     > ```go
//     > type MultiNameStruct struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     a struct {
//     >         x int
//     >         y string
//     >     }
//     >     b struct {
//     >         x int
//     >         y string
//     >     }
//     > }
//     > ```
   a, b struct {
// ^ definition 0.1.test `sg/pr95`/MultiNameStruct#a.
// documentation
// > ```go
// > struct field a struct{x int; y string}
// > ```
//    ^ definition 0.1.test `sg/pr95`/MultiNameStruct#b.
//    documentation
//    > ```go
//    > struct field b struct{x int; y string}
//    > ```
    x int
//  ^ definition 0.1.test `sg/pr95`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
//  documentation
//  > ```go
//  > struct field x int
//  > ```
    y string
//  ^ definition 0.1.test `sg/pr95`/MultiNameStruct#$anon_c0a8952b3a214f68#y.
//  documentation
//  > ```go
//  > struct field y string
//  > ```
   }
  }
  
  type ThreeNameStruct struct {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/ThreeNameStruct#
//     documentation
//     > ```go
//     > type ThreeNameStruct struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     p struct {
//     >         val int
//     >     }
//     >     q struct {
//     >         val int
//     >     }
//     >     r struct {
//     >         val int
//     >     }
//     > }
//     > ```
   p, q, r struct {
// ^ definition 0.1.test `sg/pr95`/ThreeNameStruct#p.
// documentation
// > ```go
// > struct field p struct{val int}
// > ```
//    ^ definition 0.1.test `sg/pr95`/ThreeNameStruct#q.
//    documentation
//    > ```go
//    > struct field q struct{val int}
//    > ```
//       ^ definition 0.1.test `sg/pr95`/ThreeNameStruct#r.
//       documentation
//       > ```go
//       > struct field r struct{val int}
//       > ```
    val int
//  ^^^ definition 0.1.test `sg/pr95`/ThreeNameStruct#$anon_64fbd7cc136081fc#val.
//  documentation
//  > ```go
//  > struct field val int
//  > ```
   }
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/useMultiNameFields().
  func useMultiNameFields() {
//     ^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/useMultiNameFields().
//     documentation
//     > ```go
//     > func useMultiNameFields()
//     > ```
   var m MultiNameStruct
//     ^ definition local 0
//       ^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr95`/MultiNameStruct#
   m.a.x = 1
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/MultiNameStruct#a.
//     ^ reference 0.1.test `sg/pr95`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
   m.a.y = "hello"
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/MultiNameStruct#a.
//     ^ reference 0.1.test `sg/pr95`/MultiNameStruct#$anon_c0a8952b3a214f68#y.
   m.b.x = 2
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/MultiNameStruct#b.
//     ^ reference 0.1.test `sg/pr95`/MultiNameStruct#$anon_c0a8952b3a214f68#x.
   m.b.y = "world"
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/MultiNameStruct#b.
//     ^ reference 0.1.test `sg/pr95`/MultiNameStruct#$anon_c0a8952b3a214f68#y.
  
   m.a = m.b
// ^ reference local 0
//   ^ reference 0.1.test `sg/pr95`/MultiNameStruct#a.
//       ^ reference local 0
//         ^ reference 0.1.test `sg/pr95`/MultiNameStruct#b.
  
   var t ThreeNameStruct
//     ^ definition local 1
//       ^^^^^^^^^^^^^^^ reference 0.1.test `sg/pr95`/ThreeNameStruct#
   t.p.val = 1
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/ThreeNameStruct#p.
//     ^^^ reference 0.1.test `sg/pr95`/ThreeNameStruct#$anon_64fbd7cc136081fc#val.
   t.q.val = 2
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/ThreeNameStruct#q.
//     ^^^ reference 0.1.test `sg/pr95`/ThreeNameStruct#$anon_64fbd7cc136081fc#val.
   t.r.val = 3
// ^ reference local 1
//   ^ reference 0.1.test `sg/pr95`/ThreeNameStruct#r.
//     ^^^ reference 0.1.test `sg/pr95`/ThreeNameStruct#$anon_64fbd7cc136081fc#val.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/useMultiNameFields().
  
