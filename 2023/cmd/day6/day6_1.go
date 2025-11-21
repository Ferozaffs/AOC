package day6

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day6_1Cmd = &cobra.Command{
	Use:   "day6_1",
	Short: "Day 6, challenge 1",
	Long:  `Day 6, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day6_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day6_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 1.0

	scanner := bufio.NewScanner(strings.NewReader(data))

	scanner.Scan()
	s := strings.Split(scanner.Text(), ":")
	times := strings.Fields(s[1])

	scanner.Scan()
	s = strings.Split(scanner.Text(), ":")
	distances := strings.Fields(s[1])

	for race, t := range times {
		time, _ := strconv.Atoi(t)
		midpoint := float64(time) / 2.0

		distance, _ := strconv.Atoi(distances[race])
		for i := 0.0; midpoint-i > 0.0; i++ {
			d := math.Floor(midpoint-i) * math.Ceil(midpoint+i)

			if d <= float64(distance) {
				possibleWins := i * 2
				if midpoint == math.Floor(midpoint) {
					possibleWins--
				}
				sum *= possibleWins
				break
			}
		}
	}

	return int(sum)
}
