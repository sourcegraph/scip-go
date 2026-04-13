package pr95

// Phase 1: defer-in-loop bug — multi-name field declarations.
// The `defer v.scope.pop()` inside the for-loop over field names causes
// scopes to accumulate, producing incorrect nested symbols like #a.b.x
// instead of independent #a.x and #b.x.

type MultiNameStruct struct {
	a, b struct {
		x int
		y string
	}
}

type ThreeNameStruct struct {
	p, q, r struct {
		val int
	}
}

func useMultiNameFields() {
	var m MultiNameStruct
	m.a.x = 1
	m.a.y = "hello"
	m.b.x = 2
	m.b.y = "world"

	m.a = m.b

	var t ThreeNameStruct
	t.p.val = 1
	t.q.val = 2
	t.r.val = 3
}
