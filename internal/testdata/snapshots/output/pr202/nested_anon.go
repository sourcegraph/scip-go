  package pr202
//        ^^^^^ definition 0.1.test `sg/pr202`/
  
  // Anonymous structs inside container types and nested structs.
  
  type ContainerAnon struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#
//     documentation
//     > ```go
//     > type ContainerAnon struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     items []struct {
//     >         id int
//     >     }
//     >     entries map[string]struct {
//     >         count int
//     >     }
//     >     ptr *struct {
//     >         data int
//     >     }
//     > }
//     > ```
   items   []struct{ id int }
// ^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#items.
// documentation
// > ```go
// > struct field items []struct{id int}
// > ```
//                   ^^ definition 0.1.test `sg/pr202`/ContainerAnon#$anon_71c5ea8d9342795c#id.
//                   documentation
//                   > ```go
//                   > struct field id int
//                   > ```
   entries map[string]struct{ count int }
// ^^^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#entries.
// documentation
// > ```go
// > struct field entries map[string]struct{count int}
// > ```
//                            ^^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#$anon_721f9800014370ac#count.
//                            documentation
//                            > ```go
//                            > struct field count int
//                            > ```
   ptr     *struct{ data int }
// ^^^ definition 0.1.test `sg/pr202`/ContainerAnon#ptr.
// documentation
// > ```go
// > struct field ptr *struct{data int}
// > ```
//                  ^^^^ definition 0.1.test `sg/pr202`/ContainerAnon#$anon_944f727740dfb75d#data.
//                  documentation
//                  > ```go
//                  > struct field data int
//                  > ```
  }
  
  type DeepNested struct {
//     ^^^^^^^^^^ definition 0.1.test `sg/pr202`/DeepNested#
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
// ^^^^^ definition 0.1.test `sg/pr202`/DeepNested#outer.
// documentation
// > ```go
// > struct field outer struct{inner struct{value int}}
// > ```
    inner struct {
//  ^^^^^ definition 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#inner.
//  documentation
//  > ```go
//  > struct field inner struct{value int}
//  > ```
     value int
//   ^^^^^ definition 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#$anon_77e42bf2e5c84d1a#value.
//   documentation
//   > ```go
//   > struct field value int
//   > ```
    }
   }
  }
  
  // Two fields with identical slice-of-anonymous-struct type.
  type SliceAnonShared struct {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/SliceAnonShared#
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
// ^ definition 0.1.test `sg/pr202`/SliceAnonShared#a.
// documentation
// > ```go
// > struct field a []struct{v int}
// > ```
//             ^ definition 0.1.test `sg/pr202`/SliceAnonShared#$anon_358bfde4cba1ecae#v.
//             documentation
//             > ```go
//             > struct field v int
//             > ```
   b []struct{ v int }
// ^ definition 0.1.test `sg/pr202`/SliceAnonShared#b.
// documentation
// > ```go
// > struct field b []struct{v int}
// > ```
//             ^ definition 0.1.test `sg/pr202`/SliceAnonShared#$anon_358bfde4cba1ecae#v.
//             documentation
//             > ```go
//             > struct field v int
//             > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/useContainerAnon().
  func useContainerAnon() {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/useContainerAnon().
//     documentation
//     > ```go
//     > func useContainerAnon()
//     > ```
   var c ContainerAnon
//     ^ definition local 0
//       ^^^^^^^^^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#
   if len(c.items) > 0 {
//        ^ reference local 0
//          ^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#items.
    _ = c.items[0].id
//      ^ reference local 0
//        ^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#items.
//                 ^^ reference 0.1.test `sg/pr202`/ContainerAnon#$anon_71c5ea8d9342795c#id.
   }
   entry := c.entries["key"]
// ^^^^^ definition local 1
//          ^ reference local 0
//            ^^^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#entries.
   _ = entry.count
//     ^^^^^ reference local 1
//           ^^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#$anon_721f9800014370ac#count.
   if c.ptr != nil {
//    ^ reference local 0
//      ^^^ reference 0.1.test `sg/pr202`/ContainerAnon#ptr.
    _ = c.ptr.data
//      ^ reference local 0
//        ^^^ reference 0.1.test `sg/pr202`/ContainerAnon#ptr.
//            ^^^^ reference 0.1.test `sg/pr202`/ContainerAnon#$anon_944f727740dfb75d#data.
   }
  
   var d DeepNested
//     ^ definition local 2
//       ^^^^^^^^^^ reference 0.1.test `sg/pr202`/DeepNested#
   d.outer.inner.value = 42
// ^ reference local 2
//   ^^^^^ reference 0.1.test `sg/pr202`/DeepNested#outer.
//         ^^^^^ reference 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#inner.
//               ^^^^^ reference 0.1.test `sg/pr202`/DeepNested#$anon_5ee0364e53e1abd6#$anon_77e42bf2e5c84d1a#value.
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/useContainerAnon().
  
