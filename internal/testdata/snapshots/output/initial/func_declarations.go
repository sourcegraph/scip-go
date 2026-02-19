  package initial
//        ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/UsesLater().
  func UsesLater() {
//     ^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/UsesLater().
//     documentation
//     > ```go
//     > func UsesLater()
//     > ```
   DefinedLater()
// ^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/DefinedLater().
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/UsesLater().
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/DefinedLater().
  func DefinedLater() {}
//     ^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/DefinedLater().
//     documentation
//     > ```go
//     > func DefinedLater()
//     > ```
//                     ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/initial`/DefinedLater().
  
