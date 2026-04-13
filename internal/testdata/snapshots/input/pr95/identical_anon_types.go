package pr95

// Phase 3: Identical anonymous struct types across different fields.
// Fields with the same anonymous struct type should share symbol names
// for their nested fields, enabling cross-field Find References.

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
	a struct {
		x int
		y string
	}
	b struct {
		y string
		x int
	}
}

func useFieldOrderMatters() {
	var f FieldOrderMatters
	f.a.x = 1
	f.a.y = "hello"
	f.b.y = "world"
	f.b.x = 2
}

// Different struct tags mean different types — symbols must NOT unify.
type DifferentTags struct {
	a struct {
		Name string `json:"name"`
	}
	b struct {
		Name string `json:"full_name"`
	}
}

// Exported vs unexported field names — different types.
type ExportedVsUnexported struct {
	a struct {
		X int
		Y string
	}
	b struct {
		x int
		y string
	}
}
