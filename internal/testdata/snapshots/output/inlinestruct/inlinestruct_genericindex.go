  package inlinestruct
//        ^^^^^^^^^^^^ reference sg/inlinestruct/
  
  type Processor[T any] interface {
//     ^^^^^^^^^ definition sg/inlinestruct/Processor#
//     documentation ```go
//     documentation ```go
//               ^ definition local 0
   Process(payload T)
// ^^^^^^^ definition sg/inlinestruct/Processor#Process.
// documentation ```go
//         ^^^^^^^ definition local 1
//                 ^ reference local 0
   ProcessorType() string
// ^^^^^^^^^^^^^ definition sg/inlinestruct/Processor#ProcessorType.
// documentation ```go
  }
  
  type Limit int
//     ^^^^^ definition sg/inlinestruct/Limit#
//     documentation ```go
  
  type ProcessImpl struct{}
//     ^^^^^^^^^^^ definition sg/inlinestruct/ProcessImpl#
//     documentation ```go
//     documentation ```go
  
  func (p *ProcessImpl) Process(payload Limit) { panic("not implemented") }
//      ^ definition local 2
//         ^^^^^^^^^^^ reference sg/inlinestruct/ProcessImpl#
//                      ^^^^^^^ definition sg/inlinestruct/ProcessImpl#Process().
//                      documentation ```go
//                              ^^^^^^^ definition local 3
//                                      ^^^^^ reference sg/inlinestruct/Limit#
  func (p *ProcessImpl) ProcessorType() string { panic("not implemented") }
//      ^ definition local 4
//         ^^^^^^^^^^^ reference sg/inlinestruct/ProcessImpl#
//                      ^^^^^^^^^^^^^ definition sg/inlinestruct/ProcessImpl#ProcessorType().
//                      documentation ```go
  
  var _ Processor[Limit] = &ProcessImpl{}
//      ^^^^^^^^^ reference sg/inlinestruct/Processor#
//                ^^^^^ reference sg/inlinestruct/Limit#
//                          ^^^^^^^^^^^ reference sg/inlinestruct/ProcessImpl#
  
