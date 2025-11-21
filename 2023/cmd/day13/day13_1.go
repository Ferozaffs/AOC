package day13

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day13_1Cmd = &cobra.Command{
	Use:   "day13_1",
	Short: "Day 13, challenge 1",
	Long:  `Day 13, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day13_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day13_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0

	pattern := [][]byte{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "\t") {
			s = s[1:]
		}

		if len(s) == 0 {
			sum += (FindReflection(pattern) * 100)

			pattern = Transpose(pattern)

			sum += FindReflection(pattern)

			pattern = [][]byte{}
		} else {
			bytes := []byte(s)
			pattern = append(pattern, bytes)
		}
	}

	return sum
}

func Transpose(pattern [][]byte) [][]byte {
	xl := len(pattern[0])
	yl := len(pattern)
	result := make([][]byte, xl)
	for i := range result {
		result[i] = make([]byte, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = pattern[j][i]
		}
	}
	return result
}

func FindReflection(pattern [][]byte) int {
	foundMatch := 0
	backstep := 1
	for i := 1; i < len(pattern); i++ {
		if i-backstep < 0 {
			return foundMatch
		}

		s0 := string(pattern[i-backstep])
		s1 := string(pattern[i])

		if s0 == s1 {
			if foundMatch == 0 {
				foundMatch = i
			}
			backstep += 2
		} else {
			foundMatch = 0
			backstep = 1
		}
	}

	return foundMatch
}
