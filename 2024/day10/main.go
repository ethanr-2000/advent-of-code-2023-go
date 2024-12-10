package main

import (
	"advent-of-code-go/pkg/grid"
	"strconv"

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

	g := grid.GetGrid(parsed)

	locations := grid.AllLocations(g)

	totalScore := 0
	for _, l := range locations {
		if grid.ValueAtLocation(g, l) == '0' {
			totalScore += len(trailHeadScore(g, l, []grid.Location{}))
		}
	}

	return totalScore
}

func part2(input string) int {
	parsed := parseInput(input)

	g := grid.GetGrid(parsed)

	locations := grid.AllLocations(g)

	totalScore := 0
	for _, l := range locations {
		if grid.ValueAtLocation(g, l) == '0' {
			totalScore += trailHeadRating(g, l, 0)
		}
	}

	return totalScore
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func trailHeadScore(g grid.Grid, l grid.Location, nines []grid.Location) []grid.Location {
	value := grid.ValueAtLocation(g, l)

	if value == '9' {
		if !grid.LocationInList(l, nines) {
			return append(nines, l)
		}
	}

	north := grid.MoveStepsInDirection(l, grid.North, 1)
	east := grid.MoveStepsInDirection(l, grid.East, 1)
	south := grid.MoveStepsInDirection(l, grid.South, 1)
	west := grid.MoveStepsInDirection(l, grid.West, 1)

	for _, lToCheck := range []grid.Location{north, east, south, west} {
		if canMoveToLocation(g, lToCheck, value) {
			nines = trailHeadScore(g, lToCheck, nines)
		}
	}

	return nines
}

func canMoveToLocation(g grid.Grid, l grid.Location, currentValue rune) bool {
	return !grid.LocationOutsideGrid(l, g) && isOneUp(currentValue, grid.ValueAtLocation(g, l))
}

func isOneUp(r1, r2 rune) bool {
	n1, _ := strconv.Atoi(string(r1))
	n2, _ := strconv.Atoi(string(r2))

	return n2-n1 == 1
}

// Helper functions for part 2

func trailHeadRating(g grid.Grid, l grid.Location, score int) int {
	value := grid.ValueAtLocation(g, l)

	if value == '9' {
		return score + 1
	}

	north := grid.MoveStepsInDirection(l, grid.North, 1)
	east := grid.MoveStepsInDirection(l, grid.East, 1)
	south := grid.MoveStepsInDirection(l, grid.South, 1)
	west := grid.MoveStepsInDirection(l, grid.West, 1)

	for _, lToCheck := range []grid.Location{north, east, south, west} {
		if canMoveToLocation(g, lToCheck, value) {
			score = trailHeadRating(g, lToCheck, score)
		}
	}

	return score
}
