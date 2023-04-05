  package embedded
//        ^^^^^^^^ reference 0.1.test sg/embedded/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.19 fmt/
  
  type RecentCommittersResults struct {
//     ^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#
//     documentation ```go
//     documentation ```go
   Nodes []struct {
// ^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.
// documentation ```go
    Authors struct {
//  ^^^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.
//  documentation ```go
     Nodes []struct {
//   ^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.
//   documentation ```go
      Date  string
//    ^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Date.
//    documentation ```go
      Email string
//    ^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Email.
//    documentation ```go
      Name  string
//    ^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.Name.
//    documentation ```go
      User  struct {
//    ^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.User.
//    documentation ```go
       Login string
//     ^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.User.Login.
//     documentation ```go
      }
      AvatarURL string
//    ^^^^^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#Nodes.Authors.Nodes.AvatarURL.
//    documentation ```go
     }
    }
   }
   PageInfo struct {
// ^^^^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#PageInfo.
// documentation ```go
    HasNextPage bool
//  ^^^^^^^^^^^ definition 0.1.test sg/embedded/RecentCommittersResults#PageInfo.HasNextPage.
//  documentation ```go
   }
  }
  
