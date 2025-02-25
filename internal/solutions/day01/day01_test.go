package day01

import (
	"AdventOfCode/internal/helper"
	"AdventOfCode/internal/runner"
	"testing"
)

func TestDay01_Part1(t *testing.T) {
	tests, err := helper.ReadTestData("day01", 1)

	if err != nil {
		t.Errorf("Day01: %v", err)
		return
	}
	s := Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := runner.Run(s, 1, tt.Args.InputPath)[0]; got.Result != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, got.Result, tt.Want)
			}
		})
	}
}

func TestDay01_Part2(t *testing.T) {
	tests, err := helper.ReadTestData("day01", 2)

	if err != nil {
		t.Errorf("Day01: %v", err)
		return
	}
	s := Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := runner.Run(s, 2, tt.Args.InputPath)[0]; got.Result != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, got.Result, tt.Want)
			}
		})
	}
}

func BenchmarkDay01_Part1(b *testing.B) {
	tests, err := helper.ReadTestData("day01", 1)

	if err != nil {
		b.Errorf("Day01: %v", err)
		return
	}
	s := Day01_Solution{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runner.Run(s, 1, tt.Args.InputPath)
			}
		})
	}
}

func BenchmarkDay01_Part2(b *testing.B) {
	tests, err := helper.ReadTestData("day01", 2)

	if err != nil {
		b.Errorf("Day01: %v", err)
		return
	}
	s := Day01_Solution{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runner.Run(s, 2, tt.Args.InputPath)
			}
		})
	}
}
