  package nested_internal
//        ^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded/internal`/
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "sg/embedded"
//  ^^^^^^^^^^^ reference 0.1.test `sg/embedded`/
  )
  
//⌄ enclosing_range_start 0.1.test `sg/embedded/internal`/Something().
  func Something(recent embedded.RecentCommittersResults) {
//     ^^^^^^^^^ definition 0.1.test `sg/embedded/internal`/Something().
//     documentation
//     > ```go
//     > func Something(recent RecentCommittersResults)
//     > ```
//               ^^^^^^ definition local 0
//                      ^^^^^^^^ reference 0.1.test `sg/embedded`/
//                               ^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#
   for _, commit := range recent.Nodes {
//        ^^^^^^ definition local 1
//                        ^^^^^^ reference local 0
//                               ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
    for _, author := range commit.Authors.Nodes {
//         ^^^^^^ definition local 2
//                         ^^^^^^ reference local 1
//                                ^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#$anon_d6a31d3215fd380e#Authors.
//                                        ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#$anon_d6a31d3215fd380e#$anon_ae22730d5ab4639d#Nodes.
     fmt.Println(author.Name)
//   ^^^ reference github.com/golang/go/src go1.22 fmt/
//       ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//               ^^^^^^ reference local 2
//                      ^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#$anon_d6a31d3215fd380e#$anon_ae22730d5ab4639d#$anon_9f31da10601e7ccd#Name.
    }
   }
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded/internal`/Something().
  
