package main

import (
	"advent-of-code-go/pkg/list"

	_ "embed"
	"testing"
)

var example1 = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  35,
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
			want:  46,
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

func Test_seedNumberInSeedRanges(t *testing.T) {
	tests := []struct {
		name    string
		seedNum int
		rng     [][]int
		want    bool
	}{
		{
			name:    "not in range",
			seedNum: 10,
			rng: [][]int{
				{0, 9},
				{11, 20},
			},
			want: false,
		},
		{
			name:    "at end of range",
			seedNum: 10,
			rng: [][]int{
				{0, 10},
				{14, 20},
			},
			want: true,
		},
		{
			name:    "at start of range",
			seedNum: 10,
			rng: [][]int{
				{0, 6},
				{10, 20},
			},
			want: true,
		},
		{
			name:    "middle of range",
			seedNum: 10,
			rng: [][]int{
				{0, 15},
				{100, 200},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seedNumberInSeedRanges(tt.seedNum, tt.rng); got != tt.want {
				t.Errorf("seedNumberInSeedRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSeedRanges(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]int
	}{
		{
			name:  "calculates range",
			input: "10 5 100 16",
			want: [][]int{
				{10, 15},
				{100, 116},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSeedRanges(tt.input); !list.ListOfListsAreEqual[int](got, tt.want) {
				t.Errorf("calculateSeedRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapValueReverse(t *testing.T) {
	tests := []struct {
		name string
		val  int
		maps []Map
		want int
	}{
		{
			name: "outside of either range",
			val:  10,
			maps: []Map{
				{
					source: 10,
					dest:   16,
					rng:    5,
				},
				{
					source: 50,
					dest:   65,
					rng:    10,
				},
			},
			want: 10,
		},
		{
			name: "in first range",
			val:  17,
			maps: []Map{
				{
					source: 10,
					dest:   16,
					rng:    5,
				},
				{
					source: 50,
					dest:   65,
					rng:    10,
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapValueReverse(tt.val, tt.maps); got != tt.want {
				t.Errorf("mapValueReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
