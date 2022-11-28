  package embedded
  
  import "os/exec"
//        ^^^^^^^ reference github.com/golang/go/src os/exec/
  
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
  
