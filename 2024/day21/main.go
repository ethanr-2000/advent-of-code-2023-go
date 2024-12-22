package main

import (
	"advent-of-code-go/pkg/grid"
	"advent-of-code-go/pkg/regex"
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
	codes := parseInput(input)
	InitialiseKeypads()

	totalComplexity := 0
	for _, code := range codes {
		fmt.Println("finding code", code)
		actions := InputCode(code, 2)
		c := complexity(code, actions)
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

		c := complexity(code, actions)
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
func InputCode(code string, numberOfRobots int) int {
	numericALocation, _ := FindKeyLocation(NumericKeypad, "A")

	keypads := make([]Keypad, 1 + numberOfRobots)
	keypads[0] = NumericKeypad
	for i := 0; i < numberOfRobots; i++ {
		keypads[i+1] = DirectionalKeypad
	}

	resetCache()
	actions := FindActions(keypads, numericALocation, code, "")

	fmt.Println(actions)
	return actions
	// return KeepAllShortestUnique(actions)[0]
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

var FIND_ACTIONS_CACHE = make(map[int]map[string]map[grid.Location]int) // cache[depth][code][start] -> actions

func accessCache(depth int, code string, l grid.Location) (actions int, exists bool) {
	if _, exists := FIND_ACTIONS_CACHE[depth]; !exists {
		return -1, false
	}
	if _, exists := FIND_ACTIONS_CACHE[depth][code]; !exists {
		return -1, false
	}

	actions, exists = FIND_ACTIONS_CACHE[depth][code][l]
	return
}

func addToCache(depth int, code string, l grid.Location, actionLength int) int {
	if _, exists := FIND_ACTIONS_CACHE[depth]; !exists {
		FIND_ACTIONS_CACHE[depth] = make(map[string]map[grid.Location]int)
	}
	if _, exists := FIND_ACTIONS_CACHE[depth][code]; !exists {
		FIND_ACTIONS_CACHE[depth][code] = make(map[grid.Location]int)
	}
	FIND_ACTIONS_CACHE[depth][code][l] = actionLength
	return actionLength
}

func resetCache() {
	FIND_ACTIONS_CACHE = make(map[int]map[string]map[grid.Location]int)
}

// need to split out finding the path at a depth, and going deeper

// find all shortest-distance paths 
func FindActions(keypads []Keypad, start grid.Location, codeRemaining string, actionsBefore string) int {
	if a, exists := accessCache(len(keypads), codeRemaining, start); exists {
		fmt.Println("Using cache for", codeRemaining, "from start", start, "at depth", len(keypads))
		return a
	}

	keypad := keypads[0]

	targetKey := string(codeRemaining[0])
	targetKeyLocation, exists := FindKeyLocation(keypad, targetKey)
	if !exists { panic("tried to go to key that doesn't exist") }
	codeRemaining = codeRemaining[1:]

	startState := State{
		l: start,
		path: []grid.Location{start},
		actions: actionsBefore,
		score: 0,
	}

	// minScore := math.MaxInt
	stateQueue := []State{startState}
	visited := make(map[string]int) // [keystr] = costtothere

	humanActions := []int{}

	for len(stateQueue) > 0 {
		currentState := stateQueue[0]
		stateQueue = stateQueue[1:]

		// if currentState.score > minScore {
		// 	continue
		// }

		if grid.LocationsEqual(currentState.l, targetKeyLocation) {
			// if currentState.score <= minScore {
				// minScore = currentState.score

				currentActions := currentState.actions + "A" // always finish by pressing the target button

				if len(codeRemaining) > 0 {
					// keep finding the path to the code
					// for all possible next steps, add them to the list
					actions := FindActions(keypads, currentState.l, codeRemaining, currentActions)
					humanActions = append(humanActions, actions)
					// fmt.Println(currentActions)
				} else {
					// we've done the code at this level
					if len(keypads) == 1 {
						// done!
						// fmt.Println("Human button solution", currentActions)
						humanActions = append(humanActions, len(currentActions))
					} else {
						// the path is now the code. deeper...
						startingLocation, _ := FindKeyLocation(keypads[1], "A")
						// fmt.Println("going deeper", currentActions)
						actions := FindActions(keypads[1:], startingLocation, currentActions, "")
						// fmt.Println("Back up, adding ", actions, "to humanActions")
						humanActions = append(humanActions, actions)
					}
				}
			// }
			continue
		}

		for _, nextLocation := range grid.FourAdjacentList(currentState.l) {
			keyStr, exists := keypad[nextLocation]
			if !exists { continue }
			
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
	// fmt.Println("Human actions for:", targetKey + codeRemaining, ":", humanActions, ". Depth", len(keypads))

	return addToCache(len(keypads), targetKey + codeRemaining, start, slices.Min[[]int](humanActions))
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
