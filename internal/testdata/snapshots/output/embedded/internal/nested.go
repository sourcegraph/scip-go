  package nested_internal
//        ^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded/internal`/
//        documentation
//        > package nested_internal
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "github.com/sourcegraph/scip-go/internal/testdata/embedded"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/
  )
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded/internal`/Something().
  func Something(recent embedded.RecentCommittersResults) {
//     ^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded/internal`/Something().
//     documentation
//     > ```go
//     > func Something(recent RecentCommittersResults)
//     > ```
//               ^^^^^^ definition local 0
//                      ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/
//                               ^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/RecentCommittersResults#
   for _, commit := range recent.Nodes {
//        ^^^^^^ definition local 1
//                        ^^^^^^ reference local 0
//                               ^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/RecentCommittersResults#Nodes.
    for _, author := range commit.Authors.Nodes {
//         ^^^^^^ definition local 2
//                         ^^^^^^ reference local 1
//                                ^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/RecentCommittersResults#Nodes.Authors.
//                                        ^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.
     fmt.Println(author.Name)
//   ^^^ reference github.com/golang/go/src go1.22 fmt/
//       ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//               ^^^^^^ reference local 2
//                      ^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.Name.
    }
   }
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded/internal`/Something().
  
