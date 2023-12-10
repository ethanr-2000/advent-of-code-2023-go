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

	l := traverseMap(parsed)

	return len(l) / 2
}

func part2(input string) int {
	parsed := parseInput(input)
	loop := traverseMap(parsed)

	return countSpacesWithinLoop(parsed, loop)
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
	"|": {{-1, 0}, {1, 0}},  // above or below
	"-": {{0, -1}, {0, 1}},  // left or right
	"J": {{-1, 0}, {0, -1}}, // above or left
	"F": {{1, 0}, {0, 1}},   // below or right
	"7": {{1, 0}, {0, -1}},  // below or left
	"L": {{-1, 0}, {0, 1}},  // above or right
}

func getNextLocation(pipe string, current, previous Location) Location {
	possibleSteps := PipeConnections[pipe]

	for _, step := range possibleSteps {
		possibleNextLocation := Location{
			x: current.x + step[1],
			y: current.y + step[0],
		}
		if !sameLocation(possibleNextLocation, previous) {
			return possibleNextLocation
		}
	}
	return Location{-1, -1}
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

// func getNextStepFromStart(m Grid2D, start Location) Location {
// 	possibleMoves := [][2]int{
// 		{-1, 0},
// 		{1, 0},
// 		{0, -1},
// 		{0, 1},
// 	}

// 	for _, move := range possibleMoves {
// 		possibleNextLocation := Location{
// 			x: start.x + move[1],
// 			y: start.y + move[0],
// 		}
// 		possiblePipe := m[possibleNextLocation.y][possibleNextLocation.x]

// 		for _, connection := range PipeConnections[possiblePipe] {
// 			if (connection[1]+move[1]) == 0 && (connection[0]+move[0]) == 0 {
// 				return possibleNextLocation
// 			}
// 		}
// 	}

// 	return Location{-1, -1}
// }

func replaceSWithPipe(m Grid2D, start Location) {
	possibleMoves := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	correctMoves := [][2]int{}
	for _, move := range possibleMoves {
		possibleNextLocation := Location{
			x: start.x + move[1],
			y: start.y + move[0],
		}
		possiblePipe := m[possibleNextLocation.y][possibleNextLocation.x]

		for _, connection := range PipeConnections[possiblePipe] {
			if (connection[1]+move[1]) == 0 && (connection[0]+move[0]) == 0 {
				correctMoves = append(correctMoves, connection)
			}
		}
	}

	for k, v := range PipeConnections {
		if listOf2IntEqual(v[0], correctMoves[0]) && listOf2IntEqual(v[1], correctMoves[1]) {
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

func traverseMap(m Grid2D) []Location {
	startingLocation := getStartingLocation(m)
	replaceSWithPipe(m, startingLocation)

	locations := []Location{startingLocation}

	for !sameLocation(lastValue(locations), startingLocation) || len(locations) == 1 {
		currentLocation := lastValue[Location](locations)

		currentPipe := m[currentLocation.y][currentLocation.x]
		locations = append(locations, getNextLocation(currentPipe, currentLocation, locations[len(locations)-2]))
	}
	return locations
}

// Helper functions for part 2

func countSpacesWithinLoop(grid Grid2D, loop []Location) int {
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
		if l1.x == l2.x && l1.y == l2.y {
			return true
		}
	}
	return false
}

// func padGrid(grid Grid2D, padding string) Grid2D {
// 	paddedGrid := make(Grid2D, len(grid)+2)

// 	paddedGrid[0] = padList(len(grid[0])+2, padding)
// 	paddedGrid[len(paddedGrid)-1] = padList(len(grid[0])+2, padding)

// 	for i, row := range grid {
// 		paddedGrid[i+1] = padListString(row, padding)
// 	}

// 	return paddedGrid
// }

// func padListString(list []string, padding string) []string {
// 	paddedList := make([]string, len(list)+2)
// 	paddedList[0] = padding

// 	for i, val := range list {
// 		paddedList[i+1] = val
// 	}

// 	paddedList[len(paddedList)-1] = padding
// 	return paddedList
// }

// func padList(length int, padding string) []string {
// 	paddedList := make([]string, length)
// 	for i := range paddedList {
// 		paddedList[i] = padding
// 	}
// 	return paddedList
// }

// func countFalseValuesInBoolGrid(grid [][]bool) int {
// 	count := 0

// 	for _, row := range grid {
// 		for _, value := range row {
// 			if !value {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

// func countSpacesWithinLoop(grid Grid2D, loop []Location) int {
// 	paddedGrid := padGrid(grid, ".")

// 	visited := make([][]bool, len(paddedGrid))
// 	for i := range visited {
// 		visited[i] = make([]bool, len(paddedGrid[0]))
// 	}

// 	dfs(paddedGrid, 0, 0, loop, visited)

// 	return countFalseValuesInBoolGrid(visited)
// }

// func dfs(grid Grid2D, x, y int, loop []Location, visited [][]bool) {
// 	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) || visited[y][x] || locationInList(Location{x, y}, loop) {
// 		return
// 	}

// 	visited[y][x] = true

// 	dfs(grid, x+1, y, loop, visited)
// 	dfs(grid, x-1, y, loop, visited)
// 	dfs(grid, x, y+1, loop, visited)
// 	dfs(grid, x, y-1, loop, visited)
// }
