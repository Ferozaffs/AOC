package day2

import (
	"aoc2024/cmd"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day2",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day1_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day2_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	lines := strings.Split(data, "\n")

	ans1 := 0
	ans2 := 0
	for _, line := range lines {
		parts := strings.Fields(line)

		var numbers []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers = append(numbers, num)
		}

		if Sequence(numbers, 0) {
			ans1++
		}
		if SequenceTolerance(numbers) {
			ans2++
		}
	}

	return ans1, ans2
}

func SequenceTolerance(numbers []int) bool {
	if Sequence(numbers, 0) {
		return true
	} else {
		for i := range numbers {
			prefix := append([]int{}, numbers[:i]...)
			newSequence := append(prefix, numbers[i+1:]...)

			if Sequence(newSequence, 0) {
				return true
			}
		}
	}

	return false
}

func Sequence(numbers []int, state int) bool {
	first, numbers := numbers[0], numbers[1:]

	if len(numbers) == 0 {
		return true
	}

	delta := first - numbers[0]
	deltaAbs := int(math.Abs(float64(delta)))

	if deltaAbs > 0 && deltaAbs < 4 {
		if delta > 0 && (state == 1 || state == 0) {
			return Sequence(numbers, 1)
		} else if delta < 0 && (state == -1 || state == 0) {
			return Sequence(numbers, -1)
		}
	}

	return false
}
