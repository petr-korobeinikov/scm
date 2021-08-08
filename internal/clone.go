package internal

import (
	"os"
	"os/exec"
)

// PrepareLocalWorkingCopyPath creates full working copy path.
func PrepareLocalWorkingCopyPath(scmWorkingCopyPath string, scmWorkspaceDirDefaultPermFileMode os.FileMode) error {
	return os.MkdirAll(scmWorkingCopyPath, scmWorkspaceDirDefaultPermFileMode)
}

// Clone clones remote repository.
func Clone(scmBin, scmURL, scmWorkingCopyPath string) error {
	scmCmd := exec.Command(scmBin, "clone", scmURL, scmWorkingCopyPath)
	scmCmd.Stdout = os.Stdout
	scmCmd.Stderr = os.Stderr

	return scmCmd.Run()
}
