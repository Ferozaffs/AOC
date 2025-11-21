package day21

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Coordinate struct {
	x int
	y int
}

var day21_1Cmd = &cobra.Command{
	Use:   "day21_1",
	Short: "Day 21, challenge 1",
	Long:  `Day 21, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day21_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day21_data.txt")
	ans := Solve1(string(dat), 64, false)

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string, maxStep int, repeat bool) int {
	nodes := [][]string{}
	visited := map[Coordinate]int{}

	start := Coordinate{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	x := 0
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "\t") {
			s = s[1:]
		}

		row := []string{}
		for y, r := range s {
			row = append(row, string(r))

			if r == 'S' {
				start = Coordinate{x, y}
			}
		}

		nodes = append(nodes, row)
		x++
	}

	CheckCoordinate(start, 0, maxStep, &nodes, &visited, repeat)

	numPlots := 0
	for _, p := range visited {
		if p%2 == 0 {
			numPlots++
		}
	}

	fmt.Println(numPlots)
	return numPlots
}

func CheckCoordinate(c Coordinate, step int, maxStep int, nodes *[][]string, visited *map[Coordinate]int, repeat bool) {
	directions := [4]Coordinate{
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
		{x: 0, y: -1},
	}

	rx := mod(c.x, len(*nodes))
	ry := mod(c.y, len((*nodes)[0]))

	if step > maxStep {
		return
	} else if repeat == false && (c.x < 0 || c.y < 0 || c.x >= len(*nodes) || c.y >= len((*nodes)[0])) {
		return
	} else if (*nodes)[rx][ry] == "#" {
		return
	}

	val, ok := (*visited)[c]
	if ok && val <= step {
		return
	}

	(*visited)[c] = step

	for _, d := range directions {
		newCoord := c
		newCoord.x += d.x
		newCoord.y += d.y

		CheckCoordinate(newCoord, step+1, maxStep, nodes, visited, repeat)
	}

	return
}

func mod(a, b int) int {
	return (a%b + b) % b
}
