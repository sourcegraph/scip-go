  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import "io"
//        ^^ reference v1.19 io/
  
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
//     relationship v1.19 io/Closer# implementation
//     relationship 0.1.test sg/testdata/I3# implementation
   io.Closer
// ^^ reference v1.19 io/
//    ^^^^^^ definition 0.1.test sg/testdata/TClose#Closer.
//    documentation ```go
//    ^^^^^^ reference v1.19 io/Closer#
  }
  
