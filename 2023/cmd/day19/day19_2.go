package day19

import (
	"aoc2023/cmd"

	"github.com/spf13/cobra"
)

var day19_2Cmd = &cobra.Command{
	Use:   "day19_2",
	Short: "Day 19, challenge 2",
	Long:  `Day 19, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day19_2Cmd)
}

func Run2() {
}

func Solve2() {

}
