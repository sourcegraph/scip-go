  package embedded
//        ^^^^^^^^ definition 0.1.test `sg/embedded`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start 0.1.test `sg/embedded`/RecentCommittersResults#String().
  func (r *RecentCommittersResults) String() string {
//      ^ definition local 0
//        display_name r
//        signature_documentation
//        > var r *sg/embedded.RecentCommittersResults
//         ^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#
//                                  ^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#String().
//                                         documentation
//                                         > ```go
//                                         > func (*RecentCommittersResults).String() string
//                                         > ```
//                                         relationship github.com/golang/go/src go1.22 context/stringer#String. implementation
//                                         relationship github.com/golang/go/src go1.22 fmt/Stringer#String. implementation
//                                         relationship github.com/golang/go/src go1.22 runtime/stringer#String. implementation
   return fmt.Sprintf("RecentCommittersResults{Nodes: %d}", len(r.Nodes))
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
//            ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Sprintf().
//                                                              ^ reference local 0
//                                                                ^^^^^ reference 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded`/RecentCommittersResults#String().
  
  type RecentCommittersResults struct {
//     ^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#
//                             documentation
//                             > ```go
//                             > type RecentCommittersResults struct
//                             > ```
//                             documentation
//                             > ```go
//                             > struct {
//                             >     Nodes []struct {
//                             >         Authors struct {
//                             >             Nodes []struct {
//                             >                 Date string
//                             >                 Email string
//                             >                 Name string
//                             >                 User struct {
//                             >                     Login string
//                             >                 }
//                             >                 AvatarURL string
//                             >             }
//                             >         }
//                             >     }
//                             >     PageInfo struct {
//                             >         HasNextPage bool
//                             >     }
//                             > }
//                             > ```
//                             relationship github.com/golang/go/src go1.22 context/stringer# implementation
//                             relationship github.com/golang/go/src go1.22 fmt/Stringer# implementation
//                             relationship github.com/golang/go/src go1.22 runtime/stringer# implementation
   Nodes []struct {
// ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#Nodes.
//       documentation
//       > ```go
//       > struct field Nodes []struct{Authors struct{Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}}}
//       > ```
    Authors struct {
//  ^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#Authors.
//          documentation
//          > ```go
//          > struct field Authors struct{Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}}
//          > ```
     Nodes []struct {
//   ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#Nodes.
//         documentation
//         > ```go
//         > struct field Nodes []struct{Date string; Email string; Name string; User struct{Login string}; AvatarURL string}
//         > ```
      Date  string
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#Date.
//         documentation
//         > ```go
//         > struct field Date string
//         > ```
      Email string
//    ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#Email.
//          documentation
//          > ```go
//          > struct field Email string
//          > ```
      Name  string
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#Name.
//         documentation
//         > ```go
//         > struct field Name string
//         > ```
      User  struct {
//    ^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#User.
//         documentation
//         > ```go
//         > struct field User struct{Login string}
//         > ```
       Login string
//     ^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#$anon_d4bff1f61f45b2a1#Login.
//           documentation
//           > ```go
//           > struct field Login string
//           > ```
      }
      AvatarURL string
//    ^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_90b32de18ec80596#$anon_bb290b1f6ea0cf58#$anon_b2a8a16c744b2d4b#AvatarURL.
//              documentation
//              > ```go
//              > struct field AvatarURL string
//              > ```
     }
    }
   }
   PageInfo struct {
// ^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#PageInfo.
//          documentation
//          > ```go
//          > struct field PageInfo struct{HasNextPage bool}
//          > ```
    HasNextPage bool
//  ^^^^^^^^^^^ definition 0.1.test `sg/embedded`/RecentCommittersResults#$anon_0a5c453971407ce4#HasNextPage.
//              documentation
//              > ```go
//              > struct field HasNextPage bool
//              > ```
   }
  }
  
