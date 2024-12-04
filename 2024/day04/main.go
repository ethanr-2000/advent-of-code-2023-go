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

func part1(input string) int {
	parsed, width := parseInput(input)

	return countAllXmas(parsed, width)
}

func part2(input string) int {
	parsed, width := parseInput(input)

	return countAllXMas(parsed, width)
}

func parseInput(input string) (string, int) {
	width := strings.Index(input, "\n") + 1
	return input, width
}

// Helper functions for part 1

func countAllXmas(str string, width int) int {
	count := 0
	for i, c := range str {
		if c == 'X' {
			count += checkXmasWithDiff(str, i, -width)   // N
			count += checkXmasWithDiff(str, i, -width+1) // NE
			count += checkXmasWithDiff(str, i, 1)        // E
			count += checkXmasWithDiff(str, i, width+1)  // SE
			count += checkXmasWithDiff(str, i, width)    // S
			count += checkXmasWithDiff(str, i, width-1)  // SW
			count += checkXmasWithDiff(str, i, -1)       // W
			count += checkXmasWithDiff(str, i, -width-1) // NW
		}
	}
	return count
}

func checkXmasWithDiff(str string, xIndex int, diff int) int {
	m, _ := safeAccessStr(str, xIndex+(diff))
	a, _ := safeAccessStr(str, xIndex+(2*diff))
	s, _ := safeAccessStr(str, xIndex+(3*diff))

	if m == 'M' && a == 'A' && s == 'S' {
		return 1
	} else {
		return 0
	}
}

func safeAccessStr(str string, index int) (byte, error) {
	if index < 0 || index >= len(str) {
		return byte('0'), nil
	}
	return str[index], nil
}

// Helper functions for part 2

// X-MAS
func countAllXMas(str string, width int) int {
	count := 0
	for i, c := range str {
		if c == 'A' {
			count += checkMasAroundA(str, i, width)
		}
	}
	return count
}

func checkMasAroundA(str string, aIndex int, width int) int {
	m1, _ := safeAccessStr(str, aIndex+(width+1))
	s1, _ := safeAccessStr(str, aIndex+(-width-1))
	m2, _ := safeAccessStr(str, aIndex+(-width-1))
	s2, _ := safeAccessStr(str, aIndex+(width+1))

	if !(m1 == 'M' && s1 == 'S') && !(m2 == 'M' && s2 == 'S') {
		// one diagonal is missing, don't bother checking
		return 0
	}

	m3, _ := safeAccessStr(str, aIndex+(width-1))
	s3, _ := safeAccessStr(str, aIndex+(-width+1))
	m4, _ := safeAccessStr(str, aIndex+(-width+1))
	s4, _ := safeAccessStr(str, aIndex+(width-1))

	if !(m3 == 'M' && s3 == 'S') && !(m4 == 'M' && s4 == 'S') {
		// other diagonal is missing
		return 0
	}

	return 1
}
