package main

import (
	_ "embed"
	"testing"
)

var example1 = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  46,
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
			want:  51,
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

func Test_nextDirectionForwardSlash(t *testing.T) {
	tests := []struct {
		name  string
		input Direction
		want  Direction
	}{
		{
			name:  "right",
			input: RIGHT,
			want:  UP,
		},
		{
			name:  "down",
			input: DOWN,
			want:  LEFT,
		},
		{
			name:  "left",
			input: LEFT,
			want:  DOWN,
		},
		{
			name:  "up",
			input: UP,
			want:  RIGHT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextDirectionForwardSlash(tt.input); got != tt.want {
				t.Errorf("nextDirectionForwardSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextDirectionBackSlash(t *testing.T) {
	tests := []struct {
		name  string
		input Direction
		want  Direction
	}{
		{
			name:  "right",
			input: RIGHT,
			want:  DOWN,
		},
		{
			name:  "down",
			input: DOWN,
			want:  RIGHT,
		},
		{
			name:  "left",
			input: LEFT,
			want:  UP,
		},
		{
			name:  "up",
			input: UP,
			want:  LEFT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextDirectionBackSlash(tt.input); got != tt.want {
				t.Errorf("nextDirectionBackSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
