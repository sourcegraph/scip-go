  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/
  
  import "context"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 context/
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/Target().
  func Target() interface {
//     ^^^^^^ definition 0.1.test `sg/inlinestruct`/Target().
//     kind Function
//     documentation
//     > ```go
//     > func Target() interface{AbbreviatedOID(Context) (string, error); Commit(Context) (string, error); OID(Context) (int, error); Type(Context) (int, error)}
//     > ```
   OID(context.Context) (int, error)
// ^^^ definition 0.1.test `sg/inlinestruct`/func:Target:OID().
// kind Method
// documentation
// > ```go
// > func (interface).OID(Context) (int, error)
// > ```
//     ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//             ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
   AbbreviatedOID(context.Context) (string, error)
// ^^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/func:Target:AbbreviatedOID().
// kind Method
// documentation
// > ```go
// > func (interface).AbbreviatedOID(Context) (string, error)
// > ```
//                ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                        ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
   Commit(context.Context) (string, error)
// ^^^^^^ definition 0.1.test `sg/inlinestruct`/func:Target:Commit().
// kind Method
// documentation
// > ```go
// > func (interface).Commit(Context) (string, error)
// > ```
//        ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
   Type(context.Context) (int, error)
// ^^^^ definition 0.1.test `sg/inlinestruct`/func:Target:Type().
// kind Method
// documentation
// > ```go
// > func (interface).Type(Context) (int, error)
// > ```
//      ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//              ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
  } {
   panic("not implemented")
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/Target().
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/something().
  func something() {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/something().
//     kind Function
//     documentation
//     > ```go
//     > func something()
//     > ```
   x := Target()
// ^ definition local 0
// kind Variable
//      ^^^^^^ reference 0.1.test `sg/inlinestruct`/Target().
   x.OID(context.Background())
// ^ reference local 0
//   ^^^ reference 0.1.test `sg/inlinestruct`/func:Target:OID().
//       ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//               ^^^^^^^^^^ reference github.com/golang/go/src go1.22 context/Background().
  }
//⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/something().
  
