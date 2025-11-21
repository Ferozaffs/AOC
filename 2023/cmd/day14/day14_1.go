package day14

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day14_1Cmd = &cobra.Command{
	Use:   "day14_1",
	Short: "Day 14, challenge 1",
	Long:  `Day 14, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day14_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day14_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	load := 0

	transposed := []string{}
	for i, s := range strings.Fields(data) {
		for j, r := range s {
			if i == 0 {
				transposed = append(transposed, "")
			}

			transposed[j] += string(r)
		}
	}

	for _, row := range transposed {
		block := 0
		rocks := map[int]int{}
		for i := 0; i < len(row); i++ {
			if row[i] == 'O' {
				rocks[block]++
			} else if row[i] == '#' {
				block = i + 1
			}
		}

		result := 0
		for k, v := range rocks {
			for i := 0; i < v; i++ {
				result += (len(row) - k - i)
			}
		}

		load += result
	}

	return load
}
