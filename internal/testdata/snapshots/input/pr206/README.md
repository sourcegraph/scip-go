# pr206 --- Doc comments missing from hover documentation

Regression tests for [#108][]: hover docs do not contain doc comments for
several kinds of symbols.

The `initial/child_symbols.go` test already covers the basic patterns
(standalone const/var with doc, const/var blocks, struct fields, embedded
fields, interface methods, and grouped type declarations). The test cases here
focus on scenarios that are **not** covered there or that exercise distinct edge
cases:

## Bug 1: Constants and variables in grouped declarations

Per-item doc comments on `ValueSpec` nodes inside `const(…)` / `var(…)` blocks
are ignored. Instead, the block-level `GenDecl` doc is shown (or nothing at all
if the block has no doc comment).

Unique scenarios: multi-line per-item doc (`BlockConst2`), item with no doc
inside a block that *has* a block doc (`BlockConstNoDoc` / `BlockVarNoDoc`),
item with its own doc inside a block that has *no* block doc (`OrphanConst`).

## Bug 2: Struct field doc comments are lost

Named struct fields are registered with `parent = nil`, so `extractHoverText`
returns `""`.

Unique scenario: inline trailing comment (`Score`).

## Bug 3: Interface method doc comments are lost

Interface method names are registered with `parent = name` (the same
`*ast.Ident`), causing a self-referencing delegation that returns `""`.

## Bug 4: Types in grouped `type(…)` declarations get block-level doc

Similar to Bug 1 --- `SetNewSymbol` passes `v.curDecl` (`*ast.GenDecl`) as
parent for the `TypeSpec` name ident, skipping `TypeSpec.Doc`.

Unique scenario: type with no per-type doc inside a block with a block doc
(`GammaNoDoc`).

  [#108]: https://github.com/sourcegraph/scip-go/issues/108
