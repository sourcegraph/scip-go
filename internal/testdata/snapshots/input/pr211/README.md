# pr211 --- Generic type aliases

This test covers generic type alias support, addressing [#112].

Go 1.24 introduced generic type aliases ([go#46477]), allowing type aliases to
carry their own type parameters:

```go
type Set[K comparable] = Map[K, bool]
```

The indexer must correctly:

- Emit definitions for the alias and its type parameters.
- Emit references from the alias's RHS to the aliased type and its type
  arguments.
- Render hover documentation that includes the type parameter list.

  [#112]: https://github.com/scip-code/scip-go/issues/112
  [go#46477]: https://github.com/golang/go/issues/46477
