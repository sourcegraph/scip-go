  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/
  
  import "context"
//        ^^^^^^^ reference github.com/golang/go/src go1.19 context/
  
  func Target() interface {
//     ^^^^^^ definition 0.1.test sg/inlinestruct/Target().
//     documentation ```go
   OID(context.Context) (int, error)
// ^^^ definition 0.1.test sg/inlinestruct/func:Target:OID().
// documentation ```go
//     ^^^^^^^ reference github.com/golang/go/src go1.19 context/
//             ^^^^^^^ reference github.com/golang/go/src go1.19 context/Context#
   AbbreviatedOID(context.Context) (string, error)
// ^^^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/func:Target:AbbreviatedOID().
// documentation ```go
//                ^^^^^^^ reference github.com/golang/go/src go1.19 context/
//                        ^^^^^^^ reference github.com/golang/go/src go1.19 context/Context#
   Commit(context.Context) (string, error)
// ^^^^^^ definition 0.1.test sg/inlinestruct/func:Target:Commit().
// documentation ```go
//        ^^^^^^^ reference github.com/golang/go/src go1.19 context/
//                ^^^^^^^ reference github.com/golang/go/src go1.19 context/Context#
   Type(context.Context) (int, error)
// ^^^^ definition 0.1.test sg/inlinestruct/func:Target:Type().
// documentation ```go
//      ^^^^^^^ reference github.com/golang/go/src go1.19 context/
//              ^^^^^^^ reference github.com/golang/go/src go1.19 context/Context#
  } {
   panic("not implemented")
  }
  
  func something() {
//     ^^^^^^^^^ definition 0.1.test sg/inlinestruct/something().
//     documentation ```go
   x := Target()
// ^ definition local 0
//      ^^^^^^ reference 0.1.test sg/inlinestruct/Target().
   x.OID(context.Background())
// ^ reference local 0
//   ^^^ reference 0.1.test sg/inlinestruct/func:Target:OID().
//       ^^^^^^^ reference github.com/golang/go/src go1.19 context/
//               ^^^^^^^^^^ reference github.com/golang/go/src go1.19 context/Background().
  }
  
