package main

import (
	_ "embed"
	"fmt"
	"testing"
)

var example1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var example2 = `12.......*..
+.........34
...#...-12..
..78........
..*....60...
78..........
.......23...
....90*12...
...%........
2.2......12.
.*......)..*
1.1.......56`

var example3 = `...1...5......
..7#..&.3....
...-13$.10....
10.....15.....`

var example4 = `.1.1...........
1...123........
1.*.123....123.
1...123.123.-.4
1.1.123.123...1`

var example5 = `......123.123.....
.....123...123....
.....123.+.123....
......123.123.....
......123.123.....
`

var example6 = `$..
.11
.11
$..
..$
11.
11.
..$`

var example7 = `$......$
.11..11.
.11..11.
$......$`

var example8 = `........
.24$-4..
......*.`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  4361,
		},
		{
			name:  "example2",
			input: example2,
			want:  413,
		},
		{
			name:  "example3",
			input: example3,
			want:  41,
		},
		{
			name:  "example4",
			input: example4,
			want:  123,
		},
		{
			name:  "example5",
			input: example5,
			want:  246,
		},
		{
			name:  "example6",
			input: example6,
			want:  44,
		},
		{
			name:  "example7",
			input: example7,
			want:  44,
		},
		{
			name:  "example8",
			input: example8,
			want:  28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("")
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
			want:  467835,
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

// Other tests

func Test_isAdjacent(t *testing.T) {
	tests := []struct {
		name        string
		partNumber  PartNumber
		specialChar SpecialChar
		want        bool
	}{
		{
			name: "number far far left",
			partNumber: PartNumber{
				row:        1,
				startIndex: 1,
				endIndex:   3,
				value:      10,
			},
			specialChar: SpecialChar{
				row:   1,
				index: 5,
			},
			want: false,
		},
		{
			name: "number far left",
			partNumber: PartNumber{
				row:        1,
				startIndex: 1,
				endIndex:   3,
				value:      10,
			},
			specialChar: SpecialChar{
				row:   1,
				index: 4,
			},
			want: true,
		},
		{
			name: "number left",
			partNumber: PartNumber{
				row:        0,
				startIndex: 1,
				endIndex:   3,
				value:      10,
			},
			specialChar: SpecialChar{
				row:   1,
				index: 3,
			},
			want: true,
		},
		{
			name: "number middle",
			partNumber: PartNumber{
				row:        0,
				startIndex: 1,
				endIndex:   3,
				value:      10,
			},
			specialChar: SpecialChar{
				row:   1,
				index: 2,
			},
			want: true,
		},
		{
			name: "number right",
			partNumber: PartNumber{
				row:        0,
				startIndex: 1,
				endIndex:   3,
				value:      10,
			},
			specialChar: SpecialChar{
				row:   1,
				index: 1,
			},
			want: true,
		},
		{
			name: "number far right",
			partNumber: PartNumber{
				row:        0,
				startIndex: 1,
				endIndex:   3,
				value:      10,
			},
			specialChar: SpecialChar{
				row:   1,
				index: 0,
			},
			want: true,
		},
		{
			name: "number far far right",
			partNumber: PartNumber{
				row:        0,
				startIndex: 2,
				endIndex:   4,
				value:      10,
			},
			specialChar: SpecialChar{
				row:   1,
				index: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdjacent(tt.partNumber, tt.specialChar); got != tt.want {
				t.Errorf("isAdjacent() = %v, want %v for test %s", got, tt.want, tt.name)
			}
		})
	}
}

func Test_getSpecialChars(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []SpecialChar
	}{
		{
			name:  "all",
			input: []string{"*.#.+.$./.&.%.-.@.="},
			want: []SpecialChar{
				{value: '*', row: 0, index: 0},
				{value: '#', row: 0, index: 2},
				{value: '+', row: 0, index: 4},
				{value: '$', row: 0, index: 6},
				{value: '/', row: 0, index: 8},
				{value: '&', row: 0, index: 10},
				{value: '%', row: 0, index: 12},
				{value: '-', row: 0, index: 14},
				{value: '@', row: 0, index: 16},
				{value: '=', row: 0, index: 18},
			},
		},
		{
			name:  "none",
			input: []string{"............"},
			want:  []SpecialChar{},
		},
		{
			name:  "numbers",
			input: []string{".1234567890"},
			want:  []SpecialChar{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSpecialChars(tt.input); !areSpecialCharListsEqual(tt.want, got) {
				t.Errorf("getSpecialCharsInLine() = %v, want %v for test %s", got, tt.want, tt.name)
			}
		})
	}
}

func areSpecialCharListsEqual(list1 []SpecialChar, list2 []SpecialChar) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i].value != list2[i].value {
			return false
		}
		if list1[i].index != list2[i].index {
			return false
		}
		if list1[i].row != list2[i].row {
			return false
		}
	}

	return true
}
