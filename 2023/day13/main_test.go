package main

import (
	_ "embed"
	"testing"
)

var example1 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

var example2 = `#......
###.###
#.####.
.#....#
#..##..
##.##.#
...##..
##....#
#.#..#.
#.#..#.
##....#
...##..
##.##.#

..##...####..
##.###..##..#
..#..##.#####
..#.#########
...#.#.#..#.#
...##..#..#..
....##.####.#
..#.##.#..#.#
...##########
..##.###..###
####.#..##..#`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  405,
		},
		{
			name:  "example2",
			input: example2,
			want:  901,
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
			want:  400,
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

func Test_doesReflectHorizontal(t *testing.T) {
	tests := []struct {
		name string
		g    [][]rune
		i    int
		want bool
	}{
		{
			name: "small example",
			g: [][]rune{
				{'#', '.', '.'},
				{'#', '.', '.'},
				{'.', '.', '.'},
			},
			i:    1,
			want: true,
		},
		{
			name: "bigger example",
			g: [][]rune{
				{'#', '#', '.'},
				{'#', '.', '.'},
				{'#', '.', '#'},
				{'#', '.', '#'},
				{'#', '.', '#'},
			},
			i:    4,
			want: true,
		},
		{
			name: "biggerer example",
			g: [][]rune{
				{'#', '#', '.'},
				{'#', '.', '#'},
				{'#', '#', '#'},
				{'#', '#', '#'},
				{'#', '.', '#'},
			},
			i:    3,
			want: true,
		},
		{
			name: "false example",
			g: [][]rune{
				{'.', '#', '.'},
				{'#', '.', '#'},
				{'#', '.', '.'},
				{'.', '.', '#'},
				{'#', '#', '.'},
			},
			i:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesReflectHorizontal(tt.g, tt.i, 0); got != tt.want {
				t.Errorf("doesReflectHorizontal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doesReflectVertical(t *testing.T) {
	tests := []struct {
		name string
		g    [][]rune
		i    int
		want bool
	}{
		{
			name: "small example",
			g: [][]rune{
				{'#', '#', '.'},
				{'#', '#', '.'},
				{'.', '.', '.'},
			},
			i:    1,
			want: true,
		},
		{
			name: "bigger example",
			g: [][]rune{
				{'#', '#', '.', '.', '#'},
				{'#', '.', '.', '.', '.'},
				{'#', '#', '#', '#', '#'},
			},
			i:    3,
			want: true,
		},
		{
			name: "large i",
			g: [][]rune{
				{'#', '#', '.', '.', '.', '#'},
				{'#', '.', '.', '.', '.', '.'},
				{'#', '#', '#', '#', '#', '#'},
			},
			i:    5,
			want: false,
		},
		{
			name: "biggerer example",
			g: [][]rune{
				{'#', '#', '#', '#', '.', '#', '#'},
				{'#', '.', '.', '#', '.', '.', '.'},
				{'#', '#', '#', '#', '#', '#', '#'},
				{'#', '#', '#', '#', '#', '#', '#'},
			},
			i:    2,
			want: true,
		},
		{
			name: "false example",
			g: [][]rune{
				{'.', '#', '.'},
				{'#', '.', '#'},
				{'#', '.', '.'},
				{'.', '.', '#'},
				{'#', '#', '.'},
			},
			i:    1,
			want: false,
		},
		{
			name: "big false example",
			g: [][]rune{
				{'.', '#', '.', '#', '.'},
				{'#', '.', '#', '.', '#'},
				{'#', '.', '.', '.', '.'},
				{'.', '.', '#', '.', '#'},
				{'#', '#', '.', '#', '.'},
			},
			i:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesReflectVertical(tt.g, tt.i, 0); got != tt.want {
				t.Errorf("doesReflectVertical() = %v, want %v", got, tt.want)
			}
		})
	}
}
