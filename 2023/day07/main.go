package main

import (
	"advent-of-code-go/pkg/cast"
	"advent-of-code-go/pkg/list"
	_ "embed"
	"flag"
	"fmt"
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
	parsed := parseInput(input)

	hands := getHands(parsed)
	slices.SortFunc[[]Hand](hands, compareHandStrength)

	r := 0
	for i, h := range hands {
		r += h.bid * (i + 1)
	}

	return r
}

func part2(input string) int {
	parsed := parseInput(input)

	hands := getHands(parsed)
	slices.SortFunc[[]Hand](hands, compareHandStrengthWithJokerRules)

	r := 0
	for i, h := range hands {
		r += h.bid * (i + 1)
	}

	return r
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Hand struct {
	cards []string
	bid   int
}

const (
	HighCard     = iota // 0
	OnePair             // 1
	TwoPair             // 2
	ThreeOfAKind        // 3
	FullHouse           // 4
	FourOfAKind         // 5
	FiveOfAKind         // 6
)

func getHands(input []string) []Hand {
	var hands []Hand
	for _, line := range input {
		splitLine := strings.Split(line, " ")

		cards := strings.Split(splitLine[0], "")

		hands = append(hands, Hand{
			cards: cards,
			bid:   cast.ToInt(splitLine[1]),
		})
	}
	return hands
}

// returns the count of each card type in hand, e.g. { "A": 2 }
func getCardCount(cards []string) map[string]int {
	cardCountMap := make(map[string]int)
	for _, c := range cards {
		cardCountMap[c] += 1
	}

	return cardCountMap
}

func cardCountIsFiveOfAKind(cardCount map[string]int) bool {
	for _, count := range cardCount {
		return count == 5
	}
	return false
}

func cardCountIsFourOfAKind(cardCount map[string]int) bool {
	for _, count := range cardCount {
		if count == 4 {
			return true
		}
	}
	return false
}

func cardCountIsFullHouse(cardCount map[string]int) bool {
	has2, has3 := false, false
	for _, count := range cardCount {
		if count == 3 {
			has3 = true
		}
		if count == 2 {
			has2 = true
		}
	}
	return has3 && has2
}

func cardCountIsThreeOfAKind(cardCount map[string]int) bool {
	for _, count := range cardCount {
		if count == 3 {
			return true
		}
	}
	return false
}

func cardCountIsTwoPair(cardCount map[string]int) bool {
	pairCount := 0
	for _, count := range cardCount {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

func cardCountIsOnePair(cardCount map[string]int) bool {
	for _, count := range cardCount {
		if count == 2 {
			return true
		}
	}
	return false
}

func getCardsType(cards []string) int {
	cardCount := getCardCount(cards)

	if cardCountIsFiveOfAKind(cardCount) {
		return FiveOfAKind
	}
	if cardCountIsFourOfAKind(cardCount) {
		return FourOfAKind
	}
	if cardCountIsFullHouse(cardCount) {
		return FullHouse
	}
	if cardCountIsThreeOfAKind(cardCount) {
		return ThreeOfAKind
	}
	if cardCountIsTwoPair(cardCount) {
		return TwoPair
	}
	if cardCountIsOnePair(cardCount) {
		return OnePair
	}
	return HighCard
}

func compareTwoCards(c1, c2 string, rankOrder []string) int {
	index1 := -1
	index2 := -1

	for i, r := range rankOrder {
		if r == c1 {
			index1 = i
		}
		if r == c2 {
			index2 = i
		}
	}

	return index2 - index1
}

func compareHandStrength(h1, h2 Hand) int {
	return compareCardsStrength(h1.cards, h2.cards)
}

func compareHandStrengthWithJokerRules(h1, h2 Hand) int {
	return compareCardsStrengthWithJokerRules(h1.cards, h2.cards)
}

// returns +ve if h1 is greater than h2
// returns -ve if h1 is weaker than to h2
func compareCardsStrength(cards1, cards2 []string) int {
	rankingOrder := strings.Split("AKQJT98765432", "")

	type1, type2 := getCardsType(cards1), getCardsType(cards2)
	if type1 == type2 {
		for i := range cards1 {
			cardComparison := compareTwoCards(cards1[i], cards2[i], rankingOrder)

			if cardComparison != 0 {
				return cardComparison
			}
		}
		return 1 // hands are totally equal, default h1 > h2
	} else {
		return type1 - type2
	}
}

// Helper functions for part 2

func getStrongestCardsWithJokerRules(cards []string, rankingOrder []string) []string {
	jCount := list.CountOfOccurencesOfStringInList(cards, "J")

	if jCount == 0 {
		return cards
	}

	if jCount == 5 {
		bestCard := rankingOrder[0]
		return []string{bestCard, bestCard, bestCard, bestCard, bestCard}
	}

	var possibleCardsWithJReplaced [][]string
	for c := range getCardCount(cards) {
		newCardsWithJsReplaced := make([]string, len(cards))
		copy(newCardsWithJsReplaced, cards)

		list.ReplaceAllInstancesOfStringInList(newCardsWithJsReplaced, "J", c)

		possibleCardsWithJReplaced = append(possibleCardsWithJReplaced, newCardsWithJsReplaced)
	}

	slices.SortFunc[[][]string](possibleCardsWithJReplaced, compareCardsStrength)
	slices.Reverse[[][]string](possibleCardsWithJReplaced)

	strongestCards := possibleCardsWithJReplaced[0]
	return strongestCards
}

// returns +ve if h1 is greater than h2
// returns -ve if h1 is weaker than to h2
func compareCardsStrengthWithJokerRules(cards1 []string, cards2 []string) int {
	rankingOrder := strings.Split("AKQT98765432J", "")

	cards1WithJsReplaced := getStrongestCardsWithJokerRules(cards1, rankingOrder)
	cards2WithJsReplaced := getStrongestCardsWithJokerRules(cards2, rankingOrder)

	type1, type2 := getCardsType(cards1WithJsReplaced), getCardsType(cards2WithJsReplaced)
	if type1 == type2 {
		for i := range cards1 {
			cardComparison := compareTwoCards(cards1[i], cards2[i], rankingOrder)

			if cardComparison != 0 {
				return cardComparison
			}
		}
		return 1 // hands are totally equal, default h1 > h2
	} else {
		return type1 - type2
	}
}
