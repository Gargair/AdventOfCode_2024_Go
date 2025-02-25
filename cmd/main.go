package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"AdventOfCode/internal/runner"
	"AdventOfCode/internal/solutions"
	"AdventOfCode/internal/solutions/day01"
	"AdventOfCode/internal/solutions/day02"
)

func main() {
	// Parse command-line arguments
	problemPtr := flag.String("problem", "", "Problem to solve (format: day01, day02, etc.)")
	partPtr := flag.String("part", "", "Part to solve (1 or 2)")
	inputFolderPtr := flag.String("inputs", "inputs", "Folder containing input files")
	flag.Parse()

	if *problemPtr == "" {
		fmt.Println("Please specify a problem using -problem flag")
		os.Exit(1)
	}

	if *partPtr == "" {
		fmt.Println("Please specify a part using -part flag")
		os.Exit(1)
	}

	part, err := strconv.Atoi(*partPtr)
	if err != nil || (part != 0 && part != 1 && part != 2) {
		fmt.Println("Part must be 0, 1 or 2")
		os.Exit(1)
	}

	// Ensure input folder exists
	if _, err := os.Stat(*inputFolderPtr); os.IsNotExist(err) {
		fmt.Printf("Input folder '%s' does not exist\n", *inputFolderPtr)
		os.Exit(1)
	}

	// Construct input file path
	inputFile := filepath.Join(*inputFolderPtr, fmt.Sprintf("%s.txt", *problemPtr))

	// Get the appropriate solution
	solution := getSolution(*problemPtr)
	if solution == nil {
		fmt.Printf("Unknown problem: %s\n", *problemPtr)
		os.Exit(1)
	}

	// Run the solution
	results := runner.Run(solution, part, inputFile)

	for _, result := range results {
		fmt.Printf("Name: %s\tResult: %s\tExecution time: %s\n", result.Name, result.Result, result.Elapsed)
	}
}

func getSolution(problem string) solutions.Solution {
	switch problem {
	case "day01":
		return &day01.Day01_Solution{}
	case "day02":
		return &day02.Day02_Solution{}
	default:
		return nil
	}
}
