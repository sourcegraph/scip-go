  package embedded
//        ^^^^^^^^ reference sg/embedded/
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src fmt/
   "os/exec"
//  ^^^^^^^ reference github.com/golang/go/src os/exec/
  )
  
  type osExecCommand struct {
//     ^^^^^^^^^^^^^ definition sg/embedded/osExecCommand#
//     documentation ```go
//     documentation ```go
   *exec.Cmd
//  ^^^^ reference github.com/golang/go/src os/exec/
//       ^^^ definition sg/embedded/osExecCommand#Cmd.
//       documentation ```go
//       ^^^ reference github.com/golang/go/src os/exec/Cmd#
  }
  
  func wrapExecCommand(c *exec.Cmd) {
//     ^^^^^^^^^^^^^^^ definition sg/embedded/wrapExecCommand().
//     documentation ```go
//                     ^ definition local 0
//                        ^^^^ reference github.com/golang/go/src os/exec/
//                             ^^^ reference github.com/golang/go/src os/exec/Cmd#
   _ = &osExecCommand{Cmd: c}
//      ^^^^^^^^^^^^^ reference sg/embedded/osExecCommand#
//                    ^^^ reference sg/embedded/osExecCommand#Cmd.
//                         ^ reference local 0
  }
  
  type Inner struct {
//     ^^^^^ definition sg/embedded/Inner#
//     documentation ```go
//     documentation ```go
   X int
// ^ definition sg/embedded/Inner#X.
// documentation ```go
   Y int
// ^ definition sg/embedded/Inner#Y.
// documentation ```go
   Z int
// ^ definition sg/embedded/Inner#Z.
// documentation ```go
  }
  
  type Outer struct {
//     ^^^^^ definition sg/embedded/Outer#
//     documentation ```go
//     documentation ```go
   Inner
// ^^^^^ definition sg/embedded/Outer#Inner.
// documentation ```go
// ^^^^^ reference sg/embedded/Inner#
   W int
// ^ definition sg/embedded/Outer#W.
// documentation ```go
  }
  
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition sg/embedded/useOfCompositeStructs().
//     documentation ```go
   o := Outer{
// ^ definition local 1
//      ^^^^^ reference sg/embedded/Outer#
    Inner: Inner{
//  ^^^^^ reference sg/embedded/Outer#Inner.
//         ^^^^^ reference sg/embedded/Inner#
     X: 1,
//   ^ reference sg/embedded/Inner#X.
     Y: 2,
//   ^ reference sg/embedded/Inner#Y.
     Z: 3,
//   ^ reference sg/embedded/Inner#Z.
    },
    W: 4,
//  ^ reference sg/embedded/Outer#W.
   }
  
   fmt.Printf("> %d\n", o.X)
// ^^^ reference github.com/golang/go/src fmt/
//     ^^^^^^ reference github.com/golang/go/src fmt/Printf().
//                      ^ reference local 1
//                        ^ reference sg/embedded/Inner#X.
   fmt.Println(o.Inner.Y)
// ^^^ reference github.com/golang/go/src fmt/
//     ^^^^^^^ reference github.com/golang/go/src fmt/Println().
//             ^ reference local 1
//               ^^^^^ reference sg/embedded/Outer#Inner.
//                     ^ reference sg/embedded/Inner#Y.
  }
  
