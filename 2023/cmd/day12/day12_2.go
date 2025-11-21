package day12

import (
	"aoc2023/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var day12_2Cmd = &cobra.Command{
	Use:   "day12_2",
	Short: "Day 12, challenge 2",
	Long:  `Day 12, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day12_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day12_data.txt")
	ans := Solve(string(dat), 5)

	fmt.Printf("ANSWER: %d\n", ans)
}
