package main

import (
	"advent-of-code-go/pkg/cast"
	"advent-of-code-go/pkg/regex"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
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
	cards := parseInput(input)

	totalPoints := 0
	for _, c := range cards {
		cardPoints := calculateCardPoints(c)
		totalPoints += cardPoints
	}

	return totalPoints
}

func part2(input string) int {
	cards := parseInput(input)

	updateCardCounts(cards)

	return sumOfCardCounts(cards)
}

func parseInput(input string) []Card {
	splitInput := strings.Split(input, "\n")

	var cards []Card
	for _, line := range splitInput {
		cards = append(cards, Card{
			id:             getCardIdFromLine(line),
			numbers:        getNumbersFromLine(line),
			winningNumbers: getWinningNumbersFromLine(line),
			count:          1,
		})
	}

	return cards
}

// Helper functions for part 1

type Card struct {
	id             int
	numbers        []int
	winningNumbers []int
	count          int
}

func getCardIdFromLine(line string) int {
	re := regexp.MustCompile(`Card +(\d+)`)

	match := re.FindStringSubmatch(line)

	return cast.ToInt(match[1])
}

func getWinningNumbersFromLine(line string) []int {
	stringBeforeLine := strings.Split(line, "|")[0]

	return regex.GetSpaceSeparatedNumbers(stringBeforeLine)
}

func getNumbersFromLine(line string) []int {
	stringAfterLine := strings.Split(line, "|")[1]

	return regex.GetSpaceSeparatedNumbers(stringAfterLine)
}

func calculateCardPoints(card Card) int {
	points := 0
	for _, winningNumber := range card.winningNumbers {
		if slices.Contains(card.numbers, winningNumber) {
			points *= 2
			if points == 0 {
				points = 1
			}
		}
	}
	return points
}

// Helper functions for part 2

func getMatchingNumbersCount(card Card) int {
	count := 0
	for _, winningNumber := range card.winningNumbers {
		if slices.Contains(card.numbers, winningNumber) {
			count += 1
		}
	}
	return count
}

func findCardById(target int, cards *[]Card) *Card {
	for i := range *cards {
		if (*cards)[i].id == target {
			return &(*cards)[i]
		}
	}
	return &Card{
		id: -1,
	}
}

func updateCardCounts(cards []Card) {
	for i := 0; i < len(cards); i++ {
		matchCount := getMatchingNumbersCount(cards[i])

		for j := 1; j <= matchCount; j++ {
			card := findCardById(cards[i].id+j, &cards)

			card.count += cards[i].count
		}
	}
}

func sumOfCardCounts(cards []Card) int {
	var sum int
	for _, c := range cards {
		sum += c.count
	}
	return sum
}
