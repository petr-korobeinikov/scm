package internal

import (
	"os"
	"os/exec"
)

func ExecutePostCmd(cmd *scmPostCloneCmd) error {
	if cmd.IsEmpty() {
		return nil
	}

	postCloneCmd := exec.Command(cmd.Command(), cmd.Arguments()...)
	postCloneCmd.Stdout = os.Stdout
	postCloneCmd.Stderr = os.Stderr

	return postCloneCmd.Run()
}
