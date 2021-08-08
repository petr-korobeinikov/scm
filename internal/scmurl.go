package internal

import (
	"net/url"
	"path"
)

// ExtractLocalPathFromScmURL determines local working copy path.
func ExtractLocalPathFromScmURL(scmURL string) (string, error) {
	u, err := url.Parse(scmURL)
	if err != nil {
		return "", err
	}

	return path.Join(u.Host, u.Path), nil
}
