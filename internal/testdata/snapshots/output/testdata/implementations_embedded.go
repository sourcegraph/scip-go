  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  import "io"
//        ^^ reference github.com/golang/go/src go1.22 io/
  
  type I3 interface {
//     ^^ definition 0.1.test `sg/testdata`/I3#
//     documentation
//     > ```go
//     > type I3 interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Close() error
//     > }
//     > ```
   Close() error
// ^^^^^ definition 0.1.test `sg/testdata`/I3#Close.
// documentation
// > ```go
// > func (I3).Close() error
// > ```
  }
  
  type TClose struct {
//     ^^^^^^ definition 0.1.test `sg/testdata`/TClose#
//     documentation
//     > ```go
//     > type TClose struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Closer
//     > }
//     > ```
//     relationship github.com/golang/go/src go1.22 io/Closer# implementation
//     relationship 0.1.test `sg/testdata`/I3# implementation
   io.Closer
// ^^ reference github.com/golang/go/src go1.22 io/
//    ^^^^^^ definition 0.1.test `sg/testdata`/TClose#Closer.
//    documentation
//    > ```go
//    > struct field Closer io.Closer
//    > ```
//    ^^^^^^ reference github.com/golang/go/src go1.22 io/Closer#
  }
  
