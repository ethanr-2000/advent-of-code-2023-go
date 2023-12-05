package main

import (
	"advent-of-code-go/pkg/regex"
	_ "embed"
	"flag"
	"fmt"
	"math"
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
	split := parseInput(input)

	seeds := getSeeds(split[0])
	maps := calculateMaps(split)

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		seed = getSeedDataFromMaps(seed, maps)

		if seed.location < lowestLocation {
			lowestLocation = seed.location
		}
	}

	return lowestLocation
}

func part2(input string) int {
	split := parseInput(input)

	seedRanges := calculateSeedRanges(split[0])
	maps := calculateMaps(split)

	for i := 0; i < math.MaxInt; i++ {
		seedNumber := getSeedNumberFromLocation(i, maps)
		if seedNumberInSeedRanges(seedNumber, seedRanges) {
			return i
		}
	}
	return 0
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Seed struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type Map struct {
	source int
	dest   int
	rng    int
}

func getSeeds(line string) []Seed {
	var seeds []Seed
	for _, num := range regex.GetSpaceSeparatedNumbers(line) {
		seeds = append(seeds, Seed{
			seed:        num,
			soil:        -1,
			fertilizer:  -1,
			water:       -1,
			light:       -1,
			temperature: -1,
			humidity:    -1,
			location:    -1,
		})
	}
	return seeds
}

func calculateMaps(input []string) [][]Map {
	lineNum := 2
	mapNum := 0

	maps := [][]Map{}
	for mapNum < 7 && lineNum < len(input) {
		if regex.IsEmptyString(input[lineNum]) {
			mapNum += 1
			lineNum += 1
			continue
		}

		if regex.HasText(input[lineNum]) {
			maps = append(maps, []Map{})
			lineNum += 1
			continue
		}

		singleMap := regex.GetSpaceSeparatedNumbers(input[lineNum])

		if len(singleMap) != 3 {
			return maps
		}

		maps[mapNum] = append(maps[mapNum], Map{
			dest:   singleMap[0],
			source: singleMap[1],
			rng:    singleMap[2],
		})

		lineNum += 1
	}

	return maps
}

func getSeedDataFromMaps(seed Seed, maps [][]Map) Seed {
	seed.soil = mapValue(seed.seed, maps[0])
	seed.fertilizer = mapValue(seed.soil, maps[1])
	seed.water = mapValue(seed.fertilizer, maps[2])
	seed.light = mapValue(seed.water, maps[3])
	seed.temperature = mapValue(seed.light, maps[4])
	seed.humidity = mapValue(seed.temperature, maps[5])
	seed.location = mapValue(seed.humidity, maps[6])

	return seed
}

func mapValue(val int, maps []Map) int {
	for _, m := range maps {
		if m.source <= val && val < m.source+m.rng {
			difference := m.dest - m.source
			return val + difference
		}
	}
	return val
}

// Helper functions for part 2

func mapValueReverse(val int, maps []Map) int {
	for _, m := range maps {
		if m.dest <= val && val < m.dest+m.rng {
			difference := m.dest - m.source
			return val - difference
		}
	}
	return val
}

func calculateSeedRanges(line string) [][]int {
	var seedRanges [][]int

	seedNums := regex.GetSpaceSeparatedNumbers(line)
	// for each pair of numbers
	for i := 0; i < len(seedNums); i += 2 {
		seedRanges = append(seedRanges, []int{
			seedNums[i],
			seedNums[i] + seedNums[i+1],
		})
	}
	return seedRanges
}

func seedNumberInSeedRanges(num int, ranges [][]int) bool {
	for _, r := range ranges {
		if r[0] <= num && num <= r[1] {
			return true
		}
	}

	return false
}

func getSeedNumberFromLocation(location int, maps [][]Map) int {
	humidity := mapValueReverse(location, maps[6])
	temperature := mapValueReverse(humidity, maps[5])
	light := mapValueReverse(temperature, maps[4])
	water := mapValueReverse(light, maps[3])
	fertilizer := mapValueReverse(water, maps[2])
	soil := mapValueReverse(fertilizer, maps[1])
	seedNum := mapValueReverse(soil, maps[0])

	return seedNum
}
