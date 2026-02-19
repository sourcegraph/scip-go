package nested_internal

import (
	"fmt"
	"github.com/sourcegraph/scip-go/internal/testdata/embedded"
)

func Something(recent embedded.RecentCommittersResults) {
	for _, commit := range recent.Nodes {
		for _, author := range commit.Authors.Nodes {
			fmt.Println(author.Name)
		}
	}
}
