  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
  // Const is a constant equal to 5. It's the best constant I've ever written. 😹
  const Const = 5
//      ^^^^^ definition 0.1.test `sg/initial`/Const.
//      kind Constant
//      documentation
//      > ```go
//      > const Const untyped int = 5
//      > ```
//      documentation
//      > Const is a constant equal to 5. It's the best constant I've ever written. 😹
  
  // Docs for the const block itself.
  const (
   // ConstBlock1 is a constant in a block.
   ConstBlock1 = 1
// ^^^^^^^^^^^ definition 0.1.test `sg/initial`/ConstBlock1.
// kind Constant
// documentation
// > ```go
// > const ConstBlock1 untyped int = 1
// > ```
// documentation
// > Docs for the const block itself.
  
   // ConstBlock2 is a constant in a block.
   ConstBlock2 = 2
// ^^^^^^^^^^^ definition 0.1.test `sg/initial`/ConstBlock2.
// kind Constant
// documentation
// > ```go
// > const ConstBlock2 untyped int = 2
// > ```
// documentation
// > Docs for the const block itself.
  )
  
  // Var is a variable interface.
  var Var Interface = &Struct{Field: "bar!"}
//    ^^^ definition 0.1.test `sg/initial`/Var.
//    kind Variable
//    documentation
//    > ```go
//    > var Var Interface
//    > ```
//    documentation
//    > Var is a variable interface.
//        ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                     ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                            ^^^^^ reference 0.1.test `sg/initial`/Struct#Field.
  
  // unexportedVar is an unexported variable interface.
  var unexportedVar Interface = &Struct{Field: "bar!"}
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/unexportedVar.
//    kind Variable
//    documentation
//    > ```go
//    > var unexportedVar Interface
//    > ```
//    documentation
//    > unexportedVar is an unexported variable interface.
//                  ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                               ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                                      ^^^^^ reference 0.1.test `sg/initial`/Struct#Field.
  
  // x has a builtin error type
  var x error
