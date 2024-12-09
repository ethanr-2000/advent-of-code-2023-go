package main

import (
	"advent-of-code-go/pkg/regex"
	"slices"

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
	rules, pages := parseInput(input)

	total := 0
	for _, p := range pages {
		if pagesAreCorrectlyOrdered(p, rules) {
			total += p[len(p)/2]
		}
	}

	return total
}

func part2(input string) int {
	rules, pages := parseInput(input)

	total := 0
	for _, p := range pages {
		if pagesAreCorrectlyOrdered(p, rules) {
			// already ordered
			continue
		}
		newPages := sortPages(p, rules)
		total += newPages[len(newPages)/2]
	}

	return total
}

func parseInput(input string) ([][]int, [][]int) {
	split := strings.Split(input, "\n\n")
	rulesStr := strings.Split(split[0], "\n")
	var rules [][]int
	for _, r := range rulesStr {
		rules = append(rules, regex.GetNumbers(r))
	}

	pagesStr := strings.Split(split[1], "\n")
	var pages [][]int
	for _, p := range pagesStr {
		pages = append(pages, regex.GetNumbers(p))
	}

	return rules, pages
}

// Helper functions for part 1

func pagesAreCorrectlyOrdered(pages []int, rules [][]int) bool {
	for _, r := range rules {
		if slices.Index(pages, r[0]) == -1 || slices.Index(pages, r[1]) == -1 {
			continue
		}
		if slices.Index(pages, r[0]) > slices.Index(pages, r[1]) {
			return false
		}
	}
	return true
}

// Helper functions for part 2

func sortPages(pages []int, rules [][]int) []int {
	for !pagesAreCorrectlyOrdered(pages, rules) {
		for _, r := range rules {
			if slices.Index(pages, r[0]) == -1 || slices.Index(pages, r[1]) == -1 {
				continue
			}

			p1Index := slices.Index(pages, r[0])
			p2Index := slices.Index(pages, r[1])
			if p1Index > p2Index {
				// swap them if they're incorrectly ordered
				pages[p1Index] = r[1]
				pages[p2Index] = r[0]
			}
		}
	}
	return pages
}
