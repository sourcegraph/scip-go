  package gosrc
//        ^^^^^ reference sg/testdata/duplicate_path_id/
  
  type importMeta struct{}
//     ^^^^^^^^^^ definition sg/testdata/duplicate_path_id/importMeta#
//     documentation ```go
//     documentation ```go
  
  type sourceMeta struct{}
//     ^^^^^^^^^^ definition sg/testdata/duplicate_path_id/sourceMeta#
//     documentation ```go
//     documentation ```go
  
  func fetchMeta() (string, *importMeta, *sourceMeta) {
//     ^^^^^^^^^ definition sg/testdata/duplicate_path_id/fetchMeta().
//     documentation ```go
//                           ^^^^^^^^^^ reference sg/testdata/duplicate_path_id/importMeta#
//                                        ^^^^^^^^^^ reference sg/testdata/duplicate_path_id/sourceMeta#
   panic("hmm")
  }
  
  func init() {
//     ^^^^ definition sg/testdata/duplicate_path_id/init().
//     documentation ```go
  }
  
  // two inits in the same file is legal
  func init() {
//     ^^^^ definition sg/testdata/duplicate_path_id/init().
//     documentation ```go
  }
  
  // three inits in the same file is legal
  func init() {
//     ^^^^ definition sg/testdata/duplicate_path_id/init().
//     documentation ```go
  }
  
