  package initial
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src fmt/
  
  type Inner struct {
//     ^^^^^ definition sg/initial/Inner#
   X int
// ^ definition sg/initial/Inner#X.
   Y int
// ^ definition sg/initial/Inner#Y.
   Z int
// ^ definition sg/initial/Inner#Z.
  }
  
  type Outer struct {
//     ^^^^^ definition sg/initial/Outer#
   Inner
// ^^^^^ definition sg/initial/Outer#Inner.
// ^^^^^ reference sg/initial/Inner#
   W int
// ^ definition sg/initial/Outer#W.
  }
  
  func useOfCompositeStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition sg/initial/useOfCompositeStructs().
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
//     ^^^^^^ reference github.com/golang/go/src fmt/Printf().
//                      ^ reference local 0
//                        ^ reference sg/initial/Inner#X.
   fmt.Println(o.Inner.Y)
//     ^^^^^^^ reference github.com/golang/go/src fmt/Println().
//             ^ reference local 0
//               ^^^^^ reference sg/initial/Outer#Inner.
//                     ^ reference sg/initial/Inner#Y.
  }
  
