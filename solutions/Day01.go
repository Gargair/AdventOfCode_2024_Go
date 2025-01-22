package solutions

import (
	"AdventOfCode/pgk/helper"
	"AdventOfCode/pgk/helper/intmath"
	"slices"
	"strconv"
	"strings"
)

type Day01_Solution struct{}

func (s Day01_Solution) Part1(inputPath string) string {
	leftSide, rightSide := s.PrepareInput(inputPath)

	slices.Sort(*leftSide)
	slices.Sort(*rightSide)

	sum := 0
	for i, leftValue := range *leftSide {
		rightValue := (*rightSide)[i]
		sum += intmath.AbsDiff(&leftValue, &rightValue)
	}

	return strconv.Itoa(sum)
}

func (s Day01_Solution) Part2(inputPath string) string {
	leftSide, rightSide := s.PrepareInput(inputPath)

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

func (s Day01_Solution) PrepareInput(inputPath string) (leftSide *[]int, rightSide *[]int) {
	lines, err := helper.ReadAllLines(inputPath+"/Day01.txt", 1000)

	if err != nil {
		panic(err)
	}

	localLeftSide := make([]int, 0, 1000)
	localRightSide := make([]int, 0, 1000)

	for _, line := range *lines {
		// var leftValue, rightValue int
		// _, err := fmt.Sscanf(line, "%d %d", &leftValue, &rightValue)
		// if err != nil {
		// 	panic(err)
		// }
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

func (s Day01_Solution) TestSorting(inputPath string) {
	leftSide, rightSide := s.PrepareInput(inputPath)

	slices.Sort(*leftSide)
	slices.Sort(*rightSide)
}
