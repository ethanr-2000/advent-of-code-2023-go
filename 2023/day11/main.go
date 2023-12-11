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
	universe := parseInput(input)

	locations := findGalaxies(universe)

	return sumDistances(universe, locations, 1)
}

func part2(input string) int {
	universe := parseInput(input)

	locations := findGalaxies(universe)

	return sumDistances(universe, locations, 999999)
}

func parseInput(input string) [][]rune {
	strList := strings.Split(input, "\n")
	var universe [][]rune
	for _, str := range strList {
		line := []rune(str)
		universe = append(universe, line)
	}
	return universe
}

// Helper functions for part 1

type Location struct {
	x int
	y int
}

func intBetweenInts(a, b, t int) bool {
	return (a <= t && t <= b) || (b <= t && t <= a)
}

func findEmptyHorizontalLines(universe [][]rune) []int {
	ys := []int{}

	for y, line := range universe {
		if allDots(line) {
			ys = append(ys, y)
		}
	}

	return ys
}

func findEmptyVerticalLines(universe [][]rune) []int {
	xs := []int{}

	for x := 0; x < len(universe[0]); x++ {
		empty := true
		for y := 0; y < len(universe); y++ {
			if universe[y][x] == '#' {
				empty = false
			}
		}
		if empty {
			xs = append(xs, x)
		}
	}

	return xs
}

func allDots(l []rune) bool {
	for _, r := range l {
		if r != '.' {
			return false
		}
	}
	return true
}

func findGalaxies(universe [][]rune) []Location {
	locations := []Location{}
	for y, line := range universe {
		for x, val := range line {
			if val == '#' {
				locations = append(locations, Location{x, y})
			}
		}
	}
	return locations
}

func getManhattanDistance(l1, l2 Location) int {
	return absDiff(l1.x, l2.x) + absDiff(l1.y, l2.y)
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func sumDistances(universe [][]rune, locations []Location, expansion int) int {
	emptyHorizontalLines := findEmptyHorizontalLines(universe)
	emptyVerticalLines := findEmptyVerticalLines(universe)

	sum := 0
	for _, l1 := range locations {
		for _, l2 := range locations {
			sum += getManhattanDistance(l1, l2)

			for _, y := range emptyHorizontalLines {
				if intBetweenInts(l1.y, l2.y, y) {
					sum += expansion
				}
			}
			for _, x := range emptyVerticalLines {
				if intBetweenInts(l1.x, l2.x, x) {
					sum += expansion
				}
			}
		}
	}
	return sum / 2
}
