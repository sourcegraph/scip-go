package inlinestruct

type InFuncSig struct {
	value bool
}

var rowsCloseHook = func() func(InFuncSig, *error) { return nil }
