  package testdata
//        ^^^^^^^^ reference sg/testdata/
  
  import "io"
//        ^^ reference github.com/golang/go/src io/
  
  type I3 interface {
//     ^^ definition sg/testdata/I3#
//     documentation ```go
//     documentation ```go
   Close() error
// ^^^^^ definition sg/testdata/I3#Close.
// documentation ```go
  }
  
  type TClose struct {
//     ^^^^^^ definition sg/testdata/TClose#
//     documentation ```go
//     documentation ```go
   io.Closer
// ^^ reference github.com/golang/go/src io/
//    ^^^^^^ definition sg/testdata/TClose#Closer.
//    documentation ```go
//    ^^^^^^ reference github.com/golang/go/src io/Closer#
  }
  
