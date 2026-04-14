  // Package pr199 tests package definition and documentation handling.
  package pr199
//        ^^^^^ definition 0.1.test `sg/pr199`/
//              display_name pr199
//              signature_documentation
//              > package pr199
//              documentation
//              > Package pr199 tests package definition and documentation handling.
//              documentation
//              > Additional documentation from the main file.
//              documentation
//              > Documentation for no_doc.go.
//              documentation
//              > Test file documentation for pr199.
  
//⌄ enclosing_range_start 0.1.test `sg/pr199`/FromDoc().
  func FromDoc() {}
//     ^^^^^^^ definition 0.1.test `sg/pr199`/FromDoc().
//             signature_documentation
//             > func FromDoc()
//                ⌃ enclosing_range_end 0.1.test `sg/pr199`/FromDoc().
  
