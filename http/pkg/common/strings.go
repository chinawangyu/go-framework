package common

import "strings"

func IsValidString(source string) bool {
	return source != "" && len(strings.TrimSpace(source)) > 0
}
