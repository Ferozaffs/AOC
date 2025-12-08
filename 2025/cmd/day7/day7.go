package day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"2025/cmd"

	"github.com/spf13/cobra"
)

type Coord struct {
	x int
	y int
}

type Split struct {
	coord     Coord
	timelines int
}

var DayCmd = &cobra.Command{
	Use:   "day7",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(DayCmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day7_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	rows := [][]rune{}
	splits := []Split{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		rows = append(rows, []rune(string(scanner.Bytes())))
	}

	ans2 := 0
	for i, r := range rows[0] {
		if r == 'S' {
			ans2 = TraceBeam(Coord{i, 0}, &rows, &splits)
		}
	}

	return len(splits), ans2
}

func TraceBeam(coord Coord, rows *[][]rune, splits *[]Split) int {
	newCoord := Coord{coord.x, coord.y + 1}

	if newCoord.y == len(*rows) {
		return 1
	} else if (*rows)[newCoord.y][newCoord.x] == '^' {
		for _, s := range *splits {
			if s.coord == newCoord {
				return s.timelines
			}
		}

		splitLeft := Coord{newCoord.x - 1, newCoord.y}
		splitRight := Coord{newCoord.x + 1, newCoord.y}

		timelines := 0
		if splitLeft.x >= 0 {
			timelines += TraceBeam(splitLeft, rows, splits)
		}

		if splitRight.x < len((*rows)[0]) {
			timelines += TraceBeam(splitRight, rows, splits)
		}

		(*splits) = append((*splits), Split{newCoord, timelines})

		return timelines
	}

	return TraceBeam(newCoord, rows, splits)
}
