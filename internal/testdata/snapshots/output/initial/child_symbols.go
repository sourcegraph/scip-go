  package initial
//        ^^^^^^^ reference 0.1.test sg/initial/
  
  // Const is a constant equal to 5. It's the best constant I've ever written. 😹
  const Const = 5
//      ^^^^^ definition 0.1.test sg/initial/Const.
//      documentation ```go
//      documentation Const is a constant equal to 5. It's the best constant I've ever written. 😹
  
  // Docs for the const block itself.
  const (
   // ConstBlock1 is a constant in a block.
   ConstBlock1 = 1
// ^^^^^^^^^^^ definition 0.1.test sg/initial/ConstBlock1.
// documentation ```go
// documentation Docs for the const block itself.
  
   // ConstBlock2 is a constant in a block.
   ConstBlock2 = 2
// ^^^^^^^^^^^ definition 0.1.test sg/initial/ConstBlock2.
// documentation ```go
// documentation Docs for the const block itself.
  )
  
  // Var is a variable interface.
  var Var Interface = &Struct{Field: "bar!"}
//    ^^^ definition 0.1.test sg/initial/Var.
//    documentation ```go
//    documentation Var is a variable interface.
//        ^^^^^^^^^ reference 0.1.test sg/initial/Interface#
//                     ^^^^^^ reference 0.1.test sg/initial/Struct#
//                            ^^^^^ reference 0.1.test sg/initial/Struct#Field.
  
  // unexportedVar is an unexported variable interface.
  var unexportedVar Interface = &Struct{Field: "bar!"}
//    ^^^^^^^^^^^^^ definition 0.1.test sg/initial/unexportedVar.
//    documentation ```go
//    documentation unexportedVar is an unexported variable interface.
//                  ^^^^^^^^^ reference 0.1.test sg/initial/Interface#
//                               ^^^^^^ reference 0.1.test sg/initial/Struct#
//                                      ^^^^^ reference 0.1.test sg/initial/Struct#Field.
  
  // x has a builtin error type
  var x error
//    ^ definition 0.1.test sg/initial/x.
//    documentation ```go
//    documentation x has a builtin error type
  
  var BigVar Interface = &Struct{
//    ^^^^^^ definition 0.1.test sg/initial/BigVar.
//    documentation ```go
//           ^^^^^^^^^ reference 0.1.test sg/initial/Interface#
//                        ^^^^^^ reference 0.1.test sg/initial/Struct#
   Field: "bar!",
// ^^^^^ reference 0.1.test sg/initial/Struct#Field.
   Anonymous: struct {
// ^^^^^^^^^ reference 0.1.test sg/initial/Struct#Anonymous.
    FieldA int
//  ^^^^^^ definition 0.1.test sg/initial/BigVar:FieldA.
//  documentation ```go
    FieldB int
//  ^^^^^^ definition 0.1.test sg/initial/BigVar:FieldB.
//  documentation ```go
    FieldC int
//  ^^^^^^ definition 0.1.test sg/initial/BigVar:FieldC.
//  documentation ```go
   }{FieldA: 1337},
//   ^^^^^^ reference 0.1.test sg/initial/BigVar:FieldA.
  }
  
  // What are docs, really?
  // I can't say for sure, I don't write any.
  // But look, a CAT!
  //
  //       |\      _,,,---,,_
  // ZZZzz /,`.-'`'    -.  ;-;;,_
  //      |,4-  ) )-,_. ,\ (  `'-'
  //     '---''(_/--'  `-'\_)
  //
  // It's sleeping! Some people write that as `sleeping` but Markdown
  // isn't allowed in Go docstrings, right? right?!
  var (
   // This has some docs
   VarBlock1 = "if you're reading this"
// ^^^^^^^^^ definition 0.1.test sg/initial/VarBlock1.
// documentation ```go
// documentation What are docs, really?
  
   VarBlock2 = "hi"
// ^^^^^^^^^ definition 0.1.test sg/initial/VarBlock2.
// documentation ```go
// documentation What are docs, really?
  )
  
  // Embedded is a struct, to be embedded in another struct.
  type Embedded struct {
//     ^^^^^^^^ definition 0.1.test sg/initial/Embedded#
//     documentation ```go
//     documentation Embedded is a struct, to be embedded in another struct.
//     documentation ```go
   // EmbeddedField has some docs!
   EmbeddedField string
// ^^^^^^^^^^^^^ definition 0.1.test sg/initial/Embedded#EmbeddedField.
// documentation ```go
   Field         string // conflicts with parent "Field"
// ^^^^^ definition 0.1.test sg/initial/Embedded#Field.
// documentation ```go
  }
  
  type Struct struct {
//     ^^^^^^ definition 0.1.test sg/initial/Struct#
//     documentation ```go
//     documentation ```go
//     relationship 0.1.test sg/initial/Interface# implementation
   *Embedded
//  ^^^^^^^^ definition 0.1.test sg/initial/Struct#Embedded.
//  documentation ```go
//  ^^^^^^^^ reference 0.1.test sg/initial/Embedded#
   Field     string
// ^^^^^ definition 0.1.test sg/initial/Struct#Field.
// documentation ```go
   Anonymous struct {
// ^^^^^^^^^ definition 0.1.test sg/initial/Struct#Anonymous.
// documentation ```go
    FieldA int
//  ^^^^^^ definition 0.1.test sg/initial/Struct#Anonymous.FieldA.
//  documentation ```go
    FieldB int
//  ^^^^^^ definition 0.1.test sg/initial/Struct#Anonymous.FieldB.
//  documentation ```go
    FieldC int
//  ^^^^^^ definition 0.1.test sg/initial/Struct#Anonymous.FieldC.
//  documentation ```go
   }
  }
  
  // StructMethod has some docs!
  func (s *Struct) StructMethod() {}
