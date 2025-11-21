package day10

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Coordinate struct {
	x    int
	y    int
	cost int
}

var day10_1Cmd = &cobra.Command{
	Use:   "day10_1",
	Short: "Day 10, challenge 1",
	Long:  `Day 10, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day10_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day10_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	maze := [][]rune{}
	start := Coordinate{x: 0, y: 0, cost: 0}
	i := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		row := []rune{}
		s := scanner.Text()
		if strings.Contains(s, "\t") {
			s = s[1:]
		}

		for j, r := range s {
			row = append(row, r)

			if r == 'S' {
				start = Coordinate{x: j, y: i, cost: 0}
			}
		}
		maze = append(maze, row)
		i++
	}

	visited := []Coordinate{}
	return FindFurthestSteps(start, &maze, &visited)
}

func FindFurthestSteps(start Coordinate, maze *[][]rune, visited *[]Coordinate) int {
	q := []Coordinate{start}

	highestCost := 0

	for len(q) > 0 {
		current := q[0]
		*visited = append(*visited, current)
		q = q[1:]

		highestCost = int(math.Max(float64(highestCost), float64(current.cost)))

		validCoords := CheckCoords(current, maze)
		for _, c := range validCoords {
			checked := false
			for _, v := range *visited {
				if c.x == v.x && c.y == v.y {
					checked = true
				}
			}

			if checked == false {
				q = append(q, c)
			}
		}

	}

	return highestCost
}

func CheckCoords(s Coordinate, maze *[][]rune) []Coordinate {
	directions := [4]Coordinate{
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
		{x: 0, y: -1},
	}

	validCoord := []Coordinate{}
	for _, d := range directions {
		c := Coordinate{x: s.x + d.x, y: s.y + d.y, cost: s.cost}

		if c.x < 0 || c.y < 0 || c.y == len(*maze) || c.x == len((*maze)[c.y]) {
			continue
		}

		r0 := (*maze)[s.y][s.x]
		r1 := (*maze)[c.y][c.x]

		if Compatible(r0, r1, d) {
			validCoord = append(validCoord, c)
		}
	}

	return validCoord
}

func Compatible(r0 rune, r1 rune, d Coordinate) bool {
	if d.x == 1 {
		if r0 == '-' || r0 == 'L' || r0 == 'F' || r0 == 'S' {
			if r1 == '-' || r1 == 'J' || r1 == '7' {
				return true
			}
		}
	}
	if d.x == -1 {
		if r0 == '-' || r0 == 'J' || r0 == '7' || r0 == 'S' {
			if r1 == '-' || r1 == 'L' || r1 == 'F' {
				return true
			}
		}
	}
	if d.y == -1 {
		if r0 == '|' || r0 == 'L' || r0 == 'J' || r0 == 'S' {
			if r1 == '|' || r1 == '7' || r1 == 'F' {
				return true
			}
		}
	}
	if d.y == 1 {
		if r0 == '|' || r0 == '7' || r0 == 'F' || r0 == 'S' {
			if r1 == '|' || r1 == 'L' || r1 == 'J' {
				return true
			}
		}
	}

	return false
}
