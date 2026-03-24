package client

import (
	"fmt"
	"net/url"
	"strings"
)

// assemble and parse URL, returns valid url
func assembleUrl(parts ...string) (string, error) {
	raw := strings.Join(parts, "")

	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("could not format url. raw: %s", raw)
	}

	u.RawQuery = u.Query().Encode()

	return u.String(), nil
}
