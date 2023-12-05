package main

import (
	"advent-of-code-go/pkg/cast"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"slices"
	"strings"
	"unicode"

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
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)

	var sum int = 0
	for _, str := range parsed {
		firstNum, error := getFirstNumberCharacterInString(str)
		if error != nil {
			fmt.Println("Error:", error)
			return 0
		}

		lastNum, error := getFirstNumberCharacterInString(reverseString(str))
		if error != nil {
			fmt.Println("Error:", error)
			return 0
		}

		var calibrationValue = cast.ToInt(cast.ToString(firstNum) + cast.ToString(lastNum))
		sum += calibrationValue
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	wordToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	reverseWordToInt := map[string]int{
		"eno":   1,
		"owt":   2,
		"eerht": 3,
		"ruof":  4,
		"evif":  5,
		"xis":   6,
		"neves": 7,
		"thgie": 8,
		"enin":  9,
	}

	var sum int = 0
	for _, str := range parsed {
		firstNum, error := getFirstNumberCharacterOrWordInString(str, wordToInt)
		if error != nil {
			fmt.Println("Error:", error)
			return 0
		}

		lastNum, error := getFirstNumberCharacterOrWordInString(reverseString(str), reverseWordToInt)
		if error != nil {
			fmt.Println("Error:", error)
			return 0
		}

		var calibrationValue = cast.ToInt(cast.ToString(firstNum) + cast.ToString(lastNum))
		sum += calibrationValue
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func getFirstNumberCharacterInString(str string) (rune, error) {
	for _, char := range str {
		if unicode.IsDigit(char) {
			return char, nil
		}
	}
	var emptyRune rune
	return emptyRune, errors.New(fmt.Sprintf("Could not find number in %s", str))
}

func reverseString(input string) string {
	runes := []rune(input)

	slices.Reverse[[]rune](runes)

	return string(runes)
}

func getFirstNumberCharacterOrWordInString(str string, m map[string]int) (int, error) {
	for i, char := range str {
		if unicode.IsDigit(char) {
			return cast.ToInt(cast.ToString(char)), nil
		}

		result, err := checkAndReturnIfNextRunesAreInMap(str[i:], m)
		if err != nil {
			return 0, err
		}

		if result != 0 {
			return result, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Could not find number in %s", str))
}

func checkAndReturnIfNextRunesAreInMap(str string, m map[string]int) (int, error) {
	for i := range str {
		substring := str[:i]
		// if the substring is a key in the map, e.g. "eight"
		if value, ok := m[substring]; ok {
			return value, nil
		}
	}
	return 0, nil
}
