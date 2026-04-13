  package pr95
//        ^^^^ definition 0.1.test `sg/pr95`/
  
  // Phase 3: Nested anonymous structs and container types.
  
  type DeepNested struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/pr95`/DeepNested#
//     documentation
//     > ```go
//     > type DeepNested struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     outer struct {
//     >         inner struct {
//     >             value int
//     >         }
//     >     }
//     > }
//     > ```
   outer struct {
// ^^^^^ definition 0.1.test `sg/pr95`/DeepNested#outer.
// documentation
// > ```go
// > struct field outer struct{inner struct{value int}}
// > ```
    inner struct {
//  ^^^^^ definition 0.1.test `sg/pr95`/DeepNested#$anon_30bdf47ff4254f9b#inner.
//  documentation
//  > ```go
//  > struct field inner struct{value int}
//  > ```
     value int
//   ^^^^^ definition 0.1.test `sg/pr95`/DeepNested#$anon_30bdf47ff4254f9b#$anon_25fd857a3a1203a7#value.
//   documentation
//   > ```go
//   > struct field value int
//   > ```
    }
   }
  }
  
  type SliceAnon struct {
//     ^^^^^^^^^ definition 0.1.test `sg/pr95`/SliceAnon#
//     documentation
//     > ```go
//     > type SliceAnon struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     items []struct {
//     >         id int
//     >         name string
//     >     }
//     > }
//     > ```
   items []struct {
// ^^^^^ definition 0.1.test `sg/pr95`/SliceAnon#items.
// documentation
// > ```go
// > struct field items []struct{id int; name string}
// > ```
    id   int
//  ^^ definition 0.1.test `sg/pr95`/SliceAnon#$anon_7105f7cab05b114b#id.
//  documentation
//  > ```go
//  > struct field id int
//  > ```
    name string
//  ^^^^ definition 0.1.test `sg/pr95`/SliceAnon#$anon_7105f7cab05b114b#name.
//  documentation
//  > ```go
//  > struct field name string
//  > ```
   }
  }
  
  type MapAnon struct {
//     ^^^^^^^ definition 0.1.test `sg/pr95`/MapAnon#
//     documentation
//     > ```go
//     > type MapAnon struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     entries map[string]struct {
//     >         count int
//     >         label string
//     >     }
//     > }
//     > ```
   entries map[string]struct {
// ^^^^^^^ definition 0.1.test `sg/pr95`/MapAnon#entries.
// documentation
// > ```go
// > struct field entries map[string]struct{count int; label string}
// > ```
    count int
//  ^^^^^ definition 0.1.test `sg/pr95`/MapAnon#$anon_f7ce9edfffe6705a#count.
//  documentation
//  > ```go
//  > struct field count int
//  > ```
    label string
//  ^^^^^ definition 0.1.test `sg/pr95`/MapAnon#$anon_f7ce9edfffe6705a#label.
//  documentation
//  > ```go
//  > struct field label string
//  > ```
   }
  }
  
  type PointerAnon struct {
//     ^^^^^^^^^^^ definition 0.1.test `sg/pr95`/PointerAnon#
//     documentation
//     > ```go
//     > type PointerAnon struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     ptr *struct {
//     >         data int
//     >     }
//     > }
//     > ```
   ptr *struct {
// ^^^ definition 0.1.test `sg/pr95`/PointerAnon#ptr.
// documentation
// > ```go
// > struct field ptr *struct{data int}
// > ```
    data int
//  ^^^^ definition 0.1.test `sg/pr95`/PointerAnon#$anon_9ac92b9fde9a0cad#data.
//  documentation
//  > ```go
//  > struct field data int
//  > ```
   }
  }
  
  // Two fields with identical slice-of-anonymous-struct type.
  type SliceAnonShared struct {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/SliceAnonShared#
//     documentation
//     > ```go
//     > type SliceAnonShared struct
//     > ```
//     documentation
//     > Two fields with identical slice-of-anonymous-struct type.
//     documentation
//     > ```go
//     > struct {
//     >     a []struct {
//     >         v int
//     >     }
//     >     b []struct {
//     >         v int
//     >     }
//     > }
//     > ```
   a []struct{ v int }
// ^ definition 0.1.test `sg/pr95`/SliceAnonShared#a.
// documentation
// > ```go
// > struct field a []struct{v int}
// > ```
//             ^ definition 0.1.test `sg/pr95`/SliceAnonShared#$anon_62737c874b048720#v.
//             documentation
//             > ```go
//             > struct field v int
//             > ```
   b []struct{ v int }
// ^ definition 0.1.test `sg/pr95`/SliceAnonShared#b.
// documentation
// > ```go
// > struct field b []struct{v int}
// > ```
//             ^ definition 0.1.test `sg/pr95`/SliceAnonShared#$anon_62737c874b048720#v.
//             documentation
//             > ```go
//             > struct field v int
//             > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/useNestedAnon().
  func useNestedAnon() {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/useNestedAnon().
//     documentation
//     > ```go
//     > func useNestedAnon()
//     > ```
   var d DeepNested
//     ^ definition local 0
//       ^^^^^^^^^^ reference 0.1.test `sg/pr95`/DeepNested#
   d.outer.inner.value = 42
// ^ reference local 0
//   ^^^^^ reference 0.1.test `sg/pr95`/DeepNested#outer.
//         ^^^^^ reference 0.1.test `sg/pr95`/DeepNested#$anon_30bdf47ff4254f9b#inner.
//               ^^^^^ reference 0.1.test `sg/pr95`/DeepNested#$anon_30bdf47ff4254f9b#$anon_25fd857a3a1203a7#value.
  
   var s SliceAnon
//     ^ definition local 1
//       ^^^^^^^^^ reference 0.1.test `sg/pr95`/SliceAnon#
   if len(s.items) > 0 {
//        ^ reference local 1
//          ^^^^^ reference 0.1.test `sg/pr95`/SliceAnon#items.
    _ = s.items[0].id
//      ^ reference local 1
//        ^^^^^ reference 0.1.test `sg/pr95`/SliceAnon#items.
//                 ^^ reference 0.1.test `sg/pr95`/SliceAnon#$anon_7105f7cab05b114b#id.
    _ = s.items[0].name
//      ^ reference local 1
//        ^^^^^ reference 0.1.test `sg/pr95`/SliceAnon#items.
//                 ^^^^ reference 0.1.test `sg/pr95`/SliceAnon#$anon_7105f7cab05b114b#name.
   }
  
   var m MapAnon
//     ^ definition local 2
//       ^^^^^^^ reference 0.1.test `sg/pr95`/MapAnon#
   entry := m.entries["key"]
// ^^^^^ definition local 3
//          ^ reference local 2
//            ^^^^^^^ reference 0.1.test `sg/pr95`/MapAnon#entries.
   _ = entry.count
//     ^^^^^ reference local 3
//           ^^^^^ reference 0.1.test `sg/pr95`/MapAnon#$anon_f7ce9edfffe6705a#count.
   _ = entry.label
//     ^^^^^ reference local 3
//           ^^^^^ reference 0.1.test `sg/pr95`/MapAnon#$anon_f7ce9edfffe6705a#label.
  
   var p PointerAnon
//     ^ definition local 4
//       ^^^^^^^^^^^ reference 0.1.test `sg/pr95`/PointerAnon#
   if p.ptr != nil {
//    ^ reference local 4
//      ^^^ reference 0.1.test `sg/pr95`/PointerAnon#ptr.
    _ = p.ptr.data
//      ^ reference local 4
//        ^^^ reference 0.1.test `sg/pr95`/PointerAnon#ptr.
//            ^^^^ reference 0.1.test `sg/pr95`/PointerAnon#$anon_9ac92b9fde9a0cad#data.
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/useNestedAnon().
  
