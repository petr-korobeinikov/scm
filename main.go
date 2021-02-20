package main

import (
	"log"
	"os"

	. "scm/internal"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	scmBin, scmUrl, scmOverridePostCloneCmd, err := ParseArgs(os.Args)
	if err != nil {
		return
	}

	cfg, err := Configure(scmUrl, scmOverridePostCloneCmd)
	if err != nil {
		return
	}

	err = PrepareLocalWorkingCopyPath(cfg.ScmWorkingCopyPath, cfg.ScmWorkspaceDirDefaultPerm)
	if err != nil {
		return
	}

	err = Clone(scmBin, scmUrl, cfg.ScmWorkingCopyPath)
	if err != nil {
		return
	}

	err = ExecutePostCmd(cfg)
	if err != nil {
		return
	}
}
