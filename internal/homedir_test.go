package internal_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/petr-korobeinikov/scm/internal"
)

func TestExpandHomeDir(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		expected := filepath.Join(os.Getenv("HOME"), "Workspace")

		actual, _ := ExpandHomeDir(`~/Workspace`)
		if expected != actual {
			t.Errorf(`want "%s", got "%s"`, expected, actual)
		}
	})

	t.Run(`no home dir in path`, func(t *testing.T) {
		expected := `/mnt/Volumes/Workspace`

		actual, _ := ExpandHomeDir(expected)
		if expected != actual {
			t.Errorf(`want "%s", got "%s"`, expected, actual)
		}
	})

	t.Run(`homedir not set`, func(t *testing.T) {
		saveHome := os.Getenv(`HOME`)

		_ = os.Unsetenv(`HOME`)
		_, err := ExpandHomeDir(`~/Workspace`)
		if err == nil {
			t.Error(`an error expected while expanding homedir`)
		}

		_ = os.Setenv(`HOME`, saveHome)
	})
}
