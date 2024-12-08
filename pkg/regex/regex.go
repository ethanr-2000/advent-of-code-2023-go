//nolint:gosec
package regex

import (
	"advent-of-code-go/pkg/cast"
	"regexp"
	"strings"
)

func GetNumbers(str string) []int {
	re := regexp.MustCompile(`(\d+)`)

	var numbers []int
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		numbers = append(numbers, cast.ToInt(match[1]))
	}

	return numbers
}

// "<anything> 10 11 -12 13 <anything>" -> []int{ 10, 11, -12, 13}
func GetSpaceSeparatedNumbers(str string) []int {
	// any number of digits, then any number of whitespace or nothing or EOL
	re := regexp.MustCompile(`(-?\d+((\s+|^[\s\S]|$)))`)

	matches := re.FindAllStringSubmatch(str, -1)

	var numbers []int
	for _, match := range matches {
		numStr := strings.TrimSpace(match[1])
		numbers = append(numbers, cast.ToInt(numStr))
	}
	return numbers
}

// matches empty string ""
func IsEmptyString(s string) bool {
	regex := regexp.MustCompile(`^$`)
	return regex.MatchString(s)
}

// matches line that contains any text
func HasText(s string) bool {
	regex := regexp.MustCompile(`[a-zA-Z]`)
	return regex.MatchString(s)
}

// matches line that contains given regex
func Contains(s string, regexToMatch string) bool {
	regex := regexp.MustCompile(regexp.QuoteMeta(regexToMatch))
	return regex.MatchString(s)
}

// returns index of given character
func IndicesOfCharacter(s string, charToMatch string) []int {
	regex := regexp.MustCompile(regexp.QuoteMeta(charToMatch))
	matches := regex.FindAllStringIndex(s, -1)

	indices := []int{}
	for _, match := range matches {
		indices = append(indices, match[0])
	}
	return indices
}

// "###.#.##....", "#" -> []int{3, 1, 2}
func LengthsOfGroupsOfChar(s string, c rune) []int {
	re := regexp.MustCompile(regexp.QuoteMeta(string(c)) + "+")
	matches := re.FindAllString(s, -1)

	groupLengths := []int{}
	for _, match := range matches {
		groupLengths = append(groupLengths, len(match))
	}
	return groupLengths
}

// counts instances of rune in a string
func Count(s string, c rune) int {
	re := regexp.MustCompile(regexp.QuoteMeta(string(c)))
	return len(re.FindAllString(s, -1))
}

// gets /w
func GetWords(s string) []string {
	re := regexp.MustCompile(`(([a-zA-Z]+))`)
	matches := re.FindAllString(s, -1)

	return matches
}
