package day5

import (
	"aoc2024/cmd"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day5",
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
	dat, _ := os.ReadFile("inputs/day5_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	ans1 := 0
	ans2 := 0

	pageMap := make(map[string][]string)

	stage := 0

	lines := strings.Split(data, "\n")
	for _, l := range lines {
		if len(l) == 0 {
			stage++
			continue
		}

		if stage == 0 {
			s := strings.Split(l, "|")
			pageMap[s[0]] = append(pageMap[s[0]], s[1])
		} else {
			values := strings.Split(l, ",")
			if CheckSequence(values, pageMap) {
				mv, _ := strconv.Atoi(values[len(values)/2])
				ans1 += mv
			} else {
				sortedValues := SortList(values, pageMap)

				mv, _ := strconv.Atoi(sortedValues[len(sortedValues)/2])
				ans2 += mv
			}
		}
	}

	return ans1, ans2
}

func CheckSequence(values []string, pageMap map[string][]string) bool {
	var parsedValues []string
	for _, v := range values {
		p, ok := pageMap[v]
		if ok {
			for _, pn := range p {
				if slices.Contains(parsedValues, pn) {
					return false
				}
			}
		}

		parsedValues = append(parsedValues, v)
	}

	return true
}

func SortList(values []string, pageMap map[string][]string) []string {
	var sortedList []string

	for _, v := range values {
		p := pageMap[v]

		found := false
	out:
		for i := range sortedList {
			for _, pn := range p {
				if pn == sortedList[i] {
					sortedList = append(sortedList[:i], append([]string{v}, sortedList[i:]...)...)
					found = true
					break out
				}
			}
		}

		if !found {
			sortedList = append(sortedList, v)
		}
	}

	return sortedList
}
