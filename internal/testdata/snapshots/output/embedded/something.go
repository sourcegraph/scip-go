  package embedded
//        ^^^^^^^^ reference 0.1.test `sg/embedded`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type RecentCommittersResults struct {
//     ^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#
//     documentation
//     > ```go
//     > type RecentCommittersResults struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Nodes []struct {
//     >         Authors struct {
//     >             Nodes []struct {
//     >                 Date string
//     >                 Email string
//     >                 Name string
//     >                 User struct {
//     >                     Login string
//     >                 }
//     >                 AvatarURL string
//     >             }
//     >         }
//     >     }
//     >     PageInfo struct {
//     >         HasNextPage bool
//     >     }
//     > }
//     > ```
   Nodes []struct {
// ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
// documentation
// > ```go
// > struct field Nodes []struct{Authors struct{Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}}}
// > ```
    Authors struct {
//  ^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.
//  documentation
//  > ```go
//  > struct field Authors struct{Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}}
//  > ```
     Nodes []struct {
//   ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.
//   documentation
//   > ```go
//   > struct field Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}
//   > ```
      Date  string
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.Date.
//    documentation
//    > ```go
//    > struct field Date string
//    > ```
      Email string
//    ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.Email.
//    documentation
//    > ```go
//    > struct field Email string
//    > ```
      Name  string
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.Name.
//    documentation
//    > ```go
//    > struct field Name string
//    > ```
      User  struct {
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.User.
//    documentation
//    > ```go
//    > struct field User struct{Login string}
//    > ```
       Login string
//     ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.User.Login.
//     documentation
//     > ```go
//     > struct field Login string
//     > ```
      }
      AvatarURL string
//    ^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.Authors.Nodes.AvatarURL.
//    documentation
//    > ```go
//    > struct field AvatarURL string
//    > ```
     }
    }
   }
   PageInfo struct {
// ^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#PageInfo.
// documentation
// > ```go
// > struct field PageInfo struct{HasNextPage bool}
// > ```
    HasNextPage bool
//  ^^^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#PageInfo.HasNextPage.
//  documentation
//  > ```go
//  > struct field HasNextPage bool
//  > ```
   }
  }
  
