# pr199 --- Package definition and documentation

This test covers the package symbol handling introduced in PR #199:

- Every `package` declaration is emitted as a **definition**, not just one per
  package.
- `display_name` and `signature_documentation` are set for all package symbols.
- `documentation` collects doc comments from **all** files, sorted by relevance:
  `doc.go` \> exact package name match \> other files \> test files.
- Files without a doc comment (`no_doc.go`) do not contribute to
  `documentation`.
