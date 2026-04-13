package pr95

// Phase 3: Anonymous structs in local/function scope.
// Local anonymous structs should use local symbols.
// Two locals with the same anonymous type should still share field symbols.

func localAnonymousStructs() {
	a := struct {
		x int
		y string
	}{x: 1, y: "one"}

	b := struct {
		x int
		y string
	}{x: 2, y: "two"}

	_ = a.x
	_ = a.y
	_ = b.x
	_ = b.y

	// Different type (different fields).
	c := struct {
		z int
	}{z: 3}
	_ = c.z
}

func paramAnonymousStruct(p struct{ x int }) int {
	return p.x
}
