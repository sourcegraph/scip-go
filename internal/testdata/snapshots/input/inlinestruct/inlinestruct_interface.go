package inlinestruct

import "context"

func Target() interface {
	OID(context.Context) (int, error)
	AbbreviatedOID(context.Context) (string, error)
	Commit(context.Context) (string, error)
	Type(context.Context) (int, error)
} {
	panic("not implemented")
}

func something() {
	x := Target()
	x.OID(context.Background())
}
