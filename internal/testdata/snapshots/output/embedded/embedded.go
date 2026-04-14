  package embedded
//        ^^^^^^^^ definition 0.1.test `sg/embedded`/
//                 display_name embedded
//                 signature_documentation
//                 > package embedded
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "os/exec"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
  )
  
  type osExecCommand struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/osExecCommand#
//                   signature_documentation
//                   > type osExecCommand struct {
//                   >     *Cmd
//                   > }
//                   relationship github.com/golang/go/src go1.22 context/stringer# implementation
//                   relationship github.com/golang/go/src go1.22 fmt/Stringer# implementation
//                   relationship github.com/golang/go/src go1.22 runtime/stringer# implementation
   *exec.Cmd
//  ^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
//       ^^^ definition 0.1.test `sg/embedded`/osExecCommand#Cmd.
//           signature_documentation
//           > struct field Cmd *os/exec.Cmd
//       ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
  }
  
//⌄ enclosing_range_start 0.1.test `sg/embedded`/wrapExecCommand().
  func wrapExecCommand(c *exec.Cmd) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/wrapExecCommand().
//                     signature_documentation
//                     > func wrapExecCommand(c *Cmd)
//                     ^ definition local 0
//                       display_name c
//                       signature_documentation
//                       > var c *os/exec.Cmd
//                        ^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
//                             ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
   _ = &osExecCommand{Cmd: c}
//      ^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/osExecCommand#
//                    ^^^ reference 0.1.test `sg/embedded`/osExecCommand#Cmd.
//                         ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/embedded`/wrapExecCommand().
  
  type Inner struct {
//     ^^^^^ definition 0.1.test `sg/embedded`/Inner#
//           signature_documentation
//           > type Inner struct {
//           >     X int
//           >     Y int
//           >     Z int
//           > }
   X int
// ^ definition 0.1.test `sg/embedded`/Inner#X.
//   signature_documentation
//   > struct field X int
   Y int
// ^ definition 0.1.test `sg/embedded`/Inner#Y.
//   signature_documentation
//   > struct field Y int
   Z int
// ^ definition 0.1.test `sg/embedded`/Inner#Z.
//   signature_documentation
//   > struct field Z int
  }
  
  type Outer struct {
//     ^^^^^ definition 0.1.test `sg/embedded`/Outer#
//           signature_documentation
//           > type Outer struct {
//           >     Inner
//           >     W int
//           > }
   Inner
// ^^^^^ definition 0.1.test `sg/embedded`/Outer#Inner.
//       signature_documentation
//       > struct field Inner sg/embedded.Inner
// ^^^^^ reference 0.1.test `sg/embedded`/Inner#
   W int
// ^ definition 0.1.test `sg/embedded`/Outer#W.
//   signature_documentation
//   > struct field W int
  }
  
//⌄ enclosing_range_start 0.1.test `sg/embedded`/useOfCompositeStructs().
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/useOfCompositeStructs().
//                           signature_documentation
//                           > func useOfCompositeStructs()
   o := Outer{
// ^ definition local 1
//   display_name o
//   signature_documentation
//   > var o sg/embedded.Outer
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
//⌃ enclosing_range_end 0.1.test `sg/embedded`/useOfCompositeStructs().
  
