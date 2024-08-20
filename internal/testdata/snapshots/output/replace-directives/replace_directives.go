  package replacers
//        ^^^^^^^^^ definition 0.1.test `sg/replace-directives`/
//        documentation
//        > package replacers
  
  import (
   "fmt"
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
  
   "github.com/dghubble/gologin"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference github.com/sourcegraph/gologin c6f1b62954d8 `github.com/dghubble/gologin`/
  )
  
  func Something() {
//     ^^^^^^^^^ definition 0.1.test `sg/replace-directives`/Something().
//     documentation
//     > ```go
//     > func Something()
//     > ```
   fmt.Println(gologin.DefaultCookieConfig)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^^^^^^^ reference github.com/sourcegraph/gologin c6f1b62954d8 `github.com/dghubble/gologin`/
//                     ^^^^^^^^^^^^^^^^^^^ reference github.com/sourcegraph/gologin c6f1b62954d8 `github.com/dghubble/gologin`/DefaultCookieConfig.
  }
  
