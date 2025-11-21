package day19

import (
	"aoc2024/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day19",
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
	dat, _ := os.ReadFile("inputs/day19_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	lines := strings.Split(data, "\n")

	patterns := strings.Split(lines[0], ", ")

	ans1 := 0
	ans2 := 0
	for i := 2; i < len(lines); i++ {
		cache := make(map[string]int)
		results := CheckPattern(lines[i], patterns, cache)
		if results > 0 {
			ans1++
		}

		ans2 += results
	}

	return ans1, ans2
}

func CheckPattern(s string, pattern []string, cache map[string]int) int {
	if s == "" {
		return 1
	}

	if res, ok := cache[s]; ok {
		return res
	}

	res := 0
	for _, p := range pattern {
		if strings.HasPrefix(s, p) {
			ns := s[len(p):]
			res += CheckPattern(ns, pattern, cache)
		}
	}

	cache[s] = res
	return res
}
