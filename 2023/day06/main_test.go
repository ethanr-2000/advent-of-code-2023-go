package main

import (
	_ "embed"
	"testing"
)

var example1 = `Time:      7  15   30
Distance:  9  40  200`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  288,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  71503,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchForUpperBoundary(t *testing.T) {
	tests := []struct {
		name        string
		target      int
		initialHigh int
		findLower   bool
		want        int
	}{
		{
			name:        "findUpper on test input",
			target:      9, // record
			initialHigh: 7, // max time
			want:        5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchForUpperBoundary(calculateDistanceGivenMsHeld, tt.target, tt.initialHigh); got != tt.want {
				t.Errorf("binarySearchForUpperBoundary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchForLowerBoundary(t *testing.T) {
	tests := []struct {
		name        string
		target      int
		initialHigh int
		want        int
	}{
		{
			name:        "findLowerBoundary on test input",
			target:      9,
			initialHigh: 7,
			want:        2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchForLowerBoundary(calculateDistanceGivenMsHeld, tt.target, tt.initialHigh); got != tt.want {
				t.Errorf("binarySearchForLowerBoundary() = %v, want %v", got, tt.want)
			}
		})
	}
}
