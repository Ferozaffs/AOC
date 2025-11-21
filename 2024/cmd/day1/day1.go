package day1

import (
	"aoc2024/cmd"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day1",
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
	dat, _ := os.ReadFile("inputs/day1_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	lines := strings.Split(data, "\n")

	var left []int
	var right []int
	for _, line := range lines {
		parts := strings.Fields(line)
		lv, _ := strconv.Atoi(parts[0])
		rv, _ := strconv.Atoi(parts[1])
		left = append(left, lv)
		right = append(right, rv)
	}

	sort.Ints(left)
	sort.Ints(right)

	ans1 := 0
	ans2 := 0
	countCache := make(map[int]int)
	for i, lv := range left {
		delta := lv - right[i]
		ans1 += int(math.Abs(float64(delta)))

		val, ok := countCache[lv]
		if ok {
			ans2 += lv * val
		} else {
			count := 0
			for _, rv := range right {
				if lv == rv {
					count++
				}
			}

			countCache[lv] = count
			ans2 += lv * count
		}
	}

	return ans1, ans2
}
