  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  type InterfaceWithSingleMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithSingleMethod#
//     documentation ```go
//     documentation ```go
   SingleMethod() float64
// ^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithSingleMethod#SingleMethod.
// documentation ```go
  }
  
  type StructWithMethods struct{}
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/StructWithMethods#
//     documentation ```go
//     documentation ```go
//     relationship 0.1.test sg/testdata/InterfaceWithSingleMethod# implementation
  
  func (StructWithMethods) SingleMethod() float64 { return 5.0 }
//      ^^^^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/StructWithMethods#
//                         ^^^^^^^^^^^^ definition 0.1.test sg/testdata/StructWithMethods#SingleMethod().
//                         documentation ```go
//                         relationship 0.1.test sg/testdata/InterfaceWithSingleMethod#SingleMethod. implementation
  
  type InterfaceWithSingleMethodTwoImplementers interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithSingleMethodTwoImplementers#
//     documentation ```go
//     documentation ```go
   SingleMethodTwoImpl() float64
// ^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl.
// documentation ```go
  }
  
  type TwoImplOne struct{}
//     ^^^^^^^^^^ definition 0.1.test sg/testdata/TwoImplOne#
//     documentation ```go
//     documentation ```go
//     relationship 0.1.test sg/testdata/InterfaceWithSingleMethodTwoImplementers# implementation
  
  func (TwoImplOne) SingleMethodTwoImpl() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test sg/testdata/TwoImplOne#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/TwoImplOne#SingleMethodTwoImpl().
//                  documentation ```go
//                  relationship 0.1.test sg/testdata/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
  
  type TwoImplTwo struct{}
//     ^^^^^^^^^^ definition 0.1.test sg/testdata/TwoImplTwo#
//     documentation ```go
//     documentation ```go
//     relationship 0.1.test sg/testdata/InterfaceWithSingleMethodTwoImplementers# implementation
  
  func (TwoImplTwo) SingleMethodTwoImpl() float64         { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test sg/testdata/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/TwoImplTwo#SingleMethodTwoImpl().
//                  documentation ```go
//                  relationship 0.1.test sg/testdata/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
  func (TwoImplTwo) RandomThingThatDoesntMatter() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test sg/testdata/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/TwoImplTwo#RandomThingThatDoesntMatter().
//                  documentation ```go
  
