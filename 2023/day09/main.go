package main

import (
	"advent-of-code-go/pkg/regex"
	_ "embed"
	"flag"
	"fmt"
	"slices"
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

	total := 0
	for _, line := range parsed {
		total += getNextNumber(line)
	}

	return total
}

func part2(input string) int {
	parsed := parseInput(input)

	total := 0
	for _, line := range parsed {
		slices.Reverse(line)
		total += getNextNumber(line)
	}

	return total
}

func parseInput(input string) [][]int {
	split := strings.Split(input, "\n")
	allNums := [][]int{}
	for _, line := range split {
		allNums = append(allNums, regex.GetSpaceSeparatedNumbers(line))
	}
	return allNums
}

// Helper functions for part 1

func numbersAreAllZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func getNumberHistory(nums []int) []int {
	history := []int{}
	for i := range nums[:len(nums)-1] {
		history = append(history, nums[i+1]-nums[i])
	}
	return history
}

func getNextNumber(nums []int) int {
	if numbersAreAllZero(nums) {
		return 0
	}
	return getNextNumber(getNumberHistory(nums)) + lastValueOfArray[int](nums)
}

func lastValueOfArray[T any](n []T) T {
	return n[len(n)-1]
}

// Helper functions for part 2
