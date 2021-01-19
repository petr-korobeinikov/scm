package internal

import (
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	t.Run(`complex`, func(t *testing.T) {
		type (
			expected struct {
				scmBin string
				scmUrl string
				err    error
			}

			testCase struct {
				name     string
				given    []string
				expected expected
			}
		)

		testCases := []testCase{
			{
				name:  "not enough arguments",
				given: []string{"scm"},
				expected: expected{
					scmBin: "",
					scmUrl: "",
					err:    NotEnoughArgumentsErr,
				},
			},
			{
				name:  "too long argument list",
				given: []string{"foo", "bar", "baz", "quix"},
				expected: expected{
					scmBin: "",
					scmUrl: "",
					err:    TooLongArgumentListErr,
				},
			},
			{
				name:  "git by default",
				given: []string{"scm", "https://github.com/user/repo"},
				expected: expected{
					scmBin: "git",
					scmUrl: "https://github.com/user/repo",
					err:    nil,
				},
			},
			{
				name:  "hg if needed",
				given: []string{"scm", "hg", "http://hg.robustwebserver.org/robustwebserver/"},
				expected: expected{
					scmBin: "hg",
					scmUrl: "http://hg.robustwebserver.org/robustwebserver/",
					err:    nil,
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%s", testCase.name), func(t *testing.T) {
				actualScmBin, actualScmUrl, actualErr := ParseArgs(testCase.given)

				if testCase.expected.scmBin != actualScmBin {
					t.Errorf(`want "%s", got "%s"`, testCase.expected.scmBin, actualScmBin)
				}

				if testCase.expected.scmUrl != actualScmUrl {
					t.Errorf(`want "%s", got "%s"`, testCase.expected.scmUrl, actualScmUrl)
				}

				if testCase.expected.err != actualErr {
					t.Errorf(`want "%s", got "%s"`, testCase.expected.err, actualErr)
				}
			})
		}
	})
}
