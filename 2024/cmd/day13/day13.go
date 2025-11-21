package day13

import (
	"aoc2024/cmd"
	"aoc2024/helpers"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Machine struct {
	a helpers.Point
	b helpers.Point
	d helpers.Point
}

var day1_1Cmd = &cobra.Command{
	Use:   "day13",
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
	dat, _ := os.ReadFile("inputs/day13_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	lines := strings.Split(data, "\n")
	var machines []Machine

	for i := 0; i < len(lines); {
		if len(lines[i]) == 0 {
			i++
			continue
		}

		aLine := lines[i]
		bLine := lines[i+1]
		dLine := lines[i+2]

		a := Extract(aLine)
		b := Extract(bLine)
		d := Extract(dLine)

		machines = append(machines, Machine{
			a: a,
			b: b,
			d: d,
		})

		i += 3
	}

	ans1 := 0
	for _, m := range machines {
		a, b := CramersRule(m.a.X, m.a.Y, m.b.X, m.b.Y, m.d.X, m.d.Y)

		if a > 0 && b > 0 {
			ans1 += a*3 + b
		}
	}

	ans2 := 0
	for _, m := range machines {
		a, b := CramersRule(m.a.X, m.a.Y, m.b.X, m.b.Y, m.d.X+10000000000000, m.d.Y+10000000000000)

		if a > 0 && b > 0 {
			ans2 += a*3 + b
		}
	}

	return ans1, ans2
}

func Extract(s string) helpers.Point {
	reQuery := `\d+`
	re := regexp.MustCompile(reQuery)
	matches := re.FindAllString(s, -1)

	x, _ := strconv.Atoi(matches[0])
	y, _ := strconv.Atoi(matches[1])

	return helpers.Point{X: x, Y: y}
}

func CramersRule(ax int, ay int, bx int, by int, dx int, dy int) (int, int) {
	det := ax*by - ay*bx

	if det == 0 {
		return -1, -1
	}

	d1 := dx*by - dy*bx
	d2 := ax*dy - ay*dx

	if d1%det != 0 || d2%det != 0 {
		return -1, -1
	}

	return d1 / det, d2 / det
}
