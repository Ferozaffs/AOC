package day14

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var day14_2Cmd = &cobra.Command{
	Use:   "day14_2",
	Short: "Day 14, challenge 1",
	Long:  `Day 14, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day14_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day14_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	load := 0

	r := [][]byte{}
	for _, s := range strings.Fields(data) {
		r = append(r, []byte(s))
	}

	for i := 0; i < 1000000000; i++ {
		load = 0

		r = Transpose(r, false)
		r = MoveRocks(r)

		r = Transpose(r, false)
		r = MoveRocks(r)

		r = Transpose(r, true)
		r = MoveRocks(r)

		r = Transpose(r, true)
		r = MoveRocks(r)

		for i, row := range r {
			for _, s := range row {
				if s == 'O' {
					load += i + 1
				}
			}
		}

		fmt.Printf("%d - %d\n", i, load)
	}

	return load
}

func Transpose(pattern [][]byte, invert bool) [][]byte {
	xl := len(pattern[0])
	yl := len(pattern)
	result := make([][]byte, xl)
	for i := range result {
		result[i] = make([]byte, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			if invert == true {
				result[i][j] = pattern[yl-j-1][i]
			} else {
				result[i][j] = pattern[j][i]
			}
		}
	}
	return result
}

func MoveRocks(pattern [][]byte) [][]byte {
	result := [][]byte{}
	for _, s := range pattern {
		block := 0
		rocks := map[int]int{}
		rocks[block]++
		for i := 0; i < len(s); i++ {
			if s[i] == 'O' {
				rocks[block]++
			} else if s[i] == '#' {
				block = i + 1
				rocks[block]++
			}
		}

		keys := make([]int, 0)
		for k, _ := range rocks {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		row := []byte{}
		for _, k := range keys {
			for i := len(row) + 1; i < k; i++ {
				row = append(row, '.')
			}
			if k > 0 {
				row = append(row, '#')
			}

			for i := 0; i < (rocks[k] - 1); i++ {
				row = append(row, 'O')
			}
		}

		for i := len(row); i < len(s); i++ {
			row = append(row, '.')
		}

		result = append(result, row)
	}

	return result
}