//    ^ definition 0.1.test `sg/initial`/x.
//    kind Variable
//    documentation
//    > ```go
//    > var x error
//    > ```
//    documentation
//    > x has a builtin error type
  
  var BigVar Interface = &Struct{
//    ^^^^^^ definition 0.1.test `sg/initial`/BigVar.
//    kind Variable
//    documentation
//    > ```go
//    > var BigVar Interface
//    > ```
//           ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                        ^^^^^^ reference 0.1.test `sg/initial`/Struct#
   Field: "bar!",
// ^^^^^ reference 0.1.test `sg/initial`/Struct#Field.
   Anonymous: struct {
// ^^^^^^^^^ reference 0.1.test `sg/initial`/Struct#Anonymous.
    FieldA int
//  ^^^^^^ definition 0.1.test `sg/initial`/BigVar:FieldA.
//  kind Field
//  documentation
//  > ```go
//  > struct field FieldA int
//  > ```
    FieldB int
//  ^^^^^^ definition 0.1.test `sg/initial`/BigVar:FieldB.
//  kind Field
//  documentation
//  > ```go
//  > struct field FieldB int
//  > ```
    FieldC int
//  ^^^^^^ definition 0.1.test `sg/initial`/BigVar:FieldC.
//  kind Field
//  documentation
//  > ```go
//  > struct field FieldC int
//  > ```
   }{FieldA: 1337},
//   ^^^^^^ reference 0.1.test `sg/initial`/BigVar:FieldA.
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
// ^^^^^^^^^ definition 0.1.test `sg/initial`/VarBlock1.
// kind Variable
// documentation
// > ```go
// > var VarBlock1 string
// > ```
// documentation
// > What are docs, really?
// > I can't say for sure, I don't write any.
// > But look, a CAT!
// > 
// > 	      |\      _,,,---,,_
// > 	ZZZzz /,`.-'`'    -.  ;-;;,_
// > 	     |,4-  ) )-,_. ,\ (  `'-'
// > 	    '---''(_/--'  `-'\_)
// > 
// > It's sleeping! Some people write that as `sleeping` but Markdown
// > isn't allowed in Go docstrings, right? right?!
  
   VarBlock2 = "hi"
// ^^^^^^^^^ definition 0.1.test `sg/initial`/VarBlock2.
// kind Variable
// documentation
// > ```go
// > var VarBlock2 string
// > ```
// documentation
// > What are docs, really?
// > I can't say for sure, I don't write any.
// > But look, a CAT!
// > 
// > 	      |\      _,,,---,,_
// > 	ZZZzz /,`.-'`'    -.  ;-;;,_
// > 	     |,4-  ) )-,_. ,\ (  `'-'
// > 	    '---''(_/--'  `-'\_)
// > 
// > It's sleeping! Some people write that as `sleeping` but Markdown
// > isn't allowed in Go docstrings, right? right?!
  )
  
  // Embedded is a struct, to be embedded in another struct.
  type Embedded struct {
//     ^^^^^^^^ definition 0.1.test `sg/initial`/Embedded#
//     kind Class
//     documentation
//     > ```go
//     > type Embedded struct
//     > ```
//     documentation
//     > Embedded is a struct, to be embedded in another struct.
//     documentation
//     > ```go
//     > struct {
//     >     EmbeddedField string
//     >     Field string
//     > }
//     > ```
   // EmbeddedField has some docs!
   EmbeddedField string
// ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Embedded#EmbeddedField.
// kind Field
// documentation
// > ```go
// > struct field EmbeddedField string
// > ```
   Field         string // conflicts with parent "Field"
// ^^^^^ definition 0.1.test `sg/initial`/Embedded#Field.
// kind Field
// documentation
// > ```go
// > struct field Field string
// > ```
  }
  
  type Struct struct {
//     ^^^^^^ definition 0.1.test `sg/initial`/Struct#
//     kind Class
//     documentation
//     > ```go
//     > type Struct struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     *Embedded
//     >     Field string
//     >     Anonymous struct {
//     >         FieldA int
//     >         FieldB int
//     >         FieldC int
//     >     }
//     > }
//     > ```
//     relationship 0.1.test `sg/initial`/Interface# implementation
   *Embedded
//  ^^^^^^^^ definition 0.1.test `sg/initial`/Struct#Embedded.
//  kind Field
//  documentation
//  > ```go
//  > struct field Embedded *sg/initial.Embedded
//  > ```
//  ^^^^^^^^ reference 0.1.test `sg/initial`/Embedded#
   Field     string
// ^^^^^ definition 0.1.test `sg/initial`/Struct#Field.
// kind Field
// documentation
// > ```go
// > struct field Field string
// > ```
   Anonymous struct {
// ^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#Anonymous.
// kind Field
// documentation
// > ```go
// > struct field Anonymous struct{FieldA int; FieldB int; FieldC int}
// > ```
    FieldA int
//  ^^^^^^ definition 0.1.test `sg/initial`/Struct#Anonymous.FieldA.
//  kind Field
//  documentation
//  > ```go
//  > struct field FieldA int
//  > ```
    FieldB int
//  ^^^^^^ definition 0.1.test `sg/initial`/Struct#Anonymous.FieldB.
//  kind Field
//  documentation
//  > ```go
//  > struct field FieldB int
//  > ```
    FieldC int
//  ^^^^^^ definition 0.1.test `sg/initial`/Struct#Anonymous.FieldC.
//  kind Field
//  documentation
//  > ```go
//  > struct field FieldC int
//  > ```
   }
  }
  
  // StructMethod has some docs!
//⌄ enclosing_range_start 0.1.test `sg/initial`/Struct#StructMethod().
  func (s *Struct) StructMethod() {}
//      ^ definition local 0
//      kind Variable
//         ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                 ^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#StructMethod().
//                 kind Method
//                 documentation
//                 > ```go
//                 > func (*Struct).StructMethod()
//                 > ```
//                 documentation
//                 > StructMethod has some docs!
//                                 ⌃ enclosing_range_end 0.1.test `sg/initial`/Struct#StructMethod().
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/Struct#ImplementsInterface().
  func (s *Struct) ImplementsInterface() string { return "hi!" }
