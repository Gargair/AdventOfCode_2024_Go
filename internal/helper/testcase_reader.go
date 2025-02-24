package helper

import (
	"fmt"
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
	lines, err := ReadAllLines(testInputPath, 2)

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
