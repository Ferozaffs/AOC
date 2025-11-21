package day15

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day15",
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
	dat, _ := os.ReadFile("inputs/day15_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	parts := strings.Split(data, "\n\n")

	var grid1 helpers.Grid
	grid1.Init(parts[0])

	ans1 := RunGrid(grid1, parts[1])

	var grid2 helpers.Grid
	grid2.Init(parts[0])
	modifiedGrid := ModifyGrid(grid2)
	ans2 := RunGrid(modifiedGrid, parts[1])

	return ans1, ans2
}

func RunGrid(grid helpers.Grid, instruction string) int {
	robotPos := helpers.Point{}
out:
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '@' {
				robotPos = helpers.Point{X: x, Y: y}
				grid[x][y] = '.'
				break out
			}
		}
	}

	for _, r := range instruction {
		robotPos = MoveRobot(r, robotPos, &grid)

		//for x := range grid {
		//	fmt.Print("\n")
		//	for y := range grid[x] {
		//		if robotPos.X == x && robotPos.Y == y {
		//			fmt.Print("@")
		//		} else {
		//			fmt.Print(string(grid[x][y]))
		//		}
		//	}
		//}
	}

	ans := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'O' || grid[x][y] == '[' {
				ans += 100*x + y
			}
		}
	}

	return ans
}

func ModifyGrid(grid helpers.Grid) helpers.Grid {
	var modifiedGrid helpers.Grid

	for x := range grid {
		var runes []rune
		for y := range grid[x] {
			if grid[x][y] == '#' {
				runes = append(runes, '#')
				runes = append(runes, '#')
			} else if grid[x][y] == 'O' {
				runes = append(runes, '[')
				runes = append(runes, ']')
			} else if grid[x][y] == '.' {
				runes = append(runes, '.')
				runes = append(runes, '.')
			} else {
				runes = append(runes, '@')
				runes = append(runes, '.')
			}
		}
		modifiedGrid = append(modifiedGrid, runes)
	}

	return modifiedGrid
}

func MoveRobot(r rune, p helpers.Point, g *helpers.Grid) helpers.Point {
	d := GetDirection(r)

	nr, _ := (*g).GetPointDir(p.X, p.Y, d.X, d.Y)

	if nr == '#' {
		return p
	} else if nr == '.' {
		p.X += d.X
		p.Y += d.Y

		return p
	} else if nr == 'O' {
		o := p
		o.X += d.X
		o.Y += d.Y

		if MoveObject(o, d, g, false) {
			return o
		}
	} else if nr == '[' || nr == ']' {
		o := p
		o.X += d.X
		o.Y += d.Y

		if d.X != 0 {
			firstSuccess := MoveObject(o, d, g, true)

			o2 := p
			if nr == '[' {
				o2.Y += 1
			} else {
				o2.Y -= 1
			}
			o2.X += d.X
			o2.Y += d.Y

			if firstSuccess && MoveObject(o2, d, g, true) {
				MoveObject(o, d, g, false)
				MoveObject(o2, d, g, false)
				return o
			}
		} else {
			if MoveObject(o, d, g, false) {
				return o
			}
		}

	}

	return p
}

func MoveObject(p helpers.Point, d helpers.Point, g *helpers.Grid, preCheck bool) bool {
	nr, _ := (*g).GetPointDir(p.X, p.Y, d.X, d.Y)

	if nr == '#' {
		return false
	} else if nr == '.' {
		if !preCheck {
			(*g)[p.X+d.X][p.Y+d.Y] = (*g)[p.X][p.Y]
			(*g)[p.X][p.Y] = '.'
		}

		return true
	} else if nr == 'O' {
		o := p
		o.X += d.X
		o.Y += d.Y

		if MoveObject(o, d, g, false) {
			(*g)[p.X+d.X][p.Y+d.Y] = (*g)[p.X][p.Y]
			(*g)[p.X][p.Y] = '.'

			return true
		}
	} else if nr == '[' || nr == ']' {
		o := p
		o.X += d.X
		o.Y += d.Y

		p2 := p
		if nr == '[' {
			p2.Y += 1
		} else {
			p2.Y -= 1
		}

		o2 := p2
		o2.X += d.X
		o2.Y += d.Y

		if d.X != 0 {
			if preCheck {
				firstSuccess := MoveObject(o, d, g, true)
				secondSuccess := MoveObject(o2, d, g, true)

				if firstSuccess && secondSuccess {
					return true
				} else {
					return false
				}
			} else {
				MoveObject(o, d, g, false)
				MoveObject(o2, d, g, false)

				(*g)[o.X][o.Y] = (*g)[p.X][p.Y]
				(*g)[p.X][p.Y] = '.'

				return true
			}
		} else {
			if MoveObject(o, d, g, false) {
				if !preCheck {
					(*g)[p.X+d.X][p.Y+d.Y] = (*g)[p.X][p.Y]
					(*g)[p.X][p.Y] = '.'
				}

				return true
			}
		}
	}

	return false
}

func GetDirection(r rune) helpers.Point {
	if r == '<' {
		return helpers.Point{X: 0, Y: -1}
	}
	if r == '^' {
		return helpers.Point{X: -1, Y: 0}
	}
	if r == '>' {
		return helpers.Point{X: 0, Y: 1}
	}
	if r == 'v' {
		return helpers.Point{X: 1, Y: 0}
	}

	return helpers.Point{X: 0, Y: 0}
}
