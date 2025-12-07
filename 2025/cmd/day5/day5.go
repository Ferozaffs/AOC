package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"2025/cmd"

	"github.com/spf13/cobra"
)

type Interval struct {
	start int
	end   int
}

var DayCmd = &cobra.Command{
	Use:   "day5",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(DayCmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day5_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	intervals := []Interval{}
	ingredients := []int{}

	parseIngredients := false
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		row := string(scanner.Bytes())
		if len(row) == 0 {
			parseIngredients = true
		} else if parseIngredients {
			ingredient, _ := strconv.Atoi(row)
			ingredients = append(ingredients, ingredient)
		} else {
			index := strings.Split(row, "-")
			s, _ := strconv.Atoi(index[0])
			e, _ := strconv.Atoi(index[1])

			interval := Interval{start: s, end: e}
			intervals = append(intervals, interval)
		}
	}

	merged := true
	for merged {
		merged, intervals = MergeIntervals(intervals)
	}

	ans1 := CheckFresh(ingredients, intervals)

	ans2 := 0
	for _, interval := range intervals {
		delta := interval.end - interval.start + 1
		ans2 += delta
	}

	return ans1, ans2
}

func MergeIntervals(intervals []Interval) (bool, []Interval) {
	merged := false

	i := 0
	for i < len(intervals) {
		j := i + 1
		for j < len(intervals) {
			base := intervals[i]
			other := intervals[j]

			merge := false
			if other.start >= base.start && other.start <= base.end {
				merge = true
			} else if other.end >= base.start && other.end <= base.end {
				merge = true
			} else if other.start <= base.start && other.end >= base.end {
				merge = true
			}

			if merge {
				merged = true

				intervals[i].start = int(math.Min(float64(base.start), float64(other.start)))
				intervals[i].end = int(math.Max(float64(base.end), float64(other.end)))

				intervals = append(intervals[:j], intervals[j+1:]...)
			} else {
				j++
			}
		}
		i++
	}

	return merged, intervals
}

func CheckFresh(ingredients []int, intervals []Interval) int {
	fresh := 0

	for _, ingredient := range ingredients {
		for _, interval := range intervals {
			if ingredient >= interval.start && ingredient <= interval.end {
				fresh++
				break
			}
		}
	}

	return fresh
}
