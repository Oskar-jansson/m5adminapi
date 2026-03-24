package options

import (
	"strings"
)

// internal use only.
// Inputs url and options, returns formatted url with provided options
func QueryParamMerge(url string, v []string) string {
	if len(v) == 0 {
		return url
	}

	if len(v) > 1 {
		panic("options.QueryParamMerge() accepts at most one options string")
	}

	opt := v[0]
	if opt == "" {
		return url
	}

	if strings.Contains(url, "?") {
		if strings.HasPrefix(opt, "?") {
			opt = "&" + strings.TrimPrefix(opt, "?")
		} else {
			opt = "&" + opt
		}
	} else if !strings.HasPrefix(opt, "?") {
		opt = "?" + opt
	}

	return url + opt
}
