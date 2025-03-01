package day04

import (
	"fmt"
	"strconv"

	"github.com/Gargair/AdventOfCode_2024_Go/internal/common"
)

type Day04_Solution struct{}

func (s Day04_Solution) SolvePart1(input interface{}) (string, error) {
	data, ok := input.([][]rune)
	if !ok {
		return "", fmt.Errorf("expected [][]rune, got %T", input)
	}
	matchString := "XMAS"
	match := make([]rune, 0, len(matchString))
	for _, rune := range matchString {
		match = append(match, rune)
	}

	count := 0
	for startX, line := range data {
		for startY, rune := range line {
			if rune == match[0] {
				for directionX := -1; directionX <= 1; directionX++ {
					for directionY := -1; directionY <= 1; directionY++ {
						if hasMatchingStringInDirection(data, match, startX, startY, directionX, directionY) {
							count++
						}
					}
				}
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (s Day04_Solution) SolvePart2(input interface{}) (string, error) {
	data, ok := input.([][]rune)
	if !ok {
		return "", fmt.Errorf("expected [][]rune, got %T", input)
	}
	matchString := "MAS"
	match := make([]rune, 0, len(matchString))
	for _, rune := range matchString {
		match = append(match, rune)
	}

	count := 0
	for startX, line := range data {
		for startY, rune := range line {
			if rune == match[1] {
				if hasMatchingStringInDirection(data, match, startX-1, startY-1, 1, 1) && (hasMatchingStringInDirection(data, match, startX+1, startY-1, -1, 1) || hasMatchingStringInDirection(data, match, startX-1, startY+1, 1, -1)) {
					count++
				} else if (hasMatchingStringInDirection(data, match, startX+1, startY-1, -1, 1) || hasMatchingStringInDirection(data, match, startX-1, startY+1, 1, -1)) && hasMatchingStringInDirection(data, match, startX+1, startY+1, -1, -1) {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count), nil
}

func hasMatchingStringInDirection(matrix [][]rune, match []rune, startX int, startY int, directionX int, directionY int) bool {
	currentX := startX
	currentY := startY
	for _, rune := range match {
		if currentX < 0 || currentY < 0 {
			return false
		}
		if currentX >= len(matrix) {
			return false
		}
		if currentY >= len(matrix[currentX]) {
			return false
		}
		if matrix[currentX][currentY] != rune {
			return false
		}
		currentX += directionX
		currentY += directionY
	}
	return true
}

func (s Day04_Solution) ParseInput(input string) (interface{}, error) {
	lines := common.SplitLines(input)
	matrix := make([][]rune, 0, len(lines))
	for _, line := range lines {
		runeLine := make([]rune, 0, len(line))
		for _, rune := range line {
			runeLine = append(runeLine, rune)
		}
		matrix = append(matrix, runeLine)
	}
	return matrix, nil
}
