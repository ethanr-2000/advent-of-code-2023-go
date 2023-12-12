// package main

// import (
// 	_ "embed"
// 	"slices"
// 	"testing"
// )

// var example1 = `???.### 1,1,3
// .??..??...?##. 1,1,3
// ?#?#?#?#?#?#?#? 1,3,1,6
// ????.#...#... 4,1,1
// ????.######..#####. 1,6,5
// ?###???????? 3,2,1`

// func Test_part1(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input string
// 		want  int
// 	}{
// 		{
// 			name:  "example",
// 			input: example1,
// 			want:  21,
// 		},
// 		// {
// 		// 	name:  "actual",
// 		// 	input: input,
// 		// 	want:  0,
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := part1(tt.input); got != tt.want {
// 				t.Errorf("part1() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// var example2 = `0`

// func Test_part2(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input string
// 		want  int
// 	}{
// 		{
// 			name:  "example",
// 			input: example2,
// 			want:  0,
// 		},
// 		// {
// 		// 	name:  "actual",
// 		// 	input: input,
// 		// 	want:  0,
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := part2(tt.input); got != tt.want {
// 				t.Errorf("part2() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_deleteIndicesOfSlice(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   []string
// 		indices []int
// 		want    []string
// 	}{
// 		{
// 			name:    "example",
// 			input:   []string{"A", "B", "C"},
// 			indices: []int{0, 2},
// 			want:    []string{"B"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := deleteIndicesOfSlice[string](tt.input, tt.indices); slices.Compare[[]string](got, tt.want) != 0 {
// 				t.Errorf("deleteIndicesOfSlice() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_allSpringsPlaced(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input ConditionRecord
// 		want  bool
// 	}{
// 		{
// 			name: "not enough",
// 			input: ConditionRecord{
// 				record: "?#?#?#?#?#?#?#?",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "just right",
// 			input: ConditionRecord{
// 				record: "##########???#?",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "too many",
// 			input: ConditionRecord{
// 				record: "##########??##?",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := allSpringsPlaced(tt.input); got != tt.want {
// 				t.Errorf("allSpringsPlaced() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_isValidRecord(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input ConditionRecord
// 		want  bool
// 	}{
// 		{
// 			name: "valid",
// 			input: ConditionRecord{
// 				record: "?#?###?#?######",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "invalid",
// 			input: ConditionRecord{
// 				record: "##########???#?",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "almost valid",
// 			input: ConditionRecord{
// 				record: "?###?#?#?######",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := isValidRecord(tt.input); got != tt.want {
// 				t.Errorf("isValidRecord() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func compareConditionRecords(c1, c2 ConditionRecord) int {
// 	if c1.record != c2.record || slices.Compare[[]int](c1.groups, c2.groups) != 0 {
// 		return 1
// 	}
// 	return 0
// }

// func Test_uniqueRecords(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input []ConditionRecord
// 		want  []ConditionRecord
// 	}{
// 		{
// 			name: "two the same",
// 			input: []ConditionRecord{
// 				{
// 					record: "?#?###?#?######",
// 					groups: []int{1, 3, 1, 6},
// 				},
// 				{
// 					record: "?#?###?#?######",
// 					groups: []int{1, 3, 1, 6},
// 				},
// 			},
// 			want: []ConditionRecord{
// 				{
// 					record: "?#?###?#?######",
// 					groups: []int{1, 3, 1, 6},
// 				},
// 			},
// 		},
// 		{
// 			name: "two different",
// 			input: []ConditionRecord{
// 				{
// 					record: "?#?###?#?######",
// 					groups: []int{1, 3, 1, 6},
// 				},
// 				{
// 					record: "?#?###?#?###???",
// 					groups: []int{1, 3, 1, 6},
// 				},
// 			},
// 			want: []ConditionRecord{
// 				{
// 					record: "?#?###?#?######",
// 					groups: []int{1, 3, 1, 6},
// 				},
// 				{
// 					record: "?#?###?#?###???",
// 					groups: []int{1, 3, 1, 6},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := uniqueRecords(tt.input)
// 			if !slices.EqualFunc[[]ConditionRecord](got, tt.want, conditionRecordsSame) {
// 				t.Errorf("uniqueRecords() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_invalidSoFar(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input ConditionRecord
// 		want  bool
// 	}{
// 		{
// 			name: "valid",
// 			input: ConditionRecord{
// 				record: "?#?###?#?######",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "valid so far",
// 			input: ConditionRecord{
// 				record: "?#?#?#?#?##??##",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "invalid so far",
// 			input: ConditionRecord{
// 				record: "?###?#?#?##??##",
// 				groups: []int{1, 3, 1, 6},
// 			},
// 			want: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := invalidSoFar(tt.input); got != tt.want {
// 				t.Errorf("invalidSoFar() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
