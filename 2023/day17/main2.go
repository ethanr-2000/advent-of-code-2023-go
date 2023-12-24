package main

// import (
// 	"advent-of-code-go/pkg/cast"
// 	_ "embed"
// 	"flag"
// 	"fmt"
// 	"slices"
// 	"strings"

// 	"github.com/atotto/clipboard"
// )

// //go:embed input.txt
// var input string

// func init() {
// 	// do this in init (not main) so test file has same input
// 	input = strings.TrimRight(input, "\n")
// 	if len(input) == 0 {
// 		panic("empty input.txt file")
// 	}
// }

// func main() {
// 	var part int
// 	flag.IntVar(&part, "part", 0, "part 1 or 2")
// 	flag.Parse()

// 	if part == 1 {
// 		ans := part1(input)
// 		fmt.Println("Running part 1")
// 		clipboard.WriteAll(fmt.Sprintf("%v", ans))
// 		fmt.Println("Output:", ans)
// 	} else if part == 2 {
// 		ans := part2(input)
// 		fmt.Println("Running part 2")
// 		clipboard.WriteAll(fmt.Sprintf("%v", ans))
// 		fmt.Println("Output:", ans)
// 	} else {
// 		fmt.Println("Running all")
// 		ans1 := part1(input)
// 		fmt.Println("Part 1 Output:", ans1)
// 		ans2 := part2(input)
// 		fmt.Println("Part 2 Output:", ans2)
// 	}
// }

// func part1(input string) int {
// 	parsed := parseInput(input)
// 	height := len(parsed)
// 	width := len(parsed[0])

// 	grid := getGrid(parsed)

// 	start := Location{0, 0}
// 	end := Location{width - 1, height - 1}
// 	return minHeatLoss(grid, start, end)
// }

// func part2(input string) int {
// 	parsed := parseInput(input)
// 	_ = parsed

// 	return 0
// }

// func parseInput(input string) []string {
// 	return strings.Split(input, "\n")
// }

// // Helper functions for part 1

// type Location struct {
// 	x int
// 	y int
// }

// type Node struct {
// 	value       int
// 	minHeatLoss [40]int
// 	l           Location
// }

// var Directions = []Location{
// 	Location{0, -1}, // North
// 	Location{1, 0},  // East
// 	Location{0, 1},  // South
// 	Location{-1, 0}, // West
// }

// func defaultMinLoss() [40]int {
// 	l := [40]int{}
// 	for i := 0; i < 40; i++ {
// 		l[i] = 9999999999
// 	}
// 	return l
// }

// func getGrid(parsed []string) [][]*Node {
// 	grid := make([][]*Node, len(parsed))
// 	for y, line := range parsed {
// 		grid[y] = make([]*Node, len(line))
// 		for x, val := range line {
// 			grid[y][x] = &Node{cast.ToInt(string(val)), defaultMinLoss(), Location{x, y}}
// 		}
// 	}
// 	return grid
// }

// func last3MovesStraight(history []int) bool {
// 	if len(history) < 3 {
// 		return false
// 	}

// 	return history[0] == history[1] && history[0] == history[2] && history[1] == history[2]
// }

// func considerNode() {

// }

// func nextMoves(node *Node, g [][]*Node, dir int, history []int, openSet []*Node) []*Node {
// 	moves := []*Node{}
// 	// left
// 	left := Directions[dir-1]
// 	leftNode := g[node.l.y+left.y][node.l.x+left.x]

// 	right := Directions[dir+1]
// 	rightNode := g[node.l.y+right.y][node.l.x+right.x]

// 	if !last3MovesStraight(history) {
// 		straight := Directions[dir]
// 		straightNode := g[node.l.y+straight.y][node.l.x+straight.x]
// 	}

// 	moves = append(moves)

// 	// right

// 	// straight on
// }

// func minHeatLoss(grid [][]*Node, start Location, end Location) int {
// 	grid[start.y][start.x].minHeatLoss = 0
// 	openSet := []*Node{grid[start.y][start.x]}

// 	for len(openSet) > 0 {
// 		slices.SortFunc[[]*Node](openSet, sortByHeatAscending)
// 		currentNode := openSet[0] // shortest node

// 		if currentNode.l == end {
// 			// fmt.Println(endId, nodes[endId].minimumHeatLoss, nodes[endId].locationHistory)
// 			// drawPath(nodes, len(parsed[0]), len(parsed), startId, endId)
// 			return currentNode.minHeatLoss
// 		}

// 		openSet = openSet[1:]

// 		openSet = nextMoves(currentNode, grid, openSet)

// 	// 	// up, right, down, left
// 	// 	//  0,     1,    2,    3
// 	// 	for dir, nextNodeId := range currentNode.connectedNodes {
// 	// 		if nextNodeId == "" {
// 	// 			// the current node doesn't connect in this direction
// 	// 			continue
// 	// 		}
// 	// 		if sameDirection(currentNode.pathHistory, dir) {
// 	// 			continue
// 	// 		}

// 	// 		possibleMinHeatLoss := currentNode.minimumHeatLoss + nodes[nextNodeId].value
// 	// 		if possibleMinHeatLoss < nodes[nextNodeId].minimumHeatLoss {
// 	// 			nodes[nextNodeId].minimumHeatLoss = possibleMinHeatLoss
// 	// 			nodes[nextNodeId].locationHistory = append(currentNode.locationHistory, currentNode.location)
// 	// 			nodes[nextNodeId].pathHistory = append(nodes[nextNodeId].pathHistory, dir)

// 	// 			if !slices.Contains[[]*Node](openSet, nodes[nextNodeId]) {
// 	// 				openSet = append(openSet, nodes[nextNodeId])
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }
// 	// return nodes[endId].minimumHeatLoss
// }

// // func sortByHeatAscending(n1, n2 *Node) int {
// // 	return n1.minHeatLoss - n2.minHeatLoss
// // }

// // Helper functions for part 2
