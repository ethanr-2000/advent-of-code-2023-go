package main

import (
	_ "embed"
	"testing"
)

var example1 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

var exampleNegatives = `-1 -3 -6 -10 -15`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  114,
		},
		{
			name:  "exampleNegatives",
			input: exampleNegatives,
			want:  -21,
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

var example2 = `0`

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example2,
			want:  0,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_getNumberHistory(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input []int
// 		want  []int
// 	}{
// 		{
// 			name:  "example",
// 			input: []int{0, 3, 6, 9, 12, 15},
// 			want:  []int{3, 3, 3, 3, 3},
// 		},
// 		{
// 			name:  "example2",
// 			input: []int{10, 13, 16, 21, 30, 45, 68},
// 			want:  []int{3, 3, 5, 9, 15, 23},
// 		},
// 		{
// 			name:  "negative",
// 			input: []int{0, -2, -5, -17, -38, -48, -1},
// 			want:  []int{-2, -3, -12, -21, -10, 47},
// 		},
// 		{
// 			name:  "negative2",
// 			input: []int{-1, 7, 33, 88},
// 			want:  []int{8, 26, 55},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := getNumberHistory(tt.input); slices.Compare(got, tt.want) != 0 {
// 				t.Errorf("getNumberHistory() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_getNextNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "example",
			input: []int{0, 3, 6, 9, 12, 15},
			want:  18,
		},
		{
			name:  "example2",
			input: []int{10, 13, 16, 21, 30, 45},
			want:  68,
		},
		{
			name:  "negative",
			input: []int{-1, -3, -6, -10, -15},
			want:  -21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextNumber(tt.input); got != tt.want {
				t.Errorf("getNextNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
