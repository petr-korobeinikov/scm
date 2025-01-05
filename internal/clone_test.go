package internal_test

import (
	"os"
	"testing"

	. "github.com/petr-korobeinikov/scm/internal"
)

func TestPrepareLocalWorkingCopyPath(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		err := PrepareLocalWorkingCopyPath(`/tmp/Workspace/foobar`, 0755)
		if err != nil {
			t.Error(`expected to create workspace directory without any errors`)
		}
	})
}

func TestClone(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		dest := `/tmp/github/gitignore`

		err := Clone(`git`, `https://github.com/github/gitignore`, dest)
		if err != nil {
			t.Error(`expected to clone repo without any errors`)
		}

		_ = os.RemoveAll(dest)
	})

	t.Run(`negative`, func(t *testing.T) {
		err := Clone(`unknown-scm`, `https://github.com/github/gitignore`, `/tmp/github/gitignore`)
		if err == nil {
			t.Error(`expected to fail with unknown command`)
		}
	})
}
