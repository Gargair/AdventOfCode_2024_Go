package common

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Args struct {
	InputPath string
}

type Testcase struct {
	Name string
	Args Args
	Want string
}

func ReadTestData(problem string, part int) ([]Testcase, error) {
	testInputPath := filepath.Join("../../../testdata/", fmt.Sprintf("%s.txt", problem))
	lines, err := readAllLines(testInputPath, 2)

	if err != nil {
		return nil, err
	}
	tests := make([]Testcase, 0, 2)
	for _, line := range lines {
		res := strings.Fields(line)
		want := res[part-1]
		inputPath := res[2]
		tests = append(tests, Testcase{
			Name: fmt.Sprintf("inputPath=%v", inputPath),
			Args: Args{
				InputPath: filepath.Join("../../../", inputPath, fmt.Sprintf("%s.txt", problem))},
			Want: want})
	}
	return tests, nil
}

func readAllLines(filePath string, expectedLines int) (lines []string, err error) {
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
	return localLines, nil
}
