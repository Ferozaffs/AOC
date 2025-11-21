package day12

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type PerimiterData struct {
	area    int
	fences  int
	corners int
}

var directions = []helpers.Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

var directionsDiagonals = []helpers.Point{
	{-1, -1},
	{1, -1},
	{-1, 1},
	{1, 1},
}

var day1_1Cmd = &cobra.Command{
	Use:   "day12",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day1_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day12_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	var grid helpers.Grid
	grid.Init(data)

	visited := make(map[helpers.Point]bool)
	perimiters := []PerimiterData{}
	for x := range grid {
		for y := range grid[x] {
			p := helpers.Point{X: x, Y: y}
			if visited[p] {
				continue
			}
			perimiters = append(perimiters, Search(grid[x][y], p, &grid, &visited))
		}
	}

	ans1 := 0
	ans2 := 0
	for _, pd := range perimiters {
		ans1 += pd.area * pd.fences
		ans2 += pd.area * pd.corners
	}

	return ans1, ans2
}

func Search(or rune, p helpers.Point, g *helpers.Grid, v *map[helpers.Point]bool) PerimiterData {
	(*v)[p] = true

	perimiterData := PerimiterData{area: 1, fences: 0, corners: 0}
	fenceL := false
	fenceR := false
	fenceT := false
	fenceB := false

	nL := false
	nR := false
	nT := false
	nB := false
	for i := 0; i < 4; i++ {
		np := helpers.Point{X: p.X + directions[i].X, Y: p.Y + directions[i].Y}

		r, b := g.GetPoint(np)

		if !b || r != or {
			if i == 0 {
				fenceB = true
			} else if i == 1 {
				fenceR = true
			} else if i == 2 {
				fenceT = true
			} else {
				fenceL = true
			}

			perimiterData.fences++
		} else if (*v)[np] {
			if i == 0 {
				nB = true
			} else if i == 1 {
				nR = true
			} else if i == 2 {
				nT = true
			} else {
				nL = true
			}
			continue
		} else {
			pd := Search(or, np, g, v)
			perimiterData.area += pd.area
			perimiterData.fences += pd.fences
			perimiterData.corners += pd.corners

			if i == 0 {
				nB = true
			} else if i == 1 {
				nR = true
			} else if i == 2 {
				nT = true
			} else {
				nL = true
			}
		}
	}

	if fenceL && fenceT {
		perimiterData.corners++
	}
	if fenceL && fenceB {
		perimiterData.corners++
	}
	if fenceR && fenceT {
		perimiterData.corners++
	}
	if fenceR && fenceB {
		perimiterData.corners++
	}

	if nL && nT {
		np := helpers.Point{X: p.X + directionsDiagonals[0].X, Y: p.Y + directionsDiagonals[0].Y}

		r, b := g.GetPoint(np)

		if !b || r != or {
			perimiterData.corners++
		}
	}
	if nL && nB {
		np := helpers.Point{X: p.X + directionsDiagonals[1].X, Y: p.Y + directionsDiagonals[1].Y}

		r, b := g.GetPoint(np)

		if !b || r != or {
			perimiterData.corners++
		}
	}
	if nR && nT {
		np := helpers.Point{X: p.X + directionsDiagonals[2].X, Y: p.Y + directionsDiagonals[2].Y}

		r, b := g.GetPoint(np)

		if !b || r != or {
			perimiterData.corners++
		}
	}
	if nR && nB {
		np := helpers.Point{X: p.X + directionsDiagonals[3].X, Y: p.Y + directionsDiagonals[3].Y}

		r, b := g.GetPoint(np)

		if !b || r != or {
			perimiterData.corners++
		}
	}

	return perimiterData
}
