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

var CACHE = make(map[string]int)

func part1(input string) int {
	towels, patterns := getTowelsAndPatterns(input)

	CACHE = make(map[string]int)

	possible := 0
	for _, p := range patterns {
		if countAllPossibleCombinations(p, towels) > 0 {
			possible++
		}
	}

	return possible
}

func part2(input string) int {
	towels, patterns := getTowelsAndPatterns(input)

	count := 0
	for _, p := range patterns {
		count += countAllPossibleCombinations(p, towels)
	}

	return count
}

// Helper functions for part 1

func getTowelsAndPatterns(input string) (towels []string, patterns []string) {
	parsed := strings.Split(input, "\n")
	towels = strings.Split(parsed[0], ", ")
	patterns = parsed[2:]
	return
}

func countAllPossibleCombinations(pattern string, towels []string) int {
	if len(pattern) == 0 {
		return 1 // there is one way to arrange 0 towels
	}

	if count, exists := CACHE[pattern]; exists {
		return count
	}

	total := 0
	for _, t := range towels {
		if strings.HasPrefix(pattern, t) {
			c := countAllPossibleCombinations(pattern[len(t):], towels)
			total += c
		}
	}
	CACHE[pattern] = total
	return total
}

// Helper functions for part 2
