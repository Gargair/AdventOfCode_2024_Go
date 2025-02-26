package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type DayData struct {
	Day         string
	DayNumber   string
	PackageName string
	StructName  string
}

func main() {
	// Parse command-line arguments
	dayPtr := flag.Int("day", 0, "Day number to generate (1-25)")
	overwritePtr := flag.Bool("force", false, "Force overwrite if files already exist")
	flag.Parse()

	if *dayPtr < 1 || *dayPtr > 25 {
		fmt.Println("Please provide a valid day number between 1 and 25")
		os.Exit(1)
	}

	// Format day number with leading zero
	dayNumber := fmt.Sprintf("%02d", *dayPtr)
	day := fmt.Sprintf("day%s", dayNumber)

	// Set up directory paths
	internalDir := "internal"
	solutionsDir := filepath.Join(internalDir, "solutions")
	dayDir := filepath.Join(solutionsDir, day)
	inputsDir := "inputs"
	testdataDir := "testdata"

	// Ensure directories exist
	dirs := []string{solutionsDir, dayDir, inputsDir, testdataDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}

	// Prepare data for templates
	data := DayData{
		Day:         day,
		DayNumber:   dayNumber,
		PackageName: day,
		StructName:  fmt.Sprintf("%s_Solution", strings.Title(day)),
	}

	// Template files to generate
	templates := map[string]string{
		filepath.Join(dayDir, fmt.Sprintf("%s.go", day)):      solutionTemplate,
		filepath.Join(dayDir, fmt.Sprintf("%s_test.go", day)): testTemplate,
	}

	// Input files to create if they don't exist
	inputFiles := []string{
		filepath.Join(inputsDir, fmt.Sprintf("%s.txt", day)),
		filepath.Join(testdataDir, fmt.Sprintf("%s.txt", day)),
	}

	// Generate solution and test files
	for path, content := range templates {
		if err := generateFile(path, content, data, *overwritePtr); err != nil {
			fmt.Printf("Error generating file %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// Create empty input files if they don't exist
	for _, path := range inputFiles {
		if err := createEmptyFile(path, *overwritePtr); err != nil {
			fmt.Printf("Error creating file %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// Update main.go to include the new day
	if err := updateMainFile(day, data.StructName); err != nil {
		fmt.Printf("Error updating main.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated solution files for %s\n", day)
}

func generateFile(path, content string, data DayData, overwrite bool) error {
	// Check if file already exists
	if _, err := os.Stat(path); err == nil && !overwrite {
		fmt.Printf("File %s already exists, skipping (use -force to overwrite)\n", path)
		return nil
	}

	// Create or truncate the file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse and execute the template
	tmpl, err := template.New("solution").Parse(content)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, data)
}

func createEmptyFile(path string, overwrite bool) error {
	// Check if file already exists
	if _, err := os.Stat(path); err == nil {
		if !overwrite {
			fmt.Printf("File %s already exists, skipping\n", path)
			return nil
		}
		fmt.Printf("File %s already exists, not modifying\n", path)
		return nil
	}

	// Create an empty file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if path == filepath.Join("testdata", fmt.Sprintf("%s.txt", filepath.Base(filepath.Dir(path)))) {
		// For testdata files, add a template line
		_, err = file.WriteString("0 0 inputs\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func updateMainFile(day string, structName string) error {
	// Read the main.go file
	mainFilePath := filepath.Join("cmd", "main.go")
	content, err := os.ReadFile(mainFilePath)
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Check if the day is already imported
	importLine := fmt.Sprintf(`"github.com/Gargair/AdventOfCode_2024_Go/internal/solutions/%s"`, day)
	if !strings.Contains(contentStr, importLine) {
		// Find the last import
		lastImportIndex := strings.LastIndex(contentStr, `"github.com/Gargair/AdventOfCode_2024_Go/internal/solutions/`)
		if lastImportIndex == -1 {
			return fmt.Errorf("could not find import section in main.go")
		}

		// Find the end of the last import line
		endOfLine := strings.Index(contentStr[lastImportIndex:], ")")
		if endOfLine == -1 {
			return fmt.Errorf("could not find end of import section in main.go")
		}

		insertPos := lastImportIndex + endOfLine
		newContent := contentStr[:insertPos] + "\n\t" + importLine + contentStr[insertPos:]
		contentStr = newContent
	}

	// Check if the day is already in getSolution
	caseLine := fmt.Sprintf(`case "%s":`, day)
	if !strings.Contains(contentStr, caseLine) {
		// Find the last case statement
		switchIndex := strings.LastIndex(contentStr, "switch problem {")
		if switchIndex == -1 {
			return fmt.Errorf("could not find switch statement in main.go")
		}

		defaultCaseIndex := strings.Index(contentStr[switchIndex:], "default:")
		if defaultCaseIndex == -1 {
			return fmt.Errorf("could not find default case in main.go")
		}

		insertPos := switchIndex + defaultCaseIndex
		newCase := fmt.Sprintf("\tcase \"%s\":\n\t\treturn &%s.%s{}\n\t", day, day, structName)
		newContent := contentStr[:insertPos] + newCase + contentStr[insertPos:]
		contentStr = newContent
	}

	// Write the updated content back to the file
	return os.WriteFile(mainFilePath, []byte(contentStr), 0644)
}

// Template for solution files
const solutionTemplate = `package {{.PackageName}}

import (
	"fmt"
	"strconv"

	"github.com/Gargair/AdventOfCode_2024_Go/internal/common"
)

type {{.StructName}} struct{}

func (s {{.StructName}}) SolvePart1(input interface{}) (string, error) {
	data, ok := input.([]string)
	if !ok {
		return "", fmt.Errorf("expected []string, got %T", input)
	}

	// TODO: Implement solution for Part 1
	_ = data // Remove this line once you use the data

	return "0", nil // Replace with actual result
}

func (s {{.StructName}}) SolvePart2(input interface{}) (string, error) {
	data, ok := input.([]string)
	if !ok {
		return "", fmt.Errorf("expected []string, got %T", input)
	}

	// TODO: Implement solution for Part 2
	_ = data // Remove this line once you use the data

	return "0", nil // Replace with actual result
}

func (s {{.StructName}}) ParseInput(input string) (interface{}, error) {
	// Default implementation: split the input into lines
	lines := common.SplitLines(input)
	return lines, nil

	// Modify this function based on your specific parsing needs
}
`

// Template for test files
const testTemplate = `package {{.PackageName}}

import (
	"testing"
	
	"github.com/Gargair/AdventOfCode_2024_Go/internal/helper"
	"github.com/Gargair/AdventOfCode_2024_Go/internal/runner"
)

func Test{{.StructName}}_Part1(t *testing.T) {
	tests, err := helper.ReadTestData("{{.Day}}", 1)

	if err != nil {
		t.Errorf("{{.Day}}: %v", err)
		return
	}
	s := {{.StructName}}{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res, _ := runner.Run(s, 1, tt.Args.InputPath)
			if len(res) > 0 && res[0].Result != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, res[0].Result, tt.Want)
			}
		})
	}
}

func Test{{.StructName}}_Part2(t *testing.T) {
	tests, err := helper.ReadTestData("{{.Day}}", 2)

	if err != nil {
		t.Errorf("{{.Day}}: %v", err)
		return
	}
	s := {{.StructName}}{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res, _ := runner.Run(s, 2, tt.Args.InputPath)
			if len(res) > 0 && res[0].Result != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, res[0].Result, tt.Want)
			}
		})
	}
}

func Benchmark{{.StructName}}_Part1(b *testing.B) {
	tests, err := helper.ReadTestData("{{.Day}}", 1)

	if err != nil {
		b.Errorf("{{.Day}}: %v", err)
		return
	}
	s := {{.StructName}}{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runner.Run(s, 1, tt.Args.InputPath)
			}
		})
	}
}

func Benchmark{{.StructName}}_Part2(b *testing.B) {
	tests, err := helper.ReadTestData("{{.Day}}", 2)

	if err != nil {
		b.Errorf("{{.Day}}: %v", err)
		return
	}
	s := {{.StructName}}{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runner.Run(s, 2, tt.Args.InputPath)
			}
		})
	}
}
`
