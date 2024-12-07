package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	m := getMap(parsed)
	i := parseInstructions(parsed[0])

	return traverseMap(m, i, "AAA", "ZZZ")
}

func part2(input string) int {
	parsed := parseInput(input)
	m := getMap(parsed)
	i := parseInstructions(parsed[0])
	startingNodes := getAllNodesThatEndWith(m, "A")

	minSteps := []int{}
	for _, n := range startingNodes {
		minSteps = append(minSteps, traverseMap(m, i, n, "Z"))
	}
	return findLowestCommonMultipleOfList(minSteps)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Map map[string][]string

func getMap(input []string) Map {
	m := make(map[string][]string)
	for _, line := range input[2:] {
		re := regexp.MustCompile(`(\w+)`)
		matches := re.FindAllStringSubmatch(line, 3)
		m[matches[0][1]] = []string{matches[1][1], matches[2][1]}
	}
	return m
}

func parseInstructions(line string) []int {
	instructions := []int{}
	for _, c := range line {
		if c == 'L' {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, 1)
		}
	}
	return instructions
}

func atTarget(s1, s2 string) bool {
	return strings.HasSuffix(s1, s2)
}

func traverseMap(m Map, instructions []int, startNode string, targetSuffix string) int {
	currentNode := startNode
	i := 0
	loops := 0
	for !atTarget(currentNode, targetSuffix) {
		if i == len(instructions) {
			i = 0
			loops++
		}

		currentNode = m[currentNode][instructions[i]]
		i++
	}
	return i + loops*len(instructions)
}

// Helper functions for part 2

func getAllNodesThatEndWith(m Map, s string) []string {
	nodes := []string{}
	for k := range m {
		if strings.HasSuffix(k, s) {
			nodes = append(nodes, k)
		}
	}
	return nodes
}

func greatestCommonDivisor(a, b int) int {
	// use euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lowestCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}

func findLowestCommonMultipleOfList(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for _, num := range numbers[1:] {
		result = lowestCommonMultiple(result, num)
	}
	return result
}
