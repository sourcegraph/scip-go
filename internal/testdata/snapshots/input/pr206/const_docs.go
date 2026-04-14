package pr206

// Block doc for the const group.
const (
	// BlockConst1 is the first constant in a block.
	BlockConst1 = 1

	// BlockConst2 is a multi-line doc.
	// It spans two lines.
	BlockConst2 = 2

	BlockConstNoDoc = 3

	BlockConstTrailing = 5 // trailing comment on const
)

const (
	// OrphanConst lives in a block with no block-level doc.
	OrphanConst = 99
)
