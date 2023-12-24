package main

import (
	_ "embed"
	"testing"
)

var example1 = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

// var example1 = `24134
// 32154
// 32552
// 34465
// 45466`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  102,
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

func Test_sameDirection(t *testing.T) {
	tests := []struct {
		name string
		l    []int
		d    int
		want bool
	}{
		{
			name: "standard",
			l:    []int{1, 1, 1},
			d:    1,
			want: true,
		},
		{
			name: "longer",
			l:    []int{2, 1, 1, 1},
			d:    1,
			want: true,
		},
		{
			name: "shorter",
			l:    []int{2, 2},
			d:    2,
			want: false,
		},
		{
			name: "longer invalid",
			l:    []int{2, 2, 2, 2, 1, 2},
			d:    2,
			want: false,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameDirection(tt.l, tt.d); got != tt.want {
				t.Errorf("sameDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
