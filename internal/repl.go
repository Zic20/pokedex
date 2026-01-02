package internal

import (
	"strings"
)

func CleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	str := strings.Trim(strings.ToLower(text), " ")
	return strings.Split(str, " ")
}
