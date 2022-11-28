  package initial
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src fmt/
  
  type Inner struct {
//     ^^^^^ definition sg/initial/Inner#
//     documentation ```go
//     documentation ```go
   X int
// ^ definition sg/initial/Inner#X.
// documentation ```go
   Y int
// ^ definition sg/initial/Inner#Y.
// documentation ```go
   Z int
// ^ definition sg/initial/Inner#Z.
// documentation ```go
  }
  
  type Outer struct {
//     ^^^^^ definition sg/initial/Outer#
//     documentation ```go
//     documentation ```go
   Inner
// ^^^^^ definition sg/initial/Outer#Inner.
// documentation ```go
// ^^^^^ reference sg/initial/Inner#
   W int
// ^ definition sg/initial/Outer#W.
// documentation ```go
  }
  
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition sg/initial/useOfCompositeStructs().
//     documentation ```go
   o := Outer{
// ^ definition local 0
//      ^^^^^ reference sg/initial/Outer#
    Inner: Inner{
//  ^^^^^ reference sg/initial/Outer#Inner.
//         ^^^^^ reference sg/initial/Inner#
     X: 1,
//   ^ reference sg/initial/Inner#X.
     Y: 2,
//   ^ reference sg/initial/Inner#Y.
     Z: 3,
//   ^ reference sg/initial/Inner#Z.
    },
    W: 4,
//  ^ reference sg/initial/Outer#W.
   }
  
   fmt.Printf("> %d\n", o.X)
// ^^^ reference github.com/golang/go/src fmt/
//     ^^^^^^ reference github.com/golang/go/src fmt/Printf().
//                      ^ reference local 0
//                        ^ reference sg/initial/Inner#X.
   fmt.Println(o.Inner.Y)
// ^^^ reference github.com/golang/go/src fmt/
//     ^^^^^^^ reference github.com/golang/go/src fmt/Println().
//             ^ reference local 0
//               ^^^^^ reference sg/initial/Outer#Inner.
//                     ^ reference sg/initial/Inner#Y.
  }
  
