package git

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/sourcegraph/scip-go/internal/command"
)

// FetchTagsIfMissing attempts to fetch all tags for a git repository,
// since we rely on git tags (similar to other Go tooling) to determine
// the module version number (see infer_module_version.go).
//
// Since this can change on-disk repository state, this operation should
// only be invoked during auto-indexing.
func FetchTagsIfMissing(dir string) error {
	tagsDir := path.Join(dir, ".git", "refs", "tags")
	stat, err := os.Stat(tagsDir)
	dirExists := err == nil && stat.IsDir()
	if dirExists {
		tags, err := os.ReadDir(tagsDir)
		if err != nil {
			return err
		}
		if len(tags) > 0 {
			return nil // tags already exist, don't fetch more
		}
	}
	// In an auto-indexing context, the remote is called origin.
	// See https://sourcegraph.sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/cmd/executor/internal/worker/workspace/clone.go?L31:1-31:15
	out, err := command.Run(dir, "git", "fetch", "origin", "refs/tags/*:refs/tags/*")
	if strings.Count(out, "new tag") == 0 {
		out, err := command.Run(dir, "git", "rev-parse", "--abbrev-ref", "HEAD")
		commitInfo := ""
		if err == nil {
			commitInfo = fmt.Sprintf(" (commit: %s)", strings.TrimSpace(out))
		}
		log.Warn("Found 0 tags, cross-repo navigation may not work correctly%s", commitInfo)
	}
	return err
}
