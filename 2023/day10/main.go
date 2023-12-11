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
	grid := parseInput(input)
	l := traverseGrid(grid)

	return len(l) / 2
}

func part2(input string) int {
	grid := parseInput(input)
	loop := traverseGrid(grid)

	return countSpacesWithinLoop(grid, loop)
}

func parseInput(input string) Grid2D {
	split := strings.Split(input, "\n")

	var grid Grid2D
	for _, s := range split {
		grid = append(grid, strings.Split(s, ""))
	}
	return grid
}

// Helper functions for part 1

type Grid2D [][]string

type Location struct {
	x int
	y int
}

var PipeConnections = map[string][2][2]int{
	// these are defined with y, x
	"|": {{-1, 0}, {1, 0}},  // above or below
	"-": {{0, -1}, {0, 1}},  // left or right
	"J": {{-1, 0}, {0, -1}}, // above or left
	"F": {{1, 0}, {0, 1}},   // below or right
	"7": {{1, 0}, {0, -1}},  // below or left
	"L": {{-1, 0}, {0, 1}},  // above or right
	".": {{0, 0}, {0, 0}},   // no connection
}

func getNextLocations(pipe string, current Location) []Location {
	possibleSteps := PipeConnections[pipe]

	l := []Location{}
	for _, step := range possibleSteps {
		possibleNextLocation := Location{
			x: current.x + step[1],
			y: current.y + step[0],
		}
		l = append(l, possibleNextLocation)
	}
	return l
}

func sameLocation(l1, l2 Location) bool {
	return l1.x == l2.x && l1.y == l2.y
}

func getStartingLocation(m Grid2D) Location {
	for y, line := range m {
		for x, val := range line {
			if val == "S" {
				return Location{x, y}
			}
		}
	}
	return Location{-1, -1}
}

func locationOutsideGrid(l Location, m Grid2D) bool {
	return l.x < 0 || l.x == len(m[0]) || l.y < 0 || l.y == len(m)
}

func replaceSWithPipe(m Grid2D, start Location) {
	possibleMoves := [][2]int{
		{1, 0},
		{0, 1},
		{0, -1},
		{-1, 0},
	}

	correctMoves := [][2]int{}
	for _, move := range possibleMoves {
		possibleNextLocation := Location{
			x: start.x + move[1],
			y: start.y + move[0],
		}

		// outside of grid
		if locationOutsideGrid(possibleNextLocation, m) {
			continue
		}

		possibleNextPipe := m[possibleNextLocation.y][possibleNextLocation.x]
		for _, connection := range PipeConnections[possibleNextPipe] {
			if (connection[1]+move[1]) == 0 && (connection[0]+move[0]) == 0 {
				correctMoves = append(correctMoves, move)
			}
		}
	}

	for k, v := range PipeConnections {
		if listOf2IntEqual(v[0], correctMoves[0]) && listOf2IntEqual(v[1], correctMoves[1]) ||
			listOf2IntEqual(v[1], correctMoves[0]) && listOf2IntEqual(v[0], correctMoves[1]) {
			m[start.y][start.x] = k
		}
	}
}

func listOf2IntEqual(l1, l2 [2]int) bool {
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func lastValue[T any](n []T) T {
	return n[len(n)-1]
}

func traverseGrid(m Grid2D) []Location {
	startingLocation := getStartingLocation(m)
	replaceSWithPipe(m, startingLocation)

	locations := []Location{startingLocation}

	startingPipe := m[startingLocation.y][startingLocation.x]
	locations = append(locations, getNextLocations(startingPipe, locations[0])[0])

	for !sameLocation(lastValue(locations), startingLocation) || len(locations) == 1 {
		currentLocation := lastValue[Location](locations)
		previousLocation := locations[len(locations)-2]

		currentPipe := m[currentLocation.y][currentLocation.x]

		possibleNextLocations := getNextLocations(currentPipe, currentLocation)
		for _, l := range possibleNextLocations {
			// get the next location that isn't the previous one
			if !sameLocation(l, previousLocation) {
				locations = append(locations, l)
			}
		}
	}
	return locations
}

// Helper functions for part 2

func countSpacesWithinLoop(grid Grid2D, loop []Location) int {
	// scan line by line
	// start "outside the loop"
	// if in we're in the loop and on a tile, increment
	// when we hit the loop, certain conditions flip in/out

	count := 0
	for y, line := range grid {
		inLoop := false
		x := 0
		for x < len(line) {
			if !locationInList(Location{x, y}, loop) {
				if inLoop {
					count++
				}
			} else if grid[y][x] == "|" {
				inLoop = !inLoop
			} else {
				start := grid[y][x]
				x++
				for grid[y][x] == "-" {
					// horizontal pipes don't bring us in/out
					x++
				}
				end := grid[y][x]
				if (start == "L" && end == "7") || (start == "F" && end == "J") {
					inLoop = !inLoop
				}
			}
			x++
		}
	}
	return count
}

func locationInList(l1 Location, ls []Location) bool {
	for _, l2 := range ls {
		if sameLocation(l1, l2) {
			return true
		}
	}
	return false
}
