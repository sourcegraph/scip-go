package git

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestInferRepo(t *testing.T) {
	if err := exec.Command("git", "rev-parse", "--git-dir").Run(); err != nil {
		t.Skip("not inside a git repository")
	}

	repo, err := InferRepo("")
	if err != nil {
		t.Fatalf("unexpected error inferring repo: %s", err)
	}

	if repo != "github.com/sourcegraph/scip-go" {
		t.Errorf("unexpected remote repo. want=%q have=%q", "github.com/sourcegraph/scip-go", repo)
	}
}

func TestParseRemote(t *testing.T) {
	testCases := map[string]string{
		"git@github.com:sourcegraph/scip-go.git":                                "github.com/sourcegraph/scip-go",
		"https://github.com/sourcegraph/scip-go":                                "github.com/sourcegraph/scip-go",
		"ssh://git@phabricator.company.com:2222/diffusion/COMPANY/companay.git": "phabricator.company.com/diffusion/COMPANY/companay",
	}

	for input, expectedOutput := range testCases {
		t.Run(fmt.Sprintf("input=%q", input), func(t *testing.T) {
			output, err := parseRemote(input)
			if err != nil {
				t.Fatalf("unexpected error parsing remote: %s", err)
			}

			if output != expectedOutput {
				t.Errorf("unexpected repo name. want=%q have=%q", expectedOutput, output)
			}
		})
	}
}
