package runner

import (
	"fmt"
	"os"
	"time"

	"AdventOfCode/internal/solutions"
)

// Run executes a solution with the given parameters
func Run(solution solutions.Solution, part int, inputFile string) (string, time.Duration) {
	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}
	input := string(data)

	// Measure execution time
	start := time.Now()

	// Parse the input
	parsed, err := solution.ParseInput(input)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		os.Exit(1)
	}

	// Execute the appropriate part
	var result string
	var solveErr error
	if part == 1 {
		result, solveErr = solution.SolvePart1(parsed)
	} else {
		result, solveErr = solution.SolvePart2(parsed)
	}

	if solveErr != nil {
		fmt.Printf("Error solving problem: %v\n", solveErr)
		os.Exit(1)
	}

	elapsed := time.Since(start)
	return result, elapsed
}
