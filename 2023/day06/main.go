package main

import (
	"advent-of-code-go/pkg/cast"
	"advent-of-code-go/pkg/regex"
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

	var numbers [][]int
	for _, line := range parsed {
		numbers = append(numbers, regex.GetSpaceSeparatedNumbers(line))
	}

	result := 1
	// for each race
	for i := range numbers[0] {
		r := calculateNumberOfWaysToWin(numbers[0][i], numbers[1][i])
		result *= r
	}

	return result
}

func part2(input string) int {
	parsed := parseInput(input)

	var numbers []int
	for _, line := range parsed {
		joinedNums := strings.Join(cast.IntArrayToStringArray(regex.GetSpaceSeparatedNumbers(line)), "")
		numbers = append(numbers, cast.ToInt(joinedNums))
	}

	return calculateNumberOfWaysToWin(numbers[0], numbers[1])
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func calculateDistanceGivenMsHeld(msHeld int, time int) int {
	speed := msHeld
	return speed * (time - msHeld)
}

func binarySearchForLowerBoundary(f func(int, int) int, target int, initialHigh int) int {
	high := initialHigh
	low := 0
	mid := (low + high) / 2

	for low <= high && mid != 0 {
		mid = (low + high) / 2
		midResult := f(mid, initialHigh)
		midMinusOneResult := f(mid-1, initialHigh)

		if midResult > target {
			if midMinusOneResult <= target {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return mid
}

func binarySearchForUpperBoundary(f func(int, int) int, target int, initialHigh int) int {
	high := initialHigh
	low := 0
	mid := (low + high) / 2

	for low <= high && mid != initialHigh {
		mid = (low + high) / 2
		midResult := f(mid, initialHigh)
		midPlusOneResult := f(mid+1, initialHigh)

		if midResult > target {
			if midPlusOneResult <= target {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return mid
}

func calculateNumberOfWaysToWin(time int, recordDistance int) int {
	lowestButtonTimeHeld := binarySearchForLowerBoundary(calculateDistanceGivenMsHeld, recordDistance, time)
	highestButtonTimeHeld := binarySearchForUpperBoundary(calculateDistanceGivenMsHeld, recordDistance, time)

	return highestButtonTimeHeld - lowestButtonTimeHeld + 1
}

// Helper functions for part 2
