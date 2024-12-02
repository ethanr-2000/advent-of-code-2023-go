package main

import (
	"advent-of-code-go/pkg/regex"

	_ "embed"
	"flag"
	"fmt"
	"strings"

	"advent-of-code-go/pkg/list"

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

	totalSafe := 0
	for _, line := range parsed {
		levels := getLevels(line)
		if isIncreasingSafely(levels) || isDecreasingSafely(levels) {
			totalSafe++
		}
	}
	return totalSafe
}

func part2(input string) int {
	parsed := parseInput(input)

	totalSafe := 0
	for _, line := range parsed {
		levels := getLevels(line)
		if isSafeWithDamping(levels) {
			totalSafe++
		}
	}
	return totalSafe
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func getLevels(line string) []int {
	return regex.GetSpaceSeparatedNumbers(line)
}

func isIncreasingSafely(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i] <= levels[i-1] {
			return false
		}
		if absInt(levels[i]-levels[i-1]) > 3 {
			return false
		}
	}
	return true
}

func isDecreasingSafely(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i] >= levels[i-1] {
			return false
		}
		if absInt(levels[i]-levels[i-1]) > 3 {
			return false
		}
	}
	return true
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// Helper functions for part 2

func isSafeWithDamping(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		levelsWithRemoved := list.DeleteAtIndices(levels, []int{i})
		if isIncreasingSafely(levelsWithRemoved) || isDecreasingSafely(levelsWithRemoved) {
			return true
		}
	}
	return false
}
