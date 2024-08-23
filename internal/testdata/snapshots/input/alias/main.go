package main

// Check that we don't panic
// Copied from https://github.com/golang/go/issues/68877#issuecomment-2290000187
type (
	T struct{}
	U = T
	S U
)
