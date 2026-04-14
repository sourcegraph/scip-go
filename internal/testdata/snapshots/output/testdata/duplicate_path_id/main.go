  package gosrc
//        ^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/
//              display_name gosrc
//              signature_documentation
//              > package gosrc
  
  type importMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/importMeta#
//                signature_documentation
//                > type importMeta struct{}
  
  type sourceMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/sourceMeta#
//                signature_documentation
//                > type sourceMeta struct{}
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/fetchMeta().
  func fetchMeta() (string, *importMeta, *sourceMeta) {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/fetchMeta().
//               signature_documentation
//               > func fetchMeta() (string, *importMeta, *sourceMeta)
//                           ^^^^^^^^^^ reference 0.1.test `sg/testdata/duplicate_path_id`/importMeta#
//                                        ^^^^^^^^^^ reference 0.1.test `sg/testdata/duplicate_path_id`/sourceMeta#
   panic("hmm")
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/fetchMeta().
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/init().
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//          signature_documentation
//          > func init()
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/init().
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/init().
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//          signature_documentation
//          > func init()
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/init().
//⌄ enclosing_range_start 0.1.test `sg/testdata/duplicate_path_id`/init().
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//          signature_documentation
//          > func init()
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/duplicate_path_id`/init().
  
