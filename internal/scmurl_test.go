package internal_test

import (
	"fmt"
	"testing"

	. "github.com/petr-korobeinikov/scm/internal"
)

func TestExtractLocalPathFromScmURL(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		testCases := []struct {
			given    string
			expected string
		}{
			{
				given:    `https://github.com/user/repo`,
				expected: `github.com/user/repo`,
			},
			{
				given:    `http://hg.robustwebserver.org/robustwebserver/`,
				expected: `hg.robustwebserver.org/robustwebserver`,
			},
			{
				given:    `https://gitlab.com/foo/bar/baz/itool3`,
				expected: `gitlab.com/foo/bar/baz/itool3`,
			},
			{
				given:    `https://git.postgresql.org/git/postgresql.git`,
				expected: `git.postgresql.org/git/postgresql.git`,
			},
		}

		for testCaseNo, testCase := range testCases {
			t.Run(fmt.Sprintf(`positive %d`, testCaseNo), func(t *testing.T) {
				actual, _ := ExtractLocalPathFromScmURL(testCase.given)
				if testCase.expected != actual {
					t.Error(fmt.Sprintf(`want %s got %s`, testCase.expected, actual))
				}
			})
		}
	})

	t.Run(`negative`, func(t *testing.T) {
		_, err := ExtractLocalPathFromScmURL(`malformed % url`)
		if err == nil {
			t.Error(`error expected on malformed url`)
		}
	})
}
