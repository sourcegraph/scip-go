  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  type I0 interface{}
//     ^^ definition 0.1.test `sg/testdata`/I0#
//     kind Interface
//     documentation
//     > ```go
//     > type I0 interface
//     > ```
//     documentation
//     > ```go
//     > interface{}
//     > ```
  
  type I1 interface {
//     ^^ definition 0.1.test `sg/testdata`/I1#
//     kind Interface
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
// ^^ definition 0.1.test `sg/testdata`/I1#F1.
// kind Method
// documentation
// > ```go
// > func (I1).F1()
// > ```
  }
  
  type I2 interface {
//     ^^ definition 0.1.test `sg/testdata`/I2#
//     kind Interface
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
// ^^ definition 0.1.test `sg/testdata`/I2#F2.
// kind Method
// documentation
// > ```go
// > func (I2).F2()
// > ```
  }
  
  type T1 int
//     ^^ definition 0.1.test `sg/testdata`/T1#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `sg/testdata`/I1# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//      kind Variable
//        ^^ reference 0.1.test `sg/testdata`/T1#
//            ^^ definition 0.1.test `sg/testdata`/T1#F1().
//            kind Method
//            documentation
//            > ```go
//            > func (T1).F1()
//            > ```
//            relationship 0.1.test `sg/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T1#F1().
  
  type T2 int
//     ^^ definition 0.1.test `sg/testdata`/T2#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `sg/testdata`/I1# implementation
//     relationship 0.1.test `sg/testdata`/I2# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//      kind Variable
//        ^^ reference 0.1.test `sg/testdata`/T2#
//            ^^ definition 0.1.test `sg/testdata`/T2#F1().
//            kind Method
//            documentation
//            > ```go
//            > func (T2).F1()
//            > ```
//            relationship 0.1.test `sg/testdata`/I1#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T2#F1().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//      kind Variable
//        ^^ reference 0.1.test `sg/testdata`/T2#
//            ^^ definition 0.1.test `sg/testdata`/T2#F2().
//            kind Method
//            documentation
//            > ```go
//            > func (T2).F2()
//            > ```
//            relationship 0.1.test `sg/testdata`/I2#F2. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/T2#F2().
  
  type A1 = T1
//     ^^ definition 0.1.test `sg/testdata`/A1#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
//          ^^ reference 0.1.test `sg/testdata`/T1#
  type A12 = A1
//     ^^^ definition 0.1.test `sg/testdata`/A12#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
//           ^^ reference 0.1.test `sg/testdata`/A1#
  
  type InterfaceWithNonExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#
//     kind Interface
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
// ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod.
// kind Method
// documentation
// > ```go
// > func (InterfaceWithNonExportedMethod).nonExportedMethod()
// > ```
  }
  
  type InterfaceWithExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithExportedMethod#
//     kind Interface
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
// ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod.
// kind Method
// documentation
// > ```go
// > func (InterfaceWithExportedMethod).ExportedMethod()
// > ```
  }
  
  type Foo int
//     ^^^ definition 0.1.test `sg/testdata`/Foo#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship github.com/golang/go/src go1.22 io/Closer# implementation
//     relationship 0.1.test `sg/testdata`/I3# implementation
//     relationship 0.1.test `sg/testdata`/InterfaceWithExportedMethod# implementation
//     relationship 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#nonExportedMethod().
  func (r Foo) nonExportedMethod() {}
//      ^ definition local 3
//      kind Variable
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/Foo#nonExportedMethod().
//             kind Method
//             documentation
//             > ```go
//             > func (Foo).nonExportedMethod()
//             > ```
//             relationship 0.1.test `sg/testdata`/InterfaceWithNonExportedMethod#nonExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#nonExportedMethod().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#ExportedMethod().
  func (r Foo) ExportedMethod()    {}
//      ^ definition local 4
//      kind Variable
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/Foo#ExportedMethod().
//             kind Method
//             documentation
//             > ```go
//             > func (Foo).ExportedMethod()
//             > ```
//             relationship 0.1.test `sg/testdata`/InterfaceWithExportedMethod#ExportedMethod. implementation
//                                  ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#ExportedMethod().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Foo#Close().
  func (r Foo) Close() error       { return nil }
//      ^ definition local 5
//      kind Variable
//        ^^^ reference 0.1.test `sg/testdata`/Foo#
//             ^^^^^ definition 0.1.test `sg/testdata`/Foo#Close().
//             kind Method
//             documentation
//             > ```go
//             > func (Foo).Close() error
//             > ```
//             relationship github.com/golang/go/src go1.22 io/Closer#Close. implementation
//             relationship 0.1.test `sg/testdata`/I3#Close. implementation
//                                              ⌃ enclosing_range_end 0.1.test `sg/testdata`/Foo#Close().
  
  type SharedOne interface {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#
//     kind Interface
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
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#Shared.
// kind Method
// documentation
// > ```go
// > func (SharedOne).Shared()
// > ```
   Distinct()
// ^^^^^^^^ definition 0.1.test `sg/testdata`/SharedOne#Distinct.
// kind Method
// documentation
// > ```go
// > func (SharedOne).Distinct()
// > ```
  }
  
  type SharedTwo interface {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#
//     kind Interface
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
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#Shared.
// kind Method
// documentation
// > ```go
// > func (SharedTwo).Shared()
// > ```
   Unique()
// ^^^^^^ definition 0.1.test `sg/testdata`/SharedTwo#Unique.
// kind Method
// documentation
// > ```go
// > func (SharedTwo).Unique()
// > ```
  }
  
  type Between struct{}
//     ^^^^^^^ definition 0.1.test `sg/testdata`/Between#
//     kind Class
//     documentation
//     > ```go
//     > type Between struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
//     relationship 0.1.test `sg/testdata`/SharedOne# implementation
//     relationship 0.1.test `sg/testdata`/SharedTwo# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Shared().
  func (Between) Shared()   {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^ definition 0.1.test `sg/testdata`/Between#Shared().
//               kind Method
//               documentation
//               > ```go
//               > func (Between).Shared()
//               > ```
//               relationship 0.1.test `sg/testdata`/SharedOne#Shared. implementation
//               relationship 0.1.test `sg/testdata`/SharedTwo#Shared. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Shared().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Distinct().
  func (Between) Distinct() {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^^^ definition 0.1.test `sg/testdata`/Between#Distinct().
//               kind Method
//               documentation
//               > ```go
//               > func (Between).Distinct()
//               > ```
//               relationship 0.1.test `sg/testdata`/SharedOne#Distinct. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Distinct().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Between#Unique().
  func (Between) Unique()   {}
//      ^^^^^^^ reference 0.1.test `sg/testdata`/Between#
//               ^^^^^^ definition 0.1.test `sg/testdata`/Between#Unique().
//               kind Method
//               documentation
//               > ```go
//               > func (Between).Unique()
//               > ```
//               relationship 0.1.test `sg/testdata`/SharedTwo#Unique. implementation
//                           ⌃ enclosing_range_end 0.1.test `sg/testdata`/Between#Unique().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/shouldShow().
  func shouldShow(shared SharedOne) {
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/shouldShow().
//     kind Function
//     documentation
//     > ```go
//     > func shouldShow(shared SharedOne)
//     > ```
//                ^^^^^^ definition local 6
//                kind Variable
//                       ^^^^^^^^^ reference 0.1.test `sg/testdata`/SharedOne#
   shared.Shared()
// ^^^^^^ reference local 6
//        ^^^^^^ reference 0.1.test `sg/testdata`/SharedOne#Shared.
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/shouldShow().
  
