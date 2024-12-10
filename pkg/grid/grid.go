//nolint:gosec
package grid

import "strconv"

type Grid [][]rune

type Direction int

const (
	North Direction = 0
	West  Direction = 1
	South Direction = 2
	East  Direction = 3
)

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
