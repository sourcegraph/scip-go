package pr95

// Identical anonymous struct types should share symbols for nested fields.

type IdenticalAnonFields struct {
	x struct{ t int }
	z struct{ t int }
}

func useIdenticalAnonFields() {
	var y IdenticalAnonFields
	y.x = y.z
	y.x.t = 1
	y.z.t = 2
}

// Different field order means different type — symbols must NOT unify.
type FieldOrderMatters struct {
	a struct{ x int; y string }
	b struct{ y string; x int }
}

// Different struct tags — symbols must NOT unify.
type DifferentTags struct {
	a struct{ Name string `json:"name"` }
	b struct{ Name string `json:"full_name"` }
}
