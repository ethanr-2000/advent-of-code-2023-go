package main

import (
	_ "embed"
	"testing"
)

var example0 = `..X...
.SAMX.
.A..A.
XMAS.S
.X....`

var example1 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

var test1 = `S00S00S
0A0A0A0
00MMM00
SAMXMAS
00MMM00
0A0A0A0
S00S00S
`

var test2 = `SXXS
XAXA
XXMM
SAMX
`

var test3 = `XMAS
MMXX
AXAX
SXXS
`

var test4 = `SAMX
00MM
0A0A
S00S
`

var test5 = `S00S
A0A0
MM00
XMAS
`

var test6 = `0000000
0A0A0A0
00MMM00
0AMXMA0
00MMM00
0A0A0A0
0000000
`

var test7 = `XMA
M0M
AMX`

var test8 = `XMA
S00
`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example0",
			input: example0,
			want:  4,
		},
		{
			name:  "example",
			input: example1,
			want:  18,
		},
		{
			name:  "test1",
			input: test1,
			want:  8,
		},
		{
			name:  "test2",
			input: test2,
			want:  3,
		},
		{
			name:  "test3",
			input: test3,
			want:  3,
		},
		{
			name:  "test4",
			input: test4,
			want:  3,
		},
		{
			name:  "test5",
			input: test5,
			want:  3,
		},
		{
			name:  "test6",
			input: test6,
			want:  0,
		},
		{
			name:  "test7",
			input: test7,
			want:  0,
		},
		{
			name:  "test8",
			input: test8,
			want:  0,
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

// var example2 = `0`

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  9,
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

func Test_checkXmasWithDiff(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		startI int
		diff   int
		want   int
	}{
		{
			name:   "1 diff",
			input:  "AAAXMAS",
			startI: 3,
			diff:   1,
			want:   1,
		},
		{
			name:   "2 diff",
			input:  "AAAX0M0A0S",
			startI: 3,
			diff:   2,
			want:   1,
		},
		{
			name:   "3 diff",
			input:  "XXXMXXAXXS",
			startI: 0,
			diff:   3,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkXmasWithDiff(tt.input, tt.startI, tt.diff); got != tt.want {
				t.Errorf("checkXmasWithDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_safeAccessStr(t *testing.T) {
	tests := []struct {
		name  string
		input string
		i     int
		want  byte
	}{
		{
			name:  "zero index",
			input: "example",
			i:     0,
			want:  byte('e'),
		},
		{
			name:  "one index",
			input: "example",
			i:     1,
			want:  byte('x'),
		},
		{
			name:  "negative index",
			input: "example",
			i:     -1,
			want:  byte('0'),
		},
		{
			name:  "big index",
			input: "example",
			i:     100,
			want:  byte('0'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := safeAccessStr(tt.input, tt.i); got != tt.want {
				t.Errorf("safeAccessStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
