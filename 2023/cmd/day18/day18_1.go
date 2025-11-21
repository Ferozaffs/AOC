package day18

import (
	"aoc2023/cmd"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Coordinate struct {
	x int
	y int
}

var day18_1Cmd = &cobra.Command{
	Use:   "day18_1",
	Short: "Day 18, challenge 1",
	Long:  `Day 18, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day18_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day18_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	rows := strings.Split(strings.ReplaceAll(data, "\t", ""), "\n")
	minBound := Coordinate{x: 0, y: 0}
	maxBound := Coordinate{x: 0, y: 0}
	c := Coordinate{x: 0, y: 0}
	coords := []Coordinate{}

	for _, row := range rows {
		s := strings.Fields(row)
		dir := s[0]

		for count, _ := strconv.Atoi(s[1]); count > 0; count-- {
			if dir == "R" {
				c.x++
			} else if dir == "L" {
				c.x--
			} else if dir == "D" {
				c.y--
			} else if dir == "U" {
				c.y++
			}

			minBound.x = int(math.Min(float64(c.x), float64(minBound.x)))
			minBound.y = int(math.Min(float64(c.y), float64(minBound.y)))
			maxBound.x = int(math.Max(float64(c.x), float64(maxBound.x)))
			maxBound.y = int(math.Max(float64(c.y), float64(maxBound.y)))

			coords = append(coords, c)
		}
	}

	dirs := []Coordinate{
		Coordinate{x: 1, y: 0},
		Coordinate{x: -1, y: 0},
		Coordinate{x: 0, y: 1},
		Coordinate{x: 0, y: -1},
	}

	area := 0

	visited := []Coordinate{}
	q := []Coordinate{Coordinate{x: (coords[0].x + 1), y: coords[0].y - 1}}
	for len(q) > 0 {
		c, q = q[0], q[1:]
		visited = append(visited, c)
		area++

		for _, dir := range dirs {
			nc := c
			nc.x += dir.x
			nc.y += dir.y

			if nc.x >= minBound.x && nc.y > minBound.y &&
				nc.x <= maxBound.x && nc.y <= maxBound.y &&
				Contains(q, nc) == false &&
				Contains(visited, nc) == false &&
				Contains(coords, nc) == false {
				q = append(q, nc)
			}
		}
	}

	area += len(coords)

	return area
}

func Contains(s []Coordinate, e Coordinate) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
