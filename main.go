package main

import (
	"log"
	"os"
	"scm/internal"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	scmBin, scmUrl, err := internal.ParseArgs(os.Args)
	if err != nil {
		return
	}

	cfg, err := internal.ReadCfg(scmUrl)
	if err != nil {
		return
	}

	err = internal.PrepareLocalWorkingCopyPath(cfg.ScmWorkingCopyPath, cfg.ScmWorkspaceDirDefaultPerm)
	if err != nil {
		return
	}

	err = internal.Clone(scmBin, scmUrl, cfg.ScmWorkingCopyPath)
	if err != nil {
		return
	}
}
