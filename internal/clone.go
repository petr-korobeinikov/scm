package internal

import (
	"os"
	"os/exec"
)

func PrepareLocalWorkingCopyPath(scmWorkingCopyPath string, scmWorkspaceDirDefaultPermFileMode os.FileMode) error {
	return os.MkdirAll(scmWorkingCopyPath, scmWorkspaceDirDefaultPermFileMode)
}

func Clone(scmBin, scmURL, scmWorkingCopyPath string) error {
	scmCmd := exec.Command(scmBin, "clone", scmURL, scmWorkingCopyPath)
	scmCmd.Stdout = os.Stdout
	scmCmd.Stderr = os.Stderr

	return scmCmd.Run()
}
