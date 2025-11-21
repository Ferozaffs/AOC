package day11

import (
	"aoc2024/cmd"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day11",
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
	dat, _ := os.ReadFile("inputs/day11_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	stones := strings.Fields(data)
	mappedStones := make(map[int]int)
	for _, s := range stones {
		v, _ := strconv.Atoi(s)
		mappedStones[v]++
	}

	ans1 := 0
	ans2 := 0
	for i := 0; i < 75; i++ {
		updatedStones := make(map[int]int)
		for v, c := range mappedStones {
			if v == 0 {
				updatedStones[1] += c
				continue
			}

			s := strconv.Itoa(v)
			if len(s)%2 == 0 {
				mid := len(s) / 2
				v1, _ := strconv.Atoi(s[:mid])
				v2, _ := strconv.Atoi(s[mid:])
				updatedStones[v1] += c
				updatedStones[v2] += c
				continue
			}

			updatedStones[v*2024] += c
		}

		mappedStones = updatedStones

		if i == 24 {
			for _, c := range mappedStones {
				ans1 += c
			}
		}
	}

	for _, c := range mappedStones {
		ans2 += c
	}

	return ans1, ans2
}
