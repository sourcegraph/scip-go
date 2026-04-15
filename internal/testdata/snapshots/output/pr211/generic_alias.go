  package pr211
//        ^^^^^ definition 0.1.test `sg/pr211`/
//              display_name pr211
//              signature_documentation
//              > package pr211
  
  // Map is a generic map type.
  type Map[K comparable, V any] struct {
//     ^^^ definition 0.1.test `sg/pr211`/Map#
//         display_name Map
//         signature_documentation
//         > type Map struct{ entries []entry[K, V] }
//         documentation
//         > Map is a generic map type.
//         ^ definition local 0
//           display_name K
//           signature_documentation
//           > type parameter K comparable
//                       ^ definition local 1
//                         display_name V
//                         signature_documentation
//                         > type parameter V any
   entries []entry[K, V]
// ^^^^^^^ definition 0.1.test `sg/pr211`/Map#entries.
//         display_name entries
//         signature_documentation
//         > struct field entries []entry[K, V]
//           ^^^^^ reference 0.1.test `sg/pr211`/entry#
//                 ^ reference local 0
//                    ^ reference local 1
  }
  
  type entry[K comparable, V any] struct {
//     ^^^^^ definition 0.1.test `sg/pr211`/entry#
//           display_name entry
//           signature_documentation
//           > type entry struct {
//           >     key   K
//           >     value V
//           > }
//           ^ definition local 2
//             display_name K
//             signature_documentation
//             > type parameter K comparable
//                         ^ definition local 3
//                           display_name V
//                           signature_documentation
//                           > type parameter V any
   key   K
// ^^^ definition 0.1.test `sg/pr211`/entry#key.
//     display_name key
//     signature_documentation
//     > struct field key K
//       ^ reference local 2
   value V
// ^^^^^ definition 0.1.test `sg/pr211`/entry#value.
//       display_name value
//       signature_documentation
//       > struct field value V
//       ^ reference local 3
  }
  
  // Set is a generic alias that partially instantiates Map.
  type Set[K comparable] = Map[K, bool]
//     ^^^ definition 0.1.test `sg/pr211`/Set#
//         display_name Set
//         signature_documentation
//         > type Set[K comparable] = Map[K, bool]
//         documentation
//         > Set is a generic alias that partially instantiates Map.
//         ^ definition local 4
//           display_name K
//           signature_documentation
//           > type parameter K comparable
//                         ^^^ reference 0.1.test `sg/pr211`/Map#
//                             ^ reference local 4
  
  // Alias with a tighter constraint.
  type OrderedSet[K ~int | ~string] = Set[K]
//     ^^^^^^^^^^ definition 0.1.test `sg/pr211`/OrderedSet#
//                display_name OrderedSet
//                signature_documentation
//                > type OrderedSet[K ~int | ~string] = Set[K]
//                documentation
//                > Alias with a tighter constraint.
//                ^ definition local 5
//                  display_name K
//                  signature_documentation
//                  > type parameter K ~int | ~string
//                                    ^^^ reference 0.1.test `sg/pr211`/Set#
//                                        ^ reference local 5
  
  // Alias of an alias (chained).
  type StringSet = Set[string]
//     ^^^^^^^^^ definition 0.1.test `sg/pr211`/StringSet#
//               display_name StringSet
//               signature_documentation
//               > type StringSet = Set[string]
//               documentation
//               > Alias of an alias (chained).
//                 ^^^ reference 0.1.test `sg/pr211`/Set#
  
  // Alias with all params forwarded.
  type PairMap[K comparable, V any] = Map[K, V]
//     ^^^^^^^ definition 0.1.test `sg/pr211`/PairMap#
//             display_name PairMap
//             signature_documentation
//             > type PairMap[K comparable, V any] = Map[K, V]
//             documentation
//             > Alias with all params forwarded.
//             ^ definition local 6
//               display_name K
//               signature_documentation
//               > type parameter K comparable
//                           ^ definition local 7
//                             display_name V
//                             signature_documentation
//                             > type parameter V any
//                                    ^^^ reference 0.1.test `sg/pr211`/Map#
//                                        ^ reference local 6
//                                           ^ reference local 7
  
//⌄ enclosing_range_start 0.1.test `sg/pr211`/UseAliases().
  func UseAliases() {
//     ^^^^^^^^^^ definition 0.1.test `sg/pr211`/UseAliases().
//                display_name UseAliases
//                signature_documentation
//                > func UseAliases()
   _ = Set[int]{}
//     ^^^ reference 0.1.test `sg/pr211`/Set#
   _ = OrderedSet[int]{}
//     ^^^^^^^^^^ reference 0.1.test `sg/pr211`/OrderedSet#
   _ = StringSet{}
//     ^^^^^^^^^ reference 0.1.test `sg/pr211`/StringSet#
   _ = PairMap[string, int]{}
//     ^^^^^^^ reference 0.1.test `sg/pr211`/PairMap#
  }
//⌃ enclosing_range_end 0.1.test `sg/pr211`/UseAliases().
  
