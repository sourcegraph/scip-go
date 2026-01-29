  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  type InterfaceWithSingleMethod interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethod#
//     kind Interface
//     documentation
//     > ```go
//     > type InterfaceWithSingleMethod interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     SingleMethod() float64
//     > }
//     > ```
   SingleMethod() float64
// ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethod#SingleMethod.
// kind Method
// documentation
// > ```go
// > func (InterfaceWithSingleMethod).SingleMethod() float64
// > ```
  }
  
  type StructWithMethods struct{}
//     ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructWithMethods#
//     kind Class
//     documentation
//     > ```go
//     > type StructWithMethods struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
//     relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethod# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
  func (StructWithMethods) SingleMethod() float64 { return 5.0 }
//      ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/StructWithMethods#
//                         ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
//                         kind Method
//                         documentation
//                         > ```go
//                         > func (StructWithMethods).SingleMethod() float64
//                         > ```
//                         relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethod#SingleMethod. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/testdata`/StructWithMethods#SingleMethod().
  
  type InterfaceWithSingleMethodTwoImplementers interface {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#
//     kind Interface
//     documentation
//     > ```go
//     > type InterfaceWithSingleMethodTwoImplementers interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     SingleMethodTwoImpl() float64
//     > }
//     > ```
   SingleMethodTwoImpl() float64
// ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl.
// kind Method
// documentation
// > ```go
// > func (InterfaceWithSingleMethodTwoImplementers).SingleMethodTwoImpl() float64
// > ```
  }
  
  type TwoImplOne struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplOne#
//     kind Class
//     documentation
//     > ```go
//     > type TwoImplOne struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
//     relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
  func (TwoImplOne) SingleMethodTwoImpl() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplOne#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
//                  kind Method
//                  documentation
//                  > ```go
//                  > func (TwoImplOne).SingleMethodTwoImpl() float64
//                  > ```
//                  relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplOne#SingleMethodTwoImpl().
  
  type TwoImplTwo struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#
//     kind Class
//     documentation
//     > ```go
//     > type TwoImplTwo struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
//     relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
  func (TwoImplTwo) SingleMethodTwoImpl() float64         { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
//                  kind Method
//                  documentation
//                  > ```go
//                  > func (TwoImplTwo).SingleMethodTwoImpl() float64
//                  > ```
//                  relationship 0.1.test `sg/testdata`/InterfaceWithSingleMethodTwoImplementers#SingleMethodTwoImpl. implementation
//                                                                     ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplTwo#SingleMethodTwoImpl().
//⌄ enclosing_range_start 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
  func (TwoImplTwo) RandomThingThatDoesntMatter() float64 { return 5.0 }
//      ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TwoImplTwo#
//                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
//                  kind Method
//                  documentation
//                  > ```go
//                  > func (TwoImplTwo).RandomThingThatDoesntMatter() float64
//                  > ```
//                                                                     ⌃ enclosing_range_end 0.1.test `sg/testdata`/TwoImplTwo#RandomThingThatDoesntMatter().
  
