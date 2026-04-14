  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  type I0 interface{}
//     ^^ definition 0.1.test `sg/testdata`/I0#
//        signature_documentation
//        > type I0 interface{}
  
  type I1 interface {
//     ^^ definition 0.1.test `sg/testdata`/I1#
//        signature_documentation
//        > type I1 interface {
//        >     F1()
//        > }
   F1()
// ^^ definition 0.1.test `sg/testdata`/I1#F1.
//    signature_documentation
//    > func (I1).F1()
  }
  
  type I2 interface {
//     ^^ definition 0.1.test `sg/testdata`/I2#
//        signature_documentation
//        > type I2 interface {
//        >     F2()
//        > }
   F2()
// ^^ definition 0.1.test `sg/testdata`/I2#F2.
//    signature_documentation
//    > func (I2).F2()
  }
  
  type T1 int
//     ^^ definition 0.1.test `sg/testdata`/T1#
//        signature_documentation
//        > type T1 int
//        relationship 0.1.test `sg/testdata`/I1# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//        display_name r
//        signature_documentation
//        > var r sg/testdata.T1
//        ^^ reference 0.1.test `sg/testdata`/T1#
//            ^^ definition 0.1.test `sg/testdata`/T1#F1().
//               signature_documentation
//               > func (T1).F1()
//               relationship 0.1.test `sg/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T1#F1().
  
  type T2 int
//     ^^ definition 0.1.test `sg/testdata`/T2#
//        signature_documentation
//        > type T2 int
//        relationship 0.1.test `sg/testdata`/I1# implementation
//        relationship 0.1.test `sg/testdata`/I2# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//        display_name r
//        signature_documentation
//        > var r sg/testdata.T2
//        ^^ reference 0.1.test `sg/testdata`/T2#
//            ^^ definition 0.1.test `sg/testdata`/T2#F1().
//               signature_documentation
//               > func (T2).F1()
//               relationship 0.1.test `sg/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T2#F1().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//        display_name r
//        signature_documentation
//        > var r sg/testdata.T2
//        ^^ reference 0.1.test `sg/testdata`/T2#
//            ^^ definition 0.1.test `sg/testdata`/T2#F2().
//               signature_documentation
//               > func (T2).F2()
//               relationship 0.1.test `sg/testdata`/I2#F2. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T2#F2().
  
  type A1 = T1
//     ^^ definition 0.1.test `sg/testdata`/A1#
//        signature_documentation
//        > type A1 = T1
//          ^^ reference 0.1.test `sg/testdata`/T1#
  type A12 = A1
//     ^^^ definition 0.1.test `sg/testdata`/A12#
//         signature_documentation
//         > type A12 = A1
//           ^^ reference 0.1.test `sg/testdata`/A1#
  
  type InterfaceWithNonExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#
//                                    signature_documentation
//                                    > type InterfaceWithNonExportedMethod interface {
//                                    >     nonExportedMethod()
//                                    > }
   nonExportedMethod()
// ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod.
//                   signature_documentation
//                   > func (InterfaceWithNonExportedMethod).nonExportedMethod()
  }
  
  type InterfaceWithExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithExportedMethod#
//                                 signature_documentation
//                                 > type InterfaceWithExportedMethod interface {
//                                 >     ExportedMethod()
//                                 > }
   ExportedMethod()
// ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod.
//                signature_documentation
//                > func (InterfaceWithExportedMethod).ExportedMethod()
  }
  
  type Foo int
//     ^^^ definition 0.1.test `sg/testdata`/Foo#
//         signature_documentation
//         > type Foo int
//         relationship github.com/golang/go/src go1.22 io/Closer# implementation
//         relationship 0.1.test `sg/testdata`/I3# implementation
//         relationship 0.1.test `sg/testdata`/InterfaceWithExportedMethod# implementation
//         relationship 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#nonExportedMethod().
  func (r Foo) nonExportedMethod() {}
//      ^ definition local 3
//        display_name r
//        signature_documentation
//        > var r sg/testdata.Foo
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/Foo#nonExportedMethod().
//                               signature_documentation
//                               > func (Foo).nonExportedMethod()
//                               relationship 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#nonExportedMethod().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#ExportedMethod().
  func (r Foo) ExportedMethod()    {}
//      ^ definition local 4
//        display_name r
//        signature_documentation
//        > var r sg/testdata.Foo
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/Foo#ExportedMethod().
//                            signature_documentation
//                            > func (Foo).ExportedMethod()
//                            relationship 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#ExportedMethod().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#Close().
  func (r Foo) Close() error       { return nil }
//      ^ definition local 5
//        display_name r
//        signature_documentation
//        > var r sg/testdata.Foo
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^ definition 0.1.test `sg/testdata`/Foo#Close().
//                   signature_documentation
//                   > func (Foo).Close() error
//                   relationship github.com/golang/go/src go1.22 io/Closer#Close. implementation
//                   relationship 0.1.test `sg/testdata`/I3#Close. implementation
//                                              ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#Close().
  
  type SharedOne interface {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#
//               signature_documentation
//               > type SharedOne interface {
//               >     Distinct()
//               >     Shared()
//               > }
   Shared()
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#Shared.
//        signature_documentation
//        > func (SharedOne).Shared()
   Distinct()
// ^^^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#Distinct.
//          signature_documentation
//          > func (SharedOne).Distinct()
  }
  
  type SharedTwo interface {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#
//               signature_documentation
//               > type SharedTwo interface {
//               >     Shared()
//               >     Unique()
//               > }
   Shared()
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#Shared.
//        signature_documentation
//        > func (SharedTwo).Shared()
   Unique()
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#Unique.
//        signature_documentation
//        > func (SharedTwo).Unique()
  }
  
  type Between struct{}
//     ^^^^^^^ definition 0.1.test `sg/testdata`/Between#
//             signature_documentation
//             > type Between struct{}
//             relationship 0.1.test `sg/testdata`/SharedOne# implementation
//             relationship 0.1.test `sg/testdata`/SharedTwo# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Shared().
  func (Between) Shared()   {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^ definition 0.1.test `sg/testdata`/Between#Shared().
//                      signature_documentation
//                      > func (Between).Shared()
//                      relationship 0.1.test `sg/testdata`/SharedOne#Shared. implementation
//                      relationship 0.1.test `sg/testdata`/SharedTwo#Shared. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Shared().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Distinct().
  func (Between) Distinct() {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^^^ definition 0.1.test `sg/testdata`/Between#Distinct().
//                        signature_documentation
//                        > func (Between).Distinct()
//                        relationship 0.1.test `sg/testdata`/SharedOne#Distinct. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Distinct().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Unique().
  func (Between) Unique()   {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^ definition 0.1.test `sg/testdata`/Between#Unique().
//                      signature_documentation
//                      > func (Between).Unique()
//                      relationship 0.1.test `sg/testdata`/SharedTwo#Unique. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Unique().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/shouldShow().
  func shouldShow(shared SharedOne) {
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/shouldShow().
//                signature_documentation
//                > func shouldShow(shared SharedOne)
//                ^^^^^^ definition local 6
//                       display_name shared
//                       signature_documentation
//                       > var shared sg/testdata.SharedOne
//                       ^^^^^^^^^ reference 0.1.test `sg/testdata`/SharedOne#
   shared.Shared()
// ^^^^^^ reference local 6
//        ^^^^^^ reference 0.1.test `sg/testdata`/SharedOne#Shared.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/shouldShow().
  
