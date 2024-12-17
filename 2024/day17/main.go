package main

import (
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

func part1(input string) string {
	parsed := parseInput(input)
	computer := getComputer(parsed)

	output := process(computer)

	return output
}

func part2(input string) int {
	parsed := parseInput(input)

	output := ""
	c := getComputer(parsed)

	// stringInstructions := listInttoListString(c.instructions)

	// solution := 0
	// for i := 1; i <= len(c.instructions); i++ {
	// 	for offset := 0; offset < 10; offset++ {
	// 		c = getComputer(parsed)

	// 		testRegA := solution + offset + 8*i

	// 		c.regA = testRegA
	// 		output = process(c)

	// 		if len(strings.Split(output, ",")) < i {
	// 			continue
	// 		}

	// 		fmt.Println(strings.Split(output, ","), stringInstructions[len(c.instructions)-i:])
	// 		if slices.Equal[[]string](strings.Split(output, ","), stringInstructions[len(c.instructions)-i:]) {
	// 			fmt.Println(testRegA, output)
	// 			solution = testRegA
	// 			break
	// 		}
	// 	}
	// }

	// return solution

	valid := []int{0}
	for i := 1; i < len(c.instructions)+1; i++ {
		oldValid := valid
		valid = []int{}
		for _, num := range oldValid {
			for offset := 0; offset < 8; offset++ {
				c = getComputer(parsed)

				testRegA := 8*num + offset
				c.regA = testRegA

				output = process(c)
				listOutput := strings.Split(output, ",")

				if fmt.Sprintf("%v", listOutput) == fmt.Sprintf("%v", c.instructions[len(c.instructions)-i:]) {
					valid = append(valid, testRegA)
				}
			}
		}
	}

	answer := valid[0]
	for _, v := range valid {
		if v < answer {
			answer = v
		}
	}

	return answer
}

func listInttoListString(nums []int) []string {
	s := []string{}
	for i := range nums {
		s = append(s, fmt.Sprint(nums[i]))
	}
	return s
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Computer struct {
	regA               int
	regB               int
	regC               int
	instructions       []int
	instructionPointer int
}

func getComputer(parsed []string) Computer {
	return Computer{
		regA:               regex.GetNumbers(parsed[0])[0],
		regB:               regex.GetNumbers(parsed[1])[0],
		regC:               regex.GetNumbers(parsed[2])[0],
		instructions:       regex.GetNumbers(parsed[4]),
		instructionPointer: 0,
	}
}

func process(c Computer) string {
	output := ""
	for c.instructionPointer < len(c.instructions)-1 { // -1 because we add 1 to read operand
		i := c.instructions[c.instructionPointer]
		o := c.instructions[c.instructionPointer+1]

		switch i {
		case 0:
			// adv
			c.regA = c.regA / int(math.Pow(float64(2), float64(comboValue(c, o))))
		case 1:
			// bxl
			c.regB = c.regB ^ o
		case 2:
			// bst
			c.regB = comboValue(c, o) % 8
		case 3:
			// jnz
			if c.regA != 0 {
				c.instructionPointer = o - 2 // account for +2 later
			}
		case 4:
			// bxc
			c.regB = c.regB ^ c.regC
		case 5:
			// out
			if output != "" {
				output += ","
			}
			output += fmt.Sprint(comboValue(c, o) % 8)
		case 6:
			// bdv
			c.regB = c.regA / int(math.Pow(float64(2), float64(comboValue(c, o))))
		case 7:
			// cdv
			c.regC = c.regA / int(math.Pow(float64(2), float64(comboValue(c, o))))
		}
		c.instructionPointer += 2
	}

	return output
}

func comboValue(c Computer, o int) int {
	if 0 <= o && o <= 3 {
		return o
	}
	if o == 4 {
		return c.regA
	}
	if o == 5 {
		return c.regB
	}
	if o == 6 {
		return c.regC
	}
	fmt.Println("combo value not recognised", o)
	panic("AAAAA")
}

// Helper functions for part 2

/* NOTES
The program will only exit when regA == 0, as the last instruction is jump to beginning

The penultimate is 5,5, meaning output regB % 8

Before that, 0,3 divides regA by 8

Before that, 1,3 does regB XOR 3

Before that, 4,1 regB = regB XOR regC

Before that, 7,5, divides regA by 2^regB

Before that, 1,3 does regB XOR 3

To start, 2,4 sets regB = regA % 8

--

The program must run 16 times through, as there is only one output, and there are 16 instructions

to output 2,4,1,3,7,5,4,1,1,3,0,3,5,5,3,0:

regB needs to be 2,4,1,3,7,5,4,1,1,3,0,3,5,5,3,0

*/
