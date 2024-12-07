package main

import (
	"advent-of-code-go/pkg/regex"

	_ "embed"
	"flag"
	"fmt"
	"sort"
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

	for i := range parsed {
		sort.Ints(parsed[i])
	}

	return sumDifferences(parsed[0], parsed[1])
}

func part2(input string) int {
	parsed := parseInput(input)

	return similarityScore(parsed[0], parsed[1])
}

func parseInput(input string) [][]int {
	lists := make([][]int, 2)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		for i, num := range regex.GetSpaceSeparatedNumbers(line) {
			lists[i] = append(lists[i], num)
		}
	}
	return lists
}

// Helper functions for part 1

func AbsInt(x int) int {
	if x < 0 {
			return -x
	}
	return x
}

func sumDifferences(l1 []int, l2 []int) int {
	var sum = 0
	for i := range l1 {
		sum += AbsInt(l1[i] - l2[i])
	}
	return sum
}

// Helper functions for part 2

func similarityScore(l1 []int, l2 []int) int {
	counts := make(map[int]int)

	for _, num := range l2 {
			counts[num]++
	}

	var sum = 0
	for i := range l1 {
		sum += l1[i] * counts[l1[i]]
	}
	return sum
}
