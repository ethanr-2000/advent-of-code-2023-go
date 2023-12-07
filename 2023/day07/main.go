package main

import (
	"advent-of-code-go/pkg/cast"
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

	hands := getHands(parsed)
	hands = sortHandsByStrength(hands[:])

	r := 0
	for i, h := range hands {
		r += h.bid * (i + 1)
	}

	return r
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
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
		// slices.SortFunc[[]string](cards, compareTwoCards)
		// slices.Reverse[[]string](cards)

		hands = append(hands, Hand{
			cards: cards,
			bid:   cast.ToInt(splitLine[1]),
		})
	}
	return hands
}

// returns the count of each card type in hand
func getCardCount(h Hand) []int {
	cardCountMap := make(map[string]int)
	for _, c := range h.cards {
		cardCountMap[c] += 1
	}

	var cardCountList []int
	for _, val := range cardCountMap {
		cardCountList = append(cardCountList, val)
	}
	return cardCountList
}

func cardCountIsFiveOfAKind(cardCount []int) bool {
	return cardCount[0] == 5
}

func cardCountIsFourOfAKind(cardCount []int) bool {
	return cardCount[0] == 4 || cardCount[1] == 4
}

func cardCountIsFullHouse(cardCount []int) bool {
	return slices.Contains[[]int](cardCount, 3) && slices.Contains[[]int](cardCount, 2)
}

func cardCountIsThreeOfAKind(cardCount []int) bool {
	return slices.Contains[[]int](cardCount, 3) && slices.Contains[[]int](cardCount, 1)
}

func cardCountIsTwoPair(cardCount []int) bool {
	slices.Sort[[]int](cardCount)
	slices.Reverse[[]int](cardCount)
	return len(cardCount) == 3 && cardCount[0] == 2 && cardCount[1] == 2
}

func cardCountIsOnePair(cardCount []int) bool {
	slices.Sort[[]int](cardCount)
	slices.Reverse[[]int](cardCount)
	return len(cardCount) >= 3 && cardCount[0] == 2 && cardCount[1] == 1
}

func getHandClassification(h Hand) int {
	cardCount := getCardCount(h)

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

func compareTwoCards(char1, char2 string) int {
	rankingOrder := strings.Split("AKQJT98765432", "")

	index1 := -1
	index2 := -1

	for i, r := range rankingOrder {
		if r == char1 {
			index1 = i
		}
		if r == char2 {
			index2 = i
		}
	}

	// Compare the indices to determine the rank
	if index1 > index2 {
		return -1
	} else if index1 < index2 {
		return 1
	}

	return 0
}

// returns +ve if h1 is greater than h2
// returns -ve if h1 is weaker than to h2
func compareHandStrength(h1 Hand, h2 Hand) int {
	handClass1, handClass2 := getHandClassification(h1), getHandClassification(h2)
	if handClass1 == handClass2 {
		for i := range h1.cards {
			cardComparison := compareTwoCards(h1.cards[i], h2.cards[i])

			if cardComparison != 0 {
				return cardComparison
			}
		}
	} else {
		return handClass1 - handClass2
	}
	// hands are totally equal, default h1 > h2
	return 1
}

func sortHandsByStrength(hands []Hand) []Hand {
	slices.SortFunc[[]Hand](hands, compareHandStrength)
	return hands
}

// Helper functions for part 2
