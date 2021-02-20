package internal_test

import (
	"os"
	"testing"

	. "scm/internal"
)

func TestLookupEnvOrDefault(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		expected := `foobar`
		_ = os.Setenv(`SCM_FOO_BAR`, expected)

		actual, _ := LookupEnvOrDefault(`SCM_FOO_BAR`, `foobar_default`)
		if expected != actual {
			t.Errorf(`want "%s", got "%s"`, expected, actual)
		}

		_ = os.Unsetenv(`SCM_FOO_BAR`)
	})

	t.Run(`negative`, func(t *testing.T) {
		expected := `foobarbaz`

		actual, _ := LookupEnvOrDefault(`SCM_FOO_BAR_BAZ`, expected)
		if expected != actual {
			t.Errorf(`want "%s", got "%s"`, expected, actual)
		}
	})
}
