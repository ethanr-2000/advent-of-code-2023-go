package main

import (
	"advent-of-code-go/pkg/regex"
	"regexp"

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
	mulStrings := getAllMulStrings(input)

	total := 0
	for _, m := range mulStrings {
		total += calculateMul(m)
	}

	return total
}

func part2(input string) int {
	muls := getAllMulStringsWithIndex(input)
	enabled := calculateEnabledIndices(input)
	
	return calculateAllMuls(muls, enabled)
}

// Helper functions for part 1

func getAllMulStrings(str string) []string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	return re.FindAllString(str, -1)
}

func calculateMul(mulStr string) int {
	nums := regex.GetNumbers(mulStr)
	return nums[0] * nums[1]
}

// Helper functions for part 2

type Mul struct {
	total int
	i int
}

func getAllMulStringsWithIndex(str string) []Mul {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	indices := re.FindAllStringIndex(str, -1)

	var mulStrings []Mul;
	for _, index := range indices {
		nums := regex.GetNumbers(str[index[0]:index[1]])

		mulStrings = append(mulStrings, Mul{nums[0] * nums[1], index[0]})
	}
	return mulStrings
}

func calculateEnabledIndices(str string) []bool {
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don\'t\(\)`)

	doIndices := doRe.FindAllStringIndex(str, -1)
	dontIndices := dontRe.FindAllStringIndex(str, -1)

	var indicesEnabled []bool
	enabled := true
	for i := range str {
		if len(doIndices) > 0 && doIndices[0][0] == i {
			enabled = true
			if (len(doIndices) > 1) {
				doIndices = doIndices[1:]
			}
		}
		if len(dontIndices) > 0 && dontIndices[0][0] == i {
			enabled = false
			if (len(dontIndices) > 1) {
				dontIndices = dontIndices[1:]
			}
		}
		indicesEnabled = append(indicesEnabled, enabled)
	}
	return indicesEnabled
}

func calculateAllMuls(muls []Mul, enabled []bool) int {
	sum := 0
	for _, m := range muls {
		if enabled[m.i] {
			sum += m.total
		}
	}
	return sum
}
