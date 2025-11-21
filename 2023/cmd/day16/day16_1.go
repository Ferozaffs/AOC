package day16

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Coordinate struct {
	x int
	y int
}

type DirCache struct {
	cache [4]bool
}

const (
	left  = 0
	right = 1
	down  = 2
	up    = 3
)

var day16_1Cmd = &cobra.Command{
	Use:   "day16_1",
	Short: "Day 16, challenge 1",
	Long:  `Day 16, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day16_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day16_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

var energized = map[Coordinate]DirCache{}

func Solve1(data string) int {
	s := strings.Fields(data)

	coordinate := Coordinate{x: -1, y: 0}
	MoveRay(right, coordinate, s)

	return len(energized)
}

func MoveRay(dir int, c Coordinate, s []string) {
	if dir == left {
		c.x--
	} else if dir == right {
		c.x++
	} else if dir == up {
		c.y--
	} else {
		c.y++
	}

	if c.x < 0 || c.y < 0 || c.y == len(s) || c.x == len(s[c.y]) {
		return
	}

	if energized[c].cache[dir] == true {
		return
	}

	cache := DirCache{}
	cache = energized[c]
	cache.cache[dir] = true
	energized[c] = cache

	r := s[c.y][c.x]

	if r == '.' {
		MoveRay(dir, c, s)
	} else if r == '/' {
		if dir == left {
			dir = down
		} else if dir == down {
			dir = left
		} else if dir == right {
			dir = up
		} else if dir == up {
			dir = right
		}

		MoveRay(dir, c, s)

	} else if r == '\\' {
		if dir == left {
			dir = up
		} else if dir == up {
			dir = left
		} else if dir == right {
			dir = down
		} else if dir == down {
			dir = right
		}

		MoveRay(dir, c, s)

	} else if r == '|' {
		if dir == left || dir == right {
			MoveRay(up, c, s)
			MoveRay(down, c, s)
		} else {
			MoveRay(dir, c, s)
		}

	} else if r == '-' {
		if dir == up || dir == down {
			MoveRay(left, c, s)
			MoveRay(right, c, s)
		} else {
			MoveRay(dir, c, s)
		}
	}
}
