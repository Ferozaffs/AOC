package day20

import (
	"aoc2023/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var day20_2Cmd = &cobra.Command{
	Use:   "day20_2",
	Short: "Day 20, challenge 2",
	Long:  `Day 20, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day20_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day20_data.txt")
	Solve1(string(dat), true)

	fmt.Printf("ANSWER: %d\n", pressToRX)
}
