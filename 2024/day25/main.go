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
	keys, height := GetKeys(input)
	locks := GetLocks(input)

	count := 0
	for k := range keys {
		for l := range locks {
			if KeyFitsLock(keys[k], locks[l], height) {
				count++
			}
		}
	}

	return count
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

func GetKeys(input string) (keys [][]int, height int) {
	split := strings.Split(input, "\n\n")

	keys = [][]int{}
	for _, s := range split {
		rows := strings.Split(s, "\n")

		if strings.Contains(rows[0], "#") { continue }

		numPins := len(rows[0])
		height = len(rows)
		pins := make([]int, numPins)
		for pin := range pins {
			pins[pin] = height
			for y := range rows {
				if rows[y][pin] == '#' {
					break
				}
				pins[pin]--
			}
		}
		keys = append(keys, pins)
	}
	return keys, height
}

func GetLocks(input string) [][]int {
	split := strings.Split(input, "\n\n")

	locks := [][]int{}
	for _, s := range split {
		rows := strings.Split(s, "\n")

		if strings.Contains(rows[0], ".") { continue }

		numPins := len(rows[0])
		pins := make([]int, numPins)
		for pin := range pins {
			pins[pin] = 0
			for y := range rows {
				if rows[y][pin] == '.' {
					break
				}
				pins[pin]++
			}
		}
		locks = append(locks, pins)
	}
	return locks
}

func KeyFitsLock(key, lock []int, height int) bool {
	for pin := range key {
		if key[pin] + lock[pin] > height {
			return false
		}
	}
	return true
}

// Helper functions for part 2
