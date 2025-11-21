package day15

import (
	"aoc2023/cmd"

	"github.com/spf13/cobra"
)

var day15_2Cmd = &cobra.Command{
	Use:   "day15_2",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day15_2Cmd)
}

func Run2() {
}

func Solve2() {

}
