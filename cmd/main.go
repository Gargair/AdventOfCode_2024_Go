package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/tabwriter"

	"AdventOfCode/internal/common"
	"AdventOfCode/internal/runner"
	"AdventOfCode/internal/solutions"
	"AdventOfCode/internal/solutions/day01"
	"AdventOfCode/internal/solutions/day02"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Parse command-line arguments
	problemInput := common.StringArray{}
	flag.Var(&problemInput, "problem", "Problem to solve (format: day01, day02, etc.)")
	partPtr := flag.String("part", "", "Part to solve (1 or 2)")
	inputFolderPtr := flag.String("inputs", "inputs", "Folder containing input files")
	flag.Parse()
	problems := problemInput.GetValue()
	caser := cases.Title(language.English)
	writer := tabwriter.NewWriter(os.Stdout, 1, 4, 4, ' ', 0)

	if len(problems) == 0 {
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

	for _, problem := range problems {
		// Construct input file path
		inputFile := filepath.Join(*inputFolderPtr, fmt.Sprintf("%s.txt", problem))

		// Get the appropriate solution
		solution := getSolution(problem)
		if solution == nil {
			fmt.Printf("Unknown problem: %s\n", problem)
			os.Exit(1)
		}

		// Run the solution
		results, total_elapsed := runner.Run(solution, part, inputFile)

		for _, result := range results {
			fmt.Fprintf(writer, "Problem: %s\t%s\tResult: %s\tExecution time: %s\n", caser.String(problem), result.Name, result.Result, result.Elapsed)
		}
		writer.Flush()
		fmt.Printf("Total elapsed: %s\n", total_elapsed)
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
