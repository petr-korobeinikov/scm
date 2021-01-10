package internal

import (
	"os"
	"testing"
)

func TestReadCfg(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		saveScmWorkspaceDir := os.Getenv(`SCM_WORKSPACE_DIR`)

		expected := Cfg{
			ScmWorkspaceDirDefaultPerm: 0755,
			ScmWorkingCopyPath:         `/tmp/Workspace/github.com/user/repo`,
		}

		_ = os.Setenv(`SCM_WORKSPACE_DIR`, `/tmp/Workspace`)
		actual, _ := ReadCfg(`https://github.com/user/repo`)

		if expected.ScmWorkspaceDirDefaultPerm != actual.ScmWorkspaceDirDefaultPerm {
			t.Errorf(`want "%s", got "%s"`, expected.ScmWorkspaceDirDefaultPerm, actual.ScmWorkspaceDirDefaultPerm)
		}

		if expected.ScmWorkingCopyPath != actual.ScmWorkingCopyPath {
			t.Errorf(`want "%s", got "%s"`, expected.ScmWorkingCopyPath, actual.ScmWorkingCopyPath)
		}

		_ = os.Setenv(`SCM_WORKSPACE_DIR`, saveScmWorkspaceDir)
	})

	t.Run(`homedir not detected`, func(t *testing.T) {
		saveScmWorkspaceDir := os.Getenv(`SCM_WORKSPACE_DIR`)
		saveHome := os.Getenv(`HOME`)

		_ = os.Unsetenv(`HOME`)
		_ = os.Setenv(`SCM_WORKSPACE_DIR`, `~/Workspace`)
		_, err := ReadCfg(`https://github.com/user/repo`)
		if err == nil {
			t.Errorf(`homedir detected but shouldn't'`)
		}

		_ = os.Setenv(`HOME`, saveHome)
		_ = os.Setenv(`SCM_WORKSPACE_DIR`, saveScmWorkspaceDir)
	})

	t.Run(`mailformed repo url given`, func(t *testing.T) {
		_, err := ReadCfg(`https://github % com/user/repo`)
		if err == nil {
			t.Errorf(`repo url parsed but shouldn't'`)
		}
	})
}
