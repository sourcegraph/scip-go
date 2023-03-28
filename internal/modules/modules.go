package modules

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sourcegraph/scip-go/internal/command"
	"github.com/sourcegraph/scip-go/internal/output"
	"golang.org/x/tools/go/vcs"
)

func ModuleName(dir, repo, inName string) (moduleName string, isStdLib bool, err error) {
	resolve := func() error {
		if inName != "" {
			moduleName = inName
			return nil
		}

		moduleName = repo

		if !isModule(dir) {
			log.Println("WARNING: No go.mod file found in current directory.")
		} else {
			if moduleName, err = command.Run(dir, "go", "list", "-mod=readonly", "-m"); err != nil {
				return fmt.Errorf("failed to list modules: %v\n%s", err, moduleName)
			}

			return nil
		}

		moduleName, isStdLib, err = resolveModuleName(repo, moduleName)

		return nil
	}

	err = output.WithProgress("Resolving module name", resolve)
	if err != nil {
		return "", false, err
	}

	if moduleName == "std" {
		isStdLib = true
	}

	return moduleName, isStdLib, err
}

// resolveModuleName converts the given repository and import path into a canonical
// representation of a module name usable for moniker identifiers. The base of the
// import path will be the resolved repository remote, and the given module name
// is used only to determine the path suffix.
func resolveModuleName(repo, name string) (string, bool, error) {
	// Determine path suffix relative to repository root
	var suffix string

	if nameRepoRoot, err := vcs.RepoRootForImportPath(name, false); err == nil {
		suffix = strings.TrimPrefix(name, nameRepoRoot.Root)
	} else {
		// A user-visible warning will occur on this path as the declared
		// module will be resolved as part of gomod.ListDependencies.
	}

	// Determine the canonical code host of the current repository
	repoRepoRoot, err := vcs.RepoRootForImportPath(repo, false)
	if err != nil {
		help := "Make sure your git repo has a remote (git remote add origin git@github.com:owner/repo)"
		return "", false, fmt.Errorf("%v\n\n%s", err, help)
	}

	name = repoRepoRoot.Root + suffix
	return name, name == "std", nil
}

// isModule returns true if there is a go.mod file in the given directory.
func isModule(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "go.mod"))
	return err == nil
}
