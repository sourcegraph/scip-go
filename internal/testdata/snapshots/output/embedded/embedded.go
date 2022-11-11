  package embedded
  
  import "os/exec"
//        ^^^^^^^ reference github.com/golang/go/src os/exec/
  
  type osExecCommand struct {
//     ^^^^^^^^^^^^^ definition sg/embedded/osExecCommand#
   *exec.Cmd
//       ^^^ definition sg/embedded/osExecCommand#Cmd.
//       ^^^ reference github.com/golang/go/src os/exec/Cmd#
  }
  
  func wrapExecCommand(c *exec.Cmd) {
//     ^^^^^^^^^^^^^^^ definition sg/embedded/wrapExecCommand().
//                     ^ definition local 0
//                             ^^^ reference github.com/golang/go/src os/exec/Cmd#
   _ = &osExecCommand{Cmd: c}
//      ^^^^^^^^^^^^^ reference sg/embedded/osExecCommand#
//                    ^^^ reference sg/embedded/osExecCommand#Cmd.
//                         ^ reference local 0
  }
  
