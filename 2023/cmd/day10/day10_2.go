package day10

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day10_2Cmd = &cobra.Command{
	Use:   "day10_2",
	Short: "Day 10, challenge 2",
	Long:  `Day 10, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day10_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day10_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
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

		row = append(row, '.')
		for _, r := range s {
			row = append(row, r)
		}
		row = append(row, '.')
		maze = append(maze, row)
		i++
	}

	maze = ExpandMaze(maze)

	for i, c := range maze {
		for j, r := range c {
			if r == 'S' {
				start = Coordinate{x: j, y: i, cost: 0}
			}
		}
	}

	visited := []Coordinate{}
	FindFurthestSteps(start, &maze, &visited)
	FloodFill(Coordinate{x: 0, y: 0, cost: 0}, &maze, visited)

	sum := 0
	for i, c := range maze {
		//fmt.Printf("\n")
		for j, r := range c {
			if r != '0' && r != ':' {
				blocked := false
				for _, b := range visited {
					if j == b.x && i == b.y {
						blocked = true
					}
				}
				if blocked == false {
					if i%2 == 1 && j%2 == 0 {
						sum++
						//fmt.Print("I")
					}
				} else {
					//fmt.Printf("%s", string(r))
				}
			} else {
				//fmt.Print(" ")
			}

		}
	}

	return sum
}

func ExpandMaze(maze [][]rune) [][]rune {
	nm := [][]rune{}
	for _, c := range maze {
		row := []rune{}
		for _, r := range c {
			row = append(row, r)
			row = append(row, ':')
		}
		nm = append(nm, row)

		extraRow := []rune{}
		for i := 0; i < len(row); i++ {
			extraRow = append(extraRow, ':')
		}
		nm = append(nm, extraRow)
	}

	extraRow := []rune{}
	for i := 0; i < len((nm)[0]); i++ {
		extraRow = append(extraRow, '.')
	}
	nm = append([][]rune{extraRow}, nm...)
	nm = append(nm, extraRow)

	for c := 1; c < len(nm)-1; c++ {
		for r := 1; r < len(nm[c])-1; r++ {
			if nm[c][r] == ':' {
				rr := nm[c][r+1]
				rl := nm[c][r-1]
				cd := nm[c+1][r]
				cu := nm[c-1][r]

				if (rl == '-' || rl == 'L' || rl == 'F' || rl == 'S') && (rr == '-' || rr == 'J' || rr == '7' || rr == 'S') {
					nm[c][r] = '-'
				} else if (cu == '|' || cu == 'F' || cu == '7' || cu == 'S') && (cd == '|' || cd == 'L' || cd == 'J' || cd == 'S') {
					nm[c][r] = '|'
				}
			}
		}
	}

	//for _, row := range nm {
	//	fmt.Printf("\n")
	//	for _, c := range row {
	//		fmt.Printf("%s", string(c))
	//	}
	//}

	return nm
}

func FloodFill(start Coordinate, maze *[][]rune, blockers []Coordinate) {
	q := []Coordinate{start}
	visited := []Coordinate{start}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		blocked := false
		for _, b := range blockers {
			if current.x == b.x && current.y == b.y {
				blocked = true
			}
		}

		if blocked == false {
			(*maze)[current.y][current.x] = '0'
		}

		validCoords := NextFill(current, maze, blockers)
		for _, c := range validCoords {
			checked := false
			for _, v := range visited {
				if c.x == v.x && c.y == v.y {
					checked = true
				}
			}

			if checked == false {
				visited = append(visited, c)
				q = append(q, c)
			}
		}
	}
}

func NextFill(s Coordinate, maze *[][]rune, blockers []Coordinate) []Coordinate {
	directions := [4]Coordinate{
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
		{x: 0, y: -1},
	}

	validCoord := []Coordinate{}
	for _, d := range directions {
		c := Coordinate{x: s.x + d.x, y: s.y + d.y, cost: 0}

		if c.x < 0 || c.y < 0 || c.y == len(*maze) || c.x == len((*maze)[c.y]) {
			continue
		}

		blocked := false
		for _, b := range blockers {
			if c.x == b.x && c.y == b.y {
				blocked = true
			}
		}

		if blocked == true {
			continue
		} else {
			validCoord = append(validCoord, c)
		}
	}

	return validCoord
}
