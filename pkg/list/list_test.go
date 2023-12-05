package list_test

import (
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
