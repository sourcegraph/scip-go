  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  import "io"
//        ^^ reference github.com/golang/go/src go1.22 io/
  
  type I3 interface {
//     ^^ definition 0.1.test `sg/testdata`/I3#
//        signature_documentation
//        > type I3 interface
//        > interface {
//        >     Close() error
//        > }
   Close() error
// ^^^^^ definition 0.1.test `sg/testdata`/I3#Close.
//       signature_documentation
//       > func (I3).Close() error
  }
  
  type TClose struct {
//     ^^^^^^ definition 0.1.test `sg/testdata`/TClose#
//            signature_documentation
//            > type TClose struct
//            > struct {
//            >     Closer
//            > }
//            relationship github.com/golang/go/src go1.22 io/Closer# implementation
//            relationship 0.1.test `sg/testdata`/I3# implementation
   io.Closer
// ^^ reference github.com/golang/go/src go1.22 io/
//    ^^^^^^ definition 0.1.test `sg/testdata`/TClose#Closer.
//           signature_documentation
//           > struct field Closer io.Closer
//    ^^^^^^ reference github.com/golang/go/src go1.22 io/Closer#
  }
  
