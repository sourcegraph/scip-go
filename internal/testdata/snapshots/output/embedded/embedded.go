  package embedded
//        ^^^^^^^^ reference 0.1.test sg/embedded/
  
  import (
   "fmt"
//  ^^^ reference v1.19 fmt/
   "os/exec"
//  ^^^^^^^ reference v1.19 os/exec/
  )
  
  type osExecCommand struct {
//     ^^^^^^^^^^^^^ definition 0.1.test sg/embedded/osExecCommand#
//     documentation ```go
//     documentation ```go
//     relationship v1.19 context/stringer# implementation
//     relationship v1.19 fmt/Stringer# implementation
//     relationship v1.19 runtime/stringer# implementation
   *exec.Cmd
//  ^^^^ reference v1.19 os/exec/
//       ^^^ definition 0.1.test sg/embedded/osExecCommand#Cmd.
//       documentation ```go
//       ^^^ reference v1.19 os/exec/Cmd#
  }
  
  func wrapExecCommand(c *exec.Cmd) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test sg/embedded/wrapExecCommand().
//     documentation ```go
//                     ^ definition local 0
//                        ^^^^ reference v1.19 os/exec/
//                             ^^^ reference v1.19 os/exec/Cmd#
   _ = &osExecCommand{Cmd: c}
//      ^^^^^^^^^^^^^ reference 0.1.test sg/embedded/osExecCommand#
//                    ^^^ reference 0.1.test sg/embedded/osExecCommand#Cmd.
//                         ^ reference local 0
  }
  
  type Inner struct {
//     ^^^^^ definition 0.1.test sg/embedded/Inner#
//     documentation ```go
//     documentation ```go
   X int
// ^ definition 0.1.test sg/embedded/Inner#X.
// documentation ```go
   Y int
// ^ definition 0.1.test sg/embedded/Inner#Y.
// documentation ```go
   Z int
// ^ definition 0.1.test sg/embedded/Inner#Z.
// documentation ```go
  }
  
  type Outer struct {
//     ^^^^^ definition 0.1.test sg/embedded/Outer#
//     documentation ```go
//     documentation ```go
   Inner
// ^^^^^ definition 0.1.test sg/embedded/Outer#Inner.
// documentation ```go
// ^^^^^ reference 0.1.test sg/embedded/Inner#
   W int
// ^ definition 0.1.test sg/embedded/Outer#W.
// documentation ```go
  }
  
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/embedded/useOfCompositeStructs().
//     documentation ```go
   o := Outer{
// ^ definition local 1
//      ^^^^^ reference 0.1.test sg/embedded/Outer#
    Inner: Inner{
//  ^^^^^ reference 0.1.test sg/embedded/Outer#Inner.
//         ^^^^^ reference 0.1.test sg/embedded/Inner#
     X: 1,
//   ^ reference 0.1.test sg/embedded/Inner#X.
     Y: 2,
//   ^ reference 0.1.test sg/embedded/Inner#Y.
     Z: 3,
//   ^ reference 0.1.test sg/embedded/Inner#Z.
    },
    W: 4,
//  ^ reference 0.1.test sg/embedded/Outer#W.
   }
  
   fmt.Printf("> %d\n", o.X)
// ^^^ reference v1.19 fmt/
//     ^^^^^^ reference v1.19 fmt/Printf().
//                      ^ reference local 1
//                        ^ reference 0.1.test sg/embedded/Inner#X.
   fmt.Println(o.Inner.Y)
// ^^^ reference v1.19 fmt/
//     ^^^^^^^ reference v1.19 fmt/Println().
//             ^ reference local 1
//               ^^^^^ reference 0.1.test sg/embedded/Outer#Inner.
//                     ^ reference 0.1.test sg/embedded/Inner#Y.
  }
  
