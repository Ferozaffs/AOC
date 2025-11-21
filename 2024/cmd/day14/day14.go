package day14

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

type Robot struct {
	pos helpers.Point
	vel helpers.Point
}

var sizeX = 101
var sizeY = 103

var day1_1Cmd = &cobra.Command{
	Use:   "day14",
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
	dat, _ := os.ReadFile("inputs/day14_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func SetSize(x int, y int) {
	sizeX = x
	sizeY = y
}

func Solve(data string) (int, int) {
	lines := strings.Split(data, "\n")

	re := regexp.MustCompile(`-?\d+`)

	robots := []Robot{}
	for _, l := range lines {
		parts := strings.Split(l, " ")
		posElements := re.FindAllString(parts[0], -1)
		velElements := re.FindAllString(parts[1], -1)

		px, _ := strconv.Atoi(posElements[0])
		py, _ := strconv.Atoi(posElements[1])
		vx, _ := strconv.Atoi(velElements[0])
		vy, _ := strconv.Atoi(velElements[1])

		robots = append(robots, Robot{pos: helpers.Point{X: px, Y: py}, vel: helpers.Point{X: vx, Y: vy}})
	}

	halfSizeX := sizeX / 2
	halfSizeY := sizeY / 2

	quadrant := []int{0, 0, 0, 0}
	for _, r := range robots {
		x := (r.pos.X + 100*r.vel.X) % sizeX
		y := (r.pos.Y + 100*r.vel.Y) % sizeY

		if x < 0 {
			x += sizeX
		}
		if y < 0 {
			y += sizeY
		}

		if x < halfSizeX && y < halfSizeY {
			quadrant[0]++
		} else if x > halfSizeX && y < halfSizeY {
			quadrant[1]++
		} else if x < halfSizeX && y > halfSizeY {
			quadrant[2]++
		} else if x > halfSizeX && y > halfSizeY {
			quadrant[3]++
		}
	}

	ans1 := 1
	for _, q := range quadrant {
		ans1 *= q
	}

	ans2 := 0
	//out:
	//	for {
	//		yMap := make(map[int]int)
	//		for _, r := range robots {
	//			y := (r.pos.Y + ans2*r.vel.Y) % sizeY
	//			if y < 0 {
	//				y += sizeY
	//			}
	//
	//			yMap[r.pos.Y]++
	//		}
	//
	//		ans2++
	//		for _, ym := range yMap {
	//			if ym > 11 {
	//				break out
	//			}
	//		}
	//	}

	return ans1, ans2

}
