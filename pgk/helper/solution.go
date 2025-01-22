package helper

type Solution interface {
	Part1(inputPath string, expectedLines int) string
	Part2(inputPath string, expectedLines int) string
}
