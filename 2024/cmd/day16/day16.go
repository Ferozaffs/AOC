package day16

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Edge struct {
	parent helpers.Point
	dir    helpers.Point
	weight int
}

var day1_1Cmd = &cobra.Command{
	Use:   "day16",
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
	dat, _ := os.ReadFile("inputs/day16_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	var grid helpers.Grid
	grid.Init(data)

	start := helpers.Point{}
	end := helpers.Point{}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				start = helpers.Point{X: x, Y: y}
			}

			if grid[x][y] == 'E' {
				end = helpers.Point{X: x, Y: y}
			}
		}
	}

	_, ans1 := Pathfind(start, end, &grid)

	//for x := range grid {
	//	fmt.Print("\n")
	//	for y := range grid[x] {
	//		found := false
	//		for _, p := range path {
	//			if p.X == x && p.Y == y {
	//				fmt.Print("O")
	//				found = true
	//				break
	//			}
	//		}
	//		if !found {
	//			fmt.Print(string(grid[x][y]))
	//		}
	//	}
	//}

	return ans1, 0
}

func Pathfind(start helpers.Point, end helpers.Point, g *helpers.Grid) ([]helpers.Point, int) {
	queue := make(helpers.PriorityQueue, 0)
	queue.Enqueue(start, 0)

	dist := make(map[helpers.Point]Edge)
	dist[start] = Edge{parent: helpers.Point{X: 0, Y: 0}, dir: helpers.Point{X: 0, Y: 1}, weight: 0}

	for {
		if len(queue) == 0 {
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

			r, b := (*g).GetPoint(n)

			if !b || r == '#' {
				continue
			}

			w := 1
			if helpers.Directions[i].X != dist[u].dir.X || helpers.Directions[i].Y != dist[u].dir.Y {
				w += 1000
			}

			v, exists := dist[n]
			if !exists {
				dist[n] = Edge{parent: u, dir: helpers.Directions[i], weight: dist[u].weight + w}
				queue.Enqueue(n, dist[u].weight+w)
			} else {
				if dist[u].weight+w < v.weight {
					dist[n] = Edge{parent: u, dir: helpers.Directions[i], weight: dist[u].weight + w}
					queue.Enqueue(n, dist[u].weight+w)
				}
			}
		}
	}

	path := []helpers.Point{}
	path = append(path, end)
	v := end
	weight := dist[v].weight
	for {
		_, exists := dist[v]
		if !exists {
			break
		}
		v = dist[v].parent
		path = append(path, v)
	}

	return path, weight
}
