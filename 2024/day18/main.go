package main

import (
	"advent-of-code-go/pkg/grid"
	"advent-of-code-go/pkg/regex"

	"container/heap"
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

	w, h := 71, 71
	g := grid.InitialiseGrid(w, h, '.')
	for i, line := range parsed {
		if i == 1024 {
			break
		}
		nums := regex.GetNumbers(line)
		g[nums[1]][nums[0]] = '#'
	}

	length, _ := FindOptimalPath(g, grid.Location{X: 0, Y: 0}, grid.Location{X: w - 1, Y: h - 1})
	return length
}

func part2(input string) string {
	parsed := parseInput(input)

	w, h := 71, 71
	g := grid.InitialiseGrid(w, h, '.')
	for _, line := range parsed {
		nums := regex.GetNumbers(line)
		g[nums[1]][nums[0]] = '#'

		_, found := FindOptimalPath(g, grid.Location{X: 0, Y: 0}, grid.Location{X: w - 1, Y: h - 1})
		if !found {
			return line
		}
	}
	return "could always find a path"
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type State struct {
	l              grid.Location
	path           []grid.Location
	score          int
	heuristicScore int
}

func IntLocation(l grid.Location) int {
	return l.X*10000 + l.Y
}

type QueueItem struct {
	location grid.Location
	score    int
	path     []grid.Location
}

func FindOptimalPath(g grid.Grid, start, end grid.Location) (minScore int, found bool) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	startState := QueueItem{
		location: start,
		score:    0,
		path:     []grid.Location{start},
	}
	heap.Push(pq, startState)

	visited := make(map[int]int)

	for pq.Len() > 0 {
		currentState := heap.Pop(pq).(QueueItem)

		if bestScore, exists := visited[IntLocation(currentState.location)]; exists && bestScore < currentState.score {
			continue
		}

		if grid.LocationsEqual(currentState.location, end) {
			return currentState.score, true
		}

		for _, nextLocation := range grid.FourAdjacentList(currentState.location) {
			if grid.LocationOutsideGrid(nextLocation, g) || grid.ValueAtLocation(g, nextLocation) == '#' {
				continue
			}

			newScore := currentState.score + 1

			if bestScore, exists := visited[IntLocation(nextLocation)]; exists && bestScore <= newScore {
				continue
			}

			visited[IntLocation(nextLocation)] = newScore

			nextState := QueueItem{
				location: nextLocation,
				score:    newScore,
				path:     append(append([]grid.Location{}, currentState.path...), nextLocation),
			}

			heap.Push(pq, nextState)
		}
	}

	// No path found
	return math.MaxInt, false
}

type PriorityQueue []QueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(QueueItem))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Helper functions for part 2
