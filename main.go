package main

import (
	"AdventOfCode/pgk/helper"
	"AdventOfCode/solutions"
	"fmt"
)

func main() {
	var s helper.Solution
	s = *new(solutions.Day01_Solution)
	res := s.Part1("Input")
	fmt.Println(res)
	// res = s.Part2("Input_Test")
	// fmt.Println(res)
}
