package solutions_test

import (
	"AdventOfCode/pgk/helper"
	"AdventOfCode/solutions"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type args struct {
	inputPath     string
	expectedLines int
}

func TestDay01_Part1(t *testing.T) {

	lines, err := helper.ReadAllLines("../Testcases/Day01.txt", 2)

	if err != nil {
		t.Errorf("Day01: %v", err)
		return
	}
	tests := make([]struct {
		name string
		args args
		want string
	}, 0, 2)
	for _, line := range *lines {
		res := strings.Fields(line)
		want := res[0]
		inputPath := res[2]
		expectedLines, err := strconv.Atoi(res[3])
		if err != nil {
			t.Errorf("Day01: %v", err)
			return
		}
		tests = append(tests, struct {
			name string
			args args
			want string
		}{name: fmt.Sprintf("inputPath=%v/expectedLines=%v", inputPath, expectedLines), args: args{inputPath: "../" + inputPath, expectedLines: expectedLines}, want: want})
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Part1(tt.args.inputPath, tt.args.expectedLines); got != tt.want {
				t.Errorf("%v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDay01_Part2(t *testing.T) {
	lines, err := helper.ReadAllLines("../Testcases/Day01.txt", 2)

	if err != nil {
		t.Errorf("Day01: %v", err)
		return
	}
	tests := make([]struct {
		name string
		args args
		want string
	}, 0, 2)
	for _, line := range *lines {
		res := strings.Fields(line)
		want := res[1]
		inputPath := res[2]
		expectedLines, err := strconv.Atoi(res[3])
		if err != nil {
			t.Errorf("Day01: %v", err)
			return
		}
		tests = append(tests, struct {
			name string
			args args
			want string
		}{name: fmt.Sprintf("inputPath=%v/expectedLines=%v", inputPath, expectedLines), args: args{inputPath: "../" + inputPath, expectedLines: expectedLines}, want: want})
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Part2(tt.args.inputPath, tt.args.expectedLines); got != tt.want {
				t.Errorf("%v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkDay01_Part1(b *testing.B) {
	lines, err := helper.ReadAllLines("../Testcases/Day01.txt", 2)

	if err != nil {
		b.Errorf("Day01: %v", err)
		return
	}
	tests := make([]struct {
		name string
		args args
		want string
	}, 0, 2)
	for _, line := range *lines {
		res := strings.Fields(line)
		want := res[0]
		inputPath := res[2]
		expectedLines, err := strconv.Atoi(res[3])
		if err != nil {
			b.Errorf("Day01: %v", err)
			return
		}
		tests = append(tests, struct {
			name string
			args args
			want string
		}{name: fmt.Sprintf("inputPath=%v/expectedLines=%v", inputPath, expectedLines), args: args{inputPath: "../" + inputPath, expectedLines: expectedLines}, want: want})
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Part1(tt.args.inputPath, tt.args.expectedLines)
			}
		})
	}
}

func BenchmarkDay01_Part2(b *testing.B) {
	lines, err := helper.ReadAllLines("../Testcases/Day01.txt", 2)

	if err != nil {
		b.Errorf("Day01: %v", err)
		return
	}
	tests := make([]struct {
		name string
		args args
		want string
	}, 0, 2)
	for _, line := range *lines {
		res := strings.Fields(line)
		want := res[1]
		inputPath := res[2]
		expectedLines, err := strconv.Atoi(res[3])
		if err != nil {
			b.Errorf("Day01: %v", err)
			return
		}
		tests = append(tests, struct {
			name string
			args args
			want string
		}{name: fmt.Sprintf("inputPath=%v/expectedLines=%v", inputPath, expectedLines), args: args{inputPath: "../" + inputPath, expectedLines: expectedLines}, want: want})
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Part2(tt.args.inputPath, tt.args.expectedLines)
			}
		})
	}
}
