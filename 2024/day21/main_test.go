package main

import (
	_ "embed"
	"testing"
)

var example0 = `029A`

var example1 = `029A
980A
179A
456A
379A`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example0",
			input: example0,
			want:  1972,
		},
		{
			name:  "example",
			input: example1,
			want:  126384,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_part2(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input string
// 		want  int
// 	}{
// 		{
// 			name:  "example",
// 			input: example1,
// 			want:  0,
// 		},
// 		// {
// 		// 	name:  "actual",
// 		// 	input: input,
// 		// 	want:  0,
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := part2(tt.input); got != tt.want {
// 				t.Errorf("part2() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
