  package embedded
//        ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "os/exec"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
  )
  
  type osExecCommand struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/osExecCommand#
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
//       ^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/osExecCommand#Cmd.
//       documentation
//       > ```go
//       > struct field Cmd *os/exec.Cmd
//       > ```
//       ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
  }
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/wrapExecCommand().
  func wrapExecCommand(c *exec.Cmd) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/wrapExecCommand().
//     documentation
//     > ```go
//     > func wrapExecCommand(c *Cmd)
//     > ```
//                     ^ definition local 0
//                        ^^^^ reference github.com/golang/go/src go1.22 `os/exec`/
//                             ^^^ reference github.com/golang/go/src go1.22 `os/exec`/Cmd#
   _ = &osExecCommand{Cmd: c}
//      ^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/osExecCommand#
//                    ^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/osExecCommand#Cmd.
//                         ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/wrapExecCommand().
  
  type Inner struct {
//     ^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#
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
// ^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#X.
// documentation
// > ```go
// > struct field X int
// > ```
   Y int
// ^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#Y.
// documentation
// > ```go
// > struct field Y int
// > ```
   Z int
// ^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#Z.
// documentation
// > ```go
// > struct field Z int
// > ```
  }
  
  type Outer struct {
//     ^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Outer#
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
// ^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Outer#Inner.
// documentation
// > ```go
// > struct field Inner github.com/sourcegraph/scip-go/internal/testdata/embedded.Inner
// > ```
// ^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#
   W int
// ^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Outer#W.
// documentation
// > ```go
// > struct field W int
// > ```
  }
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/useOfCompositeStructs().
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/useOfCompositeStructs().
//     documentation
//     > ```go
//     > func useOfCompositeStructs()
//     > ```
   o := Outer{
// ^ definition local 1
//      ^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Outer#
    Inner: Inner{
//  ^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Outer#Inner.
//         ^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#
     X: 1,
//   ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#X.
     Y: 2,
//   ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#Y.
     Z: 3,
//   ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#Z.
    },
    W: 4,
//  ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Outer#W.
   }
  
   fmt.Printf("> %d\n", o.X)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                      ^ reference local 1
//                        ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#X.
   fmt.Println(o.Inner.Y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^ reference local 1
//               ^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Outer#Inner.
//                     ^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/Inner#Y.
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/embedded`/useOfCompositeStructs().
  
