package day7

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
	Use:   "day7",
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
	dat, _ := os.ReadFile("inputs/day7_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	lines := strings.Split(data, "\n")

	ans1 := 0
	ans2 := 0
	for _, line := range lines {
		{
			r, s := TestEquation(line, false)
			if s {
				ans1 += r
			}
		}
		{
			r, s := TestEquation(line, true)
			if s {
				ans2 += r
			}
		}
	}

	return ans1, ans2
}

func TestEquation(data string, useBase3 bool) (int, bool) {
	s := strings.Split(data, ":")
	result, _ := strconv.Atoi(s[0])

	numberStrings := strings.Fields(s[1])

	numbers := []int{}
	for _, ns := range numberStrings {
		n, _ := strconv.Atoi(ns)
		numbers = append(numbers, n)
	}

	permutations := 0
	if useBase3 {
		permutations = int(math.Pow(3.0, float64(len(numbers)-1)))
	} else {
		permutations = int(math.Pow(2.0, float64(len(numbers)-1)))
	}

	for i := 0; i < permutations; i++ {
		bs := ""

		if useBase3 {
			bs = ToBase3(len(numbers)-1, i)
		} else {
			bs = fmt.Sprintf("%0*b", len(numbers)-1, i)
		}

		sum := numbers[0]
		for j := 1; j < len(numbers); j++ {
			if useBase3 && j <= len(bs) && bs[j-1] == '2' {
				concat := fmt.Sprintf("%d%d", sum, numbers[j])
				sum, _ = strconv.Atoi(concat)
			} else if j <= len(bs) && bs[j-1] == '1' {
				sum *= numbers[j]
			} else {
				sum += numbers[j]
			}
		}

		if sum == result {
			return result, true
		}
	}

	return result, false
}

func ToBase3(numOperator int, index int) string {
	b3 := ""
	for index > 0 {
		b3 = fmt.Sprintf("%d%s", index%3, b3)
		index /= 3
	}
	return strings.Repeat("0", numOperator-len(b3)) + b3
}
