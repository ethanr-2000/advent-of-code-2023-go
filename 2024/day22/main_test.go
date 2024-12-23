package main

import (
	_ "embed"
	"testing"
)

var example1 = `1
10
100
2024`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  37327623,
		},
		{
			name:  "actual",
			input: input,
			want:  13461553007,
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
			input: `1
2
3
2024`,
			want:  23,
		},
		{
			name:  "actual",
			input: input,
			want:  1499,
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

func Test_mix(t *testing.T) {
	tests := []struct {
		name  string
		n1 int
		n2 int
		want  int
	}{
		{
			name:  "mix 42 and 15",
			n1: 42,
			n2: 15,
			want:  37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mix(tt.n1, tt.n2); got != tt.want {
				t.Errorf("mix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prune(t *testing.T) {
	tests := []struct {
		name  string
		n1 int
		want  int
	}{
		{
			name:  "prune",
			n1: 100000000,
			want:  16113920,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prune(tt.n1); got != tt.want {
				t.Errorf("prune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_price(t *testing.T) {
	tests := []struct {
		name  string
		n int
		want  int
	}{
		{
			name:  "price 100",
			n: 100,
			want:  0,
		},
		{
			name:  "price 5",
			n: 5,
			want:  5,
		},
		{
			name:  "price 123456789",
			n: 123456789,
			want:  9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := price(tt.n); got != tt.want {
				t.Errorf("price() = %v, want %v", got, tt.want)
			}
		})
	}
}

