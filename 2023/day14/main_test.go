package main

import (
	"advent-of-code-go/pkg/list"
	_ "embed"
	"slices"
	"testing"
)

var example1 = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  136,
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
			want:  64,
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

func Test_tiltSlice(t *testing.T) {
	tests := []struct {
		name        string
		input       []rune
		want        []rune
		wantReverse []rune
	}{
		{
			name:        "1",
			input:       []rune("OO.O.O..##"),
			want:        []rune("OOOO....##"),
			wantReverse: []rune("....OOOO##"),
		},
		{
			name:        "2",
			input:       []rune("#OO.O#O.O"),
			want:        []rune("#OOO.#OO."),
			wantReverse: []rune("#.OOO#.OO"),
		},
		{
			name:        "3",
			input:       []rune("......."),
			want:        []rune("......."),
			wantReverse: []rune("......."),
		},
		{
			name:        "4",
			input:       []rune("###......O"),
			want:        []rune("###O......"),
			wantReverse: []rune("###......O"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tiltSlice(tt.input, false); !slices.Equal[[]rune](got, tt.want) {
				t.Errorf("tiltSlice() = %v, want %v", got, tt.want)
			}
			if got := tiltSlice(tt.input, true); !slices.Equal[[]rune](got, tt.wantReverse) {
				t.Errorf("reverse tiltSlice() = %v, want %v", got, tt.wantReverse)
			}
		})
	}
}

func Test_tiltWest(t *testing.T) {
	tests := []struct {
		name      string
		input     [][]rune
		direction Direction
		want      [][]rune
	}{
		{
			name: "1",
			input: [][]rune{
				[]rune("O....#...."),
				[]rune("O.OO#....#"),
				[]rune(".....##..."),
				[]rune("OO.#O....O"),
				[]rune(".O.....O#."),
				[]rune("O.#..O.#.#"),
				[]rune("..O..#O..O"),
				[]rune(".......O.."),
				[]rune("#....###.."),
				[]rune("#OO..#...."),
			},
			direction: North,
			want: [][]rune{
				[]rune("OOOO.#.O.."),
				[]rune("OO..#....#"),
				[]rune("OO..O##..O"),
				[]rune("O..#.OO..."),
				[]rune("........#."),
				[]rune("..#....#.#"),
				[]rune("..O..#.O.O"),
				[]rune("..O......."),
				[]rune("#....###.."),
				[]rune("#....#...."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// updates the given grid in place
			tiltGrid(tt.input, tt.direction)
			if !list.ListOfListsAreEqual[rune](tt.input, tt.want) {
				t.Errorf("tiltGrid() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func Test_tiltGridWest(t *testing.T) {
	tests := []struct {
		name      string
		input     [][]rune
		direction Direction
		want      [][]rune
	}{
		{
			name: "1",
			input: [][]rune{
				[]rune("O....#...."),
				[]rune("O.OO#....#"),
				[]rune(".....##..."),
				[]rune("OO.#O....O"),
				[]rune(".O.....O#."),
				[]rune("O.#..O.#.#"),
				[]rune("..O..#O..O"),
				[]rune(".......O.."),
				[]rune("#....###.."),
				[]rune("#OO..#...."),
			},
			direction: West,
			want: [][]rune{
				[]rune("O....#...."),
				[]rune("OOO.#....#"),
				[]rune(".....##..."),
				[]rune("OO.#OO...."),
				[]rune("OO......#."),
				[]rune("O.#O...#.#"),
				[]rune("O....#OO.."),
				[]rune("O........."),
				[]rune("#....###.."),
				[]rune("#OO..#...."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// updates the given grid in place
			tiltGrid(tt.input, tt.direction)
			if !list.ListOfListsAreEqual[rune](tt.input, tt.want) {
				t.Errorf("tiltGridWest() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func Test_tiltGridEast(t *testing.T) {
	tests := []struct {
		name      string
		input     [][]rune
		direction Direction
		want      [][]rune
	}{
		{
			name: "1",
			input: [][]rune{
				[]rune("O....#...."),
				[]rune("O.OO#....#"),
				[]rune(".....##..."),
				[]rune("OO.#O....O"),
				[]rune(".O.....O#."),
				[]rune("O.#..O.#.#"),
				[]rune("..O..#O..O"),
				[]rune(".......O.."),
				[]rune("#....###.."),
				[]rune("#OO..#...."),
			},
			direction: East,
			want: [][]rune{
				[]rune("....O#...."),
				[]rune(".OOO#....#"),
				[]rune(".....##..."),
				[]rune(".OO#....OO"),
				[]rune("......OO#."),
				[]rune(".O#...O#.#"),
				[]rune("....O#..OO"),
				[]rune(".........O"),
				[]rune("#....###.."),
				[]rune("#..OO#...."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// updates the given grid in place
			tiltGrid(tt.input, tt.direction)
			if !list.ListOfListsAreEqual[rune](tt.input, tt.want) {
				t.Errorf("tiltGridEast() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func Test_tiltGridSouth(t *testing.T) {
	tests := []struct {
		name      string
		input     [][]rune
		direction Direction
		want      [][]rune
	}{
		{
			name: "1",
			input: [][]rune{
				[]rune("O....#...."),
				[]rune("O.OO#....#"),
				[]rune(".....##..."),
				[]rune("OO.#O....O"),
				[]rune(".O.....O#."),
				[]rune("O.#..O.#.#"),
				[]rune("..O..#O..O"),
				[]rune(".......O.."),
				[]rune("#....###.."),
				[]rune("#OO..#...."),
			},
			direction: South,
			want: [][]rune{
				[]rune(".....#...."),
				[]rune("....#....#"),
				[]rune("...O.##..."),
				[]rune("...#......"),
				[]rune("O.O....O#O"),
				[]rune("O.#..O.#.#"),
				[]rune("O....#...."),
				[]rune("OO....OO.."),
				[]rune("#OO..###.."),
				[]rune("#OO.O#...O"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// updates the given grid in place
			tiltGrid(tt.input, tt.direction)
			if !list.ListOfListsAreEqual[rune](tt.input, tt.want) {
				t.Errorf("tiltGridEast() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func Test_hashGrid(t *testing.T) {
	tests := []struct {
		name  string
		input [][]rune
		want  string
	}{
		{
			name: "1",
			input: [][]rune{
				[]rune("O....#...."),
				[]rune("O.OO#....#"),
				[]rune(".....##..."),
				[]rune("OO.#O....O"),
				[]rune(".O.....O#."),
				[]rune("O.#..O.#.#"),
				[]rune("..O..#O..O"),
				[]rune(".......O.."),
				[]rune("#....###.."),
				[]rune("#OO..#...."),
			},
			want: `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashGrid(tt.input); got != tt.want {
				t.Errorf("hashGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKeyByValue(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		value int
		want  string
	}{
		{
			name: "1",
			input: map[string]int{
				`string1`: 1,
				`string2`: 2,
				`string3`: 3,
			},
			value: 2,
			want:  "string2",
		},
		{
			name: "2",
			input: map[string]int{
				`string1`: 1,
				`string2`: 2,
				`string3`: 3,
			},
			value: 3,
			want:  "string3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKeyByValue(tt.input, tt.value); got != tt.want {
				t.Errorf("findKeyByValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
