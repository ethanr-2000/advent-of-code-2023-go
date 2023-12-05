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

	redNum := 12
	greenNum := 13
	blueNum := 14

	sum := 0
	for _, line := range parsed {
		gameId := getGameId(line)
		red, green, blue := getMaxColourNums(line)

		if red <= redNum && green <= greenNum && blue <= blueNum {
			sum += gameId
		}
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	sum := 0
	for _, line := range parsed {
		red, green, blue := getMaxColourNums(line)

		sum += red * green * blue
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func getGameId(line string) int {
	gameIdMatch := `Game (\d+):`

	re := regexp.MustCompile(gameIdMatch)

	match := re.FindStringSubmatch(line)

	return cast.ToInt(match[1])
}

func getMaxColourNums(line string) (int, int, int) {
	maxRed := getMaxNumberBeforeString(line, "red")
	maxGreen := getMaxNumberBeforeString(line, "green")
	maxBlue := getMaxNumberBeforeString(line, "blue")

	return maxRed, maxGreen, maxBlue
}

func getMaxNumberBeforeString(str string, strToMatch string) int {
	numMatch := fmt.Sprintf(`(\d+) %s`, strToMatch)

	re := regexp.MustCompile(numMatch)

	matches := re.FindAllStringSubmatch(str, -1)

	maxNumber := 0
	for _, match := range matches {
		n := cast.ToInt(match[1])
		if n > maxNumber {
			maxNumber = n
		}
	}

	return maxNumber
}

// Helper functions for part 2
