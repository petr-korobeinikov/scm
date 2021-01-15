package internal

import "testing"

func TestExecutePostCmd(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		err := ExecutePostCmd(Cfg{
			ScmPostCloneCmd: ScmPostCloneCmd{
				Cmd:  "echo",
				Args: []string{"foo", "bar", "baz"},
			},
		})

		if err != nil {
			t.Error(err)
		}
	})

	t.Run(`negative`, func(t *testing.T) {
		err := ExecutePostCmd(Cfg{
			ScmPostCloneCmd: ScmPostCloneCmd{
				Cmd:  "/bin/nonexistentcmd",
				Args: nil,
			},
		})

		if err == nil {
			t.Fail()
		}
	})

	t.Run(`empty cmd do not crashes`, func(t *testing.T) {
		err := ExecutePostCmd(Cfg{
			ScmPostCloneCmd: ScmPostCloneCmd{
				Cmd:  "",
				Args: nil,
			},
		})

		if err != nil {
			t.Fail()
		}
	})
}
