package main

import (
	"advent-of-code-go/pkg/regex"

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
	var nums []int 
	for _, line := range parsed {
		nums = append(nums, regex.GetNumbers(line)[0])
	}

	sum := 0
	for _, n := range nums {
		allSecretNumbers := generate(n, 2000)
		sum += allSecretNumbers[len(allSecretNumbers)-1]
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	var nums []int 
	for _, line := range parsed {
		nums = append(nums, regex.GetNumbers(line)[0])
	}

	monkeys := []Monkey{}
	for _, n := range nums {
		allSecretNumbers := generate(n, 2000)
		p := prices(allSecretNumbers)
		c := diffs(p)
		sMap := make(map[[4]int]int)
		
		for i := 0; i < len(c) - 3; i++ {
			s := [4]int{c[i], c[i+1], c[i+2], c[i+3]}
			if _, exists := sMap[s]; !exists {
				sMap[s] = p[i+4] // 4 because prices starts with 1 before changes
			}
		}
		
		monkeys = append(monkeys, Monkey{
			prices: p,
			changes: c,
			sMap: sMap,
		})
	}

	r, s := findReturnOfOptimalSequence(monkeys)

	fmt.Println(s, "==>", r)
	return r
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

func generate(n int, times int) []int {
	secrets := []int{n}
	s := n
	for i := 0; i < times; i++ {
		s = prune(mix(64 * s, s))
		s = prune(mix(s / 32, s))
		s = prune(mix(2048 * s, s))
		secrets = append(secrets, s)
	}
	return secrets
}

func mix(n1, n2 int) int {
	return n1 ^ n2
}

func prune(n int) int {
	return n % 16777216
}

// Helper functions for part 2

func prices(ns []int) []int {
	ps := []int{}
	for _, n := range ns {
		ps = append(ps, price(n))
	}
	return ps
}

func price(n int) int {
	return n % 10
}

func diffs(ns []int) []int {
	ds := []int{}
	for i := 0; i < len(ns)-1; i++ {
		ds = append(ds, ns[i+1]-ns[i])
	}
	return ds
}

type Monkey struct {
	prices []int
	changes []int
	sMap map[[4]int]int
}

func findReturnOfOptimalSequence(monkeys []Monkey) (int, [4]int) {
	sequences := map[[4]int]int{}

	// for each sequence of changes in each monkey,
	// add the price to the total for that sequence
	for _, m := range monkeys {
		for k, v := range m.sMap {
			if _, exists := sequences[k]; !exists {
				sequences[k] = 0
			}
			sequences[k] += v
		}
	}
	
	bestReturn := 0
	bestSequence := [4]int{}
	for s, total := range sequences {
		if total > bestReturn {
			bestReturn = total
			bestSequence = s
		}
	}
	return bestReturn, bestSequence
}
