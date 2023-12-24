package main

import (
	_ "embed"
	"slices"
	"testing"
)

var example1 = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  2,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, 7, 27); got != tt.want {
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

func Test_getHailstones(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		ignoreZ bool
		want    []Hailstone
	}{
		{
			name: "example",
			input: []string{
				"147847636573416, 190826994408605, 140130741291716 @ 185, 49, 219",
				"287509258905812, -207449079739538, 280539021150559 @ -26, 31, 8",
			},
			ignoreZ: false,
			want: []Hailstone{
				{[3]float64{147847636573416, 190826994408605, 140130741291716}, [3]float64{185, 49, 219}},
				{[3]float64{287509258905812, -207449079739538, 280539021150559}, [3]float64{-26, 31, 8}},
			},
		},
		{
			name: "example",
			input: []string{
				"147847636573416, 190826994408605, 140130741291716 @ 185, 49, 219",
				"287509258905812, -207449079739538, 280539021150559 @ -26, 31, 8",
			},
			ignoreZ: true,
			want: []Hailstone{
				{[3]float64{147847636573416, 190826994408605, 0}, [3]float64{185, 49, 0}},
				{[3]float64{287509258905812, -207449079739538, 0}, [3]float64{-26, 31, 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHailstones(tt.input, tt.ignoreZ); !slices.Equal[[]Hailstone](got, tt.want) {
				t.Errorf("getHailstones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIntersection(t *testing.T) {
	tests := []struct {
		name         string
		h1           Hailstone
		h2           Hailstone
		intersection [3]float64
		inPast       bool
	}{
		{
			name:         "example",
			h1:           Hailstone{[3]float64{19, 13, 0}, [3]float64{-2, 1, 0}},
			h2:           Hailstone{[3]float64{18, 19, 0}, [3]float64{-1, -1, 0}},
			inPast:       false,
			intersection: [3]float64{43.0 / 3, 46.0 / 3, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, actualInPast := findIntersection(tt.h1, tt.h2); !slices.Equal[[]float64](got[:], tt.intersection[:]) || actualInPast != tt.inPast {
				t.Errorf("findIntersection() = %v, %v, want %v %v", actualInPast, got, tt.inPast, tt.intersection)
			}
		})
	}
}
