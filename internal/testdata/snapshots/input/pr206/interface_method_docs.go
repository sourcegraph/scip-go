package pr206

// Doer performs actions.
type Doer interface {
	// Do performs the action and returns an error if it fails.
	Do() error

	// Reset clears internal state.
	Reset()
}
