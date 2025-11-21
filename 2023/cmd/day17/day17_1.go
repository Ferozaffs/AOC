package day17

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Item struct {
	value    interface{}
	priority int
}

type PriorityQueue []Item

func (pq *PriorityQueue) Enqueue(val interface{}, prio int) {
	item := Item{value: val, priority: prio}
	*pq = append(*pq, item)
	pq.heapifyUp(len(*pq) - 1)
}

func (pq *PriorityQueue) Dequeue() Item {
	n := len(*pq)
	item := (*pq)[0]
	(*pq)[0] = (*pq)[n-1]
	*pq = (*pq)[:n-1]
	pq.heapifyDown(0, len(*pq))

	return item
}

func (pq *PriorityQueue) heapifyUp(index int) {
	if index == 0 {
		return
	}

	parent := (index - 1) / 2
	if (*pq)[index].priority > (*pq)[parent].priority {
		return
	}

	(*pq)[index], (*pq)[parent] = (*pq)[parent], (*pq)[index]
	pq.heapifyUp(parent)
}

func (pq *PriorityQueue) heapifyDown(index int, size int) {
	left := (2 * index) + 1
	right := (2 * index) + 2

	smallest := index
	if left < size && (*pq)[left].priority < (*pq)[smallest].priority {
		smallest = left
	}
	if right < size && (*pq)[right].priority < (*pq)[smallest].priority {
		smallest = right
	}

	if smallest == index {
		return
	}

	(*pq)[index], (*pq)[smallest] = (*pq)[smallest], (*pq)[index]
	pq.heapifyDown(smallest, size)
}

type Point struct {
	x int
	y int
}

type Move struct {
	p Point
	d Point
}

var day17_1Cmd = &cobra.Command{
	Use:   "day17_1",
	Short: "Day 17, challenge 1",
	Long:  `Day 17, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day17_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day17_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
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
		for i := -1; i > -3; i-- {
			xn := move.p.x + move.d.x*i
			yn := move.p.y + move.d.y*i

			if _, ok := points[Point{xn, yn}]; !ok {
				continue
			}

			cost += points[Point{xn, yn}]
			m1 := Move{Point{xn, yn}, Point{-move.d.y, -move.d.x}}
			m2 := Move{Point{xn, yn}, Point{move.d.y, move.d.x}}
			queue.Enqueue(m1, cost)
			queue.Enqueue(m2, cost)
		}

		cost = u.priority
		for i := 1; i <= 3; i++ {
			xn := move.p.x + move.d.x*i
			yn := move.p.y + move.d.y*i

			if _, ok := points[Point{xn, yn}]; !ok {
				continue
			}

			cost += points[Point{xn, yn}]
			m1 := Move{Point{xn, yn}, Point{-move.d.y, -move.d.x}}
			m2 := Move{Point{xn, yn}, Point{move.d.y, move.d.x}}
			queue.Enqueue(m1, cost)
			queue.Enqueue(m2, cost)
		}
	}

	return optimalCost
}
