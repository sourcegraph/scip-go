  package impls
//        ^^^^^ definition 0.1.test `sg/impls`/
//              display_name impls
//              signature_documentation
//              > package impls
  
  type I1 interface {
//     ^^ definition 0.1.test `sg/impls`/I1#
//        signature_documentation
//        > type I1 interface
//        > interface {
//        >     F1()
//        > }
   F1()
// ^^ definition 0.1.test `sg/impls`/I1#F1.
//    signature_documentation
//    > func (I1).F1()
  }
  
  type I1Clone interface {
//     ^^^^^^^ definition 0.1.test `sg/impls`/I1Clone#
//             signature_documentation
//             > type I1Clone interface
//             > interface {
//             >     F1()
//             > }
   F1()
// ^^ definition 0.1.test `sg/impls`/I1Clone#F1.
//    signature_documentation
//    > func (I1Clone).F1()
  }
  
  type IfaceOther interface {
//     ^^^^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#
//                signature_documentation
//                > type IfaceOther interface
//                > interface {
//                >     Another()
//                >     Something()
//                > }
   Something()
// ^^^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#Something.
//           signature_documentation
//           > func (IfaceOther).Something()
   Another()
// ^^^^^^^ definition 0.1.test `sg/impls`/IfaceOther#Another.
//         signature_documentation
//         > func (IfaceOther).Another()
  }
  
  type T1 int
//     ^^ definition 0.1.test `sg/impls`/T1#
//        signature_documentation
//        > type T1 int
//        relationship 0.1.test `sg/impls`/I1# implementation
//        relationship 0.1.test `sg/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//        display_name r
//        signature_documentation
//        > var r sg/impls.T1
//        ^^ reference 0.1.test `sg/impls`/T1#
//            ^^ definition 0.1.test `sg/impls`/T1#F1().
//               signature_documentation
//               > func (T1).F1()
//               relationship 0.1.test `sg/impls`/I1#F1. implementation
//               relationship 0.1.test `sg/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T1#F1().
  
  type T2 int
//     ^^ definition 0.1.test `sg/impls`/T2#
//        signature_documentation
//        > type T2 int
//        relationship 0.1.test `sg/impls`/I1# implementation
//        relationship 0.1.test `sg/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `sg/impls`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//        display_name r
//        signature_documentation
//        > var r sg/impls.T2
//        ^^ reference 0.1.test `sg/impls`/T2#
//            ^^ definition 0.1.test `sg/impls`/T2#F1().
//               signature_documentation
//               > func (T2).F1()
//               relationship 0.1.test `sg/impls`/I1#F1. implementation
//               relationship 0.1.test `sg/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T2#F1().
//⌄ enclosing_range_start 0.1.test `sg/impls`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//        display_name r
//        signature_documentation
//        > var r sg/impls.T2
//        ^^ reference 0.1.test `sg/impls`/T2#
//            ^^ definition 0.1.test `sg/impls`/T2#F2().
//               signature_documentation
//               > func (T2).F2()
//                  ⌃ enclosing_range_end 0.1.test `sg/impls`/T2#F2().
  
