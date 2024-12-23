package main

import (
	"advent-of-code-go/pkg/cache"
	"advent-of-code-go/pkg/grid"
	"advent-of-code-go/pkg/regex"
	_ "embed"
	"flag"
	"fmt"
	"math"
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
	// INITIALISATION FOR BOTH PARTS
	InitialiseKeypads()
	NumericPaths = GetAllPathsOnKeypad(NumericKeypad)
	DirectionalPaths = GetAllPathsOnKeypad(DirectionalKeypad)

	KeypadPaths = make(map[int]Paths)
	KeypadPaths[len(NumericKeypad)] = NumericPaths
	KeypadPaths[len(DirectionalKeypad)] = DirectionalPaths
  // INITIALISATION COMPLETE

	codes := parseInput(input)
	return ComplexityOfAllCodes(codes, 2)
}

func part2(input string) int {
	codes := parseInput(input)
	return ComplexityOfAllCodes(codes, 25)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

/* ======== */
/* Solution */

func ComplexityOfAllCodes(codes []string, robots int) int {
	keypads := make([]Keypad, 1 + robots)
	keypads[0] = NumericKeypad
	for i := 0; i < robots; i++ {
		keypads[i+1] = DirectionalKeypad
	}

	totalComplexity := 0
	for _, code := range codes {
		numActions := FindHumanActions(keypads, code, "A")
		totalComplexity += complexity(code, numActions)
	}
	return totalComplexity
}

func FindHumanActions(keypads []Keypad, code string, prevKey string) int {
	if a, exists := FIND_ACTIONS_CACHE.Get([]string{fmt.Sprint(len(keypads)), code, prevKey}); exists {
		return a
	}
	
	keypad := keypads[0]
	
	// get possible paths on the keypad from where we were to the first letter
	paths := KeypadPaths[len(keypad)][prevKey][string(code[0])]
	for i := 0; i < len(code)-1; i++ {
		oldPaths := paths
		paths = []string{}
		
		for _, p := range oldPaths {
			nextPaths, _ := KeypadPaths[len(keypad)][string(code[i])][string(code[i+1])]
			for _, nextP := range nextPaths {
				paths = append(paths, p+nextP)
			}
		}
	}
	paths = pruneInefficientPaths(paths)
	
	if len(keypads) == 1 {
		// we are the human! count the shortest sequence of button presses
		minLength := math.MaxInt
		for i := range paths {
			if length := len(paths[i]); length < minLength {
				minLength = length
			}
		}

		if err := FIND_ACTIONS_CACHE.Set([]string{"1", code, prevKey}, minLength); err == nil {
			return minLength
		} else {
			panic(err)
		}
	}

	humanActions := make([]int, len(paths))
	for pathNum, p := range paths {
		humanActions[pathNum] = 0
		for i := range p {
			prev := "A"
			if i > 0 { prev = string(p[i-1]) }
			// progressive calculation to take advantage of caching better
			// for every possible path at this level, the path is the new code
			humanActions[pathNum] += FindHumanActions(keypads[1:], string(p[i]), prev)
		}
	}

	minLength := slices.Min(humanActions)
	if err := FIND_ACTIONS_CACHE.Set([]string{fmt.Sprint(len(keypads)), code, prevKey}, minLength); err == nil {
		return minLength
	} else {
		panic(err)
	}
}

func complexity(code string, length int) int {
	return regex.GetNumbers(code)[0] * length
}

/* ===== */
/* Types */

type Keypad map[grid.Location]string

var NumericKeypad = make(Keypad, 11)
var DirectionalKeypad = make(Keypad, 5)

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

type Paths map[string]map[string][]string // start, end, paths

var KeypadPaths map[int]Paths // use the length of the keypad, i can't be bothered
var NumericPaths Paths
var DirectionalPaths Paths

/* ======*/
/* CACHE */

// cache[depth][code][prev] -> actions
var FIND_ACTIONS_CACHE = cache.NewCache[string, int]()

/* ============ */
/* Path Finding */

type State struct {
	l grid.Location
	path []grid.Location
	actions string
	score int
}

func GetAllPathsOnKeypad(k Keypad) Paths {
	paths := make(Paths)
	for startLocation, start := range k {
		paths[start] = make(map[string][]string)
		for endLocation, end := range k {
			paths[start][end] = pruneInefficientPaths(findPaths(k, startLocation, endLocation))
		}
	}
	return paths
}

func findPaths(k Keypad, start, end grid.Location) []string {
	startState := State{
		l: start,
		path: []grid.Location{start},
		actions: "",
		score: 0,
	}

	minScore := math.MaxInt
	stateQueue := []State{startState}
	visited := make(map[string]int) // [keystr] = costtothere
	paths := []string{}

	for len(stateQueue) > 0 {
		currentState := stateQueue[0]
		stateQueue = stateQueue[1:]

		if currentState.score > minScore {
			continue
		}

		if grid.LocationsEqual(currentState.l, end) {
			currentState.actions += "A" // Always end with input
			if currentState.score < minScore {
				minScore = currentState.score
				paths = []string{currentState.actions}
			} else if currentState.score == minScore {
				paths = append(paths, currentState.actions)
			}
			continue
		}

		for _, nextLocation := range grid.FourAdjacentList(currentState.l) {
			keyStr, exists := k[nextLocation]
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
				actions: currentState.actions + grid.DirectionCaratMap[dir],
				score: score,
			})
		}
	}
	return paths
}

func directionsAreGrouped(s string) bool {
	seen := make(map[rune]bool)
	lastDirection := rune(0)

	for _, c := range s {
		if c != 'A' {
			if c != lastDirection {
				if seen[c] {
					return false
				}
				seen[c] = true
				lastDirection = c
			}
		}
	}

	return true
}

func pruneInefficientPaths(ps []string) []string {
	result := []string{}
	for _, p := range ps {
		split := strings.Split(p, "A")
		hasChangesFlag := false
		for _, s := range split {
			if !directionsAreGrouped(s) {
				hasChangesFlag = true
			}
		}
		if !hasChangesFlag {
			result = append(result, p)
		}
	}
	return result
}

