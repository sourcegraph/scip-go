  package nested_internal
//        ^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded/internal`/
  
  import (
   "fmt"
//  ^^^ definition local 0
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "sg/embedded"
//  ^^^^^^^^^^^ reference 0.1.test `sg/embedded`/
//     ^^^^^^^^ definition local 1
  )
  
//⌄ enclosing_range_start 0.1.test `sg/embedded/internal`/Something().
  func Something(recent embedded.RecentCommittersResults) {
//     ^^^^^^^^^ definition 0.1.test `sg/embedded/internal`/Something().
//     documentation
//     > ```go
//     > func Something(recent RecentCommittersResults)
//     > ```
//               ^^^^^^ definition local 2
//                      ^^^^^^^^ reference local 1
//                               ^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#
   for _, commit := range recent.Nodes {
//        ^^^^^^ definition local 3
//                        ^^^^^^ reference local 2
//                               ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
    for _, author := range commit.Authors.Nodes {
//         ^^^^^^ definition local 4
//                         ^^^^^^ reference local 3
//                                ^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.
//                                        ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.
     fmt.Println(author.Name)
//   ^^^ reference local 0
//       ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//               ^^^^^^ reference local 4
//                      ^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.Name.
    }
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded/internal`/Something().
  
