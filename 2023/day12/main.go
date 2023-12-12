package main

import (
	"advent-of-code-go/pkg/cast"
	"advent-of-code-go/pkg/list"
	"advent-of-code-go/pkg/regex"
	"advent-of-code-go/pkg/string_util"
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

	conditionRecords := getConditionRecords(parsed, 0)
	cache := make(map[string]int, 1000000)

	sum := 0
	for _, cr := range conditionRecords {
		sum += countWays(cr, &cache)
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	conditionRecords := getConditionRecords(parsed, 4)
	cache := make(map[string]int, 1000000)

	sum := 0
	for _, cr := range conditionRecords {
		sum += countWays(cr, &cache)
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type ConditionRecord struct {
	record string
	groups []int
}

func getConditionRecords(input []string, duplication int) []ConditionRecord {
	r := []ConditionRecord{}
	for _, line := range input {
		r = append(r, ConditionRecord{
			record: string_util.Repeat(strings.Split(line, " ")[0], duplication, "?") + "..", // append two dot so that there are fewer branches
			groups: list.Repeat[int](regex.GetNumbers(line), duplication),
		})
	}
	return r
}

func countWays(cr ConditionRecord, cache *map[string]int) int {
	if val, exists := (*cache)[hashConditionRecord(cr)]; exists {
		// we've seen this one before!
		return val
	}

	if len(cr.record) < list.Sum(cr.groups) {
		return cacheResult(cache, cr, 0) // out of records
	}

	// if regex.Count(cr.record, '.') == len(cr.record) {
	// 	if len(cr.groups) != 0 {
	// 		return cacheResult(cache, cr, 0) // all dot
	// 	}
	// }

	if len(cr.groups) == 0 {
		if regex.Contains(cr.record, "#") {
			return cacheResult(cache, cr, 0) // there are springs left that aren't in a group
		} else {
			return cacheResult(cache, cr, 1) // all springs are accounted for
		}
	}

	record := []rune(cr.record)

	if record[0] == '.' { // the next one is a dot, we don't care about that
		ways := countWays(ConditionRecord{
			record: strings.TrimLeft(cr.record, "."),
			groups: cr.groups,
		}, cache)
		return cacheResult(cache, cr, ways)
	}

	if record[0] == '?' {
		ifDot := countWays(ConditionRecord{
			record: string(record[1:]), // treat as a . and skip it
			groups: cr.groups,
		}, cache)

		ifHash := countWays(ConditionRecord{
			record: string_util.ChangeRuneAtIndex(cr.record, 0, '#'), // treat next as #
			groups: cr.groups,
		}, cache)

		return cacheResult(cache, cr, ifDot+ifHash)
	}

	if record[0] == '#' {
		groupLen := regex.LengthsOfGroupsOfChar(cr.record, '#')[0]
		if groupLen > cr.groups[0] {
			return cacheResult(cache, cr, 0) // the group is too long, invalid
		}
		if groupLen == cr.groups[0] {
			ways := countWays(ConditionRecord{
				record: string(record[groupLen+1:]), // there needs to be a space after a group
				groups: cr.groups[1:],               // done this group
			}, cache)
			return cacheResult(cache, cr, ways)
		}

		if record[groupLen] == '.' { // if the next spot after the group is a .
			return cacheResult(cache, cr, 0) // the group isn't long enough and we can't make it longer
		}
		// otherwise it's a ?
		ways := countWays(ConditionRecord{
			record: string_util.ChangeRuneAtIndex(cr.record, groupLen, '#'), // the group isn't long enough, but there is a space
			groups: cr.groups,
		}, cache)
		return cacheResult(cache, cr, ways)
	}

	fmt.Println("not sure what happened", cr)
	return 0
}

func cacheResult(c *map[string]int, cr ConditionRecord, res int) int {
	(*c)[hashConditionRecord(cr)] = res
	return res
}

func hashConditionRecord(cr ConditionRecord) string {
	hashedNums := make([]string, len(cr.groups))
	for i, n := range cr.groups {
		hashedNums[i] = cast.ToString(n)
	}
	return cr.record + strings.Join(hashedNums, " ")
}

// Helper functions for part 2
