package day4

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day4_2Cmd = &cobra.Command{
	Use:   "day4_2",
	Short: "Day 4, challenge 2",
	Long:  `Day 4, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day4_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day4_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sum := 0

	i := 0
	played := map[int]int{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		played[i] += 1

		s := strings.Split(scanner.Text(), ":") //Split card
		s = strings.Split(s[1], "|")            // Split numbers

		winningNumbers := []int{}
		for _, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(s[0], -1) {
			n, _ := strconv.Atoi(s[0][m[0]:m[1]])
			winningNumbers = append(winningNumbers, n)
		}

		matches := 0
		for _, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(s[1], -1) {
			n, _ := strconv.Atoi(s[1][m[0]:m[1]])
			if slices.Contains(winningNumbers, n) {
				matches++
			}
		}
		for j := 0; j < matches; j++ {
			played[i+j+1] += 1 * played[i]
		}

		i++
	}

	for _, n := range played {
		sum += n
	}
	return sum
}
