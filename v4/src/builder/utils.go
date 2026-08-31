package builder

import (
	"strings"
)

func hasPrefix(s string) bool {
	return strings.HasPrefix(s, "urn:") || strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

func containsSlash(s string) bool {
	return strings.Contains(s, "/")
}
