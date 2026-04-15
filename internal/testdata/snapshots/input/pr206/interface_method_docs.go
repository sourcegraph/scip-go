package pr206

type Doer interface {
	// Do performs the action and returns an error if it fails.
	Do() error
}
