package day9

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day9_1Cmd = &cobra.Command{
	Use:   "day9_1",
	Short: "Day 9, challenge 1",
	Long:  `Day 9, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day9_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day9_data.txt")
	ans := Solve(string(dat), false)

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve(data string, reverse bool) int {
	sum := 0
	re := regexp.MustCompile("(-|\\d)+")

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)

		numbers := []int{}
		for _, m := range matches {
			n, _ := strconv.Atoi(m)
			numbers = append(numbers, n)
		}

		sum += Extrapolate(numbers, reverse)
	}

	return sum
}

func Extrapolate(n []int, reverse bool) int {
	deltas := []int{}
	sum := 0
	for i := 1; i < len(n); i++ {
		d := n[i] - n[i-1]
		sum += d
		deltas = append(deltas, d)
	}

	index := len(n) - 1
	if reverse {
		index = 0
	}

	if sum == 0 {
		return n[index]
	} else {
		return n[index] - Extrapolate(deltas, reverse)
	}
}
