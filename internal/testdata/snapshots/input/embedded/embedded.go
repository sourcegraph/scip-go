package embedded

import "os/exec"

type osExecCommand struct {
	*exec.Cmd
}

func wrapExecCommand(c *exec.Cmd) {
	_ = &osExecCommand{Cmd: c}
}
