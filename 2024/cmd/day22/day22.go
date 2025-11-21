package day22

import (
	"aoc2024/cmd"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day22",
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
	dat, _ := os.ReadFile("inputs/day22_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	lines := strings.Split(data, "\n")

	ans1 := 0
	bananas := make(map[[4]int]int)
	for _, s := range lines {
		seen := make(map[[4]int]bool)
		seq := [4]int{999, 999, 999, 999}

		n, _ := strconv.Atoi(s)
		for i := 0; i < 2000; i++ {
			p := n % 10

			n = (n ^ (n * 64)) % 16777216
			n = (n ^ (n / 32)) % 16777216
			n = (n ^ (n * 2048)) % 16777216

			c := n % 10

			seq[0] = seq[1]
			seq[1] = seq[2]
			seq[2] = seq[3]
			seq[3] = c - p

			if !seen[seq] {
				seen[seq] = true
				bananas[seq] += c
			}
		}

		ans1 += n
	}

	ans2 := 0
	for _, v := range bananas {
		if v > ans2 {
			ans2 = v
		}
	}

	return ans1, ans2
}
