package day2

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/spf13/cobra"
)

var day2_1Cmd = &cobra.Command{
	Use:   "day2_1",
	Short: "Day 2, challenge 1",
	Long:  `Day 2, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day2_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day2_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	cubeLimits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sumGames := 0

	reGame := regexp2.MustCompile("((?!Game )\\d+(?=\\:))", 0)
	reRed := regexp2.MustCompile("\\d+(?= red)", 0)
	reGreen := regexp2.MustCompile("\\d+(?= green)", 0)
	reBlue := regexp2.MustCompile("\\d+(?= blue)", 0)

	scanner := bufio.NewScanner(strings.NewReader(data))
OUTER:
	for scanner.Scan() {
		game := regexp2FindAllString(reGame, scanner.Text())
		numRed := regexp2FindAllString(reRed, scanner.Text())
		numGreen := regexp2FindAllString(reGreen, scanner.Text())
		numBlue := regexp2FindAllString(reBlue, scanner.Text())

		for _, val := range numRed {
			iVal, _ := strconv.Atoi(val)
			if iVal > cubeLimits["red"] {
				continue OUTER
			}
		}
		for _, val := range numGreen {
			iVal, _ := strconv.Atoi(val)
			if iVal > cubeLimits["green"] {
				continue OUTER
			}
		}
		for _, val := range numBlue {
			iVal, _ := strconv.Atoi(val)
			if iVal > cubeLimits["blue"] {
				continue OUTER
			}
		}

		i, _ := strconv.Atoi(game[0])
		sumGames += i
	}

	return sumGames
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
