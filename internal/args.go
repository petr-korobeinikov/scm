package internal

import (
	"errors"
	"strings"
)

func ParseArgs(args []string) (scmBin, scmURL, scmPostCmd string, err error) {
	switch len(args) {
	default:
		err = ErrTooLongArgumentList
		return
	case 1:
		err = ErrNotEnoughArguments
		return
	case 2:
		scmBin = defaultScmBin
		scmURL = args[1]
		return
	case 3:
		if itLooksLikeURL(args[1]) {
			scmBin = defaultScmBin
			scmURL = args[1]
			scmPostCmd = args[2]
		} else {
			scmBin = args[1]
			scmURL = args[2]
		}
		return
	case 4:
		scmBin = args[1]
		scmURL = args[2]
		scmPostCmd = args[3]
		return
	}
}

func itLooksLikeURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

var (
	ErrNotEnoughArguments  = errors.New(`need at least one argument with repo url`)
	ErrTooLongArgumentList = errors.New(`too long argument list`)
)

const defaultScmBin = "git"
