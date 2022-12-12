  package inlinestruct
//        ^^^^^^^^^^^^ reference sg/inlinestruct/
  
  import "context"
//        ^^^^^^^ reference github.com/golang/go/src context/
  
  func Target() interface {
//     ^^^^^^ definition sg/inlinestruct/Target().
//     documentation ```go
   OID(context.Context) (int, error)
// ^^^ definition sg/inlinestruct/func:Target:OID().
// documentation ```go
//     ^^^^^^^ reference github.com/golang/go/src context/
//             ^^^^^^^ reference github.com/golang/go/src context/Context#
   AbbreviatedOID(context.Context) (string, error)
// ^^^^^^^^^^^^^^ definition sg/inlinestruct/func:Target:AbbreviatedOID().
// documentation ```go
//                ^^^^^^^ reference github.com/golang/go/src context/
//                        ^^^^^^^ reference github.com/golang/go/src context/Context#
   Commit(context.Context) (string, error)
// ^^^^^^ definition sg/inlinestruct/func:Target:Commit().
// documentation ```go
//        ^^^^^^^ reference github.com/golang/go/src context/
//                ^^^^^^^ reference github.com/golang/go/src context/Context#
   Type(context.Context) (int, error)
// ^^^^ definition sg/inlinestruct/func:Target:Type().
// documentation ```go
//      ^^^^^^^ reference github.com/golang/go/src context/
//              ^^^^^^^ reference github.com/golang/go/src context/Context#
  } {
   panic("not implemented")
  }
  
  func something() {
//     ^^^^^^^^^ definition sg/inlinestruct/something().
//     documentation ```go
   x := Target()
// ^ definition local 0
//      ^^^^^^ reference sg/inlinestruct/Target().
   x.OID(context.Background())
// ^ reference local 0
//   ^^^ reference sg/inlinestruct/func:Target:OID().
//       ^^^^^^^ reference github.com/golang/go/src context/
//               ^^^^^^^^^^ reference github.com/golang/go/src context/Background().
  }
  
