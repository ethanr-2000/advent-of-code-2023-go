package main

import (
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
	parsed := parseInput(input)

	modules := getModules(parsed)

	totalLow := 0
	totalHigh := 0
	for i := 0; i < 1000; i++ {
		low, high := pressButton(modules)
		totalLow += low
		totalHigh += high
	}

	return totalLow * totalHigh
}

func part2(input string) int {
	parsed := parseInput(input)

	modules := getModules(parsed)

	cyclesToInputModules := []int{0, 0, 0, 0}
	targetModules := []string{"tr", "xm", "dr", "nh"}
	presses := 1
	for slices.Contains[[]int](cyclesToInputModules, 0) {
		highToModules := pressButtonAndCatchHighToModules(modules, targetModules)

		for i, high := range highToModules {
			if high && cyclesToInputModules[i] == 0 {
				cyclesToInputModules[i] = presses
			}
		}

		presses += 1
	}

	return findLowestCommonMultipleOfList(cyclesToInputModules)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type ModuleType byte

const (
	FlipFlop    ModuleType = '%'
	Conjunction ModuleType = '&'
	Broadcast   ModuleType = 'b'
)

type State bool

const (
	HIGH State = true
	LOW  State = false
)

type Module struct {
	name   string
	mType  ModuleType
	src    []string
	dest   []string
	state  State   // flip flop only
	memory []State // conjunction only
}

func getModules(parsed []string) map[string]*Module {
	modules := map[string]*Module{}

	// make initial modules
	for _, line := range parsed {
		moduleType := ModuleType(line[0])

		words := regex.GetWords(line)
		modules[words[0]] = &Module{
			name:  words[0],
			mType: moduleType,
			dest:  words[1:],
			state: false,
		}

		for _, destModule := range words[1:] {
			if _, ok := modules[destModule]; !ok {
				// create destination module in case it's an output
				modules[destModule] = &Module{
					name: destModule,
					src:  []string{words[0]},
				}
			}
		}
	}

	// set up sources
	for _, line := range parsed {
		words := regex.GetWords(line)
		for _, name := range words[1:] {
			// for every destination
			if module, ok := modules[name]; ok { // required for go
				// the destination has a source
				(*module).src = append((*module).src, words[0])
			}
		}
	}

	// initialise memory
	for _, m := range modules {
		if m.mType == Conjunction {
			m.memory = make([]State, len(m.src))
		}
	}
	return modules
}

type Action struct {
	src   string
	dest  string
	pulse State
}

func pressButton(modules map[string]*Module) (int, int) {
	totalLow := 1
	totalHigh := 0

	nextAction := []Action{{
		src:   "button",
		dest:  "broadcaster",
		pulse: LOW,
	}}

	for len(nextAction) > 0 {
		for _, a := range nextAction {
			destinations, nextPulse := getNextPulse(modules[a.dest], a.pulse, a.src)
			for _, d := range destinations {
				nextAction = append(nextAction, Action{
					src:   a.dest,
					dest:  d,
					pulse: nextPulse,
				})
				if nextPulse {
					totalHigh += 1
				} else {
					totalLow += 1
				}
			}
			nextAction = nextAction[1:] // remove this one
		}
	}
	return totalLow, totalHigh
}

func getNextPulse(m *Module, inputPulse State, srcModule string) ([]string, State) {
	if m.mType == FlipFlop {
		if inputPulse == LOW {
			(*m).state = !(*m).state
			return m.dest, m.state
		} else if inputPulse == HIGH {
			return []string{}, LOW
		}
	} else if m.mType == Conjunction {
		srcIndex := slices.Index[[]string](m.src, srcModule)
		(*m).memory[srcIndex] = inputPulse

		if !slices.Contains[[]State](m.memory, LOW) {
			// all high - send low pulse
			return m.dest, LOW
		} else {
			// at least one low - send high pulse
			return m.dest, HIGH
		}
	} else if m.mType == Broadcast {
		return m.dest, inputPulse
	}

	return []string{}, LOW
}

// Helper functions for part 2

func pressButtonAndCatchHighToModules(modules map[string]*Module, destModules []string) []bool {
	nextAction := []Action{{
		src:   "button",
		dest:  "broadcaster",
		pulse: LOW,
	}}

	lowToInputModules := make([]bool, len(destModules))

	for len(nextAction) > 0 {
		for _, a := range nextAction {
			destinations, nextPulse := getNextPulse(modules[a.dest], a.pulse, a.src)
			for _, d := range destinations {
				indexOfTargetModule := slices.Index[[]string](destModules, d)
				if indexOfTargetModule != -1 && nextPulse == LOW {
					// found!
					lowToInputModules[indexOfTargetModule] = true
				}

				nextAction = append(nextAction, Action{
					src:   a.dest,
					dest:  d,
					pulse: nextPulse,
				})
			}
			nextAction = nextAction[1:] // remove this one
		}
	}
	return lowToInputModules
}

func greatestCommonDivisor(a, b int) int {
	// use euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lowestCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}

func findLowestCommonMultipleOfList(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for _, num := range numbers[1:] {
		result = lowestCommonMultiple(result, num)
	}
	return result
}
