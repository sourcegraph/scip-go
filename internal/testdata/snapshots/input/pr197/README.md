# PR #197: Import symbol definitions

Test cases for https://github.com/sourcegraph/scip-go/issues/34

Covers:
- Non-aliased imports (single-segment and multi-segment paths)
- Aliased imports
- Dot imports
- Blank/side-effect imports
- Multiple files in the same package (all `package` statements should be definitions)
