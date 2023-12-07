package list_test

import (
	"slices"
	"testing"

	"advent-of-code-go/pkg/list"
)

func Test_ListOfListsOfIntAreEqual(t *testing.T) {
	tests := []struct {
		name string
		l1   [][]int
		l2   [][]int
		want bool
	}{
		{
			name: "list of lists are equal",
			l1:   [][]int{{10, 11, 12, 13}, {13, 14, 15, 16}},
			l2:   [][]int{{10, 11, 12, 13}, {13, 14, 15, 16}},
			want: true,
		},
		{
			name: "list of lists are not equal",
			l1:   [][]int{{10, 11, 12, 13}, {13, 14, 15, 16}},
			l2:   [][]int{{10, 11, 12, 13}, {13, 14, 15, 17}},
			want: false,
		},
		{
			name: "list of lists are not equal lengths",
			l1:   [][]int{{10, 11, 12, 13}},
			l2:   [][]int{{10, 11, 12, 13}, {13, 14, 15, 16}},
			want: false,
		},
		{
			name: "list in list is not equal length",
			l1:   [][]int{{10, 11, 12, 13}, {13, 14, 15, 16}},
			l2:   [][]int{{10, 11, 12, 13, 14}, {13, 14, 15, 16}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.ListOfListsOfIntAreEqual(tt.l1, tt.l2); got != tt.want {
				t.Errorf("ListOfListsOfIntAreEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CountOfOccurencesOfStringInList(t *testing.T) {
	tests := []struct {
		name string
		l    []string
		s    string
		want int
	}{
		{
			name: "contains 0",
			l:    []string{"b", "c", "d"},
			s:    "a",
			want: 0,
		},
		{
			name: "contains 1",
			l:    []string{"a", "b", "c", "d"},
			s:    "a",
			want: 1,
		},
		{
			name: "all",
			l:    []string{"a", "a", "a", "a"},
			s:    "a",
			want: 4,
		},
		{
			name: "empty string",
			l:    []string{},
			s:    "a",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.CountOfOccurencesOfStringInList(tt.l, tt.s); got != tt.want {
				t.Errorf("CountOfOccurencesOfStringInList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ReplaceAllInstancesOfStringInList(t *testing.T) {
	tests := []struct {
		name string
		l    []string
		old  string
		new  string
		want []string
	}{
		{
			name: "does not contain old",
			l:    []string{"b", "c", "d"},
			old:  "a",
			new:  "b",
			want: []string{"b", "c", "d"},
		},
		{
			name: "contains 1",
			l:    []string{"a", "b", "c", "d"},
			old:  "a",
			new:  "b",
			want: []string{"b", "b", "c", "d"},
		},
		{
			name: "all",
			l:    []string{"a", "a", "a", "a"},
			old:  "a",
			new:  "b",
			want: []string{"b", "b", "b", "b"},
		},
		{
			name: "empty string",
			l:    []string{},
			old:  "a",
			new:  "b",
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.ReplaceAllInstancesOfStringInList(tt.l, tt.old, tt.new); slices.Compare[[]string](got, tt.want) != 0 {
				t.Errorf("ReplaceAllInstancesOfStringInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
