package main

import (
	"advent-of-code-go/pkg/regex"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	_ = parsed

	conditionRecords := getConditionRecords(parsed)

	sum := 0
	for _, cr := range conditionRecords {
		arr := countPossibleArrangements(cr)
		fmt.Println(arr)
		sum += arr
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type ConditionRecord struct {
	record string
	groups []int
}

func getConditionRecords(input []string) []ConditionRecord {
	r := []ConditionRecord{}
	for _, line := range input {
		r = append(r, ConditionRecord{
			record: strings.Split(line, " ")[0],
			groups: regex.GetNumbers(line),
		})
	}
	return r
}

func countPossibleArrangements(cr ConditionRecord) int {
	allPossibleCrs := placeNextDamagedSpring([]ConditionRecord{cr})

	// fmt.Println("\n", uniqueRecords(allPossibleCrs))

	return len(uniqueRecords(allPossibleCrs))
}

func getLengthOfGroupsOfCharacter(s string, c string) []int {
	re := regexp.MustCompile(fmt.Sprintf("%s+", c))
	matches := re.FindAllString(s, -1)

	groupLengths := []int{}
	for _, match := range matches {
		groupLengths = append(groupLengths, len(match))
	}
	return groupLengths
}

func isValidRecord(cr ConditionRecord) bool {
	groupLengths := getLengthOfGroupsOfCharacter(cr.record, "#")

	for i := range groupLengths {
		if groupLengths[i] != cr.groups[i] {
			return false
		}
	}
	return true
}

func isInvalidRecord(cr ConditionRecord) bool {
	groupLengths := getLengthOfGroupsOfCharacter(cr.record, "#")

	for i := range groupLengths {
		if groupLengths[i] != cr.groups[i] {
			return true
		}
	}
	return false
}

func changeRuneAtIndexOfString(s string, i int, c rune) string {
	strRune := []rune(s)
	strRune[i] = c
	return string(strRune)
}

func allSpringsOfAllRecordsPlaced(crs []ConditionRecord) bool {
	for i := 0; i < len(crs); i++ {
		if !allSpringsPlaced(crs[i]) {
			return false
		}
	}
	return true
}

func allSpringsPlaced(cr ConditionRecord) bool {
	desiredNum := 0
	for _, n := range cr.groups {
		desiredNum += n
	}

	return desiredNum <= len(regex.IndicesOfCharacter(cr.record, "#"))
}

func uniqueRecords(crs []ConditionRecord) []ConditionRecord {
	uniqueMap := make(map[string]bool)
	var unique []ConditionRecord

	for _, cr := range crs {
		if _, exists := uniqueMap[cr.record]; !exists {
			uniqueMap[cr.record] = true
			unique = append(unique, cr)
		}
	}

	return unique
}

func deleteIndicesOfSlice[T any](s []T, is []int) []T {
	slices.Sort[[]int](is)
	slices.Reverse[[]int](is)
	newS := make([]T, len(s))
	copy(newS, s)
	for _, i := range is {
		newS = slices.Delete(newS, i, i+1)
	}
	return newS
}

func conditionRecordInList(cr1 ConditionRecord, crs []ConditionRecord) bool {
	for _, cr2 := range crs {
		if conditionRecordsSame(cr1, cr2) {
			return true
		}
	}
	return false
}

func invalidSoFar(cr ConditionRecord) bool {
	groupLengths := getLengthOfGroupsOfCharacter(cr.record, "#")

	groupNum := 0
	for i := 0; i < len(groupLengths); i++ {
		for groupNum < len(cr.groups) {
			if groupLengths[i] > cr.groups[groupNum] { // you know it's invalid if the group is too long
				if groupNum == len(cr.groups)-1 { // if it's too big and it's the last one
					return true
				}
				groupNum++
			} else {
				groupNum++
				break
			}
		}
	}
	return false
}

func placeNextDamagedSpring(crs []ConditionRecord) []ConditionRecord {
	if allSpringsOfAllRecordsPlaced(crs) {
		crs = slices.DeleteFunc[[]ConditionRecord](crs, isInvalidRecord)
		return crs
	}

	invalidRecordIndices := []int{}
	for i := 0; i < len(crs); i++ {
		possibleLocations := regex.IndicesOfCharacter(crs[i].record, "?")

		if allSpringsPlaced(crs[i]) || len(possibleLocations) == 0 { // no possible next placement. finished either valid or invalid.
			if !isValidRecord(crs[i]) { // delete if invalid
				invalidRecordIndices = append(invalidRecordIndices, i)
			}
			continue
		}

		if len(possibleLocations) > 1 {
			for _, l := range possibleLocations[1:] { // don't want to append all locations, just the ones above 1
				newCr := ConditionRecord{
					record: crs[i].record,
					groups: crs[i].groups,
				}
				newCr.record = changeRuneAtIndexOfString(newCr.record, l, '#')

				if !conditionRecordInList(newCr, crs) && !invalidSoFar(newCr) {
					crs = append(crs, newCr)
				}
			}
		}

		newCr := ConditionRecord{
			record: crs[i].record,
			groups: crs[i].groups,
		}
		newCr.record = changeRuneAtIndexOfString(newCr.record, possibleLocations[0], '#')
		if !invalidSoFar(newCr) {
			crs[i].record = newCr.record
		}

	}

	trimmedCrs := deleteIndicesOfSlice[ConditionRecord](crs, invalidRecordIndices)
	// trimmedCrs = slices.DeleteFunc[[]ConditionRecord](trimmedCrs, invalidSoFar)
	trimmedCrs = uniqueRecords(trimmedCrs)
	// end := time.Now()
	// duration := end.Sub(start)

	// fmt.Println("Duration:", duration)
	return placeNextDamagedSpring(trimmedCrs)
}

func conditionRecordsSame(c1, c2 ConditionRecord) bool {
	if c1.record != c2.record || slices.Compare[[]int](c1.groups, c2.groups) != 0 {
		return false
	}
	return true
}

// Helper functions for part 2
