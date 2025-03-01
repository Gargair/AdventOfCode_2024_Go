package day04

import (
	"testing"

	"github.com/Gargair/AdventOfCode_2024_Go/internal/common"
	"github.com/Gargair/AdventOfCode_2024_Go/internal/runner"
)

func TestDay04_Solution_Part1(t *testing.T) {
	tests, err := common.ReadTestData("day04", 1)

	if err != nil {
		t.Errorf("day04: %v", err)
		return
	}
	s := Day04_Solution{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res, _ := runner.Run(s, 1, tt.Args.InputPath)
			if len(res) > 0 && res[0].Result != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, res[0].Result, tt.Want)
			}
		})
	}
}

func TestDay04_Solution_Part2(t *testing.T) {
	tests, err := common.ReadTestData("day04", 2)

	if err != nil {
		t.Errorf("day04: %v", err)
		return
	}
	s := Day04_Solution{}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res, _ := runner.Run(s, 2, tt.Args.InputPath)
			if len(res) > 0 && res[0].Result != tt.Want {
				t.Errorf("%v = %v, want %v", tt.Name, res[0].Result, tt.Want)
			}
		})
	}
}

func BenchmarkDay04_Solution_Part1(b *testing.B) {
	tests, err := common.ReadTestData("day04", 1)

	if err != nil {
		b.Errorf("day04: %v", err)
		return
	}
	s := Day04_Solution{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runner.Run(s, 1, tt.Args.InputPath)
			}
		})
	}
}

func BenchmarkDay04_Solution_Part2(b *testing.B) {
	tests, err := common.ReadTestData("day04", 2)

	if err != nil {
		b.Errorf("day04: %v", err)
		return
	}
	s := Day04_Solution{}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runner.Run(s, 2, tt.Args.InputPath)
			}
		})
	}
}
