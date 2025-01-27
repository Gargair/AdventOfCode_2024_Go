package main

import (
	"AdventOfCode/internal/helper"
	"AdventOfCode/internal/solutions"
	"fmt"
)

func main() {
	var s helper.Solution
	s = *new(solutions.Day01_Solution)
	res := s.Part1("Input", 1000)
	fmt.Println(res)
	res = s.Part2("Input", 1000)
	fmt.Println(res)
}
