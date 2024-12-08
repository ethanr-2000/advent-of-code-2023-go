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
	antennas := getAntennas(g)

	allLocations := grid.AllLocations(g)
	count := 0
	for _, l := range allLocations {
		if isAntiNode(l, antennas) {
			count++
		}
	}

	return count
}

func part2(input string) int {
	parsed := parseInput(input)
	g := grid.GetGrid(parsed)

	return len(findAntiNodes(g))
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Antenna struct {
	Location grid.Location
	Value rune
}

func getAntennas(g grid.Grid) []Antenna {
	var antennas []Antenna
	for y, line := range g {
		for x, val := range line {
			if val != '.' {
				l := grid.Location{X: x, Y: y}
				antennas = append(antennas, Antenna{l, val})
			}
		}
	}
	return antennas
}

func isAntiNode(l grid.Location, antennas []Antenna) bool {
	for i, a1 := range antennas {
		diffToA1 := grid.LocationDiff(l, a1.Location)
		potentialDoubleLocation := grid.LocationAdd(a1.Location, diffToA1)

		potentialHalfLocation := grid.Location{X: 0, Y: 0}
		if diffToA1.X % 2 == 0 && diffToA1.Y % 2 == 0 {
			halfDiffToA1 := grid.Location{X: diffToA1.X / 2, Y: diffToA1.Y / 2}
			potentialHalfLocation = grid.LocationAdd(l, halfDiffToA1)
		}

		for _, a2 := range antennas[i+1:] {
			if a1.Value == a2.Value && (grid.LocationsEqual(a2.Location, potentialDoubleLocation) || grid.LocationsEqual(a2.Location, potentialHalfLocation)) {
				return true
			}
		}
	}
	return false
}

// Helper functions for part 2

func findAntiNodes(g grid.Grid) []grid.Location {
	antennas := getAntennas(g)
	antiNodeMap := make(map[string]grid.Location)
	var antiNodes []grid.Location
	for _, a1 := range antennas {
		for _, a2 := range antennas {
			if grid.LocationsEqual(a1.Location, a2.Location) || a1.Value != a2.Value {
				continue
			}
			a1a2Diff := grid.LocationDiff(a1.Location, a2.Location)
			antiNode := a1.Location
			for !grid.LocationOutsideGrid(antiNode, g) {
				h := grid.HashLocation(antiNode)
				_, locationIsAlreadyAntiNode := antiNodeMap[h]
				if (!locationIsAlreadyAntiNode) {
					antiNodes = append(antiNodes, antiNode)
					antiNodeMap[h] = antiNode
				}

				antiNode = grid.LocationAdd(antiNode, a1a2Diff)
			}
		}
	}
	return antiNodes
}
