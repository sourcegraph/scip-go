  package initial
//        ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/
  
  type (
   // HoverTypeList is a cool struct
   HoverTypeList struct{}
// ^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/HoverTypeList#
// documentation
// > ```go
// > type HoverTypeList struct
// > ```
// documentation
// > ```go
// > struct{}
// > ```
  )
  
  // This should show up as well
  type HoverType struct{}
//     ^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/HoverType#
//     documentation
//     > ```go
//     > type HoverType struct
//     > ```
//     documentation
//     > This should show up as well
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
