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

var day6_2Cmd = &cobra.Command{
	Use:   "day6_2",
	Short: "Day 6, challenge 2",
	Long:  `Day 6, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day6_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day6_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sum := 1.0

	scanner := bufio.NewScanner(strings.NewReader(data))

	scanner.Scan()
	s := strings.Split(scanner.Text(), ":")
	s = strings.Fields(s[1])

	timeAsString := ""
	for _, n := range s {
		timeAsString += n
	}
	time, _ := strconv.Atoi(timeAsString)

	scanner.Scan()
	s = strings.Split(scanner.Text(), ":")
	s = strings.Fields(s[1])
	distAsString := ""
	for _, n := range s {
		distAsString += n
	}
	distance, _ := strconv.Atoi(distAsString)

	midpoint := float64(time) / 2.0
	for i := 0.0; midpoint-i > 0.0; i++ {
		d := math.Floor(midpoint-i) * math.Ceil(midpoint+i)

		if d <= float64(distance) {
			possibleWins := i*2 - 1
			sum *= possibleWins
			break
		}
	}

	return int(sum)
}
