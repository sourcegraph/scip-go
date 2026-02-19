  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/
//        documentation
//        > package inlinestruct
  
  type FieldInterface interface {
//     ^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/FieldInterface#
//     documentation
//     > ```go
//     > type FieldInterface interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     SomeMethod() string
//     > }
//     > ```
   SomeMethod() string
// ^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/FieldInterface#SomeMethod.
// documentation
// > ```go
// > func (FieldInterface).SomeMethod() string
// > ```
  }
  
  var MyInline = struct {
//    ^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyInline.
//    documentation
//    > ```go
//    > var MyInline struct{privateField FieldInterface; PublicField FieldInterface}
//    > ```
   privateField FieldInterface
// ^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyInline:privateField.
// documentation
// > ```go
// > struct field privateField github.com/sourcegraph/scip-go/internal/testdata/inlinestruct.FieldInterface
// > ```
//              ^^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/FieldInterface#
   PublicField  FieldInterface
// ^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyInline:PublicField.
// documentation
// > ```go
// > struct field PublicField github.com/sourcegraph/scip-go/internal/testdata/inlinestruct.FieldInterface
// > ```
//              ^^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/FieldInterface#
  }{}
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyFunc().
  func MyFunc() {
//     ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyFunc().
//     documentation
//     > ```go
//     > func MyFunc()
//     > ```
   _ = MyInline.privateField
//     ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyInline.
//              ^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyInline:privateField.
   _ = MyInline.PublicField
//     ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyInline.
//              ^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyInline:PublicField.
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/inlinestruct`/MyFunc().
  
