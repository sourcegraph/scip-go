  package embedded
//        ^^^^^^^^ reference 0.1.test `sg/embedded`/
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "os/exec"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
  )
  
  type osExecCommand struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/osExecCommand#
//     documentation
//     > ```go
//     > type osExecCommand struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     *Cmd
//     > }
//     > ```
//     relationship github.com/golang/go/src go1.22 context/stringer# implementation
//     relationship github.com/golang/go/src go1.22 fmt/Stringer# implementation
//     relationship github.com/golang/go/src go1.22 runtime/stringer# implementation
   *exec.Cmd
//  ^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
//       ^^^ definition 0.1.test `sg/embedded`/osExecCommand#Cmd.
//       documentation
//       > ```go
//       > struct field Cmd *os/exec.Cmd
//       > ```
//       ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
  }
  
  func wrapExecCommand(c *exec.Cmd) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/wrapExecCommand().
//     documentation
//     > ```go
//     > func wrapExecCommand(c *Cmd)
//     > ```
//                     ^ definition local 0
//                        ^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
//                             ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
   _ = &osExecCommand{Cmd: c}
//      ^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/osExecCommand#
//                    ^^^ reference 0.1.test `sg/embedded`/osExecCommand#Cmd.
//                         ^ reference local 0
  }
  
  type Inner struct {
//     ^^^^^ definition 0.1.test `sg/embedded`/Inner#
//     documentation
//     > ```go
//     > type Inner struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     X int
//     >     Y int
//     >     Z int
//     > }
//     > ```
   X int
// ^ definition 0.1.test `sg/embedded`/Inner#X.
// documentation
// > ```go
// > struct field X int
// > ```
   Y int
// ^ definition 0.1.test `sg/embedded`/Inner#Y.
// documentation
// > ```go
// > struct field Y int
// > ```
   Z int
// ^ definition 0.1.test `sg/embedded`/Inner#Z.
// documentation
// > ```go
// > struct field Z int
// > ```
  }
  
  type Outer struct {
//     ^^^^^ definition 0.1.test `sg/embedded`/Outer#
//     documentation
//     > ```go
//     > type Outer struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Inner
//     >     W int
//     > }
//     > ```
   Inner
// ^^^^^ definition 0.1.test `sg/embedded`/Outer#Inner.
// documentation
// > ```go
// > struct field Inner sg/embedded.Inner
// > ```
// ^^^^^ reference 0.1.test `sg/embedded`/Inner#
   W int
// ^ definition 0.1.test `sg/embedded`/Outer#W.
// documentation
// > ```go
// > struct field W int
// > ```
  }
  
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/useOfCompositeStructs().
//     documentation
//     > ```go
//     > func useOfCompositeStructs()
//     > ```
   o := Outer{
// ^ definition local 1
//      ^^^^^ reference 0.1.test `sg/embedded`/Outer#
    Inner: Inner{
//  ^^^^^ reference 0.1.test `sg/embedded`/Outer#Inner.
//         ^^^^^ reference 0.1.test `sg/embedded`/Inner#
     X: 1,
//   ^ reference 0.1.test `sg/embedded`/Inner#X.
     Y: 2,
//   ^ reference 0.1.test `sg/embedded`/Inner#Y.
     Z: 3,
//   ^ reference 0.1.test `sg/embedded`/Inner#Z.
    },
    W: 4,
//  ^ reference 0.1.test `sg/embedded`/Outer#W.
   }
  
   fmt.Printf("> %d\n", o.X)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                      ^ reference local 1
//                        ^ reference 0.1.test `sg/embedded`/Inner#X.
   fmt.Println(o.Inner.Y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^ reference local 1
//               ^^^^^ reference 0.1.test `sg/embedded`/Outer#Inner.
//                     ^ reference 0.1.test `sg/embedded`/Inner#Y.
  }
  
