package main

import (
	"advent-of-code-go/pkg/regex"
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
		ans := part1(input, 200000000000000, 400000000000000)
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
		ans1 := part1(input, 200000000000000, 400000000000000)
		fmt.Println("Part 1 Output:", ans1)
		ans2 := part2(input)
		fmt.Println("Part 2 Output:", ans2)
	}
}

func part1(input string, min, max float64) int {
	parsed := parseInput(input)

	ignoreZ := true

	h := getHailstones(parsed, ignoreZ)

	count := 0
	for i, h1 := range h {
		for j, h2 := range h {
			if i == j {
				continue
			}
			intersection, inPast := findIntersection(h1, h2)
			fmt.Println(h1.p, h2.p, inPast, intersection)
			if !inPast && pointInRange(intersection, min, max, ignoreZ) {
				count += 1
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

func pointInRange(point [3]float64, min, max float64, ignoreZ bool) bool {
	for i, coord := range point {
		if ignoreZ && i == 2 {
			continue
		}

		if !(min <= coord && coord <= max) {
			return false
		}
	}
	return true
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Hailstone struct {
	p [3]float64
	v [3]float64
}

func getHailstones(parsed []string, ignoreZ bool) []Hailstone {
	h := []Hailstone{}
	for _, line := range parsed {
		nums := regex.GetNumbers(line)
		if ignoreZ {
			h = append(h, Hailstone{
				p: [3]float64{float64(nums[0]), float64(nums[1]), 0.0},
				v: [3]float64{float64(nums[3]), float64(nums[4]), 0.0},
			})
		} else {
			h = append(h, Hailstone{
				p: [3]float64{float64(nums[0]), float64(nums[1]), float64(nums[2])},
				v: [3]float64{float64(nums[3]), float64(nums[4]), float64(nums[5])},
			})
		}
	}
	return h
}

func findIntersection(h1, h2 Hailstone) ([3]float64, bool) {
	xDiff := (h1.p[0] - h2.p[0])
	xVDiff := (h1.v[0] - h2.v[0])

	yDiff := (h1.p[1] - h2.p[1])
	yVDiff := (h1.v[1] - h2.v[1])

	zDiff := (h1.p[2] - h2.p[2])
	zVDiff := (h1.v[2] - h2.v[2])

	t := (xDiff*xVDiff + yDiff*yVDiff + zDiff*zVDiff) /
		(sqr(h1.v[0]) + sqr(h1.v[1]) + sqr(h1.v[2]) + sqr(h2.v[0]) + sqr(h2.v[1]) + sqr(h2.v[2]))

	intersectionPoint := [3]float64{h1.p[0] + t*h1.v[0], h1.p[1] + t*h1.v[1], h1.p[2] + t*h1.v[2]}

	crossedInPast := t > 0

	return intersectionPoint, crossedInPast
}

func sqr(n float64) float64 {
	return n * n
}

// Helper functions for part 2
