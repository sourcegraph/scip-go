package pr95

// Anonymous structs inside container types and nested structs.

type ContainerAnon struct {
	items   []struct{ id int }
	entries map[string]struct{ count int }
	ptr     *struct{ data int }
}

type DeepNested struct {
	outer struct {
		inner struct {
			value int
		}
	}
}

// Two fields with identical slice-of-anonymous-struct type.
type SliceAnonShared struct {
	a []struct{ v int }
	b []struct{ v int }
}

func useContainerAnon() {
	var c ContainerAnon
	if len(c.items) > 0 {
		_ = c.items[0].id
	}
	entry := c.entries["key"]
	_ = entry.count
	if c.ptr != nil {
		_ = c.ptr.data
	}

	var d DeepNested
	d.outer.inner.value = 42
}
