  package replacers
//        ^^^^^^^^^ definition 0.1.test sg/replace-directives/
  
  import (
   "fmt"
//  ^^^ reference v1.19 fmt/
  
   "github.com/dghubble/gologin"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference v1.0.2-0.20181110030308-c6f1b62954d8 github.com/dghubble/gologin/
  )
  
  func Something() {
//     ^^^^^^^^^ definition 0.1.test sg/replace-directives/Something().
//     documentation ```go
   fmt.Println(gologin.DefaultCookieConfig)
// ^^^ reference v1.19 fmt/
//     ^^^^^^^ reference v1.19 fmt/Println().
//             ^^^^^^^ reference v1.0.2-0.20181110030308-c6f1b62954d8 github.com/dghubble/gologin/
//                     ^^^^^^^^^^^^^^^^^^^ reference v1.0.2-0.20181110030308-c6f1b62954d8 DefaultCookieConfig.
  }
  
