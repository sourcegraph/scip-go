  package gosrc
//        ^^^^^ reference 0.1.test `sg/testdata/duplicate_path_id`/
  
  type importMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/importMeta#
//     documentation
//     > ```go
//     > type importMeta struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
  type sourceMeta struct{}
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/sourceMeta#
//     documentation
//     > ```go
//     > type sourceMeta struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
  func fetchMeta() (string, *importMeta, *sourceMeta) {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/fetchMeta().
//     documentation
//     > ```go
//     > func fetchMeta() (string, *importMeta, *sourceMeta)
//     > ```
//                           ^^^^^^^^^^ reference 0.1.test `sg/testdata/duplicate_path_id`/importMeta#
//                                        ^^^^^^^^^^ reference 0.1.test `sg/testdata/duplicate_path_id`/sourceMeta#
   panic("hmm")
  }
  
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//     documentation
//     > ```go
//     > func init()
//     > ```
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//     documentation
//     > ```go
//     > func init()
//     > ```
  func init() {}
//     ^^^^ definition 0.1.test `sg/testdata/duplicate_path_id`/init().
//     documentation
//     > ```go
//     > func init()
//     > ```
  
