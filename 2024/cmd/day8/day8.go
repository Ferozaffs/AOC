package day8

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day8",
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
	dat, _ := os.ReadFile("inputs/day8_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	var grid helpers.Grid
	grid.Init(data)

	antiNodes := make(map[helpers.Point]bool)
	antiNodesRepeating := make(map[helpers.Point]bool)
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != '.' {
				FindAntiNodes(helpers.Point{X: x, Y: y}, &antiNodes, grid, false)
				FindAntiNodes(helpers.Point{X: x, Y: y}, &antiNodesRepeating, grid, true)
			}
		}
	}

	return len(antiNodes), len(antiNodesRepeating)
}

func FindAntiNodes(p helpers.Point, an *map[helpers.Point]bool, g helpers.Grid, repeating bool) {
	r, _ := g.GetPoint(p)

	oX := 0
	oY := 0
	if !repeating {
		oX = p.X
		oY = p.Y + 1
	}
	for ; oX < len(g); oX++ {
		for ; oY < len(g[0]); oY++ {
			if oX == p.X && oY == p.Y {
				continue
			}

			if g[oX][oY] == r {
				delta := helpers.Point{X: oX - p.X, Y: oY - p.Y}

				if !repeating {
					n1 := helpers.Point{X: p.X - delta.X, Y: p.Y - delta.Y}
					_, result := g.GetPoint(n1)
					if result {
						(*an)[n1] = true
					}

					n2 := helpers.Point{X: oX + delta.X, Y: oY + delta.Y}
					_, result = g.GetPoint(n2)
					if result {
						(*an)[n2] = true
					}
				} else {
					n := p
					for {
						n.X -= delta.X
						n.Y -= delta.Y
						_, result := g.GetPoint(n)
						if result {
							(*an)[n] = true
						} else {
							break
						}
					}

					n = p
					for {
						n.X += delta.X
						n.Y += delta.Y
						_, result := g.GetPoint(n)
						if result {
							(*an)[n] = true
						} else {
							break
						}
					}
				}

			}
		}
		oY = 0
	}
}
