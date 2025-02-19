package helper

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution interface {
	Part1(inputPath string, expectedLines int) string
	Part2(inputPath string, expectedLines int) string
}

type Args struct {
	InputPath     string
	ExpectedLines int
}

type Testcase struct {
	Name string
	Args Args
	Want string
}

func ReadTestData(testInputPath string, part int) ([]Testcase, error) {
	lines, err := ReadAllLines(testInputPath, 2)

	if err != nil {
		return nil, err
	}
	tests := make([]Testcase, 0, 2)
	for _, line := range lines {
		res := strings.Fields(line)
		want := res[part-1]
		inputPath := res[2]
		expectedLines, err := strconv.Atoi(res[3])
		if err != nil {
			return nil, err
		}
		tests = append(tests, Testcase{
			Name: fmt.Sprintf("inputPath=%v/expectedLines=%v", inputPath, expectedLines),
			Args: Args{
				InputPath:     "../../../" + inputPath,
				ExpectedLines: expectedLines},
			Want: want})
	}
	return tests, nil
}
