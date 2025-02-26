package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

type Day03_Solution struct {
}

var regexpPart1 *regexp.Regexp
var regexpPart2 *regexp.Regexp

func init() {
	regexpPart1 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	regexpPart2 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don't\(\))`)
}

func (s Day03_Solution) SolvePart1(input interface{}) (string, error) {
	data, ok := input.(string)
	if !ok {
		return "", fmt.Errorf("expected string, got %T", input)
	}

	sum := 0
	for _, match := range regexpPart1.FindAllStringSubmatch(data, -1) {
		// 0 Index full match, 1 CaptureGroup 1, 2 CaptureGroup 2
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		sum += left * right
	}

	return strconv.Itoa(sum), nil
}

func (s Day03_Solution) SolvePart2(input interface{}) (string, error) {
	data, ok := input.(string)
	if !ok {
		return "", fmt.Errorf("expected string, got %T", input)
	}

	sum := 0
	enabled := true
	for _, match := range regexpPart2.FindAllStringSubmatch(data, -1) {
		// 0 Index full match, 1 CaptureGroup 1, 2 CaptureGroup 2
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			left, _ := strconv.Atoi(match[1])
			right, _ := strconv.Atoi(match[2])
			sum += left * right
		}
	}

	return strconv.Itoa(sum), nil
}

func (s Day03_Solution) ParseInput(input string) (interface{}, error) {
	return input, nil
}
