package runner

import (
	"fmt"
	"os"
	"time"

	"github.com/Gargair/AdventOfCode_2024_Go/internal/solutions"
)

// Run executes a solution with the given parameters
func Run(solution solutions.Solution, part int, inputFile string) ([]RunResult, time.Duration) {
	total_start := time.Now()

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}
	input := string(data)

	// Parse the input
	parsed, err := solution.ParseInput(input)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		os.Exit(1)
	}

	results := make([]RunResult, 0)
	// Execute the appropriate part
	if part == 0 || part == 1 {
		start := time.Now()
		res, solveErr := solution.SolvePart1(parsed)
		if solveErr != nil {
			fmt.Printf("Error solving problem: %v\n", solveErr)
			os.Exit(1)
		}
		elapsed := time.Since(start)
		results = append(results, RunResult{"Part: 1", res, elapsed})
	}

	if part == 0 || part == 2 {
		start := time.Now()
		res, solveErr := solution.SolvePart2(parsed)
		if solveErr != nil {
			fmt.Printf("Error solving problem: %v\n", solveErr)
			os.Exit(1)
		}
		elapsed := time.Since(start)
		results = append(results, RunResult{"Part: 2", res, elapsed})
	}

	total_elapsed := time.Since(total_start)

	return results, total_elapsed
}

type RunResult struct {
	Name    string
	Result  string
	Elapsed time.Duration
}
