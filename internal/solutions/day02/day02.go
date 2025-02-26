package day02

import (
	"AdventOfCode/internal/common"
	"AdventOfCode/internal/intmath"
	"fmt"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type Day02_Solution struct{}

func (s Day02_Solution) SolvePart1(input interface{}) (string, error) {
	reports, ok := input.([][]int)
	if !ok {
		return "", fmt.Errorf("expected [][]int, got %T", input)
	}

	count := 0
	for _, report := range reports {
		if IsValidReportPart1(report) {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func (s Day02_Solution) SolvePart2(input interface{}) (string, error) {
	reports, ok := input.([][]int)
	if !ok {
		return "", fmt.Errorf("expected [][]int, got %T", input)
	}

	reporter := make(chan bool, len(reports))
	numWorkers := runtime.NumCPU()
	jobs := make(chan []int, len(reports))

	for w := 0; w < numWorkers; w++ {
		go func() {
			for job := range jobs {
				reporter <- IsValidReportPart2(job)
			}
		}()
	}

	for _, report := range reports {
		jobs <- report
	}
	close(jobs)

	count := 0
	for range reports {
		if <-reporter {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func IsValidReportPart1(report []int) bool {
	sign := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if intmath.Abs(diff) > 3 || diff == 0 {
			return false
		}
		if sign == 0 {
			if diff < 0 {
				sign = -1
			} else if diff > 0 {
				sign = 1
			}
		} else {
			if diff > 0 && sign < 0 || diff < 0 && sign > 0 {
				return false
			}
		}
	}
	return true
}

func IsValidReportPart2(report []int) bool {
	if IsValidReportPart1(report) {
		return true
	}
	for index := range report {
		if IsValidReportPart1(slices.Delete(slices.Clone(report), index, index+1)) {
			return true
		}
	}
	return false
}

func (s Day02_Solution) ParseInput(input string) (ret interface{}, err error) {
	lines := common.SplitLines(input)

	reports := make([][]int, 0, len(lines))

	for _, line := range lines {
		res := strings.Fields(line)
		report := make([]int, len(res))
		for i, v := range res {
			report[i], err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		}
		reports = append(reports, report)
	}

	return reports, nil
}
