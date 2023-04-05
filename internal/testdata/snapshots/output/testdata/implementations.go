  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  type I0 interface{}
//     ^^ definition 0.1.test sg/testdata/I0#
//     documentation ```go
//     documentation ```go
  
  type I1 interface {
//     ^^ definition 0.1.test sg/testdata/I1#
//     documentation ```go
//     documentation ```go
   F1()
// ^^ definition 0.1.test sg/testdata/I1#F1.
// documentation ```go
  }
  
  type I2 interface {
//     ^^ definition 0.1.test sg/testdata/I2#
//     documentation ```go
//     documentation ```go
   F2()
// ^^ definition 0.1.test sg/testdata/I2#F2.
// documentation ```go
  }
  
  type T1 int
//     ^^ definition 0.1.test sg/testdata/T1#
//     documentation ```go
//     relationship 0.1.test sg/testdata/I1# implementation
  
  func (r T1) F1() {}
//      ^ definition local 0
//        ^^ reference 0.1.test sg/testdata/T1#
//            ^^ definition 0.1.test sg/testdata/T1#F1().
//            documentation ```go
//            relationship 0.1.test sg/testdata/I1#F1. implementation
//            relationship 0.1.test sg/testdata/I1#F1. implementation
//            relationship 0.1.test sg/testdata/I1#F1. implementation
  
  type T2 int
//     ^^ definition 0.1.test sg/testdata/T2#
//     documentation ```go
//     relationship 0.1.test sg/testdata/I1# implementation
//     relationship 0.1.test sg/testdata/I2# implementation
  
  func (r T2) F1() {}
//      ^ definition local 1
//        ^^ reference 0.1.test sg/testdata/T2#
//            ^^ definition 0.1.test sg/testdata/T2#F1().
//            documentation ```go
//            relationship 0.1.test sg/testdata/I1#F1. implementation
  func (r T2) F2() {}
//      ^ definition local 2
//        ^^ reference 0.1.test sg/testdata/T2#
//            ^^ definition 0.1.test sg/testdata/T2#F2().
//            documentation ```go
//            relationship 0.1.test sg/testdata/I2#F2. implementation
  
  type A1 = T1
//     ^^ definition 0.1.test sg/testdata/A1#
//     documentation ```go
//     relationship 0.1.test sg/testdata/I1# implementation
//          ^^ reference 0.1.test sg/testdata/T1#
  type A12 = A1
//     ^^^ definition 0.1.test sg/testdata/A12#
//     documentation ```go
//     relationship 0.1.test sg/testdata/I1# implementation
//           ^^ reference 0.1.test sg/testdata/A1#
  
  type InterfaceWithNonExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithNonExportedMethod#
//     documentation ```go
//     documentation ```go
   nonExportedMethod()
// ^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithNonExportedMethod#nonExportedMethod.
// documentation ```go
  }
  
  type InterfaceWithExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithExportedMethod#
//     documentation ```go
//     documentation ```go
   ExportedMethod()
// ^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithExportedMethod#ExportedMethod.
// documentation ```go
  }
  
  type Foo int
//     ^^^ definition 0.1.test sg/testdata/Foo#
//     documentation ```go
//     relationship github.com/golang/go/src go1.19 io/Closer# implementation
//     relationship 0.1.test sg/testdata/I3# implementation
//     relationship 0.1.test sg/testdata/InterfaceWithExportedMethod# implementation
//     relationship 0.1.test sg/testdata/InterfaceWithNonExportedMethod# implementation
  
  func (r Foo) nonExportedMethod() {}
//      ^ definition local 3
//        ^^^ reference 0.1.test sg/testdata/Foo#
//             ^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/Foo#nonExportedMethod().
//             documentation ```go
//             relationship 0.1.test sg/testdata/InterfaceWithNonExportedMethod#nonExportedMethod. implementation
  func (r Foo) ExportedMethod()    {}
//      ^ definition local 4
//        ^^^ reference 0.1.test sg/testdata/Foo#
//             ^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/Foo#ExportedMethod().
//             documentation ```go
//             relationship 0.1.test sg/testdata/InterfaceWithExportedMethod#ExportedMethod. implementation
  func (r Foo) Close() error       { return nil }
//      ^ definition local 5
//        ^^^ reference 0.1.test sg/testdata/Foo#
//             ^^^^^ definition 0.1.test sg/testdata/Foo#Close().
//             documentation ```go
//             relationship github.com/golang/go/src go1.19 io/Closer#Close. implementation
//             relationship 0.1.test sg/testdata/I3#Close. implementation
  
  type SharedOne interface {
//     ^^^^^^^^^ definition 0.1.test sg/testdata/SharedOne#
//     documentation ```go
//     documentation ```go
   Shared()
// ^^^^^^ definition 0.1.test sg/testdata/SharedOne#Shared.
// documentation ```go
   Distinct()
// ^^^^^^^^ definition 0.1.test sg/testdata/SharedOne#Distinct.
// documentation ```go
  }
  
  type SharedTwo interface {
//     ^^^^^^^^^ definition 0.1.test sg/testdata/SharedTwo#
//     documentation ```go
//     documentation ```go
   Shared()
// ^^^^^^ definition 0.1.test sg/testdata/SharedTwo#Shared.
// documentation ```go
   Unique()
// ^^^^^^ definition 0.1.test sg/testdata/SharedTwo#Unique.
// documentation ```go
  }
  
  type Between struct{}
//     ^^^^^^^ definition 0.1.test sg/testdata/Between#
//     documentation ```go
//     documentation ```go
//     relationship 0.1.test sg/testdata/SharedOne# implementation
//     relationship 0.1.test sg/testdata/SharedTwo# implementation
  
  func (Between) Shared()   {}
//      ^^^^^^^ reference 0.1.test sg/testdata/Between#
//               ^^^^^^ definition 0.1.test sg/testdata/Between#Shared().
//               documentation ```go
//               relationship 0.1.test sg/testdata/SharedOne#Shared. implementation
//               relationship 0.1.test sg/testdata/SharedTwo#Shared. implementation
  func (Between) Distinct() {}
//      ^^^^^^^ reference 0.1.test sg/testdata/Between#
//               ^^^^^^^^ definition 0.1.test sg/testdata/Between#Distinct().
//               documentation ```go
//               relationship 0.1.test sg/testdata/SharedOne#Distinct. implementation
  func (Between) Unique()   {}
//      ^^^^^^^ reference 0.1.test sg/testdata/Between#
//               ^^^^^^ definition 0.1.test sg/testdata/Between#Unique().
//               documentation ```go
//               relationship 0.1.test sg/testdata/SharedTwo#Unique. implementation
  
  func shouldShow(shared SharedOne) {
//     ^^^^^^^^^^ definition 0.1.test sg/testdata/shouldShow().
//     documentation ```go
//                ^^^^^^ definition local 6
//                       ^^^^^^^^^ reference 0.1.test sg/testdata/SharedOne#
   shared.Shared()
// ^^^^^^ reference local 6
//        ^^^^^^ reference 0.1.test sg/testdata/SharedOne#Shared.
  }
  
