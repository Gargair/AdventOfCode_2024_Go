package solutions

// Solution defines the interface for all problem solutions
type Solution interface {
	// ParseInput parses the input string into a problem-specific format
	ParseInput(input string) (interface{}, error)

	// SolvePart1 solves the first part of the problem
	SolvePart1(parsed interface{}) (string, error)

	// SolvePart2 solves the second part of the problem
	SolvePart2(parsed interface{}) (string, error)
}
