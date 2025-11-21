package day8

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

type Direction struct {
	left  string
	right string
}

var day8_1Cmd = &cobra.Command{
	Use:   "day8_1",
	Short: "Day 8, challenge 1",
	Long:  `Day 8, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day8_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day8_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	re := regexp.MustCompile("([A-Z])\\w+")
	s := re.FindAllString(data, -1)

	instruction := s[0]
	mapData := map[string]Direction{}

	for i := 1; i < len(s); i += 3 {
		mapData[s[i]] = Direction{left: s[i+1], right: s[i+2]}
	}

	destination := "AAA"
	steps := 0
	for destination != "ZZZ" {
		for _, instruction := range instruction {
			steps++

			if instruction == 'R' {
				destination = mapData[destination].right
			} else {
				destination = mapData[destination].left
			}

			if destination == "ZZZ" {
				break
			}
		}
	}

	return steps
}
