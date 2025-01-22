package solutions_test

import (
	"AdventOfCode/solutions"
	"testing"
)

func TestDay01_Solution_Part1(t *testing.T) {
	type args struct {
		inputPath     string
		expectedLines int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Input_Test", args: args{inputPath: "../Input_Test", expectedLines: 6}, want: "11"},
		{name: "Input", args: args{inputPath: "../Input", expectedLines: 1000}, want: "1873376"},
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Part1(tt.args.inputPath, tt.args.expectedLines); got != tt.want {
				t.Errorf("Day01_Solution.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay01_Solution_Part2(t *testing.T) {
	type args struct {
		inputPath     string
		expectedLines int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Input_Test", args: args{inputPath: "../Input_Test", expectedLines: 6}, want: "31"},
		{name: "Input", args: args{inputPath: "../Input", expectedLines: 1000}, want: "18997088"},
	}
	s := solutions.Day01_Solution{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Part2(tt.args.inputPath, tt.args.expectedLines); got != tt.want {
				t.Errorf("Day01_Solution.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDay01_Solution_Part1(b *testing.B) {
	s := solutions.Day01_Solution{}
	b.Run("inputPath=Input", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Part1("../Input", 1000)
		}
	})
}

func BenchmarkDay01_Solution_Part2(b *testing.B) {
	s := solutions.Day01_Solution{}
	b.Run("inputPath=Input", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Part2("../Input", 1000)
		}
	})
}
