package server

import "testing"

func TestStuff(t *testing.T) {
	wd := "hello"
	repo := "world"

	runCmd(t, wd, "git", "init", repo)
}
