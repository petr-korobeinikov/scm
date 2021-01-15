package internal

import (
	"os"
	"os/exec"
)

func ExecutePostCmd(cfg Cfg) error {
	if cfg.ScmPostCloneCmd.Cmd == "" {
		return nil
	}

	postCloneCmd := exec.Command(cfg.ScmPostCloneCmd.Cmd, cfg.ScmPostCloneCmd.Args...)
	postCloneCmd.Stdout = os.Stdout
	postCloneCmd.Stderr = os.Stderr

	return postCloneCmd.Run()
}
