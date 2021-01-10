package internal

import (
	"errors"
)

func ParseArgs(args []string) (scmBin, scmUrl string, err error) {
	switch len(args) {
	default:
		err = TooLongArgumentListErr
		return
	case 1:
		err = NotEnoughArgumentsErr
		return
	case 2:
		scmBin = "git"
		scmUrl = args[1]
		return
	case 3:
		scmBin = args[1]
		scmUrl = args[2]
		return
	}
}

var (
	NotEnoughArgumentsErr  = errors.New(`need at least one argument with repo url`)
	TooLongArgumentListErr = errors.New(`too long argument list`)
)
