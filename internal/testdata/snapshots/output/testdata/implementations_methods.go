  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  type InterfaceWithSingleMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethod#
//                               signature_documentation
//                               > type InterfaceWithSingleMethod interface {
//                               >     SingleMethod() float64
//                               > }
   SingleMethod() float64
// ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethod#SingleMethod.
//              signature_documentation
//              > func (InterfaceWithSingleMethod).SingleMethod() float64
  }
  
  type StructWithMethods struct{}
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructWithMethods#
//                       signature_documentation
//                       > type StructWithMethods struct{}
//                       relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethod# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
  func (StructWithMethods) SingleMethod() float64 { return 5.0 }
//      ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/StructWithMethods#
//                         ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
//                                      signature_documentation
//                                      > func (StructWithMethods).SingleMethod() float64
//                                      relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethod#SingleMethod. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
  
  type InterfaceWithSingleMethodTwoImplementers interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#
//                                              signature_documentation
//                                              > type InterfaceWithSingleMethodTwoImplementers interface {
//                                              >     SingleMethodTwoImpl() float64
//                                              > }
   SingleMethodTwoImpl() float64
// ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl.
//                     signature_documentation
//                     > func (InterfaceWithSingleMethodTwoImplementers).SingleMethodTwoImpl() float64
  }
  
  type TwoImplOne struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplOne#
//                signature_documentation
//                > type TwoImplOne struct{}
//                relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
  func (TwoImplOne) SingleMethodTwoImpl() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplOne#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
//                                      signature_documentation
//                                      > func (TwoImplOne).SingleMethodTwoImpl() float64
//                                      relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
  
  type TwoImplTwo struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#
//                signature_documentation
//                > type TwoImplTwo struct{}
//                relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
  func (TwoImplTwo) SingleMethodTwoImpl() float64         { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
//                                      signature_documentation
//                                      > func (TwoImplTwo).SingleMethodTwoImpl() float64
//                                      relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
//                                                                     ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
  func (TwoImplTwo) RandomThingThatDoesntMatter() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
//                                              signature_documentation
//                                              > func (TwoImplTwo).RandomThingThatDoesntMatter() float64
//                                                                     ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
  
