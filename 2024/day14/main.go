package main

import (
	"advent-of-code-go/pkg/grid"
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
	var robots []Robot
	for _, line := range parsed {
		robots = append(robots, getRobot(line))
	}

	seconds := 100
	// width, height := 11, 7
	width, height := 101, 103
	robots = simulateRobots(robots, seconds, width, height)

	return safetyFactor(robots, width, height)
}

func part2(input string) int {
	parsed := parseInput(input)
	var robots []Robot
	for _, line := range parsed {
		robots = append(robots, getRobot(line))
	}
	width, height := 101, 103
	
	start := 6243
	robots = simulateRobots(robots, start, width, height)
	fmt.Println(start)
	grid.PrintGrid(grid.GetGridFromLocations(robotLocations(robots), width, height, '.', '#'))

	return start
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Robot struct {
	p grid.Location
	v grid.Location
}

func getRobot(line string) Robot {
	nums := regex.GetNumbersWithNegative(line)

	return Robot{
		grid.Location{X: nums[0], Y: nums[1]},
		grid.Location{X: nums[2], Y: nums[3]},
	}
}

func simulateRobot(r Robot, seconds, width, height int) Robot {
	return Robot{
		grid.Location{X: ModularAdd(r.p.X, seconds*r.v.X, width), Y: ModularAdd(r.p.Y, seconds*r.v.Y, height)},
		r.v,
	}
}

func ModularAdd(a, b, mod int) int {
	return ((a % mod) + (b % mod) + mod) % mod
}

func safetyFactor(robots []Robot, width, height int) int {
	neCount, seCount, swCount, nwCount := 0, 0, 0, 0

	halfWidth := width / 2
	halfHeight := height / 2

	
	for i := range robots {
		if robots[i].p.X < halfWidth && robots[i].p.Y < halfHeight {
			nwCount++
		}
		if robots[i].p.X > halfWidth && robots[i].p.Y < halfHeight {
			neCount++
		}
		if robots[i].p.X < halfWidth && robots[i].p.Y > halfHeight {
			swCount++
		}
		if robots[i].p.X > halfWidth && robots[i].p.Y > halfHeight {
			seCount++
		}
	}
	return neCount * seCount * swCount * nwCount
}

// Helper functions for part 2

func simulateRobots(robots []Robot, seconds, width, height int) []Robot {
	for i, r := range robots {
		robots[i] = simulateRobot(r, seconds, width, height)
	}
	return robots
}

func robotLocations(robots []Robot) []grid.Location {
	ls := []grid.Location{}
	for _, r := range robots {
		ls = append(ls, r.p)
	}
	return ls
}