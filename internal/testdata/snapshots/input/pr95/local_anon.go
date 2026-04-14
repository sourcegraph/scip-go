package pr95

// Anonymous structs in local/function scope.

func localAnonymousStructs() {
	a := struct{ x int }{x: 1}
	b := struct{ x int }{x: 2}
	_ = a.x
	_ = b.x
}

func paramAnonymousStruct(p struct{ x int }) int {
	return p.x
}
