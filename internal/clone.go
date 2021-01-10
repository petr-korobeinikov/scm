package internal

import (
	"os"
	"os/exec"
)

func PrepareLocalWorkingCopyPath(scmWorkingCopyPath string, scmWorkspaceDirDefaultPermFileMode os.FileMode) error {
	return os.MkdirAll(scmWorkingCopyPath, scmWorkspaceDirDefaultPermFileMode)
}

func Clone(scmBin, scmUrl, scmWorkingCopyPath string) error {
	scmCmd := exec.Command(scmBin, "clone", scmUrl, scmWorkingCopyPath)
	scmCmd.Stdout = os.Stdout
	scmCmd.Stderr = os.Stderr

	if err := scmCmd.Run(); err != nil {
		return err
	}

	return nil
}
