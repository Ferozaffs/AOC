package day8

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var day8_2Cmd = &cobra.Command{
	Use:   "day8_2",
	Short: "Day 8, challenge 2",
	Long:  `Day 8, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day8_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day8_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	re := regexp.MustCompile("([A-Z]|[0-9])\\w+")
	s := re.FindAllString(data, -1)

	instruction := s[0]
	mapData := map[string]Direction{}

	for i := 1; i < len(s); i += 3 {
		mapData[s[i]] = Direction{left: s[i+1], right: s[i+2]}
	}

	destinations := []string{}
	for s, _ := range mapData {
		if s[2] == 'A' {
			destinations = append(destinations, s)
		}
	}

	foundSteps := []int{}
	for i, _ := range destinations {
		steps := 0
		for i == len(foundSteps) {
			for _, instruction := range instruction {
				steps++
				if instruction == 'R' {
					destinations[i] = mapData[destinations[i]].right
				} else {
					destinations[i] = mapData[destinations[i]].left
				}

				if destinations[i][2] == 'Z' {
					foundSteps = append(foundSteps, steps)
					break
				}
			}
		}
	}

	return LCM(foundSteps...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(numbers ...int) int {
	result := numbers[0] * numbers[1] / GCD(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		result = LCM(result, numbers[i])
	}

	return result
}
