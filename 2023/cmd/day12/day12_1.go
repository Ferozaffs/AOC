package day12

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

var day12_1Cmd = &cobra.Command{
	Use:   "day12_1",
	Short: "Day 12, challenge 1",
	Long:  `Day 12, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day12_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day12_data.txt")
	ans := Solve(string(dat), 1)

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve(data string, multiplier int) int {
	sum := 0

	reIndex := regexp.MustCompile("\\d+")

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {

		index := []int{}

		indexScanned := reIndex.FindAllString(scanner.Text(), -1)
		for _, i := range indexScanned {
			n, _ := strconv.Atoi(i)
			index = append(index, n)
		}

		split := strings.Fields(scanner.Text())
		springs := split[0] + "?"

		if multiplier > 1 {
			orgIndex := index
			orgSprings := springs

			for i := 0; i < multiplier-1; i++ {
				for _, idx := range orgIndex {
					index = append(index, idx)
				}

				springs += orgSprings
			}
		}

		result := Count(springs, index)

		sum += result
	}

	return sum
}

var cache = map[string]int{}

func Count(s string, i []int) (sum int) {
	if sum, ok := cache[fmt.Sprint(s, i)]; ok {
		return sum
	}
	defer func() { recover(); cache[fmt.Sprint(s, i)] = sum }()

	if len(s) == 0 {
		if len(i) == 0 {
			return 1
		}
		return 0
	}

	if s[0] == '?' || s[0] == '.' {
		sum += Count(s[1:], i)
	}
	if (s[0] == '?' || s[0] == '#') &&
		s[i[0]] != '#' &&
		strings.Contains(s[:i[0]], ".") == false {
		sum += Count(s[i[0]+1:], i[1:])
	}

	return
}
