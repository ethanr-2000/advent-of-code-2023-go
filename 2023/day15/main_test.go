package main

import (
	_ "embed"
	"testing"
)

var example1 = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  1320,
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
			want:  145,
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

func Test_hash(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "1",
			input: "rn=1",
			want:  30,
		},
		{
			name:  "2",
			input: "cm-",
			want:  253,
		},
		{
			name:  "3",
			input: "qp=3",
			want:  97,
		},
		{
			name:  "4",
			input: "cm=2",
			want:  47,
		},
		{
			name:  "5",
			input: "qp-",
			want:  14,
		},
		{
			name:  "6",
			input: "HASH",
			want:  52,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.input); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_label(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "1",
			input: "rn=1",
			want:  "rn",
		},
		{
			name:  "3",
			input: "qp=3",
			want:  "qp",
		},
		{
			name:  "4",
			input: "cm=2",
			want:  "cm",
		},
		{
			name:  "5",
			input: "cm-",
			want:  "cm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := label(tt.input); got != tt.want {
				t.Errorf("label() = %v, want %v", got, tt.want)
			}
		})
	}
}
