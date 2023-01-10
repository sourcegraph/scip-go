  package replacers
//        ^^^^^^^^^ definition sg/replace-directives/
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src fmt/
  
   "github.com/dghubble/gologin"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference github.com/sourcegraph/gologin github.com/dghubble/gologin/
  )
  
  func Something() {
//     ^^^^^^^^^ definition sg/replace-directives/Something().
//     documentation ```go
   fmt.Println(gologin.DefaultCookieConfig)
// ^^^ reference github.com/golang/go/src fmt/
//     ^^^^^^^ reference github.com/golang/go/src fmt/Println().
//             ^^^^^^^ reference github.com/sourcegraph/gologin github.com/dghubble/gologin/
//                     ^^^^^^^^^^^^^^^^^^^ reference github.com/sourcegraph/gologin DefaultCookieConfig.
  }
  
