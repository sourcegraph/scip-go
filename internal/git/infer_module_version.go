package git

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"

	"github.com/sourcegraph/scip-go/internal/command"
)

// InferModuleVersion returns the version of the module declared in the given
// directory. This will be either the work tree commit's tag, or it will be the
// short revhash of the HEAD commit.
func InferModuleVersion(dir string) (string, error) {
	tags, err := command.Run(dir, "git", "tag", "-l", "--points-at", "HEAD")
	if err != nil {
		return "", fmt.Errorf("failed to tags for current commit: %v\n%s", err, tags)
	}
	if tags != "" {
		lines := strings.Split(tags, "\n")
		for _, tag := range lines {
			if semver.IsValid(tag) {
				return tag, nil
			}
		}

		// None of the tags look like a version, but return one of them anyway.
		return lines[0], nil
	}

	commit, err := command.Run(dir, "git", "rev-parse", "HEAD")
	if err != nil {
		return "", fmt.Errorf("failed to get current commit: %v\n%s", err, commit)
	}

	return commit[:12], nil
}
