package day18

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Edge struct {
	parent helpers.Point
	weight int
}

var sizeX = 70
var sizeY = 70
var num = 1024

var day1_1Cmd = &cobra.Command{
	Use:   "day18",
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
	dat, _ := os.ReadFile("inputs/day18_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %s\n", ans2)
}

func SetSize(x int, y int, n int) {
	sizeX = x
	sizeY = y
	num = n
}

func Solve(data string) (int, string) {
	var grid1 helpers.Grid
	var grid2 helpers.Grid

	for x := 0; x <= sizeX; x++ {
		var runes []rune
		for y := 0; y <= sizeY; y++ {
			runes = append(runes, '.')
		}
		grid1 = append(grid1, runes)
		grid2 = append(grid1, runes)
	}

	lines := strings.Split(data, "\n")
	for i, s := range lines {
		if i == num {
			break
		}

		vals := strings.Split(s, ",")
		y, _ := strconv.Atoi(vals[0])
		x, _ := strconv.Atoi(vals[1])

		grid1[x][y] = '#'
	}

	path, _ := Pathfind(helpers.Point{X: 0, Y: 0}, helpers.Point{X: sizeX, Y: sizeY}, grid1)

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

	ans2 := "0,0"
	for _, s := range lines {
		vals := strings.Split(s, ",")
		y, _ := strconv.Atoi(vals[0])
		x, _ := strconv.Atoi(vals[1])

		grid2[x][y] = '#'

		_, r := Pathfind(helpers.Point{X: 0, Y: 0}, helpers.Point{X: sizeX, Y: sizeY}, grid2)

		if !r {
			ans2 = s
			break
		}
	}

	return len(path) - 2, ans2
}

func Pathfind(start helpers.Point, end helpers.Point, g helpers.Grid) ([]helpers.Point, bool) {
	foundPath := true

	queue := make(helpers.PriorityQueue, 0)
	queue.Enqueue(start, 0)

	dist := make(map[helpers.Point]Edge)
	dist[start] = Edge{parent: helpers.Point{X: -1, Y: -1}, weight: 0}

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
