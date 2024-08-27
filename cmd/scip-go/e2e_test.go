package main

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/sourcegraph/scip-go/internal/command"
	"github.com/stretchr/testify/require"
)

var scipGoPath = ""

func TestMain(m *testing.M) {
	tempDir := os.TempDir()
	scipGoPath = path.Join(tempDir, "scip-go")
	_, err := command.Run(".", "go", "build", "-o", scipGoPath, ".")
	if err != nil {
		panic(fmt.Sprintf("failed to build scip-go: %s", err.Error()))
	}
	os.Exit(m.Run())
}

func TestAutoIndexFlagFetchesTags(t *testing.T) {
	tempDir := os.TempDir()
	concDir := path.Join(tempDir, "conc")
	t.Cleanup(func() { os.RemoveAll(concDir) })
	_, err := command.Run(".", "git", "clone", "--depth=1", "--no-tags", "https://github.com/sourcegraph/conc", concDir)
	require.NoError(t, err)
	tagsDir := path.Join(concDir, ".git", "refs", "tags")
	_, err = os.Stat(path.Join(tagsDir, "v0.3.0"))
	require.ErrorIs(t, err, os.ErrNotExist)

	outErr, err := command.Run(concDir, scipGoPath, "--auto-index")
	require.NoError(t, err, outErr)
	dirEntries, err := os.ReadDir(tagsDir)
	require.NoError(t, err)
	tagNames := map[string]struct{}{}
	for _, dirEntry := range dirEntries {
		tagNames[dirEntry.Name()] = struct{}{}
	}
	require.Contains(t, tagNames, "v0.1.0")
	require.Contains(t, tagNames, "v0.2.0")
	require.Contains(t, tagNames, "v0.3.0")
}
