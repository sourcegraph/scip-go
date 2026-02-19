  package testdata
//        ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/
  
  type I0 interface{}
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I0#
//     documentation
//     > ```go
//     > type I0 interface
//     > ```
//     documentation
//     > ```go
//     > interface{}
//     > ```
  
  type I1 interface {
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I1#
//     documentation
//     > ```go
//     > type I1 interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     F1()
//     > }
//     > ```
   F1()
// ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I1#F1.
// documentation
// > ```go
// > func (I1).F1()
// > ```
  }
  
  type I2 interface {
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I2#
//     documentation
//     > ```go
//     > type I2 interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     F2()
//     > }
//     > ```
   F2()
// ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I2#F2.
// documentation
// > ```go
// > func (I2).F2()
// > ```
  }
  
  type T1 int
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T1#
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I1# implementation
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//        ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T1#
//            ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T1#F1().
//            documentation
//            > ```go
//            > func (T1).F1()
//            > ```
//            relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T1#F1().
  
  type T2 int
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I1# implementation
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I2# implementation
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//        ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#
//            ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#F1().
//            documentation
//            > ```go
//            > func (T2).F1()
//            > ```
//            relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#F1().
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//        ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#
//            ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#F2().
//            documentation
//            > ```go
//            > func (T2).F2()
//            > ```
//            relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I2#F2. implementation
//                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T2#F2().
  
  type A1 = T1
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/A1#
//     documentation
//     > ```go
//     > int
//     > ```
//          ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/T1#
  type A12 = A1
//     ^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/A12#
//     documentation
//     > ```go
//     > int
//     > ```
//           ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/A1#
  
  type InterfaceWithNonExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithNonExportedMethod#
//     documentation
//     > ```go
//     > type InterfaceWithNonExportedMethod interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     nonExportedMethod()
//     > }
//     > ```
   nonExportedMethod()
// ^^^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod.
// documentation
// > ```go
// > func (InterfaceWithNonExportedMethod).nonExportedMethod()
// > ```
  }
  
  type InterfaceWithExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithExportedMethod#
//     documentation
//     > ```go
//     > type InterfaceWithExportedMethod interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     ExportedMethod()
//     > }
//     > ```
   ExportedMethod()
// ^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithExportedMethod#ExportedMethod.
// documentation
// > ```go
// > func (InterfaceWithExportedMethod).ExportedMethod()
// > ```
  }
  
  type Foo int
//     ^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship github.com/golang/go/src go1.22 io/Closer# implementation
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I3# implementation
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithExportedMethod# implementation
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithNonExportedMethod# implementation
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#nonExportedMethod().
  func (r Foo) nonExportedMethod() {}
//      ^ definition local 3
//        ^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#
//             ^^^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#nonExportedMethod().
//             documentation
//             > ```go
//             > func (Foo).nonExportedMethod()
//             > ```
//             relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#nonExportedMethod().
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#ExportedMethod().
  func (r Foo) ExportedMethod()    {}
//      ^ definition local 4
//        ^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#
//             ^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#ExportedMethod().
//             documentation
//             > ```go
//             > func (Foo).ExportedMethod()
//             > ```
//             relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/InterfaceWithExportedMethod#ExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#ExportedMethod().
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#Close().
  func (r Foo) Close() error       { return nil }
//      ^ definition local 5
//        ^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#
//             ^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#Close().
//             documentation
//             > ```go
//             > func (Foo).Close() error
//             > ```
//             relationship github.com/golang/go/src go1.22 io/Closer#Close. implementation
//             relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/I3#Close. implementation
//                                              ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Foo#Close().
  
  type SharedOne interface {
//     ^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne#
//     documentation
//     > ```go
//     > type SharedOne interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Distinct()
//     >     Shared()
//     > }
//     > ```
   Shared()
// ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne#Shared.
// documentation
// > ```go
// > func (SharedOne).Shared()
// > ```
   Distinct()
// ^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne#Distinct.
// documentation
// > ```go
// > func (SharedOne).Distinct()
// > ```
  }
  
  type SharedTwo interface {
//     ^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedTwo#
//     documentation
//     > ```go
//     > type SharedTwo interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Shared()
//     >     Unique()
//     > }
//     > ```
   Shared()
// ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedTwo#Shared.
// documentation
// > ```go
// > func (SharedTwo).Shared()
// > ```
   Unique()
// ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedTwo#Unique.
// documentation
// > ```go
// > func (SharedTwo).Unique()
// > ```
  }
  
  type Between struct{}
//     ^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#
//     documentation
//     > ```go
//     > type Between struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne# implementation
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedTwo# implementation
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Shared().
  func (Between) Shared()   {}
//      ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#
//               ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Shared().
//               documentation
//               > ```go
//               > func (Between).Shared()
//               > ```
//               relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne#Shared. implementation
//               relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedTwo#Shared. implementation
//                           ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Shared().
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Distinct().
  func (Between) Distinct() {}
//      ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#
//               ^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Distinct().
//               documentation
//               > ```go
//               > func (Between).Distinct()
//               > ```
//               relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne#Distinct. implementation
//                           ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Distinct().
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Unique().
  func (Between) Unique()   {}
//      ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#
//               ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Unique().
//               documentation
//               > ```go
//               > func (Between).Unique()
//               > ```
//               relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedTwo#Unique. implementation
//                           ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Between#Unique().
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/shouldShow().
  func shouldShow(shared SharedOne) {
//     ^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/shouldShow().
//     documentation
//     > ```go
//     > func shouldShow(shared SharedOne)
//     > ```
//                ^^^^^^ definition local 6
//                       ^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne#
   shared.Shared()
// ^^^^^^ reference local 6
//        ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SharedOne#Shared.
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/shouldShow().
  
