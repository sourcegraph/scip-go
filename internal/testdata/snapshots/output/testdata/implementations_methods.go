  package testdata
//        ^^^^^^^^ reference sg/testdata/
  
  type InterfaceWithSingleMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^ definition sg/testdata/InterfaceWithSingleMethod#
//     documentation ```go
//     documentation ```go
   SingleMethod() float64
// ^^^^^^^^^^^^ definition sg/testdata/InterfaceWithSingleMethod#SingleMethod.
// documentation ```go
  }
  
  type StructWithMethods struct{}
//     ^^^^^^^^^^^^^^^^^ definition sg/testdata/StructWithMethods#
//     documentation ```go
//     documentation ```go
//     relationship sg/testdata/InterfaceWithSingleMethod# implementation
  
  func (StructWithMethods) SingleMethod() float64 { return 5.0 }
//      ^^^^^^^^^^^^^^^^^ reference sg/testdata/StructWithMethods#
//                         ^^^^^^^^^^^^ definition sg/testdata/StructWithMethods#SingleMethod().
//                         documentation ```go
//                         relationship sg/testdata/InterfaceWithSingleMethod#SingleMethod. implementation
  
  type InterfaceWithSingleMethodTwoImplementers interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition sg/testdata/InterfaceWithSingleMethodTwoImplementers#
//     documentation ```go
//     documentation ```go
   SingleMethodTwoImpl() float64
// ^^^^^^^^^^^^^^^^^^^ definition sg/testdata/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl.
// documentation ```go
  }
  
  type TwoImplOne struct{}
//     ^^^^^^^^^^ definition sg/testdata/TwoImplOne#
//     documentation ```go
//     documentation ```go
//     relationship sg/testdata/InterfaceWithSingleMethodTwoImplementers# implementation
  
  func (TwoImplOne) SingleMethodTwoImpl() float64 { return 5.0 }
//      ^^^^^^^^^^ reference sg/testdata/TwoImplOne#
//                  ^^^^^^^^^^^^^^^^^^^ definition sg/testdata/TwoImplOne#SingleMethodTwoImpl().
//                  documentation ```go
//                  relationship sg/testdata/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
  
  type TwoImplTwo struct{}
//     ^^^^^^^^^^ definition sg/testdata/TwoImplTwo#
//     documentation ```go
//     documentation ```go
//     relationship sg/testdata/InterfaceWithSingleMethodTwoImplementers# implementation
  
  func (TwoImplTwo) SingleMethodTwoImpl() float64         { return 5.0 }
//      ^^^^^^^^^^ reference sg/testdata/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^ definition sg/testdata/TwoImplTwo#SingleMethodTwoImpl().
//                  documentation ```go
//                  relationship sg/testdata/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
  func (TwoImplTwo) RandomThingThatDoesntMatter() float64 { return 5.0 }
//      ^^^^^^^^^^ reference sg/testdata/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition sg/testdata/TwoImplTwo#RandomThingThatDoesntMatter().
//                  documentation ```go
  
