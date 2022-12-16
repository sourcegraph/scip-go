package embedded

import "fmt"

type RecentCommittersResults struct {
	Nodes []struct {
		Authors struct {
			Nodes []struct {
				Date  string
				Email string
				Name  string
				User  struct {
					Login string
				}
				AvatarURL string
			}
		}
	}
	PageInfo struct {
		HasNextPage bool
	}
}
