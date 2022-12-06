  package testdata
//        ^^^^^^^^ reference sg/testdata/
  
  type I0 interface{}
//     ^^ definition sg/testdata/I0#
//     documentation ```go
//     documentation ```go
  
  type I1 interface {
//     ^^ definition sg/testdata/I1#
//     documentation ```go
//     documentation ```go
   F1()
// ^^ definition sg/testdata/I1#F1.
// documentation ```go
  }
  
  type I2 interface {
//     ^^ definition sg/testdata/I2#
//     documentation ```go
//     documentation ```go
   F2()
// ^^ definition sg/testdata/I2#F2.
// documentation ```go
  }
  
  type T1 int
//     ^^ definition sg/testdata/T1#
//     documentation ```go
  
  func (r T1) F1() {}
//      ^ definition local 0
//        ^^ reference sg/testdata/T1#
//            ^^ definition sg/testdata/T1#F1().
//            documentation ```go
  
  type T2 int
//     ^^ definition sg/testdata/T2#
//     documentation ```go
  
  func (r T2) F1() {}
//      ^ definition local 1
//        ^^ reference sg/testdata/T2#
//            ^^ definition sg/testdata/T2#F1().
//            documentation ```go
  func (r T2) F2() {}
//      ^ definition local 2
//        ^^ reference sg/testdata/T2#
//            ^^ definition sg/testdata/T2#F2().
//            documentation ```go
  
  type A1 = T1
//     ^^ definition sg/testdata/A1#
//     documentation ```go
//          ^^ reference sg/testdata/T1#
  type A12 = A1
//     ^^^ definition sg/testdata/A12#
//     documentation ```go
//           ^^ reference sg/testdata/A1#
  
  type InterfaceWithNonExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition sg/testdata/InterfaceWithNonExportedMethod#
//     documentation ```go
//     documentation ```go
   nonExportedMethod()
// ^^^^^^^^^^^^^^^^^ definition sg/testdata/InterfaceWithNonExportedMethod#nonExportedMethod.
// documentation ```go
  }
  
  type InterfaceWithExportedMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition sg/testdata/InterfaceWithExportedMethod#
//     documentation ```go
//     documentation ```go
   ExportedMethod()
// ^^^^^^^^^^^^^^ definition sg/testdata/InterfaceWithExportedMethod#ExportedMethod.
// documentation ```go
  }
  
  type Foo int
//     ^^^ definition sg/testdata/Foo#
//     documentation ```go
  
  func (r Foo) nonExportedMethod() {}
//      ^ definition local 3
//        ^^^ reference sg/testdata/Foo#
//             ^^^^^^^^^^^^^^^^^ definition sg/testdata/Foo#nonExportedMethod().
//             documentation ```go
  func (r Foo) ExportedMethod()    {}
//      ^ definition local 4
//        ^^^ reference sg/testdata/Foo#
//             ^^^^^^^^^^^^^^ definition sg/testdata/Foo#ExportedMethod().
//             documentation ```go
  func (r Foo) Close() error       { return nil }
//      ^ definition local 5
//        ^^^ reference sg/testdata/Foo#
//             ^^^^^ definition sg/testdata/Foo#Close().
//             documentation ```go
  
  type SharedOne interface {
//     ^^^^^^^^^ definition sg/testdata/SharedOne#
//     documentation ```go
//     documentation ```go
   Shared()
// ^^^^^^ definition sg/testdata/SharedOne#Shared.
// documentation ```go
   Distinct()
// ^^^^^^^^ definition sg/testdata/SharedOne#Distinct.
// documentation ```go
  }
  
  type SharedTwo interface {
//     ^^^^^^^^^ definition sg/testdata/SharedTwo#
//     documentation ```go
//     documentation ```go
   Shared()
// ^^^^^^ definition sg/testdata/SharedTwo#Shared.
// documentation ```go
   Unique()
// ^^^^^^ definition sg/testdata/SharedTwo#Unique.
// documentation ```go
  }
  
  type Between struct{}
//     ^^^^^^^ definition sg/testdata/Between#
//     documentation ```go
//     documentation ```go
  
  func (Between) Shared()   {}
//      ^^^^^^^ reference sg/testdata/Between#
//               ^^^^^^ definition sg/testdata/Between#Shared().
//               documentation ```go
  func (Between) Distinct() {}
//      ^^^^^^^ reference sg/testdata/Between#
//               ^^^^^^^^ definition sg/testdata/Between#Distinct().
//               documentation ```go
  func (Between) Unique()   {}
//      ^^^^^^^ reference sg/testdata/Between#
//               ^^^^^^ definition sg/testdata/Between#Unique().
//               documentation ```go
  
  func shouldShow(shared SharedOne) {
//     ^^^^^^^^^^ definition sg/testdata/shouldShow().
//     documentation ```go
//                ^^^^^^ definition local 6
//                       ^^^^^^^^^ reference sg/testdata/SharedOne#
   shared.Shared()
// ^^^^^^ reference local 6
//        ^^^^^^ reference sg/testdata/SharedOne#Shared.
  }
  
