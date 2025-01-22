package solutions

import (
	"AdventOfCode/pgk/helper"
	"AdventOfCode/pgk/helper/intmath"
	"slices"
	"strconv"
	"strings"
)

type Day01_Solution struct{}

func (s Day01_Solution) Part1(inputPath string, expectedLines int) string {
	leftSide, rightSide := s.PrepareInput(inputPath, expectedLines)

	slices.Sort(*leftSide)
	slices.Sort(*rightSide)

	sum := 0
	for i, leftValue := range *leftSide {
		rightValue := (*rightSide)[i]
		sum += intmath.AbsDiff(&leftValue, &rightValue)
	}

	return strconv.Itoa(sum)
}

func (s Day01_Solution) Part2(inputPath string, expectedLines int) string {
	leftSide, rightSide := s.PrepareInput(inputPath, expectedLines)

	rightDict := make(map[int]int)

	for _, rightValue := range *rightSide {
		rightDict[rightValue] = rightDict[rightValue] + 1
	}

	sum := 0
	for _, leftValue := range *leftSide {
		sum += leftValue * rightDict[leftValue]
	}
	return strconv.Itoa(sum)
}

func (s Day01_Solution) PrepareInput(inputPath string, expectedLines int) (leftSide *[]int, rightSide *[]int) {
	lines, err := helper.ReadAllLines(inputPath+"/Day01.txt", expectedLines)

	if err != nil {
		panic(err)
	}

	localLeftSide := make([]int, 0, expectedLines)
	localRightSide := make([]int, 0, expectedLines)

	for _, line := range *lines {
		res := strings.Fields(line)
		leftValue, leftErr := strconv.Atoi(res[0])
		if leftErr != nil {
			panic(leftErr)
		}
		rightValue, rightErr := strconv.Atoi(res[1])
		if rightErr != nil {
			panic(rightErr)
		}
		localLeftSide = append(localLeftSide, leftValue)
		localRightSide = append(localRightSide, rightValue)
	}

	return &localLeftSide, &localRightSide
}
