package main

import (
	_ "embed"
	"testing"
)

var example1 = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var example2 = `Register A: 0
Register B: 0
Register C: 9

Program: 2,6,5,5`

var example3 = `Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`

var example4 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var example5 = `Register A: 0
Register B: 29
Register C: 0

Program: 1,7,5,5`

var example6 = `Register A: 0
Register B: 2024
Register C: 43690

Program: 4,0,5,5`

// func Test_part1(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input string
// 		want  string
// 	}{
// 		{
// 			name:  "example",
// 			input: example1,
// 			want:  "4,6,3,5,6,3,5,2,1,0",
// 		},
// 		{
// 			name:  "example2",
// 			input: example2,
// 			want:  "1",
// 		},
// 		{
// 			name:  "example3",
// 			input: example3,
// 			want:  "0,1,2",
// 		},
// 		{
// 			name:  "example4",
// 			input: example4,
// 			want:  "4,2,5,6,7,7,7,7,3,1,0",
// 		},
// 		{
// 			name:  "example5",
// 			input: example5,
// 			want:  "2", // 26 mod 8
// 		},
// 		{
// 			name:  "example6",
// 			input: example6,
// 			want:  "2", // 44354 mod 8
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := part1(tt.input); got != tt.want {
// 				t.Errorf("part1() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "actual",
			input: input,
			want:  100,
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
