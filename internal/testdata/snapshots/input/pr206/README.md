# pr206 — Doc comments missing from hover documentation

Regression tests for [#108](https://github.com/scip-code/scip-go/issues/108):
hover docs do not contain doc comments for several kinds of symbols.

The `initial/child_symbols.go` test already covers the basic patterns (standalone
const/var with doc, const/var blocks with per-item docs, struct fields with doc
and inline comments, embedded fields, interface methods without docs, and grouped
type declarations).  The test cases here cover only **unique edge cases** not
present there:

- `const_docs.go`: multi-line per-item doc (`BlockConst1`), item with no doc
  falling back to block doc (`BlockConstNoDoc`), trailing comment on a const
  (`BlockConstTrailing`), item with doc in a block without block-level doc
  (`OrphanConst`).
- `embedded_field_docs.go`: doc comment on an embedded struct field.
- `interface_method_docs.go`: doc comment on an interface method.
