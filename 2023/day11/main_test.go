package main

import (
	_ "embed"
	"testing"
)

var example1 = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  374,
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
			input: example1,
			want:  82000210,
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

// func Test_expandUniverseVertically(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input [][]rune
// 		want  [][]rune
// 	}{
// 		{
// 			name: "example",
// input: [][]rune{
// 	{'.', '.', '.', '.'},
// 	{'.', '.', '.', '.'},
// 	{'.', '#', '.', '.'},
// 	{'.', '.', '.', '.'},
// 	{'.', '.', '.', '.'},
// 	{'.', '.', '#', '.'},
// },
// want: [][]rune{
// 	{'.', '.', '.', '.'},
// 	{'.', '.', '.', '.'},
// 	{'.', '#', '.', '.'},
// 	{'.', '.', '.', '.'},
// 	{'.', '.', '.', '.'},
// 	{'.', '.', '#', '.'},
// },
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := expandUniverseVertically(tt.input); got != tt.want {
// 				t.Errorf("expandUniverseVertically() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
