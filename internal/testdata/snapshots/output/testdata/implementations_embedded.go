  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import "io"
//        ^^ reference github.com/golang/go/src go1.19 io/
  
  type I3 interface {
//     ^^ definition 0.1.test sg/testdata/I3#
//     documentation ```go
//     documentation ```go
   Close() error
// ^^^^^ definition 0.1.test sg/testdata/I3#Close.
// documentation ```go
  }
  
  type TClose struct {
//     ^^^^^^ definition 0.1.test sg/testdata/TClose#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src go1.19 io/Closer# implementation
//     relationship 0.1.test sg/testdata/I3# implementation
   io.Closer
// ^^ reference github.com/golang/go/src go1.19 io/
//    ^^^^^^ definition 0.1.test sg/testdata/TClose#Closer.
//    documentation ```go
//    ^^^^^^ reference github.com/golang/go/src go1.19 io/Closer#
  }
  
