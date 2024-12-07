package main

// import (
// 	"advent-of-code-go/pkg/regex"
// 	"strconv"

// 	_ "embed"
// 	"flag"
// 	"fmt"
// 	"strings"

// 	"github.com/atotto/clipboard"
// 	"gonum.org/v1/gonum/stat/combin"
// )

// //go:embed input.txt
// var input string

// func init() {
// 	// do this in init (not main) so test file has same input
// 	input = strings.TrimRight(input, "\n")
// 	if len(input) == 0 {
// 		panic("empty input.txt file")
// 	}
// }

// func main() {
// 	var part int
// 	flag.IntVar(&part, "part", 0, "part 1 or 2")
// 	flag.Parse()

// 	if part == 1 {
// 		ans := part1(input)
// 		fmt.Println("Running part 1")
// 		clipboard.WriteAll(fmt.Sprintf("%v", ans))
// 		fmt.Println("Output:", ans)
// 	} else if part == 2 {
// 		ans := part2(input)
// 		fmt.Println("Running part 2")
// 		clipboard.WriteAll(fmt.Sprintf("%v", ans))
// 		fmt.Println("Output:", ans)
// 	} else {
// 		fmt.Println("Running all")
// 		ans1 := part1(input)
// 		fmt.Println("Part 1 Output:", ans1)
// 		ans2 := part2(input)
// 		fmt.Println("Part 2 Output:", ans2)
// 	}
// }

// func part1(input string) int {
// 	parsed := parseInput(input)
// 	var equations []Equation
// 	for _, l := range parsed {
// 		equations = append(equations, getEquation(l))
// 	}

// 	totalValidEquations := 0
// 	for _, e := range equations {
// 		possibleOperators := generateAllPossibleOperators(len(e.operators))
// 		for _, os := range possibleOperators {
// 			if testEquation(Equation{e.testValue, e.nums, os}) == 0 {
// 				totalValidEquations+=e.testValue
// 				break
// 			}
// 		}
// 	}

// 	return totalValidEquations
// }

// func part2(input string) int {
// 	parsed := parseInput(input)
// 	var equations []Equation
// 	for _, l := range parsed {
// 		equations = append(equations, getEquation(l))
// 	}

// 	totalValidEquations := 0
// 	for i, e := range equations {
// 		fmt.Println(i)
// 		possibleOperators := generateAllPossibleOperatorsWithConcat(len(e.operators))
// 		for _, os := range possibleOperators {
// 			if testEquation(Equation{e.testValue, e.nums, os}) == 0 {
// 				totalValidEquations+=e.testValue
// 				break
// 			}
// 		}
// 	}

// 	return totalValidEquations
// }

// func parseInput(input string) []string {
// 	return strings.Split(input, "\n")
// }

// type Operator rune

// const (
// 	Plus Operator = '+'
// 	Multiply Operator = '*'
// 	Concat Operator = '|'
// )

// // Helper functions for part 1
// type Equation struct {
// 	testValue int
// 	nums []int
// 	operators []Operator
// }

// func getEquation(line string) Equation {
// 	nums := regex.GetNumbers(line)
// 	var operators []Operator
// 	for range nums[2:] {
// 		operators = append(operators, Plus)
// 	}
// 	return Equation{testValue: nums[0], nums: nums[1:], operators: operators}
// }

// // func tryFindOperators(e Equation) Equation {
// // 	// numberOfMult := countMult(e.operators)
// // 	result := testEquation(e)
// // 	if result < 0 {
// // 		// test value too big

// // 	} else if result > 0 {
// // 		// test value too small
// // 	} else {
// // 		// correct
// // 		return e
// // 	}
// // 	return e
// // }

// func generateAllPossibleOperators(length int) [][]Operator {
// 	possibleOperators := make([][]Operator, 0)

// 	allPlusOperators := make([]Operator, length)
// 	for i := range allPlusOperators {
// 		allPlusOperators[i] = Plus
// 	}
// 	for i := 0; i <= length; i++ {
// 		// i is the number of mults
// 		for _, combs := range combin.Combinations(length, i) {
// 			operators := make([]Operator, length)
// 			copy(operators, allPlusOperators)
// 			for _, c := range combs {
// 				operators[c] = Multiply
// 			}
// 			possibleOperators = append(possibleOperators, operators)
// 		}
// 	}
// 	return possibleOperators
// }

// // returns whether there was a match, the largest and smallest result
// // func tryAllCombinationsWithNMultiply(e Equation, n int) (bool, int, int) {
// // 	length := len(e.operators)
// // 	operators := make([]Operator, length)
// // 	for i := 0; i < n; i++ {
// // 			operators[i] = Multiply
// // 	}
// // 	for i := n; i < length; i++ {
// // 			operators[i] = Plus
// // 	}

// // 	perms := combin.Permutations(length, length)
// // 	for _, perm := range perms {
// // 		result := make([]Operator, length)
// // 		for i, idx := range perm {
// // 				result[i] = operators[idx]
// // 		}
// // 		fmt.Println(result)
// // 	}
// // 	return false, 0, 0
// // }

// // func countMult(operators []Operator) int {
// // 	count := 0
// // 	for i := range operators {
// // 		if (operators[i] == Multiply) {
// // 			count++
// // 		}
// // 	}
// // 	return count
// // }

// // -ve -> testValue >  total
// // +ve -> testValue <  total
// // 0 	 -> testValue == total
// func testEquation(e Equation) int {
// 	total := e.nums[0]
// 	for i := range e.operators {
// 		if e.operators[i] == Plus {
// 			total += e.nums[i+1]
// 		}
// 		if e.operators[i] == Multiply {
// 			total *= e.nums[i+1]
// 		}
// 		if e.operators[i] == Concat {
// 			total, _ = strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(e.nums[i+1]))
// 		}
// 	}
// 	return total - e.testValue
// }

// // Helper functions for part 2

// func generateAllPossibleOperatorsWithConcat(length int) [][]Operator {
// 	possibleOperators := make([][]Operator, 0)

// 	// Start with all Plus operators as the base
// 	allPlusOperators := make([]Operator, length)
// 	for i := range allPlusOperators {
// 		allPlusOperators[i] = Plus
// 	}

// 	// Iterate through possible Multiply replacements
// 	for multCount := 0; multCount <= length; multCount++ {
// 		// Get all combinations of where Multiply can be placed
// 		for _, multCombs := range combin.Combinations(length, multCount) {
// 			// Iterate through possible Concat replacements
// 			for concatCount := 0; concatCount <= length-multCount; concatCount++ {
// 				// Get all combinations of where Concat can be placed
// 				for _, concatCombs := range combin.Combinations(length, concatCount) {
// 					// Create a copy of the base Plus operators
// 					operators := make([]Operator, length)
// 					copy(operators, allPlusOperators)

// 					// Replace some operators with Multiply
// 					for _, c := range multCombs {
// 						operators[c] = Multiply
// 					}

// 					// Replace some remaining operators with Concat
// 					for _, c := range concatCombs {
// 						// Ensure we don't overwrite existing Multiply operators
// 						if operators[c] == Plus {
// 							operators[c] = Concat
// 						}
// 					}

// 					possibleOperators = append(possibleOperators, operators)
// 				}
// 			}
// 		}
// 	}

// 	return possibleOperators
// }