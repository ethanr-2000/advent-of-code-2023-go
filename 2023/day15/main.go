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

	sum := 0
	for _, s := range parsed {
		sum += hash(s)
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	boxes := Boxes{}
	// for i := range boxes {
	// 	boxes[i] = make([]string, 0)
	// }

	for _, s := range parsed {
		if strings.HasSuffix(s, "-") {
			removeFromBox(s, &boxes)
		} else {
			addToBox(s, &boxes)
		}
	}

	return totalFocussingPower(boxes)
}

func parseInput(input string) []string {
	return strings.Split(input, ",")
}

// Helper functions for part 1

func hash(s string) int {
	h := 0
	for _, r := range []rune(s) {
		h += int(r)
		h *= 17
		h %= 256
	}
	return h
}

// Helper functions for part 2

type Boxes [256][]string

func label(s string) string {
	return strings.Split(strings.Split(s, "=")[0], "-")[0]
}

func focalLength(s string) int {
	return cast.ToInt(strings.Split(s, "=")[1])
}

func removeFromBox(s string, boxes *Boxes) {
	boxContents := (*boxes)[hash(label(s))]

	for i, b := range boxContents {
		if label(b) == label(s) {
			(*boxes)[hash(label(s))] = slices.Delete[[]string](boxContents, i, i+1)
		}
	}
}

func addToBox(s string, boxes *Boxes) {
	l := label(s)
	i := indexOfLabelInBox((*boxes)[hash(l)], l)

	if i != -1 {
		(*boxes)[hash(l)][i] = s
	} else {
		(*boxes)[hash(label(s))] = append((*boxes)[hash(label(s))], s)
	}
}

func indexOfLabelInBox(box []string, l string) int {
	for i := range box {
		if label(box[i]) == l {
			return i
		}
	}
	return -1
}

func totalFocussingPower(boxes Boxes) int {
	sum := 0
	for i := range boxes {
		for j := range boxes[i] {
			sum += focalLength(boxes[i][j]) * (i + 1) * (j + 1)
		}
	}
	return sum
}
