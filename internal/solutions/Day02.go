package solutions

import (
	"AdventOfCode/internal/helper"
	"AdventOfCode/internal/intmath"
	"slices"
	"strconv"
	"strings"
)

type Day02_Solution struct{}

func (s Day02_Solution) Part1(inputPath string, expectedLines int) string {
	reports, err := s.PrepareInput(inputPath, expectedLines)

	if err != nil {
		return err.Error()
	}

	count := 0
	for _, report := range reports {
		if IsValidReportPart1(report) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func (s Day02_Solution) Part2(inputPath string, expectedLines int) string {
	reports, err := s.PrepareInput(inputPath, expectedLines)

	if err != nil {
		return err.Error()
	}

	count := 0
	for _, report := range reports {
		if IsValidReportPart2(report) {
			count++
		}
	}

	return strconv.Itoa(count)
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

func (s Day02_Solution) PrepareInput(inputPath string, expectedLines int) (reports [][]int, err error) {
	lines, err := helper.ReadAllLines(inputPath+"/Day02.txt", expectedLines)

	if err != nil {
		return nil, err
	}

	reports = make([][]int, 0, expectedLines)

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
