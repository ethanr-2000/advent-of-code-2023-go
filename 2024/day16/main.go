package main

import (
	"advent-of-code-go/pkg/grid"

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
	g := grid.GetGrid(parsed)

	path, _ := FindOptimalPath(g, grid.GetLocationsOfCharacter(g, 'S')[0], grid.GetLocationsOfCharacter(g, 'E')[0])

	return path
}

func part2(input string) int {
	parsed := parseInput(input)
	g := grid.GetGrid(parsed)

	_, seats := FindOptimalPath(g, grid.GetLocationsOfCharacter(g, 'S')[0], grid.GetLocationsOfCharacter(g, 'E')[0])

	return seats
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func IntLocation(l grid.Location) int {
	return l.X * 10000 + l.Y
}

type Node struct {
	location grid.Location
	direction grid.Direction
}

type State struct {
	reindeer Node
	path []grid.Location
	score int
}

func FindOptimalPath(g grid.Grid, start, end grid.Location) (minScore int, bestSeatCount int) {
	startState := State{
		reindeer: Node{start, grid.East},
		path: []grid.Location{start},
		score: 0,
	}

	minScore = math.MaxInt
	stateQueue := []State{startState}
	visited := make(map[int]map[int]int) // [location][direction] = visited
	scoreToPath := make(map[int][]grid.Location) // for each final score, what locations led there?

	for len(stateQueue) > 0 {
		currentState := stateQueue[0]
		stateQueue = stateQueue[1:]

		if currentState.score > minScore {
			continue
		}

		if grid.LocationsEqual(currentState.reindeer.location, end) {
			if currentState.score <= minScore {
				minScore = currentState.score
			}
			scoreToPath[currentState.score] = append(scoreToPath[currentState.score], currentState.path...)
			continue
		}

		for _, nextLocation := range grid.FourAdjacentList(currentState.reindeer.location) {
			if grid.ValueAtLocation(g, nextLocation) == '#' { continue }

			nextDir := grid.DirectionBetweenLocations(currentState.reindeer.location, nextLocation)
			if grid.OppositeDirections(currentState.reindeer.direction, nextDir) { continue }

			score := currentState.score + 1
			if currentState.reindeer.direction != nextDir {
				score += 1000
			}
			if previousScore, has := accessVisited(visited, nextLocation, nextDir); has {
				if previousScore < score { continue }
			}

			addToVisited(visited, nextLocation, nextDir, score)
			nextPath := make([]grid.Location, len(currentState.path))
			copy(nextPath, currentState.path)

			stateQueue = append(stateQueue, State{
				reindeer: Node{nextLocation, nextDir},
				path: append(nextPath, nextLocation),
				score: score,
			})
		}
	}

	return minScore, numberOfUniqueLocations(scoreToPath[minScore])
}

func numberOfUniqueLocations(ls []grid.Location) int {
	locationMap := make(map[int]bool)
	for _, l := range ls {
		locationMap[IntLocation(l)] = true
	}
	return len(locationMap)
}

func accessVisited(v map[int]map[int]int, l grid.Location, d grid.Direction) (val int, exists bool) {
	if _, e := v[IntLocation(l)]; !e {
		return -1, false
	}
	val, exists = v[IntLocation(l)][int(d)]
	return
}

func addToVisited(v map[int]map[int]int, l grid.Location, d grid.Direction, score int) {
	if _, exists := v[IntLocation(l)]; !exists {
		v[IntLocation(l)] = make(map[int]int)
	}
	v[IntLocation(l)][int(d)] = score
}

// Helper functions for part 2
