package main

import (
	"advent-of-code-go/pkg/cast"
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

	instructions := getInstructions(parsed)
	area := areaFromInstructions(instructions)

	return area
}

func part2(input string) int {
	parsed := parseInput(input)

	instructions := getInstructionsPt2(parsed)
	area := areaFromInstructions(instructions)

	return area
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Location struct {
	x int
	y int
}

type Direction rune

const (
	RIGHT Direction = 'R'
	DOWN  Direction = 'D'
	LEFT  Direction = 'L'
	UP    Direction = 'U'
)

var DirectionList = []Direction{RIGHT, DOWN, LEFT, UP}

type Instruction struct {
	direction Direction
	distance  int
}

func getInstructions(parsed []string) []Instruction {
	instructions := []Instruction{}
	for _, line := range parsed {
		d := Direction(rune(line[0]))
		num := cast.ToInt(strings.TrimRight(line[2:4], " "))
		instructions = append(instructions, Instruction{d, num})
	}
	return instructions
}

func areaFromInstructions(instructions []Instruction) int {
	perimeter := 0
	area := 0
	l := Location{0, 0}
	for _, i := range instructions {
		newl := move(l, i)
		dy := newl.y - l.y

		area += (newl.x) * (dy)
		perimeter += i.distance

		l = newl
	}

	return area + perimeter/2 + 1
}

func move(l Location, i Instruction) Location {
	if i.direction == UP {
		return Location{l.x, l.y - i.distance}
	} else if i.direction == RIGHT {
		return Location{l.x + i.distance, l.y}
	} else if i.direction == DOWN {
		return Location{l.x, l.y + i.distance}
	} else if i.direction == LEFT {
		return Location{l.x - i.distance, l.y}
	}
	return Location{}
}

// Helper functions for part 2

func getInstructionsPt2(parsed []string) []Instruction {
	instructions := []Instruction{}
	for _, line := range parsed {
		numStr := strings.TrimRight(line[6:], ")")
		numStr = strings.TrimLeft(numStr, "#")

		hexStr, dirNum := numStr[:len(numStr)-1], cast.ToInt(string((numStr[len(numStr)-1])))
		num := getIntFromHexString(hexStr)
		d := DirectionList[dirNum]

		instructions = append(instructions, Instruction{d, num})
	}
	return instructions
}

func getIntFromHexString(s string) int {
	result, _ := strconv.ParseInt(s, 16, 64)
	return int(result)
}
