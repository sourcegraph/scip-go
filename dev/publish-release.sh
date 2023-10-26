#!/usr/bin/env bash

set -euo pipefail

{
if [ -z "${NEW_VERSION:-}" ]; then
  echo "error: Missing value for environment variable NEW_VERSION"
  echo "hint: Invoke this script as NEW_VERSION=M.N.P ./dev/publish-release.sh"
  exit 1
fi

if ! grep -q "## v$NEW_VERSION" CHANGELOG.md; then
  echo "error: Missing CHANGELOG entry for v$NEW_VERSION"
  echo "note: CHANGELOG entries are required for publishing releases"
  exit 1
fi

VERSION_FILE_PATH="internal/index/version.txt"
if ! grep -q "$NEW_VERSION" "$VERSION_FILE_PATH"; then
  echo "error: scip-go version in $VERSION_FILE_PATH doesn't match NEW_VERSION=$NEW_VERSION"
  exit 1
fi

if ! git diff --quiet; then
  echo "error: Found unstaged changes; aborting."
  exit 1
fi

if ! git diff --quiet --cached; then
  echo "error: Found staged-but-uncommitted changes; aborting."
  exit 1
fi

if ! git rev-parse --abbrev-ref HEAD | grep -q "main"; then
  echo "error: Releases should be published from main but HEAD is on a different branch" >&2
  exit 1
fi
} >&2

TAG="v$NEW_VERSION"
git tag "$TAG"
git push origin "$TAG"
