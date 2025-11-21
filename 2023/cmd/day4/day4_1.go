package day4

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day4_1Cmd = &cobra.Command{
	Use:   "day4_1",
	Short: "Day 4, challenge 1",
	Long:  `Day 4, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day4_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day4_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {

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

		sum += int(math.Pow(2, float64(matches-1)))

	}

	return sum
}
