package git

import (
	"fmt"

	"github.com/sourcegraph/scip-go/internal/command"
)

// InferModuleVersion returns the version of the module declared in the given
// directory. This will be either the work tree commit's tag, or it will be the
// short revhash of the HEAD commit.
func InferModuleVersion(dir string) (string, error) {
	version, err := command.Run(dir, "git", "tag", "-l", "--points-at", "HEAD")
	if err != nil {
		return "", fmt.Errorf("failed to tags for current commit: %v\n%s", err, version)
	}
	if version != "" {
		return version, nil
	}

	commit, err := command.Run(dir, "git", "rev-parse", "HEAD")
	if err != nil {
		return "", fmt.Errorf("failed to get current commit: %v\n%s", err, commit)
	}

	return commit[:12], nil
}
