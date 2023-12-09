package regex_test

import (
	"testing"

	"advent-of-code-go/pkg/regex"

	"slices"
)

func Test_GetNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "normal list of numbers",
			input: " 10 11 12 13 ",
			want:  []int{10, 11, 12, 13},
		},
		{
			name:  "text with number and colon and list of numbers",
			input: "A thing 1: 10 11 12 13",
			want:  []int{1, 10, 11, 12, 13},
		},
		{
			name:  "some text",
			input: "A thing",
			want:  []int{},
		},
		{
			name:  "numbers then some text",
			input: "10 11 12 13 | A thing",
			want:  []int{10, 11, 12, 13},
		},
		{
			name:  "one number",
			input: "10",
			want:  []int{10},
		},
		{
			name:  "empty string",
			input: "",
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regex.GetNumbers(tt.input); slices.Compare[[]int](got, tt.want) != 0 {
				t.Errorf("GetNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetSpaceSeparatedNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "normal list of numbers",
			input: " 10 11 12 13 ",
			want:  []int{10, 11, 12, 13},
		},
		{
			name:  "text with number and colon and list of numbers",
			input: "A thing 1: -10 11 12 13",
			want:  []int{-10, 11, 12, 13},
		},
		{
			name:  "some text",
			input: "A thing",
			want:  []int{},
		},
		{
			name:  "numbers then some text",
			input: "10 11 -12 13 | A thing",
			want:  []int{10, 11, -12, 13},
		},
		{
			name:  "one number",
			input: "10",
			want:  []int{10},
		},
		{
			name:  "empty string",
			input: "",
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regex.GetSpaceSeparatedNumbers(tt.input); slices.Compare[[]int](got, tt.want) != 0 {
				t.Errorf("GetSpaceSeparatedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsEmptyString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  true,
		},
		{
			name:  "numbers",
			input: "123",
			want:  false,
		},
		{
			name:  "letters",
			input: "aaaaa",
			want:  false,
		},
		{
			name:  "special chars",
			input: "!@£$%^&*()",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regex.IsEmptyString(tt.input); got != tt.want {
				t.Errorf("IsEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HasText(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "numbers",
			input: "123",
			want:  false,
		},
		{
			name:  "lower case",
			input: "abc",
			want:  true,
		},
		{
			name:  "upper case",
			input: "ABC",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regex.HasText(tt.input); got != tt.want {
				t.Errorf("HasText() = %v, want %v", got, tt.want)
			}
		})
	}
}
