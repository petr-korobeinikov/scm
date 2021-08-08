package internal

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandHomeDir expands home directory path.
func ExpandHomeDir(d string) (string, error) {
	if strings.HasPrefix(d, `~/`) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		return filepath.Join(homeDir, d[2:]), nil
	}

	return d, nil
}
