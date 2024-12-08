package grid_test

import (
	"advent-of-code-go/pkg/grid"
	"advent-of-code-go/pkg/list"
	"slices"

	"testing"
)

func Test_GetGrid(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  grid.Grid
	}{
		{
			name:  "single line",
			input: []string{".#X+[]"},
			want:  grid.Grid{[]rune{'.', '#', 'X', '+', '[', ']'}},
		},
		{
			name:  "two lines",
			input: []string{".'@", "QWE"},
			want:  grid.Grid{[]rune{'.', '\'', '@'}, []rune{'Q', 'W', 'E'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grid.GetGrid(tt.input); !list.ListOfListsAreEqual[rune](got, tt.want) {
				t.Errorf("getGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_VerticalSlice(t *testing.T) {
	tests := []struct {
		name  string
		input grid.Grid
		index int
		want  []rune
	}{
		{
			name:  "single line",
			input: grid.Grid{[]rune{'.', '#', 'X', '+', '[', ']'}},
			index: 0,
			want:  []rune{'.'},
		},
		{
			name:  "two lines",
			input: grid.Grid{[]rune{'.', '\'', '@'}, []rune{'Q', 'W', 'E'}},
			index: 1,
			want:  []rune{'\'', 'W'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grid.VerticalSlice(tt.input, tt.index); !slices.Equal[[]rune](got, tt.want) {
				t.Errorf("VerticalSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}