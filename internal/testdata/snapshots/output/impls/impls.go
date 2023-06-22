  package impls
//        ^^^^^ definition 0.1.test sg/impls/
//        documentation package impls
  
  type I1 interface {
//     ^^ definition 0.1.test sg/impls/I1#
//     documentation ```go
//     documentation ```go
   F1()
// ^^ definition 0.1.test sg/impls/I1#F1.
// documentation ```go
  }
  
  type I1Clone interface {
//     ^^^^^^^ definition 0.1.test sg/impls/I1Clone#
//     documentation ```go
//     documentation ```go
   F1()
// ^^ definition 0.1.test sg/impls/I1Clone#F1.
// documentation ```go
  }
  
  type IfaceOther interface {
//     ^^^^^^^^^^ definition 0.1.test sg/impls/IfaceOther#
//     documentation ```go
//     documentation ```go
   Something()
// ^^^^^^^^^ definition 0.1.test sg/impls/IfaceOther#Something.
// documentation ```go
   Another()
// ^^^^^^^ definition 0.1.test sg/impls/IfaceOther#Another.
// documentation ```go
  }
  
  type T1 int
//     ^^ definition 0.1.test sg/impls/T1#
//     documentation ```go
//     relationship 0.1.test sg/impls/I1# implementation
//     relationship 0.1.test sg/impls/I1Clone# implementation
  
  func (r T1) F1() {}
//      ^ definition local 0
//        ^^ reference 0.1.test sg/impls/T1#
//            ^^ definition 0.1.test sg/impls/T1#F1().
//            documentation ```go
//            relationship 0.1.test sg/impls/I1#F1. implementation
//            relationship 0.1.test sg/impls/I1Clone#F1. implementation
  
  type T2 int
//     ^^ definition 0.1.test sg/impls/T2#
//     documentation ```go
//     relationship 0.1.test sg/impls/I1# implementation
//     relationship 0.1.test sg/impls/I1Clone# implementation
  
  func (r T2) F1() {}
//      ^ definition local 1
//        ^^ reference 0.1.test sg/impls/T2#
//            ^^ definition 0.1.test sg/impls/T2#F1().
//            documentation ```go
//            relationship 0.1.test sg/impls/I1#F1. implementation
//            relationship 0.1.test sg/impls/I1Clone#F1. implementation
  func (r T2) F2() {}
//      ^ definition local 2
//        ^^ reference 0.1.test sg/impls/T2#
//            ^^ definition 0.1.test sg/impls/T2#F2().
//            documentation ```go
  
