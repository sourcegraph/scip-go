package testdata

import "fmt"

type TypeContainingAnonymousStructs struct {
	a, b struct {
		x int
		y string
	}
	c struct {
		X int
		Y string
	}
}

func funcContainingAnonymousStructs() {
	d := struct {
		x int
		y string
	}{
		x: 1,
		y: "one",
	}

	var e struct {
		x int
		y string
	}

	e.x = 2
	e.y = "two"

	var f TypeContainingAnonymousStructs
	f.a.x = 3
	f.a.y = "three"
	f.b.x = 4
	f.b.y = "four"
	f.c.X = 5
	f.c.Y = "five"

	fmt.Printf("> %s, %s\n", d.x, d.y)
	fmt.Printf("> %s, %s\n", e.x, e.y)

	fmt.Printf("> %s, %s\n", f.a.x, f.a.y)
	fmt.Printf("> %s, %s\n", f.b.x, f.b.y)
	fmt.Printf("> %s, %s\n", f.c.X, f.c.Y)
}
