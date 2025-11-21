package day1

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

var day1_1Cmd = &cobra.Command{
	Use:   "day1_1",
	Short: "Day 1, challenge 1",
	Long:  `Day 1, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day1_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day1_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0
	re := regexp.MustCompile("[0-9]")

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		numbers := re.FindAllString(scanner.Text(), -1)
		combined := numbers[0] + numbers[len(numbers)-1]
		i, _ := strconv.Atoi(combined)
		sum += i
	}

	return sum
}
