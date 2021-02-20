package internal

import (
	"errors"
	"strings"
)

func ParseArgs(args []string) (scmBin, scmUrl, scmPostCmd string, err error) {
	switch len(args) {
	default:
		err = TooLongArgumentListErr
		return
	case 1:
		err = NotEnoughArgumentsErr
		return
	case 2:
		scmBin = defaultScmBin
		scmUrl = args[1]
		return
	case 3:
		if itLooksLikeUrl(args[1]) {
			scmBin = defaultScmBin
			scmUrl = args[1]
			scmPostCmd = args[2]
		} else {
			scmBin = args[1]
			scmUrl = args[2]
		}
		return
	case 4:
		scmBin = args[1]
		scmUrl = args[2]
		scmPostCmd = args[3]
		return
	}
}

func itLooksLikeUrl(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

var (
	NotEnoughArgumentsErr  = errors.New(`need at least one argument with repo url`)
	TooLongArgumentListErr = errors.New(`too long argument list`)
)

const defaultScmBin = "git"
