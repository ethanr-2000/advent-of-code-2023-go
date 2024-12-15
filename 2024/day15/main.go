package main

import (
	"advent-of-code-go/pkg/grid"

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
	g, instructions := parseInput(input)
	
	robotLocation := getRobotLocation(g)

	for i := range instructions {
		dirMap := map[rune]grid.Direction{
			'>': grid.East,
			'<': grid.West,
			'^': grid.North,
			'v': grid.South,
		}

		g, robotLocation = move(g, robotLocation, dirMap[rune(instructions[i])])
	}

	return sumBoxGps(g)
}

func part2(input string) int {
	g, instructions := parseInput(input)
	g = widenGrid(g)
	robotLocation := getRobotLocation(g)
	
	for i := range instructions {
		dirMap := map[rune]grid.Direction{
			'>': grid.East,
			'<': grid.West,
			'^': grid.North,
			'v': grid.South,
		}

		d := dirMap[rune(instructions[i])]
		if d == grid.East || d == grid.West {
			g, robotLocation = move(g, robotLocation, d)
		} else {
			g, robotLocation = moveWide(g, robotLocation, d)
		}
	}

	return sumBoxGps(g)
}

func parseInput(input string) (grid.Grid, string) {
	split := strings.Split(input, "\n\n")
	g := grid.GetGrid(strings.Split(split[0], "\n"))
	instructions := strings.ReplaceAll(split[1], "\n", "")
	return g, instructions
}

// Helper functions for part 1

func getRobotLocation(g grid.Grid) grid.Location {
	return grid.GetLocationsOfCharacter(g, '@')[0]
}

func move(g grid.Grid, l grid.Location, dir grid.Direction) (grid.Grid, grid.Location) {
	potentialNewLocation :=	grid.MoveStepsInDirection(l, dir, 1)
	
	if grid.ValueAtLocation(g, potentialNewLocation) == '.' {
		// moves freely
		g[potentialNewLocation.Y][potentialNewLocation.X] = g[l.Y][l.X]
		g[l.Y][l.X] = '.'

		return g, potentialNewLocation
	}

	if grid.ValueAtLocation(g, potentialNewLocation) == '#' {
		// can't move
		return g, l
	}

	// try move the box
	newG, potentialBoxMove := move(g, potentialNewLocation, dir)

	if !grid.LocationsEqual(potentialBoxMove, potentialNewLocation) {
		newG[potentialNewLocation.Y][potentialNewLocation.X] = newG[l.Y][l.X]
		newG[l.Y][l.X] = '.'

		return newG, potentialNewLocation
	}
	return g, l
}

func sumBoxGps(g grid.Grid) int {
	total := 0
	ls := grid.AllLocations(g)

	for _, l := range ls {
		if grid.ValueAtLocation(g, l) == 'O' || grid.ValueAtLocation(g, l) == '[' {
			total += l.X + l.Y*100
		}
	}

	return total
}

// Helper functions for part 2

func widenGrid(g grid.Grid) grid.Grid {
	ls := grid.AllLocations(g)

	wide := grid.InitialiseGrid(len(g[0])*2, len(g), '.')

	wideMap := map[rune][2]rune{
		'#': {'#', '#'},
		'@': {'@', '.'},
		'O': {'[', ']'},
		'.': {'.', '.'},
	}

	for _, l := range ls {
		wide[l.Y][2*l.X] = wideMap[g[l.Y][l.X]][0]
		wide[l.Y][2*l.X+1] = wideMap[g[l.Y][l.X]][1]
	}
	return wide
}

func moveWide(g grid.Grid, l grid.Location, dir grid.Direction) (grid.Grid, grid.Location) {
	// if dir == grid.East || dir == grid.West {
	// 	panic("moveWide should not be called for east and west moves")
	// }

	currentValue := grid.ValueAtLocation(g, l)
	coupledLocations := []grid.Location{l}
	if currentValue == '[' {
		coupledLocations = append(coupledLocations, grid.Location{X: l.X+1, Y: l.Y})
	} else if currentValue == ']' {
		coupledLocations = append(coupledLocations, grid.Location{X: l.X-1, Y: l.Y})
	}

	newG := grid.DeepCopyGrid(g)
	for _, cL := range coupledLocations {
		potentialNewLocation :=	grid.MoveStepsInDirection(cL, dir, 1)

		if grid.ValueAtLocation(newG, potentialNewLocation) == '#' {
			// can't move
			return g, l
		}

		if grid.ValueAtLocation(newG, potentialNewLocation) == '[' || grid.ValueAtLocation(newG, potentialNewLocation) == ']' {
			// try move the box in the way
			potentialNewG, potentialBoxMove := moveWide(newG, potentialNewLocation, dir)
	
			if grid.LocationsEqual(potentialBoxMove, potentialNewLocation) {
				// were unable to move the box, so return
				return g, l
			}

			newG = potentialNewG
		}
	}

	for _, cL := range coupledLocations {
		potentialNewLocation :=	grid.MoveStepsInDirection(cL, dir, 1)
		newG[potentialNewLocation.Y][potentialNewLocation.X] = newG[cL.Y][cL.X]
		newG[cL.Y][cL.X] = '.'
	}

	return newG, grid.MoveStepsInDirection(l, dir, 1)
}
