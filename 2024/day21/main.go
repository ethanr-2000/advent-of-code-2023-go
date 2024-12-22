package main

import (
	"advent-of-code-go/pkg/grid"
	"advent-of-code-go/pkg/regex"
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
	codes := parseInput(input)
	InitialiseKeypads()

	totalComplexity := 0
	for _, code := range codes {
		fmt.Println("finding code", code)
		actions := InputCode(code, 2)
		c := complexity(code, len(actions))
		fmt.Println(actions, c)
		totalComplexity += c
	}

	return totalComplexity
}

func part2(input string) int {
	codes := parseInput(input)
	InitialiseKeypads()

	totalComplexity := 0
	for _, code := range codes {
		fmt.Println("finding code", code)
		actions := InputCode(code, 25)

		c := complexity(code, len(actions))
		fmt.Println(actions, c)
		totalComplexity += c
	}

	return totalComplexity
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Keypad map[grid.Location]string

var NumericKeypad = make(Keypad, 11)
var DirectionalKeypad = make(Keypad, 5)

var DirectionMap = map[string]grid.Direction {
	"^": grid.North,
	"<": grid.West,
	"v": grid.South,
	">": grid.East,
}

var DirectionMapReverse = map[grid.Direction]string {
	grid.North: "^",
	grid.West: "<",
	grid.South: "v",
	grid.East: ">",
}

func FindKeyLocation(kp Keypad, key string) (grid.Location, bool) {
	for loc, val := range kp {
		if val == key {
			return loc, true
		}
	}
	return grid.Location{}, false
}

func InitialiseKeypads() {
	NumericKeypad[grid.Location{X: 0, Y: 0}] = "7"
	NumericKeypad[grid.Location{X: 1, Y: 0}] = "8"
	NumericKeypad[grid.Location{X: 2, Y: 0}] = "9"
	NumericKeypad[grid.Location{X: 0, Y: 1}] = "4"
	NumericKeypad[grid.Location{X: 1, Y: 1}] = "5"
	NumericKeypad[grid.Location{X: 2, Y: 1}] = "6"
	NumericKeypad[grid.Location{X: 0, Y: 2}] = "1"
	NumericKeypad[grid.Location{X: 1, Y: 2}] = "2"
	NumericKeypad[grid.Location{X: 2, Y: 2}] = "3"
	NumericKeypad[grid.Location{X: 1, Y: 3}] = "0"
	NumericKeypad[grid.Location{X: 2, Y: 3}] = "A"

	DirectionalKeypad[grid.Location{X: 1, Y: 0}] = "^"
	DirectionalKeypad[grid.Location{X: 2, Y: 0}] = "A"
	DirectionalKeypad[grid.Location{X: 0, Y: 1}] = "<"
	DirectionalKeypad[grid.Location{X: 1, Y: 1}] = "v"
	DirectionalKeypad[grid.Location{X: 2, Y: 1}] = ">"
}

type System struct {
	keypadPointer grid.Location
	robot1Pointer grid.Location
	robot2Pointer grid.Location
	humanActions []string
}

// for a given code, what's the optimal button presses to input that?
func InputCode(code string, numberOfRobots int) string {
	numericA, _ := FindKeyLocation(NumericKeypad, "A")
	directionalA, _ := FindKeyLocation(DirectionalKeypad, "A")

	resetCache()
	possibleNumericKeypadPaths := FindActions(NumericKeypad, numericA, code)

	resetCache() // numeric is no good to us here, we're going directional
	prevPossibleActions := possibleNumericKeypadPaths
	for i := 0; i < numberOfRobots; i++ {
		fmt.Println("considering robot", i+1, "/", numberOfRobots, ". With", len(prevPossibleActions), "possible paths")

		possibleNextActions := []string{}
		for _, numericPath := range prevPossibleActions {
			possibleNextActions = append(possibleNextActions, FindActions(DirectionalKeypad, directionalA, numericPath)...)
		}

		prevPossibleActions = KeepAllShortestUnique(possibleNextActions)
	}

	// fmt.Println(len(prevPossibleActions), "final possible actions")

	return prevPossibleActions[0]
}

func complexity(code string, length int) int {
	return regex.GetNumbers(code)[0] * length
}

type State struct {
	l grid.Location
	path []grid.Location
	actions string
	score int
}

var FIND_ACTIONS_CACHE = make(map[string]map[grid.Location][]string) // cache[code][start] -> actions

func accessCache(code string, l grid.Location) (actions []string, exists bool) {
	if _, exists := FIND_ACTIONS_CACHE[code]; !exists {
		return []string{}, false
	}

	actions, exists = FIND_ACTIONS_CACHE[code][l]
	return
}

func addToCache(code string, l grid.Location, actions []string) []string {
	if _, exists := FIND_ACTIONS_CACHE[code]; !exists {
		FIND_ACTIONS_CACHE[code] = make(map[grid.Location][]string)
	}
	FIND_ACTIONS_CACHE[code][l] = actions
	return actions
}

func resetCache() {
	FIND_ACTIONS_CACHE = make(map[string]map[grid.Location][]string)
}

// find all shortest-distance paths 
func FindActions(keypad Keypad, start grid.Location, codeRemaining string) ([]string) {
	if a, exists := accessCache(codeRemaining, start); exists {
		// fmt.Println("Using cache for", codeRemaining, "from start", start)
		return a
	}

	targetKey := string(codeRemaining[0])
	targetKeyLocation, exists := FindKeyLocation(keypad, targetKey)
	if !exists {
		panic("tried to go to key that doesn't exist")
	}
	codeRemaining = codeRemaining[1:]

	startState := State{
		l: start,
		path: []grid.Location{start},
		actions: "",
		score: 0,
	}

	minScore := math.MaxInt
	stateQueue := []State{startState}
	visited := make(map[string]int) // [keystr] = costtothere
	shortestActions := []string{}

	for len(stateQueue) > 0 {
		currentState := stateQueue[0]
		stateQueue = stateQueue[1:]

		if currentState.score > minScore {
			continue
		}

		if grid.LocationsEqual(currentState.l, targetKeyLocation) {
			if currentState.score <= minScore {
				minScore = currentState.score

				currentActions := currentState.actions + "A" // always finish by pressing the target button

				if len(codeRemaining) > 0 {
					// keep finding the path to the code
					// for all possible next steps, add them to the list
					actionsFromHere := FindActions(keypad, currentState.l, codeRemaining)
					for _, a := range actionsFromHere {
						shortestActions = append(shortestActions, currentActions + a)
					}
				} else {
					// code is input
					shortestActions = append(shortestActions, currentActions) 
				}
			}
			continue
		}

		for _, nextLocation := range grid.FourAdjacentList(currentState.l) {
			keyStr, exists := keypad[nextLocation]
			if !exists {
				continue
			}
			
			score := currentState.score + 1
			
			if previousScore, has := visited[keyStr]; has {
				if previousScore < score { continue }
			}
			
			visited[keyStr] = score
			nextPath := make([]grid.Location, len(currentState.path))
			copy(nextPath, currentState.path)
			
			dir := grid.DirectionBetweenLocations(currentState.l, nextLocation)

			stateQueue = append(stateQueue, State{
				l: nextLocation,
				path: append(nextPath, nextLocation),
				actions: currentState.actions + DirectionMapReverse[dir],
				score: score,
			})
		}
	}

	// shortest actions may contain sub-optimal paths
	// need to filter to only those with the min length
	shortestActions = KeepAllShortestUnique(shortestActions)

	return addToCache(targetKey + codeRemaining, start, shortestActions)
}

func KeepAllShortestUnique(actions []string) []string {
	TOLERANCE := 0

	if len(actions) == 0 {
		return nil
	}

	shortest := findShortest(actions)

	filteredActions := []string{shortest}
	seen := map[string]bool{shortest: true}

	for _, a := range actions {
		if seen[a] { continue }
		if len(a) - len(shortest) <= TOLERANCE {
			filteredActions = append(filteredActions, a)
			seen[a] = true
		}
	}

	return filteredActions
}

func findShortest(ss []string) string {
	shortest := ss[0]
	for _, s := range ss {
		if len(s) < len(shortest) {
			shortest = s
		}
	}
	return shortest
}

// Helper functions for part 2
