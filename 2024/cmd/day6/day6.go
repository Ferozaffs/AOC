package day6

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Visit struct {
	d helpers.Point
	p helpers.Point
}

var day1_1Cmd = &cobra.Command{
	Use:   "day6",
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
	dat, _ := os.ReadFile("inputs/day6_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {

	var grid helpers.Grid
	grid.Init(data)

	var start helpers.Point
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '^' {
				start.X = x
				start.Y = y
			}
		}
	}

	direction := helpers.Point{X: -1, Y: 0}
	vists := []Visit{}
	ans1, uv, _ := Walk(start, direction, vists, &grid, false)

	ans2 := 0
	for _, v := range uv {
		if grid[v.X][v.Y] == '.' {
			d := helpers.Point{X: -1, Y: 0}
			nv := []Visit{}
			grid[v.X][v.Y] = '#'
			_, _, loop := Walk(start, d, nv, &grid, true)
			if loop {
				ans2++
			}
			grid[v.X][v.Y] = '.'
		}
	}
	return ans1, ans2
}

func Walk(position helpers.Point, direction helpers.Point, visits []Visit, grid *helpers.Grid, checkLoop bool) (int, []helpers.Point, bool) {
	currentVisit := Visit{d: direction, p: position}
	for _, v := range visits {
		if v == currentVisit {
			if checkLoop {
				return -1, []helpers.Point{}, true
			}
			l, uv := CalculateUniquePoints(visits)
			return l, uv, true
		}
	}

	visits = append(visits, currentVisit)

	nextVisit := currentVisit

	r, e := grid.GetPointOffset(nextVisit.p.X, nextVisit.p.Y, direction)
	for i := 0; i < 4; i++ {
		if !e {
			if checkLoop {
				return -1, []helpers.Point{}, false
			}
			l, uv := CalculateUniquePoints(visits)
			return l, uv, false
		}

		if r == '#' {
			direction.RotateCW()
		} else {
			break
		}

		r, e = grid.GetPointOffset(nextVisit.p.X, nextVisit.p.Y, direction)
	}

	nextVisit.p.X += direction.X
	nextVisit.p.Y += direction.Y

	return Walk(nextVisit.p, direction, visits, grid, checkLoop)
}

func CalculateUniquePoints(visits []Visit) (int, []helpers.Point) {
	uniqueVisits := []helpers.Point{}
	for _, v := range visits {
		isUnique := true
		for _, u := range uniqueVisits {
			if u == v.p {
				isUnique = false
				break
			}
		}

		if isUnique {
			uniqueVisits = append(uniqueVisits, v.p)
		}
	}

	return len(uniqueVisits), uniqueVisits
}
