package main

import (
	_ "embed"
	"testing"
)

var example0 = `AAAA
BBCD
BBCC
EEEC`

var example1 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

var example2 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example0",
			input: example0,
			want:  140,
		},
		{
			name:  "example1",
			input: example1,
			want:  772,
		},
		{
			name:  "example2",
			input: example2,
			want:  1930,
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
			name:  "example0",
			input: example0,
			want:  80,
		},
		{
			name:  "example1",
			input: example1,
			want:  436,
		},
		{
			name:  "example2",
			input: example2,
			want:  1206,
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
