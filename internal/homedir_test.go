package internal

import (
	"os"
	"testing"
)

func TestExpandHomeDir(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		actual, _ := ExpandHomeDir(`~/Workspace`)
		// fixme hardcoded value
		if `/Users/pkorobeinikov/Workspace` != actual {
			t.Fail()
		}
	})

	t.Run(`no home dir in path`, func(t *testing.T) {
		actual, _ := ExpandHomeDir(`/mnt/Volumes/Workspace`)
		if `/mnt/Volumes/Workspace` != actual {
			t.Fail()
		}
	})

	t.Run(`homedir not set`, func(t *testing.T) {
		saveHome := os.Getenv(`HOME`)

		_ = os.Unsetenv(`HOME`)
		_, err := ExpandHomeDir(`~/Workspace`)
		if err == nil {
			t.Fail()
		}

		_ = os.Setenv(`HOME`, saveHome)
	})
}
