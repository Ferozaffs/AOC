package day15

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day15_1Cmd = &cobra.Command{
	Use:   "day15_1",
	Short: "Day 15, challenge 1",
	Long:  `Day 15, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day15_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day15_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0

	sequence := strings.Split(data, ",")

	for _, s := range sequence {
		value := 0
		for _, r := range s {
			value += int(r)
			value *= 17
			value = value % 256
		}

		sum += value
	}

	return sum
}
