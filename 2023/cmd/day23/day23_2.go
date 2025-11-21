package day23

import (
	"aoc2023/cmd"

	"github.com/spf13/cobra"
)

var day23_2Cmd = &cobra.Command{
	Use:   "day23_2",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day23_2Cmd)
}

func Run2() {
}

func Solve2() {

}
