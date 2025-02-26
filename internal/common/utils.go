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

type StringArray struct {
	arr []string
}

func (s *StringArray) Set(input string) error {
	s.arr = strings.Split(input, ",")
	return nil
}

func (s *StringArray) String() string {
	return strings.Join(s.arr, ",")
}

func (s *StringArray) GetValue() []string {
	return s.arr
}
