package main

import (
	"advent-of-code-go/pkg/set"
	"sort"

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
	
	triples := findTriplets(constructLanMap(parsed))
	
	triplesContainingT := [][]string{}
	for _, t := range triples {
		if t[0][0] == 't' || t[1][0] == 't' || t[2][0] == 't' {
			triplesContainingT = append(triplesContainingT, t)
		}
	}

	return len(triplesContainingT)
}

func part2(input string) string {
	parsed := parseInput(input)
	return largestInterconnectedNetwork(constructLanMap(parsed))
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Graph map[string]set.Set

func constructLanMap(parsed []string) Graph {
	m := make(Graph)
	for i := range parsed {
		computers := strings.Split(parsed[i], "-")
		if _, has := m[computers[0]]; !has {
			m[computers[0]] = make(set.Set)
		}
		if _, has := m[computers[1]]; !has {
			m[computers[1]] = make(set.Set)
		}
		m[computers[0]].Add(computers[1])
		m[computers[1]].Add(computers[0])
	}
	return m
}

func findTriplets(m Graph) [][]string {
	triples := [][]string{}
	for c1, c1Connections := range m {
		for c2 := range c1Connections {
			c2Connections := m[c2]
			for c3 := range c2Connections {
				c3Connections := m[c3]
				if _, has := c3Connections[c1]; has {
					triples = append(triples, []string{c1, c2, c3})
				}
			}	
		}
	}
	return uniqueNetwork(triples)
}

func password(network set.Set) string {
	sorted := network.ToSlice()
	sort.Strings(sorted)
	return strings.Join(sorted, ",")
}

func uniqueNetwork(network [][]string) [][]string {
	seen := make(map[string]bool)
	var result [][]string

	for _, n := range network {
		key := password(set.NewSetFromSlice(n))

		if !seen[key] {
			seen[key] = true
			result = append(result, n)
		}
	}

	return result
}

// Helper functions for part 2

func largestInterconnectedNetwork(m Graph) string {
	results := make(chan set.Set)

	potential := make(set.Set)
	for k := range m {
		potential.Add(k)
	}

	go func() {
		BronKerbosch(make(set.Set), potential, make(set.Set), m, results)
		close(results)
	}()

	longestPassword := ""
	for clique := range results {
		p := password(clique)
		if len(p) > len(longestPassword) {
			longestPassword = p
		}
	}

	return longestPassword
}

// https://www.geeksforgeeks.org/maximal-clique-problem-recursive-solution/
func BronKerbosch(current, potential, excluded set.Set, m Graph, results chan<- set.Set) {
	if len(potential) == 0 && len(excluded) == 0 {
		results <- current
		return
	}
	for len(potential) > 0 {
		computer := potential.Pop()
		BronKerbosch(
			current.Union(set.NewSetFromSlice([]string{computer})),
			potential.Intersection(m[computer]),
			excluded.Intersection(m[computer]),
			m,
			results,
		)
		excluded.Add(computer)
	}
}
