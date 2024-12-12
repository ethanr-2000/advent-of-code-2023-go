package main

import (
	"advent-of-code-go/pkg/grid"

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

	totalPrice := 0
	alreadyFenced := []grid.Location{}
	for _, l := range grid.AllLocations(g) {
		if !grid.LocationInList(l, alreadyFenced) {
			newlyFenced := fenceRegion(g, l, []grid.Location{})
			alreadyFenced = append(alreadyFenced, newlyFenced...)
			totalPrice += fencePrice(newlyFenced)
		}
	}

	return totalPrice
}

func part2(input string) int {
	parsed := parseInput(input)
	g := grid.GetGrid(parsed)

	totalPrice := 0
	alreadyFenced := []grid.Location{}
	for _, l := range grid.AllLocations(g) {
		if !grid.LocationInList(l, alreadyFenced) {
			newlyFenced := fenceRegion(g, l, []grid.Location{})
			alreadyFenced = append(alreadyFenced, newlyFenced...)
			totalPrice += fencePriceSides(newlyFenced)
		}
	}

	return totalPrice
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

// returns locationsFenced
// assumes the starting point is not already a fenced region
func fenceRegion(g grid.Grid, l grid.Location, alreadyFenced []grid.Location) []grid.Location {
	value := grid.ValueAtLocation(g, l)

	alreadyFenced = append(alreadyFenced, l)

	for _, lToCheck := range grid.FourAdjacentList(l) {
		if !grid.LocationInList(lToCheck, alreadyFenced) && locationIsPartOfRegion(g, lToCheck, value) {
			alreadyFenced = fenceRegion(g, lToCheck, alreadyFenced)
		}
	}

	return alreadyFenced
}

func fencePrice(ls []grid.Location) int {
	area := len(ls)

	perimeter := 0
	for _, l := range ls {
		adj := numberOfAdjacent(l, ls)

		perimeter += 4 - adj
	}
	return area * perimeter
}

func numberOfAdjacent(l grid.Location, ls []grid.Location) int {
	adjacentCount := 0
	for _, lToCheck := range grid.FourAdjacentList(l) {
		if grid.LocationInList(lToCheck, ls) {
			adjacentCount++
		}
	}
	return adjacentCount
}

func locationIsPartOfRegion(g grid.Grid, l grid.Location, currentValue rune) bool {
	return !grid.LocationOutsideGrid(l, g) && currentValue == grid.ValueAtLocation(g, l)
}

// Helper functions for part 2

func fencePriceSides(ls []grid.Location) int {
	area := len(ls)

	sides := 0
	for _, l := range ls {
		north, northEast, east, southEast, south, southWest, west, northWest := grid.EightAdjacent(l)

		if isCorner(north, west, northWest, ls) {
			sides++
		}
		if isCorner(south, west, southWest, ls) {
			sides++
		}
		if isCorner(south, east, southEast, ls) {
			sides++
		}
		if isCorner(north, east, northEast, ls) {
			sides++
		}
	}
	return area * sides
}

func isCorner(l1, l2, l12 grid.Location, ls []grid.Location) bool {
	return !(grid.LocationInList(l1, ls) || grid.LocationInList(l2, ls)) ||
		(grid.LocationInList(l1, ls) &&
			grid.LocationInList(l2, ls) &&
			!grid.LocationInList(l12, ls))
}
