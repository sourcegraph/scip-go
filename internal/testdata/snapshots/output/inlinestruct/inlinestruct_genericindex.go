  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/
  
  type Processor[T any] interface {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#
//     documentation
//     > ```go
//     > type Processor interface
//     > ```
//     documentation
//     > ```go
//     > interface {
//     >     Process(payload T)
//     >     ProcessorType() string
//     > }
//     > ```
//               ^ definition local 0
   Process(payload T)
// ^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#Process.
// documentation
// > ```go
// > func (Processor[T any]).Process(payload T)
// > ```
//         ^^^^^^^ definition local 1
//                 ^ reference local 0
   ProcessorType() string
// ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#ProcessorType.
// documentation
// > ```go
// > func (Processor[T any]).ProcessorType() string
// > ```
  }
  
  type Limit int
//     ^^^^^ definition 0.1.test `sg/inlinestruct`/Limit#
//     documentation
//     > ```go
//     > int
//     > ```
  
  type ProcessImpl struct{}
//     ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#
//     documentation
//     > ```go
//     > type ProcessImpl struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
  func (p *ProcessImpl) Process(payload Limit) { panic("not implemented") }
//      ^ definition local 2
//         ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
//                      ^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
//                      documentation
//                      > ```go
//                      > func (*ProcessImpl).Process(payload Limit)
//                      > ```
//                              ^^^^^^^ definition local 3
//                                      ^^^^^ reference 0.1.test `sg/inlinestruct`/Limit#
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
  func (p *ProcessImpl) ProcessorType() string { panic("not implemented") }
//      ^ definition local 4
//         ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
//                      ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
//                      documentation
//                      > ```go
//                      > func (*ProcessImpl).ProcessorType() string
//                      > ```
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
  
  var _ Processor[Limit] = &ProcessImpl{}
//      ^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/Processor#
//                ^^^^^ reference 0.1.test `sg/inlinestruct`/Limit#
//                          ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
  
