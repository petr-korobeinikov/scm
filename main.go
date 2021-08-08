package main

import (
	"log"
	"os"

	scm "scm/internal"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	scmBin, scmURL, scmOverridePostCloneCmd, err := scm.ParseArgs(os.Args)
	if err != nil {
		return
	}

	cfg, err := scm.Configure(scmURL, scmOverridePostCloneCmd)
	if err != nil {
		return
	}

	err = scm.PrepareLocalWorkingCopyPath(cfg.ScmWorkingCopyPath, cfg.ScmWorkspaceDirDefaultPerm)
	if err != nil {
		return
	}

	err = scm.Clone(scmBin, scmURL, cfg.ScmWorkingCopyPath)
	if err != nil {
		return
	}

	err = scm.ExecutePostCmd(cfg.ScmPostCloneCmd)
	if err != nil {
		return
	}
}
