package day9

import (
	"aoc2023/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var day9_2Cmd = &cobra.Command{
	Use:   "day9_2",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day9_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day9_data.txt")
	ans := Solve(string(dat), true)

	fmt.Printf("ANSWER: %d\n", ans)
}
