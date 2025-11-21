package day13

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day13_2Cmd = &cobra.Command{
	Use:   "day13_2",
	Short: "Day 13, challenge 2",
	Long:  `Day 13, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day13_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day13_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sum := 0

	pattern := [][]byte{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "\t") {
			s = s[1:]
		}

		if len(s) == 0 {
			colsPattern := Transpose(pattern)

			sum += FindReflectionSmudge(colsPattern)

			sum += (FindReflectionSmudge(pattern) * 100)

			pattern = [][]byte{}
		} else {
			bytes := []byte(s)
			pattern = append(pattern, bytes)
		}
	}

	return sum
}

func FindReflectionSmudge(pattern [][]byte) int {
	foundMatch := 0
	backstep := 1
	for i := 1; i < len(pattern); i++ {
		if i-backstep < 0 {
			return foundMatch
		}

		s0 := string(pattern[i-backstep])
		s1 := string(pattern[i])

		if SmudgeMatch(s0, s1) {
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

func SmudgeMatch(a string, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff <= 1
}