//      ^ definition local 0
//         ^^^^^^ reference 0.1.test sg/initial/Struct#
//                 ^^^^^^^^^^^^ definition 0.1.test sg/initial/Struct#StructMethod().
//                 documentation ```go
//                 documentation StructMethod has some docs!
  
  func (s *Struct) ImplementsInterface() string { return "hi!" }
//      ^ definition local 1
//         ^^^^^^ reference 0.1.test sg/initial/Struct#
//                 ^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/initial/Struct#ImplementsInterface().
//                 documentation ```go
//                 relationship 0.1.test sg/initial/Interface#ImplementsInterface. implementation
  
  func (s *Struct) MachineLearning(
//      ^ definition local 2
//         ^^^^^^ reference 0.1.test sg/initial/Struct#
//                 ^^^^^^^^^^^^^^^ definition 0.1.test sg/initial/Struct#MachineLearning().
//                 documentation ```go
   param1 float32, // It's ML, I can't describe what this param is.
// ^^^^^^ definition local 3
  
   // We call the below hyperparameters because, uhh, well:
   //
   //    ,-.       _,---._ __  / \
   //   /  )    .-'       `./ /   \
   //   (  (   ,'            `/    /|
   //    \  `-"             \'\   / |
   //     `.              ,  \ \ /  |
   //   /`.          ,'-`----Y   |
   //     (            ;        |   '
   //     |  ,-.    ,-'         |  /
   //     |  | (   |        hjw | /
   //     )  |  \  `.___________|/
   //     `--'   `--'
   //
   hyperparam2 float32,
// ^^^^^^^^^^^ definition local 4
   hyperparam3 float32,
// ^^^^^^^^^^^ definition local 5
  ) float32 {
   // varShouldNotHaveDocs is in a function, should not have docs emitted.
   var varShouldNotHaveDocs int32
//     ^^^^^^^^^^^^^^^^^^^^ definition local 6
  
   // constShouldNotHaveDocs is in a function, should not have docs emitted.
   const constShouldNotHaveDocs = 5
//       ^^^^^^^^^^^^^^^^^^^^^^ definition local 7
  
   // typeShouldNotHaveDocs is in a function, should not have docs emitted.
   type typeShouldNotHaveDocs struct{ a string }
//      ^^^^^^^^^^^^^^^^^^^^^ definition local 8
//                                    ^ definition local 9
  
   // funcShouldNotHaveDocs is in a function, should not have docs emitted.
   funcShouldNotHaveDocs := func(a string) string { return "hello" }
// ^^^^^^^^^^^^^^^^^^^^^ definition local 10
//                               ^ definition local 11
  
   return param1 + (hyperparam2 * *hyperparam3) // lol is this all ML is? I'm gonna be rich
//        ^^^^^^ reference local 3
//                  ^^^^^^^^^^^ reference local 4
//                                 ^^^^^^^^^^^ reference local 5
  }
  
  // Interface has docs too
  type Interface interface {
//     ^^^^^^^^^ definition 0.1.test sg/initial/Interface#
//     documentation ```go
//     documentation Interface has docs too
//     documentation ```go
   ImplementsInterface() string
// ^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/initial/Interface#ImplementsInterface.
// documentation ```go
  }
  
  func NewInterface() Interface { return nil }
//     ^^^^^^^^^^^^ definition 0.1.test sg/initial/NewInterface().
//     documentation ```go
//                    ^^^^^^^^^ reference 0.1.test sg/initial/Interface#
  
  var SortExportedFirst = 1
//    ^^^^^^^^^^^^^^^^^ definition 0.1.test sg/initial/SortExportedFirst.
//    documentation ```go
  
  var sortUnexportedSecond = 2
//    ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/initial/sortUnexportedSecond.
//    documentation ```go
  
  var _sortUnderscoreLast = 3
//    ^^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/initial/_sortUnderscoreLast.
//    documentation ```go
  
  // Yeah this is some Go magic incantation which is common.
  //
  //  ,_     _
  //  |\\_,-~/
  //  / _  _ |    ,--.
  // (  @  @ )   / ,-'
  //  \  _T_/-._( (
  // /         `. \
  // |         _  \ |
  // \ \ ,  /      |
  //  || |-_\__   /
  // ((_/`(____,-'
  var _ = Interface(&Struct{})
//        ^^^^^^^^^ reference 0.1.test sg/initial/Interface#
//                   ^^^^^^ reference 0.1.test sg/initial/Struct#
  
  type _ = struct{}
  
  // crypto/tls/common_string.go uses this pattern..
  func _() {
  }
  
  // Go can be fun
  type (
   // And confusing
   X struct {
// ^ definition 0.1.test sg/initial/X#
// documentation ```go
// documentation Go can be fun
// documentation ```go
    bar string
//  ^^^ definition 0.1.test sg/initial/X#bar.
//  documentation ```go
   }
  
   Y struct {
// ^ definition 0.1.test sg/initial/Y#
// documentation ```go
// documentation Go can be fun
// documentation ```go
    baz float64
//  ^^^ definition 0.1.test sg/initial/Y#baz.
//  documentation ```go
   }
  )
  
