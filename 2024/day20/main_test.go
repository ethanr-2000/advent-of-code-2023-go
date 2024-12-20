package main

import (
	_ "embed"
	"testing"
)

var example1 = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// This test reuires modifyign the code to look for savings of at least 2 picoseconds rather than 100
		// {
		// 	name:  "example",
		// 	input: example1,
		// 	want:  44,
		// },
		{
			name:  "actual",
			input: input,
			want:  1358,
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
		// {
		// 	name:  "example",
		// 	input: example1,
		// 	want:  0,
		// },
		{
			name:  "actual",
			input: input,
			want:  1005856,
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
