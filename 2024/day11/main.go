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

var STONE_CACHE = make(map[int]map[int]int)

func part1(input string) int {
	parsed := parseInput(input)
	stones := regex.GetNumbers(parsed[0])
	return blinkAllStones(stones, 25)
}

func part2(input string) int {
	parsed := parseInput(input)
	stones := regex.GetNumbers(parsed[0])

	return blinkAllStones(stones, 75)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func blinkAllStones(stones []int, blinks int) int {
	var totalStones int
	for _, s := range stones {
		totalStones += blinkStone(s, blinks)
	}
	return totalStones
}

func changeStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	stringStone := strconv.Itoa(stone)
	if len(stringStone)%2 == 0 {
		// split the stone in 2, e.g. 2024 -> 20, 24
		a, _ := strconv.Atoi(stringStone[0 : len(stringStone)/2])
		b, _ := strconv.Atoi(stringStone[len(stringStone)/2:])

		return []int{a, b}
	}

	return []int{stone * 2024}
}

func blinkStone(stone int, blinksRemaining int) int { //returns either the stones or the number of stones
	cacheResult := getFromCacheResult(stone, blinksRemaining)
	if cacheResult >= 0 {
		return cacheResult
	}

	changedStone := changeStone(stone)

	if blinksRemaining == 1 {
		return len(changedStone)
	}

	totalStones := 0
	for _, changed := range changedStone {
		totalStones += blinkStone(changed, blinksRemaining-1)
	}

	return addToCache(stone, blinksRemaining, totalStones)
}

// Helper functions for part 2

func addToCache(stone int, blinks int, numStones int) int {
	if _, exists := STONE_CACHE[stone]; !exists {
		STONE_CACHE[stone] = make(map[int]int)
	}

	STONE_CACHE[stone][blinks] = numStones
	return numStones
}

func getFromCacheResult(stone int, iterations int) int {
	stoneMap, exists := STONE_CACHE[stone]
	if !exists {
		return -1
	}

	numStones, exists := stoneMap[iterations]
	if !exists {
		return -1
	}
	// fmt.Println("cache hit", stone, iterations, stones)
	return numStones
}
