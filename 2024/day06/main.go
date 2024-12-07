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
	m := parseInput(input)
	g := getInitialGuard(m)

	errorGuard := Guard{Location{-1, -1}, North}
	for g != errorGuard {
		prevG := g
		g = moveGuard(m, prevG)
		if g.location != prevG.location {
			m[prevG.location.y][prevG.location.x] = 'X'
		}

		// prettyPrint(m)
	}

	return countX(m)
}

func part2(input string) int {
	m := parseInput(input)
	g := getInitialGuard(m)
	
	var possibleObstructionLocations []Location
	for y, line := range m {
		for x, value := range line {
			if value == '.' {
				possibleObstructionLocations = append(possibleObstructionLocations, Location{x, y})
			}
		}
	}

	count := 0
	for _, l := range possibleObstructionLocations {
		m[l.y][l.x] = '#'

		if guardCannotEscape(m, g) {
			count++
		}

		m[l.y][l.x] = '.'
	}

	return count
}

func parseInput(input string) [][]rune {
	strList := strings.Split(input, "\n")
	var m [][]rune
	for _, str := range strList {
		line := []rune(str)
		m = append(m, line)
	}
	return m
}

// Helper functions for part 1


func prettyPrint(grid [][]rune) {
	for _, row := range grid {
		for _, r := range row {
			fmt.Printf("%c ", r) // Print each rune with a space
		}
		fmt.Println() // Newline after each row
	}
}

type Location struct {
	x int
	y int
}

type Direction int

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

type Guard struct {
	location Location
	direction Direction
}

func getInitialGuard(m [][]rune) Guard {
	for y, line := range m {
		for x, value := range line {
			if value == '^' {
				return Guard{Location{x, y}, North}
			}
		}
	}
	return Guard{Location{-1, -1}, North}
}

func moveGuard(m [][]rune, guard Guard) Guard {
	var nextLocation Location
	if (guard.direction == North) {
		nextLocation = Location{guard.location.x, guard.location.y - 1}
	}
	if (guard.direction == East) {
		nextLocation = Location{guard.location.x + 1, guard.location.y}
	}
	if (guard.direction == South) {
		nextLocation = Location{guard.location.x, guard.location.y + 1}
	}
	if (guard.direction == West) {
		nextLocation = Location{guard.location.x - 1, guard.location.y}
	}

	if nextLocation.x < 0 || nextLocation.x >= len(m[0]) || nextLocation.y < 0 || nextLocation.y >= len(m) {
		return Guard{Location{-1, -1}, North}
	}

	if (m[nextLocation.y][nextLocation.x] == '#') {
		return Guard{guard.location, turnRight(guard.direction)}
	}

	return Guard{nextLocation, guard.direction}
}

func turnRight(d Direction) Direction {
	return (d + 1) % 4
}

func countX(m [][]rune) int {
	count := 0
	for _, line := range m {
		for _, value := range line {
			if value == 'X' {
				count++
			}
		}
	}
	return count
}

// Helper functions for part 2

func guardCannotEscape(m [][]rune, g Guard) bool {
	errorGuard := Guard{Location{-1, -1}, North}

	var guardHistory []Guard
	// initialG := g
	for (g != errorGuard && !guardInGuards(guardHistory, g)) {
		prevG := g
		g = moveGuard(m, prevG)

		guardHistory = append(guardHistory, prevG)
		m[prevG.location.y][prevG.location.x] = 'X'

		// prettyPrint(m)/
	}
	// if error guard was returned, it means they escaped
	return g != errorGuard
}

func locationsEqual(a, b Location) bool {
	return a.x == b.x && a.y == b.y
}

func guardInGuards(guards []Guard, target Guard) bool {
	for _, g := range guards {
		if g.direction == target.direction && locationsEqual(g.location, target.location) {
			return true
		}
	}
	return false
}

// func locationInLocations(locations []Location, target Location) bool {
// 	for _, l := range locations {
// 		if locationsEqual(l, target) {
// 			return true
// 		}
// 	}
// 	return false
// }
