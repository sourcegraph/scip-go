# fingerprint

This package is a verbatim copy of
[`golang.org/x/tools/gopls/internal/util/fingerprint`], which is not importable
because it lives under an `internal/` path.

It is used by the implementation-indexing code to produce canonical
method-signature strings (fingerprints) for the 64-bit bitmask filter that
quickly rejects non-matching type/interface pairs before calling the expensive
`types.Implements`.

## Keeping in sync

The copy is pinned to the `golang.org/x/tools` version declared in `go.mod`
(currently **v0.43.0**). Run the sync script whenever you bump that dependency:

``` bash
./internal/implementations/fingerprint/sync.sh
```

The script fetches the file at the matching tag and overwrites the local copy.
Any diff will show up in `git status`.

  [`golang.org/x/tools/gopls/internal/util/fingerprint`]: https://github.com/golang/tools/tree/master/gopls/internal/util/fingerprint
