package internal

import (
	"os"
	"testing"
)

func TestLookupEnvOrDefault(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		_ = os.Setenv(`SCM_FOO_BAR`, `foobar`)

		actual := LookupEnvOrDefault(`SCM_FOO_BAR`, `foobar_default`)
		if `foobar` != actual {
			t.Fail()
		}

		_ = os.Unsetenv(`SCM_FOO_BAR`)
	})

	t.Run(`negative`, func(t *testing.T) {
		actual := LookupEnvOrDefault(`SCM_FOO_BAR_BAZ`, `foobarbaz`)
		if `foobarbaz` != actual {
			t.Fail()
		}
	})
}
