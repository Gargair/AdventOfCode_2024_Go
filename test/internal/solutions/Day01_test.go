package solutions_test

import (
	"AdventOfCode/internal/helper"
	"AdventOfCode/internal/solutions"
	"testing"
)

func TestDay01_Part1(t *testing.T) {
	tests, err := helper.ReadTestData("../../testdata/Day01.txt", 1)

	if err != nil {
		t.Errorf("Day01: %v", err)
		return
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := s.Part1(tt.Args.InputPath, tt.Args.ExpectedLines); got != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, got, tt.Want)
			}
		})
	}
}

func TestDay01_Part2(t *testing.T) {
	tests, err := helper.ReadTestData("../../testdata/Day01.txt", 2)

	if err != nil {
		t.Errorf("Day01: %v", err)
		return
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := s.Part2(tt.Args.InputPath, tt.Args.ExpectedLines); got != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, got, tt.Want)
			}
		})
	}
}

func BenchmarkDay01_Part1(b *testing.B) {
	tests, err := helper.ReadTestData("../../testdata/Day01.txt", 1)

	if err != nil {
		b.Errorf("Day01: %v", err)
		return
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Part1(tt.Args.InputPath, tt.Args.ExpectedLines)
			}
		})
	}
}

func BenchmarkDay01_Part2(b *testing.B) {
	tests, err := helper.ReadTestData("../../testdata/Day01.txt", 2)

	if err != nil {
		b.Errorf("Day01: %v", err)
		return
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Part2(tt.Args.InputPath, tt.Args.ExpectedLines)
			}
		})
	}
}
