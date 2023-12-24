package main

import (
	"advent-of-code-go/pkg/regex"
	_ "embed"
	"flag"
	"fmt"
	"reflect"
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

	sep := slices.Index[[]string](parsed, "")
	workflowLines, partLines := parsed[:sep], parsed[sep+1:]

	workflows := getWorkflows(workflowLines)
	parts := getParts(partLines)

	total := 0
	for _, p := range parts {
		accepted := partAccepted(p, workflows)
		if accepted {
			total += p.x + p.m + p.a + p.s
		}
	}

	return total
}

func part2(input string) int {
	parsed := parseInput(input)
	sep := slices.Index[[]string](parsed, "")
	workflowLines, _ := parsed[:sep], parsed[sep+1:]

	workflows := getWorkflows(workflowLines)

	ranges := []Range{}
	// initialise with one range of all possible values, starting at "in"
	narrowDownRange(&workflows, "in", Range{1, 4000, 1, 4000, 1, 4000, 1, 4000}, &ranges)

	return calculateUniquePossibilities(ranges)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Property rune

const (
	X          Property = 'x'
	M          Property = 'm'
	A          Property = 'a'
	S          Property = 's'
	NoProperty Property = 0
)

type Comparison struct {
	property  Property
	condition rune
	value     int
}

type Rule struct {
	comparison  Comparison
	destination string
}

type Workflow struct {
	name  string
	rules []Rule
}

type Part struct {
	x int
	m int
	a int
	s int
}

func getWorkflows(parsed []string) map[string]Workflow {
	workflows := map[string]Workflow{}
	for _, line := range parsed {
		w := getWorkflow(line)
		workflows[w.name] = w
	}
	return workflows
}

func getWorkflow(line string) Workflow {
	name, line, _ := strings.Cut(line, "{")
	line, _, _ = strings.Cut(line, "}")

	rawRules := strings.Split(line, ",")
	rules := []Rule{}
	for _, r := range rawRules {
		words := regex.GetWords(r)
		numbers := regex.GetNumbers(r)

		condition := ' '
		if regex.Contains(r, ">") {
			condition = '>'
		} else if regex.Contains(r, "<") {
			condition = '<'
		}

		if condition == ' ' {
			rules = append(rules, Rule{
				destination: words[0],
			})
		} else {
			rules = append(rules, Rule{
				comparison: Comparison{
					property:  Property(words[0][0]),
					condition: condition,
					value:     numbers[0],
				},
				destination: words[1],
			})
		}

	}

	return Workflow{name, rules}
}

func getParts(parsed []string) []Part {
	parts := []Part{}

	for _, line := range parsed {
		parts = append(parts, getPart(line))

	}
	return parts
}

func getPart(line string) Part {
	nums := regex.GetNumbers(line)

	return Part{nums[0], nums[1], nums[2], nums[3]}
}

func partAccepted(p Part, workflows map[string]Workflow) bool {
	nextWorkflow := "in"
	for (nextWorkflow != "A") && (nextWorkflow != "R") {
		w := workflows[nextWorkflow]

		for _, r := range w.rules {
			if r.comparison.condition == 0 {
				nextWorkflow = r.destination
				break
			} else if makeComparison(p, r.comparison) {
				nextWorkflow = r.destination
				break
			}
		}
	}

	return nextWorkflow == "A"
}

func makeComparison(p Part, c Comparison) bool {
	if c.condition == '>' {
		return accessProperty(p, c.property) > c.value
	}
	return accessProperty(p, c.property) < c.value
}

func accessProperty(p Part, property Property) int {
	r := reflect.ValueOf(p)
	field := reflect.Indirect(r).FieldByName(string(property))
	return int(field.Int())
}

// Helper functions for part 2

type Range struct {
	minx int
	maxx int
	minm int
	maxm int
	mina int
	maxa int
	mins int
	maxs int
}

func narrowDownRange(workflows *map[string]Workflow, nextWorkflowName string, activeRange Range, ranges *[]Range) {
	if nextWorkflowName == "A" {
		// the path we went down was accepted! The range represents good values
		*ranges = append(*ranges, activeRange)
		return
	} else if nextWorkflowName == "R" {
		// the path we went down was rejected
		return
	}

	w := (*workflows)[nextWorkflowName]

	for _, r := range w.rules {
		if r.comparison.condition == 0 {
			// this comparison did not change the accepted range, and no more comparisons should be made in this workflow
			narrowDownRange(workflows, r.destination, activeRange, ranges)
			break
		} else {
			// go to destination with accepted range
			acceptedRange := updateRangeAccepted(activeRange, r.comparison)
			narrowDownRange(workflows, r.destination, acceptedRange, ranges)

			// carry on in this workdlow with rejected range
			activeRange = updateRangeRejected(activeRange, r.comparison)
		}
	}
}

func updateRangeAccepted(r Range, c Comparison) Range {
	switch {
	case c.condition == '>':
		max := getMaxField(r, c.property)
		min := getMinField(r, c.property)
		if max < c.value {
			return Range{}
		} else if min > c.value {
			return r
		}
		return setField(r, c.property, false, c.value+1)
	case c.condition == '<':
		max := getMaxField(r, c.property)
		min := getMinField(r, c.property)
		if min > c.value {
			return Range{}
		} else if max < c.value {
			return r
		}
		return setField(r, c.property, true, c.value-1)
	}
	return Range{}
}

func updateRangeRejected(r Range, c Comparison) Range {
	switch {
	case c.condition == '<':
		max := getMaxField(r, c.property)
		min := getMinField(r, c.property)
		if max < c.value {
			return Range{}
		} else if min > c.value {
			return r
		}
		return setField(r, c.property, false, c.value)
	case c.condition == '>':
		max := getMaxField(r, c.property)
		min := getMinField(r, c.property)
		if min > c.value {
			return Range{}
		} else if max < c.value {
			return r
		}
		return setField(r, c.property, true, c.value)
	}
	return Range{}
}

func getMaxField(r Range, prop Property) int {
	switch prop {
	case X:
		return r.maxx
	case M:
		return r.maxm
	case A:
		return r.maxa
	case S:
		return r.maxs
	}
	return 0
}

func getMinField(r Range, prop Property) int {
	switch prop {
	case X:
		return r.minx
	case M:
		return r.minm
	case A:
		return r.mina
	case S:
		return r.mins
	}
	return 0
}

func setField(r Range, prop Property, max bool, value int) Range {
	if max {
		switch prop {
		case X:
			return Range{r.minx, value, r.minm, r.maxm, r.mina, r.maxa, r.mins, r.maxs}
		case M:
			return Range{r.minx, r.maxx, r.minm, value, r.mina, r.maxa, r.mins, r.maxs}
		case A:
			return Range{r.minx, r.maxx, r.minm, r.maxm, r.mina, value, r.mins, r.maxs}
		case S:
			return Range{r.minx, r.maxx, r.minm, r.maxm, r.mina, r.maxa, r.mins, value}
		}
	}
	switch prop {
	case X:
		return Range{value, r.maxx, r.minm, r.maxm, r.mina, r.maxa, r.mins, r.maxs}
	case M:
		return Range{r.minx, r.maxx, value, r.maxm, r.mina, r.maxa, r.mins, r.maxs}
	case A:
		return Range{r.minx, r.maxx, r.minm, r.maxm, value, r.maxa, r.mins, r.maxs}
	case S:
		return Range{r.minx, r.maxx, r.minm, r.maxm, r.mina, r.maxa, value, r.maxs}
	}
	return r
}

func calculateUniquePossibilities(ranges []Range) int {
	total := 0
	for _, r1 := range ranges {
		total += combinationsOfRange(r1)
	}
	return total
}

func combinationsOfRange(r Range) int {
	return (r.maxx - r.minx + 1) * (r.maxm - r.minm + 1) * (r.maxa - r.mina + 1) * (r.maxs - r.mins + 1)
}
