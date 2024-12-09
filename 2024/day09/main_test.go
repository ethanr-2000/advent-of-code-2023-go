package main

import (
	_ "embed"
	"testing"
)

var example1 = `2333133121414131402
`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  1928,
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
			want:  2858,
		},
		// {
		// 	name:  "reddit",
		// 	input: "12345",
		// 	want:  132,
		// },
		// {
		// 	name:  "reddit2",
		// 	input: "1010101010101010101010",
		// 	want:  385,
		// },
		// {
		// 	name:  "252",
		// 	input: "252",
		// 	want:  5,
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
