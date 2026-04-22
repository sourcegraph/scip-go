#!/usr/bin/env bash
# Sync fingerprint.go from the upstream golang.org/x/tools repository.
#
# The version tag is read from go.mod so the copy stays pinned to the
# same golang.org/x/tools release that scip-go depends on.
set -euo pipefail

VERSION=$(go list -m -f '{{.Version}}' golang.org/x/tools)
SRC="https://raw.githubusercontent.com/golang/tools/${VERSION}/gopls/internal/util/fingerprint/fingerprint.go"
DST="$(dirname "$0")/fingerprint.go"

echo "Fetching fingerprint.go at golang.org/x/tools ${VERSION}..."
curl -fsSL "$SRC" -o "$DST"
echo "Updated $DST"
