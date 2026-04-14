  package inlinestruct
//        ^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/
  
  type Processor[T any] interface {
//     ^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#
//               signature_documentation
//               > type Processor interface {
//               >     Process(payload T)
//               >     ProcessorType() string
//               > }
//               ^ definition local 0
//                 display_name T
//                 signature_documentation
//                 > T T
   Process(payload T)
// ^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#Process.
//         signature_documentation
//         > func (Processor[T any]).Process(payload T)
//         ^^^^^^^ definition local 1
//                 display_name payload
//                 signature_documentation
//                 > var payload T
//                 ^ reference local 0
   ProcessorType() string
// ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/Processor#ProcessorType.
//               signature_documentation
//               > func (Processor[T any]).ProcessorType() string
  }
  
  type Limit int
//     ^^^^^ definition 0.1.test `sg/inlinestruct`/Limit#
//           signature_documentation
//           > type Limit int
  
  type ProcessImpl struct{}
//     ^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#
//                 signature_documentation
//                 > type ProcessImpl struct{}
  
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
  func (p *ProcessImpl) Process(payload Limit) { panic("not implemented") }
//      ^ definition local 2
//        display_name p
//        signature_documentation
//        > var p *sg/inlinestruct.ProcessImpl
//         ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
//                      ^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
//                              signature_documentation
//                              > func (*ProcessImpl).Process(payload Limit)
//                              ^^^^^^^ definition local 3
//                                      display_name payload
//                                      signature_documentation
//                                      > var payload sg/inlinestruct.Limit
//                                      ^^^^^ reference 0.1.test `sg/inlinestruct`/Limit#
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/ProcessImpl#Process().
//⌄ enclosing_range_start 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
  func (p *ProcessImpl) ProcessorType() string { panic("not implemented") }
//      ^ definition local 4
//        display_name p
//        signature_documentation
//        > var p *sg/inlinestruct.ProcessImpl
//         ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
//                      ^^^^^^^^^^^^^ definition 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
//                                    signature_documentation
//                                    > func (*ProcessImpl).ProcessorType() string
//                                                                        ⌃ enclosing_range_end 0.1.test `sg/inlinestruct`/ProcessImpl#ProcessorType().
  
  var _ Processor[Limit] = &ProcessImpl{}
//      ^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/Processor#
//                ^^^^^ reference 0.1.test `sg/inlinestruct`/Limit#
//                          ^^^^^^^^^^^ reference 0.1.test `sg/inlinestruct`/ProcessImpl#
  
