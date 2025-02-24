package common

import (
	"strings"
)

// SplitLines splits the input into lines, trimming whitespace
func SplitLines(input string) []string {
	lines := strings.Split(input, "\n")
	var result []string
	for _, line := range lines {
		trimmed := strings.Trim(strings.TrimSpace(line), "\ufeff")
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
