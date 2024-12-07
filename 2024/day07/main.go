package main

import (
	"advent-of-code-go/pkg/regex"
	"strconv"

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
	var equations []Equation
	for _, l := range parsed {
		equations = append(equations, getEquation(l))
	}

	totalValidEquations := 0
	for _, e := range equations {
		if equationIsPossible(Equation{e.testValue, e.nums}, false) {
			totalValidEquations+=e.testValue
			continue
		}
	}

	return totalValidEquations
}

func part2(input string) int {
	parsed := parseInput(input)
	var equations []Equation
	for _, l := range parsed {
		equations = append(equations, getEquation(l))
	}

	totalValidEquations := 0
	for _, e := range equations {
		if equationIsPossible(Equation{e.testValue, e.nums}, true) {
			totalValidEquations+=e.testValue
			continue
		}
	}

	return totalValidEquations
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1
type Equation struct {
	testValue int
	nums []int
}

func getEquation(line string) Equation {
	nums := regex.GetNumbers(line)
	return Equation{testValue: nums[0], nums: nums[1:]}
}

func equationIsPossible(e Equation, checkConcat bool) bool {
	l := len(e.nums)
	if l == 1 {
		return e.nums[l-1] == e.testValue
	}

	if e.testValue % e.nums[l-1] == 0 && equationIsPossible(Equation{e.testValue / e.nums[l-1], e.nums[0:l-1]}, checkConcat){
		return true
	}

	trimmedTestValue, _ := strconv.Atoi(strings.TrimSuffix(strconv.Itoa(e.testValue), strconv.Itoa(e.nums[l-1])))
	if checkConcat && strings.HasSuffix(strconv.Itoa(e.testValue), strconv.Itoa(e.nums[l-1])) && equationIsPossible(Equation{trimmedTestValue, e.nums[0:l-1]}, checkConcat) {
		return true
	}
	if e.testValue > e.nums[l-1] && equationIsPossible(Equation{e.testValue - e.nums[l-1], e.nums[0:l-1]}, checkConcat) {
		return true
	}
	return false
}

// Helper functions for part 2
