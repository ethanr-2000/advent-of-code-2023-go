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

	grid := getGrid(parsed)

	positions := []Position{{Location{-1, 0}, RIGHT}}

	traverseGrid(grid, &positions)

	return len(getUniqueLocations(positions)) - 1
}

func part2(input string) int {
	parsed := parseInput(input)

	grid := getGrid(parsed)

	maxEnergy := 0
	for _, startPosition := range getStartPositions(grid) {
		positions := []Position{startPosition}
		traverseGrid(grid, &positions)
		energy := len(getUniqueLocations(positions)) - 1
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}

	return maxEnergy
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

type Location struct {
	x int
	y int
}

type Direction struct {
	x int
	y int
}

type Position struct {
	l Location
	d Direction
}

var UP Direction = Direction{0, -1}
var DOWN Direction = Direction{0, 1}
var LEFT Direction = Direction{-1, 0}
var RIGHT Direction = Direction{1, 0}

func traverseGrid(g Grid, positions *[]Position) {
	nextTile := ' '
	for !beenHereBefore(positions) {
		currentPosition := (*positions)[len(*positions)-1]
		currentDirection := currentPosition.d
		currentLocation := currentPosition.l

		nextTile = getNextTile(g, currentPosition)
		if nextTile == 'X' {
			return
		}

		nextLocation := Location{currentLocation.x + int(currentDirection.x), currentLocation.y + int(currentDirection.y)}
		nextDirection := currentDirection

		if nextTile == '/' {
			nextDirection = nextDirectionForwardSlash(currentDirection)
		} else if nextTile == '\\' {
			nextDirection = nextDirectionBackSlash(currentDirection)
		} else if nextTile == '-' {
			if currentDirection == UP || currentDirection == DOWN {
				*positions = append(*positions, Position{nextLocation, LEFT})
				traverseGrid(g, positions)

				nextDirection = RIGHT
			}
		} else if nextTile == '|' {
			if currentDirection == LEFT || currentDirection == RIGHT {
				*positions = append(*positions, Position{nextLocation, UP})
				traverseGrid(g, positions)

				nextDirection = DOWN
			}
		}
		*positions = append(*positions, Position{nextLocation, nextDirection})
	}
}

func getNextTile(g Grid, p Position) rune {
	l := p.l
	d := p.d

	if l.y+int(d.y) >= len(g) || l.y+int(d.y) < 0 || l.x+int(d.x) >= len(g[0]) || l.x+int(d.x) < 0 {
		return 'X'
	}

	return g[l.y+int(d.y)][l.x+int(d.x)]
}

func nextDirectionForwardSlash(d Direction) Direction {
	if d == UP {
		return RIGHT
	} else if d == RIGHT {
		return UP
	} else if d == DOWN {
		return LEFT
	} else if d == LEFT {
		return DOWN
	}
	return Direction{-100, -100}
}

func nextDirectionBackSlash(d Direction) Direction {
	if d == UP {
		return LEFT
	} else if d == RIGHT {
		return DOWN
	} else if d == DOWN {
		return RIGHT
	} else if d == LEFT {
		return UP
	}
	return Direction{-100, -100}
}

func beenHereBefore(positions *[]Position) bool {
	currentP := (*positions)[len(*positions)-1]

	for i := range (*positions)[:len(*positions)-1] {
		if (*positions)[i] == currentP {
			return true
		}
	}
	return false
}

func getUniqueLocations(positions []Position) []Position {
	seen := make(map[Location]bool)
	result := []Position{}

	for _, p := range positions {
		if _, exists := seen[p.l]; !exists {
			result = append(result, p)
			seen[p.l] = true
		}
	}

	return result
}

// Helper functions for part 2

func getStartPositions(g Grid) []Position {
	startingPositions := []Position{}
	for x := range g[0] {
		startingPositions = append(startingPositions, Position{Location{x, -1}, DOWN})
		startingPositions = append(startingPositions, Position{Location{x, len(g)}, UP})
	}

	for y := range g {
		startingPositions = append(startingPositions, Position{Location{-1, y}, RIGHT})
		startingPositions = append(startingPositions, Position{Location{len(g[0]), y}, LEFT})
	}
	return startingPositions
}
