package day22

import (
	"aoc2023/cmd"

	"github.com/spf13/cobra"
)

var day22_1Cmd = &cobra.Command{
	Use:   "day22_1",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day22_1Cmd)
}

func Run1() {

}

func Solve1() {

}
