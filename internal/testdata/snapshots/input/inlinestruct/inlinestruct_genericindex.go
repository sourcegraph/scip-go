package inlinestruct

type Processor[T any] interface {
	Process(payload T)
	ProcessorType() string
}

type Limit int

type ProcessImpl struct{}

func (p *ProcessImpl) Process(payload Limit) { panic("not implemented") }
func (p *ProcessImpl) ProcessorType() string { panic("not implemented") }

var _ Processor[Limit] = &ProcessImpl{}
