//nolint:gosec
package grid

import (
	"fmt"
	"math"
	"strconv"
)

type Grid [][]rune

type Direction int

const (
	North Direction = 0
	West  Direction = 1
	South Direction = 2
	East  Direction = 3
)

var CaratDirectionMap = map[string]Direction {
	"^": North,
	"<": West,
	"v": South,
	">": East,
}

var DirectionCaratMap = map[Direction]string {
	North: "^",
	West: "<",
	South: "v",
	East: ">",
}

type Location struct {
	X, Y int
}

// Takes lines and turns it into grid
func GetGrid(input []string) Grid {
	g := Grid{}
	for _, line := range input {
		g = append(g, []rune(line))
	}
	return g
}

func VerticalSlice(g Grid, i int) []rune {
	s := make([]rune, len(g))
	for row := range g {
		s[row] = g[row][i]
	}
	return s
}

func HashGrid(g Grid) string {
	// one for each cell, plus a new line for each row, except the last
	hash := make([]rune, (len(g)+1)*len(g[0])-1)
	for y, row := range g {
		startIndex := len(row)*y + y
		endIndex := len(row)*(y+1) + y
		copy(hash[startIndex:endIndex], row)

		if endIndex != len(hash) {
			hash[endIndex] = '\n'
		}
	}
	return string(hash)
}

func PrintGrid(g Grid) {
	for _, line := range g {
		fmt.Println(string(line))
	}
	fmt.Println("")
}

func GetGridFromLocations(locations []Location, width, height int, background, foreground rune) Grid {
	g := InitialiseGrid(width, height, background)

	for _, l := range locations {
		g[l.Y][l.X] = foreground
	}
	return g
}

func GetLocationsOfCharacter(g Grid, r rune) []Location {
	var ls []Location
	for y, line := range g {
		for x := range line {
			if g[y][x] == r {
				ls = append(ls, Location{x, y})
			}
		}
	}
	return ls
}

func InitialiseGrid(width, height int, background rune) Grid {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[y][x] = background
		}
	}

	return grid
}

func DeepCopyGrid(g Grid) Grid {
	c := InitialiseGrid(len(g[0]), len(g), '.')
	for y := range g {
		copy(c[y], g[y])
	}
	return c
}

func MoveStepsInDirection(l Location, d Direction, steps int) Location {
	if d == North {
		return Location{l.X, l.Y - steps}
	}
	if d == East {
		return Location{l.X + steps, l.Y}
	}
	if d == South {
		return Location{l.X, l.Y + steps}
	}
	if d == West {
		return Location{l.X - steps, l.Y}
	}
	return l
}

func DirectionBetweenLocations(l1, l2 Location) Direction {
	if l1.X-l2.X == 1 && l1.Y-l2.Y == 0 {
		return West
	}
	if l1.X-l2.X == -1 && l1.Y-l2.Y == 0 {
		return East
	}
	if l1.X-l2.X == 0 && l1.Y-l2.Y == 1 {
		return North
	}
	if l1.X-l2.X == 0 && l1.Y-l2.Y == -1 {
		return South
	}
	panic("Locations are not adjacent!")
}

func OppositeDirections(d1, d2 Direction) bool {
	if (d1 == North && d2 == South) || (d1 == South && d2 == North) {
		return true
	}
	if (d1 == East && d2 == West) || (d1 == West && d2 == East) {
		return true
	}
	return false
}

// north, east, south, west
func FourAdjacent(l Location) (Location, Location, Location, Location) {
	north := MoveStepsInDirection(l, North, 1)
	east := MoveStepsInDirection(l, East, 1)
	south := MoveStepsInDirection(l, South, 1)
	west := MoveStepsInDirection(l, West, 1)
	return north, east, south, west
}

func FourAdjacentList(l Location) []Location {
	north, east, south, west := FourAdjacent(l)
	return []Location{north, east, south, west}
}

// north, northEast, east, southEast, south, southWest, west, northWest
func EightAdjacent(l Location) (Location, Location, Location, Location, Location, Location, Location, Location) {
	north := MoveStepsInDirection(l, North, 1)
	northEast := MoveStepsInDirection(north, East, 1)
	east := MoveStepsInDirection(l, East, 1)
	southEast := MoveStepsInDirection(east, South, 1)
	south := MoveStepsInDirection(l, South, 1)
	southWest := MoveStepsInDirection(south, West, 1)
	west := MoveStepsInDirection(l, West, 1)
	northWest := MoveStepsInDirection(west, North, 1)
	return north, northEast, east, southEast, south, southWest, west, northWest
}

func EightAdjacentList(l Location) []Location {
	north, northEast, east, southEast, south, southWest, west, northWest := EightAdjacent(l)
	return []Location{north, northEast, east, southEast, south, southWest, west, northWest}
}

func AllLocations(g Grid) []Location {
	var ls []Location
	for y, line := range g {
		for x := range line {
			ls = append(ls, Location{x, y})
		}
	}
	return ls
}

func ValueAtLocation(g Grid, l Location) rune {
	return g[l.Y][l.X]
}

func LocationsEqual(l1, l2 Location) bool {
	return l1.X == l2.X && l1.Y == l2.Y
}

func ListOfLocationsEqual(ls1, ls2 []Location) bool {
	if len(ls1) != len(ls2) {
		return false
	}
	for i := range ls1 {
		if !LocationsEqual(ls1[i], ls2[i]) {
			return false
		}
	}
	return true
}

func LocationOutsideGrid(l Location, m Grid) bool {
	return l.X < 0 || l.X >= len(m[0]) || l.Y < 0 || l.Y >= len(m)
}

// steps going from l1 -> l2
func LocationDiff(l1, l2 Location) Location {
	return Location{l2.X - l1.X, l2.Y - l1.Y}
}

// step diff from l1
func LocationAdd(l, diff Location) Location {
	return Location{l.X + diff.X, l.Y + diff.Y}
}

func MultiplyDistance(l Location, multiplier int) Location {
	return Location{l.X * multiplier, l.Y * multiplier}
}

func ManhattanDistance(l1, l2 Location) int {
	return AbsDiff(l1.X, l2.X) + AbsDiff(l1.Y, l2.Y)
}

// truncates
func PythagorasDistance(l1, l2 Location) int {
	return int(math.Pow(float64(AbsDiff(l1.X, l2.X)), 2) + math.Pow(float64(AbsDiff(l1.Y, l2.Y)), 2))
}

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func HashLocation(l Location) string {
	return strconv.Itoa(l.X) + " " + strconv.Itoa(l.Y)
}

// e.g. diff = {2,4} -> {1,2}
func SimplifyLocation(diff Location) Location {
	gcd := greatestCommonDivisor(diff.X, diff.Y)
	if gcd == 0 {
		return diff
	}
	return Location{diff.X / gcd, diff.Y / gcd}
}

func LocationInList(l Location, ls []Location) bool {
	for i := range ls {
		if ls[i].X == l.X && ls[i].Y == l.Y {
			return true
		}
	}
	return false
}
