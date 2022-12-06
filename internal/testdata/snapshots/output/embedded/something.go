  package embedded
//        ^^^^^^^^ reference sg/embedded/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src fmt/
  
  type RecentCommittersResults struct {
//     ^^^^^^^^^^^^^^^^^^^^^^^ definition sg/embedded/RecentCommittersResults#
//     documentation ```go
//     documentation ```go
   Nodes []struct {
// ^^^^^ definition sg/embedded/RecentCommittersResults#Nodes.
// documentation ```go
    Authors struct {
//  ^^^^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.
//  documentation ```go
     Nodes []struct {
//   ^^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.
//   documentation ```go
      Date  string
//    ^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Date.
//    documentation ```go
      Email string
//    ^^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Email.
//    documentation ```go
      Name  string
//    ^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Name.
//    documentation ```go
      User  struct {
//    ^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.User.
//    documentation ```go
       Login string
//     ^^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.User.Login.
//     documentation ```go
      }
      AvatarURL string
//    ^^^^^^^^^ definition sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.AvatarURL.
//    documentation ```go
     }
    }
   }
   PageInfo struct {
// ^^^^^^^^ definition sg/embedded/RecentCommittersResults#PageInfo.
// documentation ```go
    HasNextPage bool
//  ^^^^^^^^^^^ definition sg/embedded/RecentCommittersResults#PageInfo.HasNextPage.
//  documentation ```go
   }
  }
  
