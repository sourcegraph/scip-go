package pr206

// Block doc for the const group.
const (
	// BlockConst1 is a multi-line doc.
	// It spans two lines.
	BlockConst1 = 1

	BlockConstNoDoc = 2

	BlockConstTrailing = 3 // trailing comment on const
)

const (
	// OrphanConst lives in a block with no block-level doc.
	OrphanConst = 99
)
