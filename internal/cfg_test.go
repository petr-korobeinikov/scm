package internal

import (
	"os"
	"reflect"
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

	t.Run(`post clone cmd set`, func(t *testing.T) {
		saveScmWorkspaceDir := os.Getenv(`SCM_WORKSPACE_DIR`)
		saveScmPostCloneCmdStr := os.Getenv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_WORKSPACE_DIR`, `~/Workspace`)
		expected := Cfg{
			ScmPostCloneCmd: ScmPostCloneCmd{
				Cmd:  "idea",
				Args: []string{"/Users/pkorobeinikov/Workspace/github.com/user/repo"},
			},
		}

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, `idea {{.ScmWorkingCopyPath}}`)
		actual, _ := ReadCfg(`https://github.com/user/repo`)

		if !reflect.DeepEqual(expected.ScmPostCloneCmd, actual.ScmPostCloneCmd) {
			t.Errorf("want %#v, got %#v", expected.ScmPostCloneCmd, actual.ScmPostCloneCmd)
		}

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr)
		_ = os.Setenv(`SCM_WORKSPACE_DIR`, saveScmWorkspaceDir)
	})

	t.Run(`empty post clone cmd`, func(t *testing.T) {
		saveScmPostCloneCmdStr := os.Getenv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, ``)
		_, err := ReadCfg(`https://github.com/user/repo`)

		if err != nil {
			t.Error(`empty post clone cmd should not cause an error`)
		}

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr)
	})

	t.Run(`incorrect post clone cmd option`, func(t *testing.T) {
		saveScmPostCloneCmdStr := os.Getenv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, `editor {{.WrongTmplAttr}}`)
		_, err := ReadCfg(`https://github.com/user/repo`)

		if err == nil {
			t.Error(`expected template error`)
		}

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr)
	})

	t.Run(`template error in post clone cmd option`, func(t *testing.T) {
		_ = os.Setenv(`SCM_POST_CLONE_CMD`, `editor {{.WrongTmplAttr}`)
		_, err := ReadCfg(`https://github.com/user/repo`)

		if err == nil {
			t.Error(`expected template error`)
		}
	})
}
