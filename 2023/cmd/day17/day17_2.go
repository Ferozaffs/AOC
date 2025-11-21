package day17

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day17_2Cmd = &cobra.Command{
	Use:   "day17_2",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day17_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day17_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	s := strings.Fields(data)

	dest := Point{}
	points := map[Point]int{}
	for x, row := range s {
		for y := 0; y < len(row); y++ {
			points[Point{x, y}] = int(s[x][y] - '0')
			dest = Point{x, y}
		}
	}

	queue := make(PriorityQueue, 0)
	m1 := Move{Point{0, 0}, Point{1, 0}}
	m2 := Move{Point{0, 0}, Point{0, 1}}
	queue.Enqueue(m1, 0)
	queue.Enqueue(m2, 0)

	seenMoves := map[Move]bool{}

	optimalCost := 0
	for len(queue) > 0 {

		u := queue.Dequeue()

		move := u.value.(Move)
		if move.p == dest {
			optimalCost = u.priority
			break
		}

		if _, ok := seenMoves[move]; ok {
			continue
		}

		seenMoves[move] = true

		cost := u.priority
		for i := -1; i > -10; i-- {
			xn := move.p.x + move.d.x*i
			yn := move.p.y + move.d.y*i

			if _, ok := points[Point{xn, yn}]; !ok {
				continue
			}

			cost += points[Point{xn, yn}]

			if i <= -4 {
				m1 := Move{Point{xn, yn}, Point{-move.d.y, -move.d.x}}
				m2 := Move{Point{xn, yn}, Point{move.d.y, move.d.x}}
				queue.Enqueue(m1, cost)
				queue.Enqueue(m2, cost)
			}
		}

		cost = u.priority
		for i := 1; i <= 10; i++ {
			xn := move.p.x + move.d.x*i
			yn := move.p.y + move.d.y*i

			if _, ok := points[Point{xn, yn}]; !ok {
				continue
			}

			cost += points[Point{xn, yn}]

			if i >= 4 {
				m1 := Move{Point{xn, yn}, Point{-move.d.y, -move.d.x}}
				m2 := Move{Point{xn, yn}, Point{move.d.y, move.d.x}}
				queue.Enqueue(m1, cost)
				queue.Enqueue(m2, cost)
			}
		}
	}

	return optimalCost
}
