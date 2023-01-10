  package gosrc
//        ^^^^^ reference 0.1.test sg/testdata/duplicate_path_id/
  
  type importMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test sg/testdata/duplicate_path_id/importMeta#
//     documentation ```go
//     documentation ```go
  
  type sourceMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test sg/testdata/duplicate_path_id/sourceMeta#
//     documentation ```go
//     documentation ```go
  
  func fetchMeta() (string, *importMeta, *sourceMeta) {
//     ^^^^^^^^^ definition 0.1.test sg/testdata/duplicate_path_id/fetchMeta().
//     documentation ```go
//                           ^^^^^^^^^^ reference 0.1.test sg/testdata/duplicate_path_id/importMeta#
//                                        ^^^^^^^^^^ reference 0.1.test sg/testdata/duplicate_path_id/sourceMeta#
   panic("hmm")
  }
  
  func init() {}
//     ^^^^ definition 0.1.test sg/testdata/duplicate_path_id/init().
//     documentation ```go
  func init() {}
//     ^^^^ definition 0.1.test sg/testdata/duplicate_path_id/init().
//     documentation ```go
  func init() {}
//     ^^^^ definition 0.1.test sg/testdata/duplicate_path_id/init().
//     documentation ```go
  
