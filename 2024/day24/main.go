package main

import (
	"advent-of-code-go/pkg/regex"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"math/rand"
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
	circuit := getCircuit(input)
  return circuit.Simulate()
}

func part2(input string) string {
	circuit := getCircuit(input)
	circuit.Simulate()
	
	swapped := circuit.Fix()
	return strings.Join(swapped, ",")
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// Helper functions for part 1

type Wire struct {
	name string
	state bool
	on bool
}

type GateType string

const (
	AND    GateType = "AND"
	OR GateType = "OR"
	XOR   GateType = "XOR"
)

type Gate struct {
	gateType GateType
	inputs []*Wire
	output *Wire
}

type Circuit struct {
	wires map[string]*Wire
	gates []Gate
}

func getCircuit(input string) Circuit {
	wires := make(map[string]*Wire)
	gates := []Gate{}
	split := strings.Split(input, "\n\n")

	inputWires := strings.Split(split[0], "\n")
	inputGates := strings.Split(split[1], "\n")

	for _, inputWire := range inputWires {
		info := strings.Split(inputWire, ": ")
		state := false
		if info[1] == "1" {
			state = true
		}
		wires[info[0]] = &Wire{
			name: info[0],
			state: state,
			on: true,
		}
	}

	for _, gate := range inputGates {
		info := strings.Split(gate, " ")

		// might be some new wires here
		for _, w := range []string{info[0], info[2], info[4]} {
			if _, exists := wires[w]; !exists {
				wires[w] = &Wire{
					name: w,
					state: false, // arbitrarily select false
					on: false,
				}
			}
		}

		gateType := AND
		if info[1] == "OR" {
			gateType = OR
		} else if info[1] == "XOR" {
			gateType = XOR
		}

		gates = append(gates, Gate{
			gateType: gateType,
			inputs: []*Wire{
				wires[info[0]],
				wires[info[2]],
			},
			output: wires[info[4]],
		})
	}

	return Circuit{
		wires: wires,
		gates: gates,
	}
}

func (c *Circuit) Simulate() int {
	onWires := c.GetInputWires()

	// xWires := c.GetWiresWithPrefix("x")
	// fmt.Println("In simulate:")
	// for i := range xWires {
	// 	fmt.Print(xWires[len(xWires)-1-i].state, " ")
	// }
	// fmt.Println("=========")
	for i := range onWires {
		fmt.Println(&onWires[i], onWires[i].name, onWires[i].state)
	}
	
	gates := c.gates
	for len(gates) > 0 {
		g := gates[0]
		gates = gates[1:]

		if slices.Contains(onWires, g.inputs[0]) && slices.Contains(onWires, g.inputs[1]) {
			c.UpdateGate(&g)
			onWires = append(onWires, g.output)
		} else {
			// move to the end, we'll come back to it once the wires are on
			gates = append(gates, g)
		}
	}
	return c.GetOutput()
}

func (c *Circuit) GetOutput() int {
	z := 0
	for _, wire := range c.GetOutputWires() {
		val := 0
		if wire.state { val = 1 }
		bits := regex.GetNumbers(wire.name)[0]

		z += val << bits
	}
	return z
}

func (c *Circuit) GetWiresWithPrefix(p string) []*Wire {
	outWires := []*Wire{}
	for name := range c.wires {
		if strings.HasPrefix(name, p) {
			outWires = append(outWires, c.wires[name])
		}
	}
	slices.SortFunc(outWires, func(a, b *Wire) int {return strings.Compare(a.name, b.name)})
	return outWires
}

func (c *Circuit) GetInputWires() []*Wire {
	outWires := []*Wire{}
	outWires = append(outWires, c.GetWiresWithPrefix("x")...)
	outWires = append(outWires, c.GetWiresWithPrefix("y")...)
	return outWires
}

func (c *Circuit) GetOutputWires() []*Wire {
	return c.GetWiresWithPrefix("z")
}

func (c *Circuit) FindGateWithInput(wire *Wire) *Gate {
	for i, gate := range c.gates {
		if gate.inputs[0] == wire || gate.inputs[1] == wire {
			return &c.gates[i]
		}
	}
	return nil
}

func (c *Circuit) UpdateGate(g *Gate) {
	if g.output.on { return }
	for _, input := range g.inputs {
		if !input.on { return }
	}
	
	g.output.on = true
	switch g.gateType {
	case AND:
		g.output.state = g.inputs[0].state && g.inputs[1].state
	case OR:
		g.output.state = g.inputs[0].state || g.inputs[1].state
	case XOR:
		g.output.state = g.inputs[0].state != g.inputs[1].state
	}
}

// Helper functions for part 2

func (c *Circuit) GetInputs() []int {
	x := 0
	y := 0
	for name, wire := range c.wires {
		if strings.HasPrefix(name, "x") {
			val := 0
			if wire.state { val = 1 }
			bits := regex.GetNumbers(name)[0]
	
			x += val << bits
		}
		if strings.HasPrefix(name, "y") {
			val := 0
			if wire.state { val = 1 }
			bits := regex.GetNumbers(name)[0]
	
			y += val << bits
		}
	}
	return []int{x, y}
}

func (c *Circuit) CircuitIsCorrect() bool {
	inputs := c.GetInputs()
	return inputs[0] + inputs[1] == c.GetOutput()
}

func (c *Circuit) Fix() []string {
	// swappedGates := []string{}
	// outWires := c.GetOutputWires()

	// inputs := c.GetInputs()
	// output := c.GetOutput()

	// wiresSwapped := []string{}

	for i := range c.gates {
		for j := range c.gates {
			if i == j { continue }
			for k := range c.gates {
				if i == k || j == k { continue }
				for l := range c.gates {
					if i == l || j == l || k == l { continue }
					c.gates[i].output, c.gates[j].output = c.gates[j].output, c.gates[i].output
					c.gates[k].output, c.gates[l].output = c.gates[l].output, c.gates[k].output

					if c.Test() {
						return []string{
							c.gates[i].output.name,
							c.gates[j].output.name,
							c.gates[k].output.name,
							c.gates[l].output.name,
						}
					}
				}
			}
		}
	}

	return []string{}



	// // starting with LSB
	// for z := range outWires {
	// 	if c.TestBit(z) { continue } // we good

	// 	found := false
	// 	for i := 0; i < len(c.gates); i++ {
	// 		if found { break }
	// 		if slices.Contains(wiresSwapped, c.gates[i].output.name) { continue } // already swapped
	// 		for j := i+1; j < len(c.gates); j++ {
	// 			if found { break }
	// 			if slices.Contains(wiresSwapped, c.gates[j].output.name) { continue } // already swapped

	// 			// swap
	// 			c.gates[i].output, c.gates[j].output = c.gates[j].output, c.gates[i].output

	// 			swapWorked := true
	// 			for testZ := 0; testZ <= z; testZ++ {
	// 				if !c.TestBit(testZ) {
	// 					// swap back and break
	// 					c.gates[i].output, c.gates[j].output = c.gates[j].output, c.gates[i].output
	// 					swapWorked = false
	// 					break
	// 				}
	// 			}
	// 			if swapWorked {
	// 				wiresSwapped = append(wiresSwapped, c.gates[i].output.name, c.gates[j].output.name)
	// 				found = true
	// 			}
	// 		}
	// 	}
	// }

	// slices.Sort(wiresSwapped)
	// return wiresSwapped
}

func (c *Circuit) Test() bool {
	testCases := 100
	numbers := make([]int, testCases)
	maxOut := int(math.Pow(float64(2), float64(len(c.GetOutputWires())-1)))-1
	for i := 0; i < testCases; i++ {
		numbers[i] = rand.Intn(maxOut) // Generate number in range [0, 1,000,000,000]
	}

	for i, x := range numbers[:testCases/2] {
		y := numbers[i*2]
		wrongBits := c.TestCase(x, y) // wrong bits should be 0 if correct

		if wrongBits != 0 {
			return false
		}
	}
	return true
}

func (c *Circuit) TestBit(z int) bool {
	testCases := 10000
	numbers := make([]int, testCases)
	for i := 0; i < testCases; i++ {
		numbers[i] = rand.Intn(1_000_000_0) // Generate number in range [0, 1,000,000,000]
	}

	for i, x := range numbers[:testCases/2] {
		y := numbers[i*2]
		wrongBits := c.TestCase(x, y) // wrong bits should be 0 if correct
		if wrongBits & (1 << z) != 0 {
			return false
		}
	}
	return true
}

func (c *Circuit) TestCase(x int, y int) int {
	xWires := c.GetWiresWithPrefix("x")
	yWires := c.GetWiresWithPrefix("y")

	// set inputs
	for i := range xWires {
		c.wires[xWires[i].name].state = x & (1 << i) != 0
	}
	for i := range yWires {
		c.wires[yWires[i].name].state = y & (1 << i) != 0
	}

	for i := range xWires {
		fmt.Println(&xWires[i], xWires[i].name, xWires[i].state)
	}
	for i := range yWires {
		fmt.Println(&yWires[i], yWires[i].name, yWires[i].state)
	}

	fmt.Println("Output:", c.Simulate(), ". Should be", x+y)

	// return 0 if the output is correct
	return c.Simulate() ^ (x + y)
}

