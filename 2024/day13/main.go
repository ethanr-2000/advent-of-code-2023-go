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

var NOT_FOUND_COST = 9999999999999999

func part1(input string) int {
	machines := getMachines(input, 0)

	totalCost := 0
	for i := range machines {
		solution := getOptimalSolutionCost(machines[i])
		if solution != NOT_FOUND_COST {
			totalCost += solution
		}
	}

	return totalCost
}

func part2(input string) int {
	machines := getMachines(input, 10000000000000)
	totalCost := 0
	for i := range machines {
		solution := getOptimalSolutionCost(machines[i])
		if solution != NOT_FOUND_COST {
			totalCost += solution
		}
	}

	return totalCost
}

// Helper functions for part 1

type Machine struct {
	a     grid.Location
	b     grid.Location
	prize grid.Location
}

func getMachines(input string, error int) []Machine {
	lines := strings.Split(input, "\n")
	numLines := len(lines)

	machines := []Machine{}
	for i := 0; i < numLines; i += 4 {
		aNums := regex.GetNumbers(lines[i])
		bNums := regex.GetNumbers(lines[i+1])
		prizeNums := regex.GetNumbers(lines[i+2])

		machines = append(machines, Machine{
			a:     grid.Location{X: aNums[0], Y: aNums[1]},
			b:     grid.Location{X: bNums[0], Y: bNums[1]},
			prize: grid.Location{X: prizeNums[0] + error, Y: prizeNums[1] + error},
		})
	}
	return machines
}

func getOptimalSolutionCost(m Machine) int {
	// x*a + y*b = p
	minCost := NOT_FOUND_COST
	for x := 0; x < 100; x++ {
		if moveCost(x, 0) > minCost {
			return minCost
		}
		aMove := grid.MultiplyDistance(m.a, x)

		for y := 0; y < 100; y++ {
			if moveCost(x, y) > minCost {
				// it's only getting more expensive
				break
			}

			bMove := grid.MultiplyDistance(m.b, y)

			if grid.LocationsEqual(m.prize, grid.LocationAdd(aMove, bMove)) {
				minCost = moveCost(x, y)
			}
		}
	}
	return minCost
}

func moveCost(x, y int) int {
	return 3*x + y
}

// Helper functions for part 2
