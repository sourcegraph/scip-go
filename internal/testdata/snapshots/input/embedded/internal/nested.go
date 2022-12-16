package nested_internal

import (
	"fmt"
	"sg/embedded"
)

func Something(recent embedded.RecentCommittersResults) {
	for _, commit := range recent.Nodes {
		for _, author := range commit.Authors.Nodes {
			fmt.Println(author.Name)
		}
	}
}
