  package impls
//        ^^^^^ definition 0.1.test `sg/impls`/
//        documentation
//        > package impls
  
  type I1 interface {
//     ^^ definition 0.1.test `sg/impls`/I1#
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
// ^^ definition 0.1.test `sg/impls`/I1#F1.
// kind Method
// documentation
// > ```go
// > func (I1).F1()
// > ```
  }
  
  type I1Clone interface {
//     ^^^^^^^ definition 0.1.test `sg/impls`/I1Clone#
//     kind Interface
//     documentation
//     > ```go
//     > type I1Clone interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     F1()
//     > }
//     > ```
   F1()
// ^^ definition 0.1.test `sg/impls`/I1Clone#F1.
// kind Method
// documentation
// > ```go
// > func (I1Clone).F1()
// > ```
  }
  
  type IfaceOther interface {
//     ^^^^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#
//     kind Interface
//     documentation
//     > ```go
//     > type IfaceOther interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Another()
//     >     Something()
//     > }
//     > ```
   Something()
// ^^^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#Something.
// kind Method
// documentation
// > ```go
// > func (IfaceOther).Something()
// > ```
   Another()
// ^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#Another.
// kind Method
// documentation
// > ```go
// > func (IfaceOther).Another()
// > ```
  }
  
  type T1 int
//     ^^ definition 0.1.test `sg/impls`/T1#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `sg/impls`/I1# implementation
//     relationship 0.1.test `sg/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//      kind Variable
//        ^^ reference 0.1.test `sg/impls`/T1#
//            ^^ definition 0.1.test `sg/impls`/T1#F1().
//            kind Method
//            documentation
//            > ```go
//            > func (T1).F1()
//            > ```
//            relationship 0.1.test `sg/impls`/I1#F1. implementation
//            relationship 0.1.test `sg/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T1#F1().
  
  type T2 int
//     ^^ definition 0.1.test `sg/impls`/T2#
//     kind Type
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `sg/impls`/I1# implementation
//     relationship 0.1.test `sg/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//      kind Variable
//        ^^ reference 0.1.test `sg/impls`/T2#
//            ^^ definition 0.1.test `sg/impls`/T2#F1().
//            kind Method
//            documentation
//            > ```go
//            > func (T2).F1()
//            > ```
//            relationship 0.1.test `sg/impls`/I1#F1. implementation
//            relationship 0.1.test `sg/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T2#F1().
//⌄ enclosing_range_start 0.1.test `sg/impls`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//      kind Variable
//        ^^ reference 0.1.test `sg/impls`/T2#
//            ^^ definition 0.1.test `sg/impls`/T2#F2().
//            kind Method
//            documentation
//            > ```go
//            > func (T2).F2()
//            > ```
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T2#F2().
  
