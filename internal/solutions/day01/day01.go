package day01

import (
	"AdventOfCode/internal/common"
	"AdventOfCode/internal/intmath"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day01_Solution struct{}

func (s Day01_Solution) SolvePart1(input interface{}) (string, error) {
	p, ok := input.(parsed)
	if !ok {
		return "", fmt.Errorf("expected %T, got %T", parsed{}, input)
	}

	leftSide := p.leftSide
	rightSide := p.rightSide

	slices.Sort(leftSide)
	slices.Sort(rightSide)

	sum := 0
	for i, leftValue := range leftSide {
		rightValue := rightSide[i]
		sum += intmath.AbsDiff(leftValue, rightValue)
	}

	return strconv.Itoa(sum), nil
}

func (s Day01_Solution) SolvePart2(input interface{}) (string, error) {
	p, ok := input.(parsed)
	if !ok {
		return "", fmt.Errorf("expected %T, got %T", parsed{}, input)
	}

	leftSide := p.leftSide
	rightSide := p.rightSide

	rightDict := make(map[int]int)

	for _, rightValue := range rightSide {
		rightDict[rightValue] = rightDict[rightValue] + 1
	}

	sum := 0
	for _, leftValue := range leftSide {
		sum += leftValue * rightDict[leftValue]
	}
	return strconv.Itoa(sum), nil
}

func (s Day01_Solution) ParseInput(input string) (ret interface{}, err error) {
	lines := common.SplitLines(input)

	localLeftSide := make([]int, 0, len(lines))
	localRightSide := make([]int, 0, len(lines))

	for _, line := range lines {
		res := strings.Fields(line)
		leftValue, err := strconv.Atoi(res[0])
		if err != nil {
			return parsed{}, err
		}
		rightValue, err := strconv.Atoi(res[1])
		if err != nil {
			return parsed{}, err
		}
		localLeftSide = append(localLeftSide, leftValue)
		localRightSide = append(localRightSide, rightValue)
	}

	return parsed{leftSide: localLeftSide, rightSide: localRightSide}, nil
}

type parsed struct {
	leftSide  []int
	rightSide []int
}
