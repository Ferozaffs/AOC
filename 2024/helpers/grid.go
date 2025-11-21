package helpers

import "strings"

type Point struct {
	X int
	Y int
}

var Directions = []Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func (p *Point) RotateCW() {
	if p.X == 1 {
		p.X = 0
		p.Y = -1
	} else if p.Y == -1 {
		p.X = -1
		p.Y = 0
	} else if p.X == -1 {
		p.X = 0
		p.Y = 1
	} else if p.Y == 1 {
		p.X = 1
		p.Y = 0
	}
}

func (p *Point) RotateCCW() {
	if p.X == 1 {
		p.X = 0
		p.Y = 1
	} else if p.Y == 1 {
		p.X = -1
		p.Y = 0
	} else if p.X == -1 {
		p.X = 0
		p.Y = -1
	} else if p.Y == -1 {
		p.X = 1
		p.Y = 0
	}
}

type Grid [][]rune

func (g *Grid) Init(input string) bool {
	lines := strings.Split(input, "\n")

	if len(lines) == 0 {
		return false
	}

	for _, line := range lines {
		var runes []rune
		for _, r := range line {
			runes = append(runes, r)
		}
		(*g) = append((*g), runes)
	}

	return true
}

func (g Grid) GetPointDir(x int, y int, dx int, dy int) (rune, bool) {
	return g.GetPointOffset(x, y, Point{dx, dy})
}

func (g Grid) GetPointOffset(x int, y int, p Point) (rune, bool) {
	x += p.X
	y += p.Y
	return g.GetPointXY(x, y)
}

func (g Grid) GetPoint(p Point) (rune, bool) {
	return g.GetPointXY(p.X, p.Y)
}

func (g Grid) GetPointXY(x int, y int) (rune, bool) {
	if x < 0 || y < 0 || x >= len(g) || y >= len(g[0]) {
		return ' ', false
	}

	return g[x][y], true
}

func (g Grid) DeepCopy() Grid {
	newGrid := make(Grid, len(g))
	for i := range g {
		newGrid[i] = make([]rune, len(g[i]))
		copy(newGrid[i], g[i])
	}
	return newGrid
}
