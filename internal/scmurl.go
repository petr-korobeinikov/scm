package internal

import (
	"net/url"
	"path"
)

func ExtractLocalPathFromScmURL(scmUrl string) (string, error) {
	u, err := url.Parse(scmUrl)
	if err != nil {
		return "", err
	}

	return path.Join(u.Host, u.Path), nil
}
