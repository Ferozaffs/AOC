package day21

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day21",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func init() {
	cmd.RootCmd.AddCommand(day1_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day21_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	numpad := make(map[rune]helpers.Point)
	direction := make(map[helpers.Point]helpers.Point)

	numpad['7'] = helpers.Point{X: 0, Y: 3}
	numpad['8'] = helpers.Point{X: 1, Y: 3}
	numpad['9'] = helpers.Point{X: 2, Y: 3}
	numpad['4'] = helpers.Point{X: 0, Y: 2}
	numpad['5'] = helpers.Point{X: 1, Y: 2}
	numpad['6'] = helpers.Point{X: 2, Y: 2}
	numpad['1'] = helpers.Point{X: 0, Y: 1}
	numpad['2'] = helpers.Point{X: 1, Y: 1}
	numpad['3'] = helpers.Point{X: 2, Y: 1}
	numpad['0'] = helpers.Point{X: 1, Y: 0}
	numpad['A'] = helpers.Point{X: 2, Y: 0}

	direction[helpers.Point{X: 0, Y: 1}] = helpers.Point{X: 1, Y: 1}
	direction[helpers.Point{X: 0, Y: 0}] = helpers.Point{X: 2, Y: 1}
	direction[helpers.Point{X: -1, Y: 0}] = helpers.Point{X: 0, Y: 0}
	direction[helpers.Point{X: 0, Y: -1}] = helpers.Point{X: 1, Y: 0}
	direction[helpers.Point{X: 1, Y: 0}] = helpers.Point{X: 2, Y: 0}

	lines := strings.Split(data, "\n")
	ans1 := 0
	for _, s := range lines {
		dirs := SolveNumpad(s, numpad)
		dirs = SolveDirection(dirs, direction)
		dirs = SolveDirection(dirs, direction)

		re := regexp.MustCompile(`\d+`)
		numStr := re.FindString(s)
		num, _ := strconv.Atoi(numStr)

		ins := len(dirs)
		ans1 += ins * num
	}

	return ans1, 0
}

func SolveNumpad(s string, np map[rune]helpers.Point) []helpers.Point {
	dirs := []helpers.Point{}

	cp := helpers.Point{X: 2, Y: 0}
	for _, r := range s {
		np := np[r]
		HandleDelta(&cp, np, helpers.Point{X: 0, Y: 0}, &dirs)
	}

	return dirs
}

func SolveDirection(pts []helpers.Point, d map[helpers.Point]helpers.Point) []helpers.Point {
	dirs := []helpers.Point{}

	cp := helpers.Point{X: 2, Y: 1}
	for _, p := range pts {
		np := d[p]
		HandleDelta(&cp, np, helpers.Point{X: 0, Y: 1}, &dirs)
	}

	return dirs
}

func HandleDelta(cp *helpers.Point, np helpers.Point, block helpers.Point, dirs *[]helpers.Point) {
	dx := np.X - cp.X
	dy := np.Y - cp.Y

	mod := 1
	if dx < 0 {
		mod = -1
	}
	for i := 0; i < Abs(dx); i++ {

		*dirs = append(*dirs, helpers.Point{X: 1 * mod, Y: 0})
	}

	mod = 1
	if dy < 0 {
		mod = -1
	}
	for i := 0; i < Abs(dy); i++ {
		*dirs = append(*dirs, helpers.Point{X: 0, Y: 1 * mod})
	}

	cp.X += dx
	cp.Y += dy

	*dirs = append(*dirs, helpers.Point{X: 0, Y: 0})
}
