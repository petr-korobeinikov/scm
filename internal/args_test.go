package internal

import (
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	t.Run(`complex`, func(t *testing.T) {
		testCases := []struct {
			name     string
			given    []string
			expected struct {
				scmBin string
				scmUrl string
				err    error
			}
		}{
			{
				name:  "not enough arguments",
				given: []string{"scm"},
				expected: struct {
					scmBin string
					scmUrl string
					err    error
				}{
					scmBin: "",
					scmUrl: "",
					err:    NotEnoughArgumentsErr,
				},
			},
			{
				name:  "too long argument list",
				given: []string{"foo", "bar", "baz", "quix"},
				expected: struct {
					scmBin string
					scmUrl string
					err    error
				}{
					scmBin: "",
					scmUrl: "",
					err:    TooLongArgumentListErr,
				},
			},
			{
				name:  "git by default",
				given: []string{"scm", "https://github.com/user/repo"},
				expected: struct {
					scmBin string
					scmUrl string
					err    error
				}{
					scmBin: "git",
					scmUrl: "https://github.com/user/repo",
					err:    nil,
				},
			},
			{
				name:  "hg if needed",
				given: []string{"scm", "hg", "http://hg.robustwebserver.org/robustwebserver/"},
				expected: struct {
					scmBin string
					scmUrl string
					err    error
				}{
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
					t.Fail()
				}

				if testCase.expected.scmUrl != actualScmUrl {
					t.Fail()
				}

				if testCase.expected.err != actualErr {
					t.Fail()
				}
			})
		}
	})
}
