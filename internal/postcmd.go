package internal

import (
	"os"
	"os/exec"
)

// todo embrace LoD: cfg -> postCloneCmd
func ExecutePostCmd(cfg Cfg) error {
	if cfg.ScmPostCloneCmd.IsEmpty() {
		return nil
	}

	postCloneCmd := exec.Command(cfg.ScmPostCloneCmd.Command(), cfg.ScmPostCloneCmd.Arguments()...)
	postCloneCmd.Stdout = os.Stdout
	postCloneCmd.Stderr = os.Stderr

	return postCloneCmd.Run()
}
