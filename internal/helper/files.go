package helper

import (
	"bufio"
	"os"
	"strings"
)

func ReadAllLines(filePath string, expectedLines int) (lines *[]string, err error) {
	localLines := make([]string, 0, expectedLines)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		localLines = append(localLines, strings.Trim(strings.TrimSpace(scanner.Text()), "\ufeff"))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &localLines, nil
}
