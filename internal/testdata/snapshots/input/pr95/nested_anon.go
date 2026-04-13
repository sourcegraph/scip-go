package pr95

// Phase 3: Nested anonymous structs and container types.

type DeepNested struct {
	outer struct {
		inner struct {
			value int
		}
	}
}

type SliceAnon struct {
	items []struct {
		id   int
		name string
	}
}

type MapAnon struct {
	entries map[string]struct {
		count int
		label string
	}
}

type PointerAnon struct {
	ptr *struct {
		data int
	}
}

// Two fields with identical slice-of-anonymous-struct type.
type SliceAnonShared struct {
	a []struct{ v int }
	b []struct{ v int }
}

func useNestedAnon() {
	var d DeepNested
	d.outer.inner.value = 42

	var s SliceAnon
	if len(s.items) > 0 {
		_ = s.items[0].id
		_ = s.items[0].name
	}

	var m MapAnon
	entry := m.entries["key"]
	_ = entry.count
	_ = entry.label

	var p PointerAnon
	if p.ptr != nil {
		_ = p.ptr.data
	}
}
