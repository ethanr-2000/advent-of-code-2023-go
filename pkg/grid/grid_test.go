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

func Test_OppositeDirections(t *testing.T) {
	tests := []struct {
		name  string
		d1 grid.Direction
		d2 grid.Direction
		want  bool
	}{
		{
			name: "east west",
			d1: grid.East,
			d2: grid.West,
			want: true,
		},
		{
			name: "west east",
			d1: grid.West,
			d2: grid.East,
			want: true,
		},
		{
			name: "north south",
			d1: grid.North,
			d2: grid.South,
			want: true,
		},
		{
			name: "north north",
			d1: grid.North,
			d2: grid.North,
			want: false,
		},
		{
			name: "west west",
			d1: grid.West,
			d2: grid.West,
			want: false,
		},
		{
			name: "north east",
			d1: grid.North,
			d2: grid.East,
			want: false,
		},
		{
			name: "west south",
			d1: grid.West,
			d2: grid.South,
			want: false,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grid.OppositeDirections(tt.d1, tt.d2); got != tt.want {
				t.Errorf("OppositeDirections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DirectionBetweenLocations(t *testing.T) {
	tests := []struct {
		name  string
		l1 grid.Location
		l2 grid.Location
		want  grid.Direction
	}{
		{
			name: "east",
			l1: grid.Location{10, 10},
			l2: grid.Location{11, 10},
			want: grid.East,
		},
		{
			name: "west",
			l1: grid.Location{11, 10},
			l2: grid.Location{10, 10},
			want: grid.West,
		},
		{
			name: "south",
			l1: grid.Location{10, 9},
			l2: grid.Location{10, 10},
			want: grid.South,
		},
		{
			name: "north",
			l1: grid.Location{10, 10},
			l2: grid.Location{10, 9},
			want: grid.North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grid.DirectionBetweenLocations(tt.l1, tt.l2); got != tt.want {
				t.Errorf("DirectionBetweenLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}