//      ^ definition local 1
//      kind Variable
//         ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                 ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#ImplementsInterface().
//                 kind Method
//                 documentation
//                 > ```go
//                 > func (*Struct).ImplementsInterface() string
//                 > ```
//                 relationship 0.1.test `sg/initial`/Interface#ImplementsInterface. implementation
//                                                             ⌃ enclosing_range_end 0.1.test `sg/initial`/Struct#ImplementsInterface().
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/Struct#MachineLearning().
  func (s *Struct) MachineLearning(
//      ^ definition local 2
//      kind Variable
//         ^^^^^^ reference 0.1.test `sg/initial`/Struct#
//                 ^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Struct#MachineLearning().
//                 kind Method
//                 documentation
//                 > ```go
//                 > func (*Struct).MachineLearning(param1 float32, hyperparam2 float32, hyperparam3 float32) float32
//                 > ```
   param1 float32, // It's ML, I can't describe what this param is.
// ^^^^^^ definition local 3
// kind Variable
  
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
// kind Variable
   hyperparam3 float32,
// ^^^^^^^^^^^ definition local 5
// kind Variable
  ) float32 {
   // varShouldNotHaveDocs is in a function, should not have docs emitted.
   var varShouldNotHaveDocs int32
//     ^^^^^^^^^^^^^^^^^^^^ definition local 6
//     kind Variable
  
   // constShouldNotHaveDocs is in a function, should not have docs emitted.
   const constShouldNotHaveDocs = 5
//       ^^^^^^^^^^^^^^^^^^^^^^ definition local 7
//       kind Constant
  
   // typeShouldNotHaveDocs is in a function, should not have docs emitted.
   type typeShouldNotHaveDocs struct{ a string }
//      ^^^^^^^^^^^^^^^^^^^^^ definition local 8
//      kind Class
//                                    ^ definition local 9
//                                    kind Field
  
   // funcShouldNotHaveDocs is in a function, should not have docs emitted.
   funcShouldNotHaveDocs := func(a string) string { return "hello" }
// ^^^^^^^^^^^^^^^^^^^^^ definition local 10
// kind Variable
//                               ^ definition local 11
//                               kind Variable
  
   return param1 + (hyperparam2 * *hyperparam3) // lol is this all ML is? I'm gonna be rich
//        ^^^^^^ reference local 3
//                  ^^^^^^^^^^^ reference local 4
//                                 ^^^^^^^^^^^ reference local 5
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/Struct#MachineLearning().
  
  // Interface has docs too
  type Interface interface {
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/Interface#
//     kind Interface
//     documentation
//     > ```go
//     > type Interface interface
//     > ```
//     documentation
//     > Interface has docs too
//     documentation
//     > ```go
//     > interface {
//     >     ImplementsInterface() string
//     > }
//     > ```
   ImplementsInterface() string
// ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/Interface#ImplementsInterface.
// kind Method
// documentation
// > ```go
// > func (Interface).ImplementsInterface() string
// > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/NewInterface().
  func NewInterface() Interface { return nil }
//     ^^^^^^^^^^^^ definition 0.1.test `sg/initial`/NewInterface().
//     kind Function
//     documentation
//     > ```go
//     > func NewInterface() Interface
//     > ```
//                    ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                                           ⌃ enclosing_range_end 0.1.test `sg/initial`/NewInterface().
  
  var SortExportedFirst = 1
//    ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/SortExportedFirst.
//    kind Variable
//    documentation
//    > ```go
//    > var SortExportedFirst int
//    > ```
  
  var sortUnexportedSecond = 2
//    ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/sortUnexportedSecond.
//    kind Variable
//    documentation
//    > ```go
//    > var sortUnexportedSecond int
//    > ```
  
  var _sortUnderscoreLast = 3
//    ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/_sortUnderscoreLast.
//    kind Variable
//    documentation
//    > ```go
//    > var _sortUnderscoreLast int
//    > ```
  
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
//        ^^^^^^^^^ reference 0.1.test `sg/initial`/Interface#
//                   ^^^^^^ reference 0.1.test `sg/initial`/Struct#
  
  type _ = struct{}
  
  // crypto/tls/common_string.go uses this pattern..
  func _() {
  }
  
  // Go can be fun
  type (
   // And confusing
   X struct {
// ^ definition 0.1.test `sg/initial`/X#
// kind Class
// documentation
// > ```go
// > type X struct
// > ```
// documentation
// > Go can be fun
// documentation
// > ```go
// > struct {
// >     bar string
// > }
// > ```
    bar string
//  ^^^ definition 0.1.test `sg/initial`/X#bar.
//  kind Field
//  documentation
//  > ```go
//  > struct field bar string
//  > ```
   }
  
   Y struct {
// ^ definition 0.1.test `sg/initial`/Y#
// kind Class
// documentation
// > ```go
// > type Y struct
// > ```
// documentation
// > Go can be fun
// documentation
// > ```go
// > struct {
// >     baz float64
// > }
// > ```
    baz float64
//  ^^^ definition 0.1.test `sg/initial`/Y#baz.
//  kind Field
//  documentation
//  > ```go
//  > struct field baz float64
//  > ```
   }
  )
  
