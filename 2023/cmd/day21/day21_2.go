package day21

import (
	"aoc2023/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var day21_2Cmd = &cobra.Command{
	Use:   "day21_2",
	Short: "Day 21, challenge 2",
	Long:  `Day 21, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day21_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day21_data.txt")
	ans := Solve1(string(dat), 26501365, true)

	fmt.Printf("ANSWER: %d\n", ans)
}
