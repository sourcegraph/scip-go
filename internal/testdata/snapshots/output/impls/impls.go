  package impls
//        ^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/
//        documentation
//        > package impls
  
  type I1 interface {
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1#
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
// ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1#F1.
// documentation
// > ```go
// > func (I1).F1()
// > ```
  }
  
  type I1Clone interface {
//     ^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1Clone#
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
// ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1Clone#F1.
// documentation
// > ```go
// > func (I1Clone).F1()
// > ```
  }
  
  type IfaceOther interface {
//     ^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/IfaceOther#
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
// ^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/IfaceOther#Something.
// documentation
// > ```go
// > func (IfaceOther).Something()
// > ```
   Another()
// ^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/IfaceOther#Another.
// documentation
// > ```go
// > func (IfaceOther).Another()
// > ```
  }
  
  type T1 int
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T1#
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1# implementation
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T1#F1().
  func (r T1) F1() {}
//      ^ definition local 0
//        ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T1#
//            ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T1#F1().
//            documentation
//            > ```go
//            > func (T1).F1()
//            > ```
//            relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1#F1. implementation
//            relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T1#F1().
  
  type T2 int
//     ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#
//     documentation
//     > ```go
//     > int
//     > ```
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1# implementation
//     relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1Clone# implementation
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#F1().
  func (r T2) F1() {}
//      ^ definition local 1
//        ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#
//            ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#F1().
//            documentation
//            > ```go
//            > func (T2).F1()
//            > ```
//            relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1#F1. implementation
//            relationship 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/I1Clone#F1. implementation
//                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#F1().
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#F2().
  func (r T2) F2() {}
//      ^ definition local 2
//        ^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#
//            ^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#F2().
//            documentation
//            > ```go
//            > func (T2).F2()
//            > ```
//                  ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/impls`/T2#F2().
  
