package pr95

// Phase 2: lookup.Set conflict scenarios.
// When multiple fields share the same anonymous struct type, their nested
// fields map to the same token.Pos. The current lookup.Set silently
// overwrites, masking the conflict.

type ConflictingFields struct {
	first struct {
		shared int
	}
	second struct {
		shared int
	}
}

func useConflictingFields() {
	var c ConflictingFields
	c.first.shared = 1
	c.second.shared = 2
}

// Multi-name fields also trigger the same-pos conflict.
type MultiNameConflict struct {
	x, y struct {
		field int
	}
}

func useMultiNameConflict() {
	var m MultiNameConflict
	m.x.field = 10
	m.y.field = 20
}
