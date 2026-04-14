# pr201 --- Import alias resolution

This test covers import alias handling introduced in PR #201, addressing [#34].

Import aliases are resolved as **global references** to the original package
symbol rather than local definitions. This means:

- `h` in `import h "net/http"` emits a reference to `net/http`, not a local
  definition.
- All subsequent usages of `h` (e.g., `h.StatusOK`) also reference `net/http`
  directly.
- A "Find references" lookup on the original `package http` statement returns
  every individual usage site, not just the import line.

This applies consistently to all aliased imports: same-name aliases, renamed
aliases, and aliases used across multiple files in the same package.

While this approach differs from gopls, which treats import aliases as local
definitions, SCIP is a protocol serving a greater number of languages. Resolving
aliases as global references is a more universal approach --- it matches how
most languages treat imports and ensures consistent cross-language behavior.

  [#34]: https://github.com/sourcegraph/scip-go/issues/34
