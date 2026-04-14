package pr202

// Multi-name field declarations: a, b share a type and must be siblings.

type MultiNameStruct struct {
	a, b struct {
		x int
		y string
	}
}

func useMultiNameFields() {
	var m MultiNameStruct
	m.a.x = 1
	m.b.x = 2
	m.a = m.b
}
