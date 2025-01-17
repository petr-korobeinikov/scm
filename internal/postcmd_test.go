package internal_test

import (
	"testing"

	. "github.com/petr-korobeinikov/scm/internal"
)

func TestExecutePostCmd(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		err := ExecutePostCmd(NewPostCloneCmd("echo", []string{"foo", "bar", "baz"}))

		if err != nil {
			t.Error(err)
		}
	})

	t.Run(`negative`, func(t *testing.T) {
		err := ExecutePostCmd(NewPostCloneCmd("/bin/nonexistentcmd", nil))

		if err == nil {
			t.Error(`an error expected when calling not installed binary`)
		}
	})

	t.Run(`empty cmd do not crashes`, func(t *testing.T) {
		err := ExecutePostCmd(EmptyPostCloneCmd)

		if err != nil {
			t.Error(`empty cmd must not crash`)
		}
	})
}
