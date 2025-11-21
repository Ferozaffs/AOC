package day20

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Edge struct {
	parent helpers.Point
	weight int
}

var cutoff int = 100

var day1_1Cmd = &cobra.Command{
	Use:   "day20",
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
	dat, _ := os.ReadFile("inputs/day20_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func SetCutoff(c int) {
	cutoff = c
}

func Solve(data string) (int, int) {
	var grid helpers.Grid
	grid.Init(data)

	ans1 := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '#' {
				r1, b1 := grid.GetPointXY(x-1, y)
				r2, b2 := grid.GetPointXY(x+1, y)

				if b1 && b2 && r1 != '#' && r2 != '#' {
					path, b := Pathfind(helpers.Point{X: x - 1, Y: y}, helpers.Point{X: x + 1, Y: y}, grid)
					if b && len(path)-3 >= cutoff {
						ans1++
					}
				}

				r1, b1 = grid.GetPointXY(x, y-1)
				r2, b2 = grid.GetPointXY(x, y+1)

				if b1 && b2 && r1 != '#' && r2 != '#' {
					path, b := Pathfind(helpers.Point{X: x, Y: y - 1}, helpers.Point{X: x, Y: y + 1}, grid)
					if b && len(path)-3 >= cutoff {
						ans1++
					}
				}
			}

		}
	}

	return ans1, 0
}

func Pathfind(start helpers.Point, end helpers.Point, g helpers.Grid) ([]helpers.Point, bool) {
	foundPath := true

	queue := make(helpers.PriorityQueue, 0)
	queue.Enqueue(start, 0)

	dist := make(map[helpers.Point]Edge)
	dist[start] = Edge{parent: helpers.Point{X: 0, Y: 0}, weight: 0}

	for {
		if len(queue) == 0 {
			foundPath = false
			break
		}

		u := queue.Dequeue().(helpers.Point)

		if u.X == end.X && u.Y == end.Y {
			break
		}

		for i := 0; i < 4; i++ {
			n := u
			n.X += helpers.Directions[i].X
			n.Y += helpers.Directions[i].Y

			r, b := g.GetPoint(n)

			if !b || r == '#' {
				continue
			}

			w := 1

			v, exists := dist[n]
			if !exists {
				dist[n] = Edge{parent: u, weight: dist[u].weight + w}
				queue.Enqueue(n, dist[u].weight+w)
			} else {
				if dist[u].weight+w < v.weight {
					dist[n] = Edge{parent: u, weight: dist[u].weight + w}
					queue.Enqueue(n, dist[u].weight+w)
				}
			}
		}
	}

	path := []helpers.Point{}
	path = append(path, end)
	v := end
	for {
		_, exists := dist[v]
		if !exists {
			break
		}
		v = dist[v].parent
		path = append(path, v)
	}

	return path, foundPath
}
