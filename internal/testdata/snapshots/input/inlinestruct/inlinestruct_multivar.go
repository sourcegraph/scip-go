package inlinestruct

type Params struct{}
type HighlightedCode struct{}

var Mocks, emptyMocks struct {
	Code func(p Params) (response *HighlightedCode, aborted bool, err error)
}

var MocksSingle struct {
	Code func(p Params) (response *HighlightedCode, aborted bool, err error)
}

var (
	okReply   interface{} = "OK"
	pongReply interface{} = "PONG"
)
