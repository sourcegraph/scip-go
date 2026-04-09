package pr198

import "github.com/example/dep"

// Foo is an interface defined downstream of the type that implements it.
// The dep.T type (from a dependency) implements Foo, and scip-go should
// emit an external symbol for dep.T with an IsImplementation relationship
// pointing to Foo.
type Foo interface {
	Bar()
}

func UseFoo(f Foo) {}

func Example() {
	UseFoo(&dep.T{})
}
