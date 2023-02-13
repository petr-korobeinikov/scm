package internal_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	. "github.com/pkorobeinikov/scm/internal"
)

func TestReadCfg(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		saveScmWorkspaceDir, found := os.LookupEnv(`SCM_WORKSPACE_DIR`)

		expected := Cfg{
			ScmWorkspaceDirDefaultPerm: 0755,
			ScmWorkingCopyPath:         `/tmp/Workspace/github.com/user/repo`,
		}

		_ = os.Setenv(`SCM_WORKSPACE_DIR`, `/tmp/Workspace`)
		actual, _ := Configure(`https://github.com/user/repo`, ``)

		if expected.ScmWorkspaceDirDefaultPerm != actual.ScmWorkspaceDirDefaultPerm {
			t.Errorf(`want "%s", got "%s"`, expected.ScmWorkspaceDirDefaultPerm, actual.ScmWorkspaceDirDefaultPerm)
		}

		if expected.ScmWorkingCopyPath != actual.ScmWorkingCopyPath {
			t.Errorf(`want "%s", got "%s"`, expected.ScmWorkingCopyPath, actual.ScmWorkingCopyPath)
		}

		restoreEnvIfItWasFound(`SCM_WORKSPACE_DIR`, saveScmWorkspaceDir, found)
	})

	t.Run(`homedir not detected`, func(t *testing.T) {
		saveScmWorkspaceDir, foundScmWorkspaceDir := os.LookupEnv(`SCM_WORKSPACE_DIR`)
		saveHome, foundHome := os.LookupEnv(`HOME`)

		_ = os.Unsetenv(`HOME`)
		_ = os.Setenv(`SCM_WORKSPACE_DIR`, `~/Workspace`)
		_, err := Configure(`https://github.com/user/repo`, ``)
		if err == nil {
			t.Errorf(`homedir detected but shouldn't'`)
		}

		restoreEnvIfItWasFound(`HOME`, saveHome, foundHome)
		restoreEnvIfItWasFound(`SCM_WORKSPACE_DIR`, saveScmWorkspaceDir, foundScmWorkspaceDir)
	})

	t.Run(`invalid workspace dir perm`, func(t *testing.T) {
		_ = os.Setenv(`SCM_WORKSPACE_DIR_DEFAULT_PERM`, "invalid_file_mode")

		if _, err := Configure(`https://github.com/user/repo`, ``); err == nil {
			t.Error(`expected error while reading invalid file mode`)
		}

		_ = os.Unsetenv(`SCM_WORKSPACE_DIR_DEFAULT_PERM`)
	})

	t.Run(`mailformed repo url given`, func(t *testing.T) {
		_, err := Configure(`https://github % com/user/repo`, ``)
		if err == nil {
			t.Errorf(`repo url parsed but shouldn't'`)
		}
	})

	t.Run(`post clone cmd set`, func(t *testing.T) {
		saveScmWorkspaceDir, foundScmWorkspaceDir := os.LookupEnv(`SCM_WORKSPACE_DIR`)
		saveScmPostCloneCmdStr, foundScmPostCloneCmd := os.LookupEnv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_WORKSPACE_DIR`, `~/Workspace`)
		expected := Cfg{
			ScmPostCloneCmd: NewPostCloneCmd("idea", []string{filepath.Join(os.Getenv("HOME"), "Workspace", "github.com", "user", "repo")}),
		}

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, `idea {{.ScmWorkingCopyPath}}`)
		actual, _ := Configure(`https://github.com/user/repo`, ``)

		if !reflect.DeepEqual(expected.ScmPostCloneCmd, actual.ScmPostCloneCmd) {
			t.Errorf("want %#v, got %#v", expected.ScmPostCloneCmd, actual.ScmPostCloneCmd)
		}

		restoreEnvIfItWasFound(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr, foundScmPostCloneCmd)
		restoreEnvIfItWasFound(`SCM_WORKSPACE_DIR`, saveScmWorkspaceDir, foundScmWorkspaceDir)
	})

	t.Run(`unset post clone cmd`, func(t *testing.T) {
		saveScmPostCloneCmdStr, foundScmPostCloneCmd := os.LookupEnv(`SCM_POST_CLONE_CMD`)

		_ = os.Unsetenv(`SCM_POST_CLONE_CMD`)
		_, err := Configure(`https://github.com/user/repo`, ``)

		if err != nil {
			t.Error(`unset post clone cmd should not cause an error`)
		}

		restoreEnvIfItWasFound(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr, foundScmPostCloneCmd)
	})

	t.Run(`empty post clone cmd`, func(t *testing.T) {
		saveScmPostCloneCmdStr, foundScmPostCloneCmd := os.LookupEnv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, ``)
		_, err := Configure(`https://github.com/user/repo`, ``)

		if err != nil {
			t.Error(`empty post clone cmd should not cause an error`)
		}

		restoreEnvIfItWasFound(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr, foundScmPostCloneCmd)
	})

	t.Run(`incorrect post clone cmd option`, func(t *testing.T) {
		saveScmPostCloneCmdStr, foundScmPostCloneCmd := os.LookupEnv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, `editor {{.WrongTmplAttr}}`)
		_, err := Configure(`https://github.com/user/repo`, ``)

		if err == nil {
			t.Error(`expected template error`)
		}

		restoreEnvIfItWasFound(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr, foundScmPostCloneCmd)
	})

	t.Run(`template error in post clone cmd option`, func(t *testing.T) {
		saveScmPostCloneCmdStr, foundScmPostCloneCmd := os.LookupEnv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, `editor {{.WrongTmplAttr}`)
		_, err := Configure(`https://github.com/user/repo`, ``)

		if err == nil {
			t.Error(`expected template error`)
		}

		restoreEnvIfItWasFound(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr, foundScmPostCloneCmd)
	})

	t.Run(`override post clone cmd`, func(t *testing.T) {
		saveScmPostCloneCmdStr, foundScmPostCloneCmd := os.LookupEnv(`SCM_POST_CLONE_CMD`)

		_ = os.Setenv(`SCM_POST_CLONE_CMD`, ``)
		cfg, _ := Configure(`https://github.com/user/repo`, `-`)

		if !cfg.ScmPostCloneCmd.IsEmpty() {
			t.Errorf(`want empty post clone cmd, got %s`, cfg.ScmPostCloneCmd.Command())
		}

		restoreEnvIfItWasFound(`SCM_POST_CLONE_CMD`, saveScmPostCloneCmdStr, foundScmPostCloneCmd)
	})
}

func restoreEnvIfItWasFound(key, value string, restore bool) {
	if restore {
		_ = os.Setenv(key, value)
	}
}
