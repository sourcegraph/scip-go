  package nested_internal
//        ^^^^^^^^^^^^^^^ definition sg/embedded/internal/
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src fmt/
   "sg/embedded"
//  ^^^^^^^^^^^ reference sg/embedded/
  )
  
  func Something(recent embedded.RecentCommittersResults) {
//     ^^^^^^^^^ definition sg/embedded/internal/Something().
//     documentation ```go
//               ^^^^^^ definition local 0
//                      ^^^^^^^^ reference sg/embedded/
//                               ^^^^^^^^^^^^^^^^^^^^^^^ reference sg/embedded/RecentCommittersResults#
   for _, commit := range recent.Nodes {
//        ^^^^^^ definition local 1
//                        ^^^^^^ reference local 0
//                               ^^^^^ reference sg/embedded/RecentCommittersResults#Nodes.
    for _, author := range commit.Authors.Nodes {
//         ^^^^^^ definition local 2
//                         ^^^^^^ reference local 1
//                                ^^^^^^^ reference sg/embedded/RecentCommittersResults#Nodes.Authors.
//                                        ^^^^^ reference sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.
     fmt.Println(author.Name)
//   ^^^ reference github.com/golang/go/src fmt/
//       ^^^^^^^ reference github.com/golang/go/src fmt/Println().
//               ^^^^^^ reference local 2
//                      ^^^^ reference sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Name.
    }
   }
  }
  
