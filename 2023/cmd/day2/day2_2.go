package day2

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/spf13/cobra"
)

var day2_2Cmd = &cobra.Command{
	Use:   "day2_2",
	Short: "Day 2, challenge 2",
	Long:  `Day 2, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day2_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day2_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sumGames := 0

	reRed := regexp2.MustCompile("\\d+(?= red)", 0)
	reGreen := regexp2.MustCompile("\\d+(?= green)", 0)
	reBlue := regexp2.MustCompile("\\d+(?= blue)", 0)

	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		numRedS := regexp2FindAllString(reRed, scanner.Text())
		numGreenS := regexp2FindAllString(reGreen, scanner.Text())
		numBlueS := regexp2FindAllString(reBlue, scanner.Text())

		var numRedI = []int{}
		var numGreenI = []int{}
		var numBlueI = []int{}

		for _, val := range numRedS {
			i, _ := strconv.Atoi(val)
			numRedI = append(numRedI, i)
		}
		for _, val := range numGreenS {
			i, _ := strconv.Atoi(val)
			numGreenI = append(numGreenI, i)
		}
		for _, val := range numBlueS {
			i, _ := strconv.Atoi(val)
			numBlueI = append(numBlueI, i)
		}

		sort.Ints(numRedI)
		sort.Ints(numGreenI)
		sort.Ints(numBlueI)

		sumGames += numRedI[len(numRedI)-1] * numGreenI[len(numGreenI)-1] * numBlueI[len(numBlueI)-1]
	}

	return sumGames
}
