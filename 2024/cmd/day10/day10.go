package day10

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var directions = []helpers.Point{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

var day1_1Cmd = &cobra.Command{
	Use:   "day10",
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
	dat, _ := os.ReadFile("inputs/day10_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	var grid helpers.Grid
	grid.Init(data)

	ans1 := 0
	ans2 := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '0' {
				paths := Search(helpers.Point{X: x, Y: y}, 0, &grid)
				ans1 += FindUniquePaths(paths)
				ans2 += len(paths)
			}
		}
	}

	return ans1, ans2
}

func Search(p helpers.Point, index int, g *helpers.Grid) []helpers.Point {
	if index == 9 {
		return []helpers.Point{p}
	}

	result := []helpers.Point{}
	for i := 0; i < 4; i++ {
		np := helpers.Point{X: p.X + directions[i].X, Y: p.Y + directions[i].Y}

		r, b := g.GetPoint(np)
		if b && int(r-'0') == index+1 {
			result = append(result, Search(np, index+1, g)...)
		}
	}

	return result
}

func FindUniquePaths(p []helpers.Point) int {
	uniqueMap := make(map[helpers.Point]bool)
	for _, v := range p {
		uniqueMap[v] = true
	}
	return len(uniqueMap)
}
