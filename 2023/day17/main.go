package main

import (
	"advent-of-code-go/pkg/cast"
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
	height := len(parsed)
	width := len(parsed[0])

	nodes := getNodes(parsed)

	startId := hashLocation(Location{0, 0})
	endId := hashLocation(Location{width - 1, height - 1})
	// endId := hashLocation(Location{3, 12})
	return minHeatLoss(nodes, startId, endId, parsed)
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

type Location struct {
	x int
	y int
}

type Node struct {
	id              string
	location        Location
	value           int
	minimumHeatLoss []int // length 40
	// visited         bool
	connectedNodes  [4]string
	pathHistory     []int
	locationHistory []Location
}

func defaultMinLoss(val int) []int {
	l := []int{}
	for i := 0; i < 40; i++ {
		l = append(l, val)
	}
	return l
}

func getNodes(parsed []string) map[string]*Node {
	nodes := map[string]*Node{}
	for y, line := range parsed {
		for x, num := range []rune(line) {
			connectedNodes := [4]string{"", "", "", ""}
			if y-1 >= 0 {
				up := hashLocation(Location{x, y - 1})
				connectedNodes[0] = up
			}
			if x+1 < len(parsed[0]) {
				right := hashLocation(Location{x + 1, y})
				connectedNodes[1] = right
			}
			if y+1 < len(parsed) {
				down := hashLocation(Location{x, y + 1})
				connectedNodes[2] = down
			}
			if x-1 >= 0 {
				left := hashLocation(Location{x - 1, y})
				connectedNodes[3] = left
			}

			n := Node{
				id:              hashLocation(Location{x, y}),
				location:        Location{x, y},
				value:           cast.ToInt(string(num)),
				minimumHeatLoss: defaultMinLoss(99999999999),
				// visited:         false,
				connectedNodes:  connectedNodes,
				pathHistory:     []int{},      // the directions 0,1,2,3 that were moved to get to the current node
				locationHistory: []Location{}, // the locations leading up to the current node
			}

			nodes[hashLocation(Location{x, y})] = &n
		}
	}
	return nodes
}

func hashLocation(l Location) string {
	return cast.ToString(l.x) + "_" + cast.ToString(l.y)
}

func movesStraightInDirection(history []int, dir int) int {
	length := len(history)

	num := 0
	for i := range history {
		if history[length-i-1] == dir {
			num += 1
		} else {
			return num
		}
	}
	return num
}

func possibleMinHeatLoss(minHeatLoss []int, value int, lossIndex int) int {
	return minHeatLoss[lossIndex] + value
}

func geLossIndex(minHeatLoss []int, value int, history []int, direction int) int {
	movesInLine := movesStraightInDirection(history, direction)
	return movesInLine*4 + direction
}

func minHeatLoss(nodes map[string]*Node, startId string, endId string, parsed []string) int {
	nodes[startId].minimumHeatLoss = defaultMinLoss(0)
	openSet := []*Node{nodes[startId]}

	for len(openSet) > 0 {
		slices.SortFunc[[]*Node](openSet, sortByHeatAscending)
		currentNode := openSet[0] // shortest node

		if currentNode.id == endId {
			drawPath(nodes, len(parsed[0]), len(parsed), startId, endId)
			return slices.Min[[]int](currentNode.minimumHeatLoss)
		}

		openSet = openSet[1:]

		// up, right, down, left
		//  0,     1,    2,    3
		for dir, nextNodeId := range currentNode.connectedNodes {
			if nextNodeId == "" {
				// the current node doesn't connect in this direction
				continue
			}
			if sameDirection(currentNode.pathHistory, dir) {
				continue
			}

			lossIndex := geLossIndex(currentNode.minimumHeatLoss, nodes[nextNodeId].value, currentNode.pathHistory, dir)
			possLoss := possibleMinHeatLoss(currentNode.minimumHeatLoss, nodes[nextNodeId].value, lossIndex)
			if possLoss < nodes[nextNodeId].minimumHeatLoss[lossIndex] {
				nodes[nextNodeId].minimumHeatLoss[lossIndex] = possLoss
				nodes[nextNodeId].locationHistory = append(currentNode.locationHistory, currentNode.location)
				nodes[nextNodeId].pathHistory = append(nodes[nextNodeId].pathHistory, dir)

				// if !slices.Contains[[]*Node](openSet, nodes[nextNodeId]) {
				openSet = append(openSet, nodes[nextNodeId])
				// }
			}
		}
	}
	fmt.Println(endId, nodes[endId].minimumHeatLoss, nodes[endId].locationHistory)
	drawPath(nodes, len(parsed[0]), len(parsed), startId, endId)
	return slices.Min[[]int](nodes[endId].minimumHeatLoss)
}

// func minHeatLoss(nodes map[string]*Node, startId string, endId string, parsed []string) int {
// 	currentNode := nodes[startId]
// 	currentNode.minimumHeatLoss = 0

// 	unvisitedNodes := []*Node{}
// 	for y, line := range parsed {
// 		for x := range line {
// 			unvisitedNodes = append(unvisitedNodes, nodes[hashLocation(Location{x, y})])
// 		}
// 	}

// 	// nextNodes := []*Node{}
// 	for len(unvisitedNodes) > 0 {
// 		// fmt.Println(nodes["0_9"].visited, nodes["0_9"].pathHistory)
// 		fmt.Println(currentNode.id, currentNode.minimumHeatLoss, currentNode.locationHistory)
// 		// up, right, down, left
// 		//  0,     1,    2,    3
// 		for direction, id := range currentNode.connectedNodes {
// 			if id == "" {
// 				continue
// 			}
// 			if sameDirection(currentNode.pathHistory, direction) {
// 				// don't consider this one if going there would result in moving in the same direction as the previous three steps
// 				continue
// 			}
// 			// it actually won't ever be visited
// 			// if nodes[possibleNextNodeId].visited {
// 			// 	return false
// 			// }
// 			next := nodes[id]

// 			if next.minimumHeatLoss > currentNode.minimumHeatLoss+next.value {
// 				// update the next node's heat loss and history if this path is better
// 				next.minimumHeatLoss = currentNode.minimumHeatLoss + next.value
// 				next.pathHistory = append(currentNode.pathHistory, direction)
// 				next.locationHistory = append(currentNode.locationHistory, currentNode.location)
// 			}

// 			// TODO it's probably to do with discounted nodes as visited too soon - this algorithm really favours going in a straight line until it can't anymore
// 			// currentNode.visited = true
// 		}

// 		slices.SortFunc[[]*Node](unvisitedNodes, sortByHeatAscending)
// 		currentNode = unvisitedNodes[0] // the one with the lowest heat loss
// 		unvisitedNodes = unvisitedNodes[1:]
// 	}
// 	fmt.Println(endId, nodes[endId].minimumHeatLoss, nodes[endId].locationHistory)
// 	drawPath(nodes, len(parsed[0]), len(parsed), startId, endId)
// 	return nodes[endId].minimumHeatLoss
// }

func sameDirection(d []int, d4 int) bool {
	l := len(d)

	if l < 3 {
		return false
	}

	// if (d[l-1]+2)%4 == d4 {
	// 	// if going backwards, continue
	// 	return false
	// }

	return d4 == d[l-3] && d4 == d[l-2] && d4 == d[l-1]
}

// draw
func drawPath(nodes map[string]*Node, height, width int, start, end string) {
	overwritePath(nodes, end)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(nodes[hashLocation(Location{x, y})].value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func overwritePath(nodes map[string]*Node, endId string) {
	// currentNode := nodes[endId]

	currentNode := nodes[endId]
	currentNode.value = 0
	// finalPathHistory := make([]int, len(nodes[endId].pathHistory))
	// copy(finalPathHistory, nodes[endId].pathHistory)
	// slices.Reverse[[]int](finalPathHistory)

	// for i := 0; i < len(currentNode.locationHistory); i++ {
	// 	fmt.Println(currentNode.locationHistory[i], currentNode.pathHistory[i])
	// }

	for _, l := range currentNode.locationHistory {
		// if currentNode == nil {
		// 	fmt.Println("Went out of bounds while drawing!")
		// 	return
		// }
		nodes[hashLocation(l)].value = 0
		// currentNode.value = 0

		// move in the reverse direction for each step
		// if finalPathHistory[0] == 0 {
		// 	// up
		// 	currentNode = nodes[hashLocation(Location{currentNode.location.x, currentNode.location.y + 1})]
		// } else if finalPathHistory[0] == 1 {
		// 	// right
		// 	currentNode = nodes[hashLocation(Location{currentNode.location.x - 1, currentNode.location.y})]
		// } else if finalPathHistory[0] == 2 {
		// 	// down
		// 	currentNode = nodes[hashLocation(Location{currentNode.location.x, currentNode.location.y - 1})]
		// } else if finalPathHistory[0] == 3 {
		// 	// left
		// 	currentNode = nodes[hashLocation(Location{currentNode.location.x + 1, currentNode.location.y})]
		// } else {
		// 	return
		// }
		// finalPathHistory = finalPathHistory[1:]
	}
}

// func deleteVisited(n *Node) bool {
// 	return n.visited
// }

func sortByHeatAscending(n1, n2 *Node) int {
	return slices.Min[[]int](n1.minimumHeatLoss) - slices.Min[[]int](n2.minimumHeatLoss)
}

// Helper functions for part 2
