  package nested_internal
//        ^^^^^^^^^^^^^^^ definition 0.1.test sg/embedded/internal/
//        documentation package nested_internal
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.19 fmt/
   "sg/embedded"
//  ^^^^^^^^^^^ reference 0.1.test sg/embedded/
  )
  
  func Something(recent embedded.RecentCommittersResults) {
//     ^^^^^^^^^ definition 0.1.test sg/embedded/internal/Something().
//     documentation ```go
//               ^^^^^^ definition local 0
//                      ^^^^^^^^ reference 0.1.test sg/embedded/
//                               ^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test sg/embedded/RecentCommittersResults#
   for _, commit := range recent.Nodes {
//        ^^^^^^ definition local 1
//                        ^^^^^^ reference local 0
//                               ^^^^^ reference 0.1.test sg/embedded/RecentCommittersResults#Nodes.
    for _, author := range commit.Authors.Nodes {
//         ^^^^^^ definition local 2
//                         ^^^^^^ reference local 1
//                                ^^^^^^^ reference 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.
//                                        ^^^^^ reference 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.
     fmt.Println(author.Name)
//   ^^^ reference github.com/golang/go/src go1.19 fmt/
//       ^^^^^^^ reference github.com/golang/go/src go1.19 fmt/Println().
//               ^^^^^^ reference local 2
//                      ^^^^ reference 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Name.
    }
   }
  }
  
