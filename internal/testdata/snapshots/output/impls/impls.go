  package impls
  
  type I1 interface {
//     ^^ definition sg/impls/I1#
//     documentation ```go
//     documentation ```go
   F1()
// ^^ definition sg/impls/I1#F1.
// documentation ```go
  }
  
  type I1Clone interface {
//     ^^^^^^^ definition sg/impls/I1Clone#
//     documentation ```go
//     documentation ```go
   F1()
// ^^ definition sg/impls/I1Clone#F1.
// documentation ```go
  }
  
  type IfaceOther interface {
//     ^^^^^^^^^^ definition sg/impls/IfaceOther#
//     documentation ```go
//     documentation ```go
   Something()
// ^^^^^^^^^ definition sg/impls/IfaceOther#Something.
// documentation ```go
   Another()
// ^^^^^^^ definition sg/impls/IfaceOther#Another.
// documentation ```go
  }
  
  type T1 int
//     ^^ definition sg/impls/T1#
//     documentation ```go
//     relationship sg/impls/I1# implementation
//     relationship sg/impls/I1Clone# implementation
  
  func (r T1) F1() {}
//      ^ definition local 0
//        ^^ reference sg/impls/T1#
//            ^^ definition sg/impls/T1#F1().
//            documentation ```go
  
  type T2 int
//     ^^ definition sg/impls/T2#
//     documentation ```go
//     relationship sg/impls/I1# implementation
//     relationship sg/impls/I1Clone# implementation
  
  func (r T2) F1() {}
//      ^ definition local 1
//        ^^ reference sg/impls/T2#
//            ^^ definition sg/impls/T2#F1().
//            documentation ```go
  func (r T2) F2() {}
//      ^ definition local 2
//        ^^ reference sg/impls/T2#
//            ^^ definition sg/impls/T2#F2().
//            documentation ```go
  
