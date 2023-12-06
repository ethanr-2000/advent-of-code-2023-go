package main

import (
	"advent-of-code-go/pkg/cast"
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

	partNumbers := getPartNumbers(parsed)
	specialChars := getSpecialChars(parsed)

	sum := 0
	for _, partNumber := range partNumbers {
		if isAdjacentWithAtLeastOneSpecial(partNumber, specialChars) {
			sum += partNumber.value
		}
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	partNumbers := getPartNumbers(parsed)
	specialChars := getSpecialChars(parsed)

	sum := 0
	for _, specialChar := range specialChars {
		gearRatio := getGearRatio(specialChar, partNumbers)
		sum += gearRatio
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type PartNumber struct {
	value      int
	row        int
	startIndex int
	endIndex   int
}

type SpecialChar struct {
	value byte
	row   int
	index int
}

func getPartNumbers(input []string) []PartNumber {
	re := regexp.MustCompile(`(\d+)`)

	var partNumbers []PartNumber

	for row, line := range input {
		matches := re.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			startIndex, endIndex := match[0], match[1]
			numberStr := line[startIndex:endIndex]

			partNumbers = append(partNumbers, PartNumber{
				row:        row,
				value:      cast.ToInt(numberStr),
				startIndex: startIndex,
				endIndex:   endIndex - 1,
			})
		}
	}
	return partNumbers
}

func getSpecialChars(input []string) []SpecialChar {
	re := regexp.MustCompile(`[^\w\d.]`)

	var specialChars []SpecialChar

	for row, line := range input {
		matches := re.FindAllStringSubmatchIndex(line, -1)

		for _, match := range matches {
			specialChars = append(specialChars, SpecialChar{
				value: line[match[0]],
				row:   row,
				index: match[0],
			})
		}
	}

	return specialChars
}

func isAdjacentWithAtLeastOneSpecial(partNumber PartNumber, specialChars []SpecialChar) bool {
	for _, specialChar := range specialChars {
		if isAdjacent(partNumber, specialChar) {
			return true
		}
	}
	return false
}

func isAdjacent(partNumber PartNumber, specialChar SpecialChar) bool {
	return (partNumber.row == specialChar.row-1 || partNumber.row == specialChar.row || partNumber.row == specialChar.row+1) &&
		(partNumber.startIndex-1 <= specialChar.index && specialChar.index <= partNumber.endIndex+1)
}

// Helper functions for part 2

func getGearRatio(specialChar SpecialChar, partNumbers []PartNumber) int {
	if specialChar.value != byte('*') {
		return 0
	}

	adjacentParts := getAdjacentPartNumbers(specialChar, partNumbers)
	if len(adjacentParts) != 2 {
		return 0
	}
	return adjacentParts[0].value * adjacentParts[1].value
}

func getAdjacentPartNumbers(specialChar SpecialChar, partNumbers []PartNumber) []PartNumber {
	adjacentParts := []PartNumber{}
	for _, partNumber := range partNumbers {
		if isAdjacent(partNumber, specialChar) {
			adjacentParts = append(adjacentParts, partNumber)
		}
	}
	return adjacentParts
}
