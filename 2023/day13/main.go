package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		ans := part1(input)
		fmt.Println("Running part 1")
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(input)
		fmt.Println("Running part 2")
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		fmt.Println("Running all")
		ans1 := part1(input)
		fmt.Println("Part 1 Output:", ans1)
		ans2 := part2(input)
		fmt.Println("Part 2 Output:", ans2)
	}
}

func part1(input string) int {
	parsed := parseInput(input)

	grids := getGrids(parsed)

	sum := 0
	for _, g := range grids {
		sum += 100*findHorizontalReflectionLines(g, 0) + findVerticalReflectionLines(g, 0)
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	grids := getGrids(parsed)

	sum := 0
	for _, g := range grids {
		sum += 100*findHorizontalReflectionLines(g, 1) + findVerticalReflectionLines(g, 1)
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func getGrids(input []string) []Grid {
	grids := []Grid{}
	g := Grid{}
	for _, line := range input {
		if line == "" {
			grids = append(grids, g)
			g = Grid{}
		} else {
			g = append(g, []rune(line))
		}
	}
	return append(grids, g) // add the last one
}

// Helper functions for part 1

type Grid [][]rune

func findHorizontalReflectionLines(g Grid, smudges int) int {
	for i := 1; i < len(g); i++ {
		if doesReflectHorizontal(g, i, smudges) {
			return i
		}
	}
	return 0
}

func doesReflectHorizontal(g Grid, i int, smudges int) bool {
	numDiffs := 0
	for line := range g[:i] {
		if i+line == len(g) {
			return numDiffs == smudges
		}
		numDiffs += getDifferences(g[i+line], g[i-line-1])
		if numDiffs > smudges {
			return false
		}
	}
	return numDiffs == smudges
}

func verticalSlice(g Grid, i int) []rune {
	s := make([]rune, len(g))
	for row := range g {
		s[row] = g[row][i]
	}
	return s
}

func findVerticalReflectionLines(g Grid, smudges int) int {
	for i := 1; i < len(g[0]); i++ {
		if doesReflectVertical(g, i, smudges) {
			return i
		}
	}
	return 0
}

func doesReflectVertical(g Grid, i int, smudges int) bool {
	numDiffs := 0
	for line := range g[0][:i] {
		if i+line == len(g[0]) {
			return numDiffs == smudges
		}
		numDiffs += getDifferences(verticalSlice(g, i+line), verticalSlice(g, i-line-1))
		if numDiffs > smudges {
			return false
		}
	}
	return numDiffs == smudges
}

// Helper functions for part 2

// assume same length
func getDifferences(s1, s2 []rune) int {
	numDiffs := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			numDiffs++
		}
	}
	return numDiffs
}
