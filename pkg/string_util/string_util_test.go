package string_util_test

import (
	"advent-of-code-go/pkg/string_util"
	"testing"
)

func Test_ChangeRuneAtIndex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		i     int
		r     rune
		want  string
	}{
		{
			name:  "0 index",
			input: "ABCDEFG",
			i:     0,
			r:     '#',
			want:  "#BCDEFG",
		},
		{
			name:  "3 index",
			input: "ABCDEFG",
			i:     3,
			r:     'P',
			want:  "ABCPEFG",
		},
		{
			name:  "end index",
			input: "ABCDEFG",
			i:     6,
			r:     '+',
			want:  "ABCDEF+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := string_util.ChangeRuneAtIndex(tt.input, tt.i, tt.r); got != tt.want {
				t.Errorf("ChangeRuneAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Repeat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		times int
		sep   string
		want  string
	}{
		{
			name:  "0 times",
			input: "ABCDEFG",
			times: 0,
			sep:   "#",
			want:  "ABCDEFG",
		},
		{
			name:  "1 times",
			input: "ABCDEFG",
			times: 1,
			sep:   "#",
			want:  "ABCDEFG#ABCDEFG",
		},
		{
			name:  "2 times with no sep",
			input: "ABCDEFG",
			times: 2,
			sep:   "",
			want:  "ABCDEFGABCDEFGABCDEFG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := string_util.Repeat(tt.input, tt.times, tt.sep); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
