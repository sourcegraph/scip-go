package embedded

import "fmt"

func (r *RecentCommittersResults) String() string {
	return fmt.Sprintf("RecentCommittersResults{Nodes: %d}", len(r.Nodes))
}

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
