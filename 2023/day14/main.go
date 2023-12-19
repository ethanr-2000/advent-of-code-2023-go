package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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

	grid := getGrid(parsed)
	tiltGrid(grid, North)

	return calculateLoadNorth(grid)
}

func part2(input string) int {
	parsed := parseInput(input)

	grid := getGrid(parsed)

	gridMap := make(map[string]int, 0)
	targetCycles := 1000000000
	for cycles := 0; cycles < targetCycles; cycles++ {
		tiltGrid(grid, North)
		tiltGrid(grid, West)
		tiltGrid(grid, South)
		tiltGrid(grid, East)

		hash := hashGrid(grid)
		if cycleStart, exists := gridMap[hash]; exists {
			cycleLength := (cycles + 1) - cycleStart
			finalStateIndex := ((targetCycles - cycleStart) % cycleLength) + cycleStart

			finalGrid := getGrid(parseInput(findKeyByValue(gridMap, finalStateIndex)))
			return calculateLoadNorth(finalGrid)
		} else {
			gridMap[hash] = cycles + 1
		}
	}
	return calculateLoadNorth(grid)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func getGrid(input []string) Grid {
	g := Grid{}
	for _, line := range input {
		g = append(g, []rune(line))
	}
	return g
}

type Grid [][]rune

func verticalSlice(g Grid, i int) []rune {
	s := make([]rune, len(g))
	for row := range g {
		s[row] = g[row][i]
	}
	return s
}

// Helper functions for part 2

type Direction int

const (
	North Direction = 0
	West  Direction = 1
	South Direction = 2
	East  Direction = 3
)

func tiltGrid(g Grid, direction Direction) {
	switch direction {
	case North:
		for column := range g[0] {
			tilted := tiltSlice(verticalSlice(g, column), false)

			for y := range tilted {
				g[y][column] = tilted[y]
			}
		}
	case West:
		for row := range g {
			g[row] = tiltSlice(g[row], false)
		}
	case South:
		for column := range g[0] {
			tilted := tiltSlice(verticalSlice(g, column), true)

			for y := range tilted {
				g[y][column] = tilted[y]
			}
		}
	case East:
		for row := range g {
			g[row] = tiltSlice(g[row], true)
		}
	default:
		fmt.Println("unknown direction")
	}
}

func calculateLoadNorth(g Grid) int {
	totalLoad := 0
	for column := range g[0] {
		totalLoad += calculateLoadOnSlice(verticalSlice(g, column))
	}
	return totalLoad
}

// tilts a list
func tiltSlice(s []rune, down bool) []rune {
	if down {
		slices.Reverse[[]rune](s)
	}

	lastCubeIndex := -1
	for i := range s {
		if s[i] == '#' {
			lastCubeIndex = i
			continue
		}
		if s[i] == 'O' {
			if i-lastCubeIndex > 1 { // if it's not already settled
				s[i] = '.'
				s[lastCubeIndex+1] = 'O'
			}
			lastCubeIndex++
		}
	}

	if down {
		slices.Reverse[[]rune](s)
	}
	return s
}

func calculateLoadOnSlice(s []rune) int {
	totalLoad := 0
	for i := range s {
		if s[i] == 'O' {
			totalLoad += len(s) - i
		}
	}
	return totalLoad
}

func hashGrid(g Grid) string {
	// one for each cell, plus a new line for each row, except the last
	hash := make([]rune, (len(g)+1)*len(g[0])-1)
	for y, row := range g {
		startIndex := len(row)*y + y
		endIndex := len(row)*(y+1) + y
		copy(hash[startIndex:endIndex], row)

		if endIndex != len(hash) {
			hash[endIndex] = '\n'
		}
	}
	return string(hash)
}

func findKeyByValue(m map[string]int, value int) string {
	for key, val := range m {
		if val == value {
			return key
		}
	}
	return "" // Not found
}
