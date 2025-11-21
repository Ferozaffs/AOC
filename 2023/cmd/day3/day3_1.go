package day3

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/spf13/cobra"
)

var day3_1Cmd = &cobra.Command{
	Use:   "day3_1",
	Short: "Day 3, challenge 1",
	Long:  `Day 3, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day3_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day3_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0

	re := regexp2.MustCompile("\\d+", 0)
	rows := []string{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\t", "")
		rows = append(rows, line)
	}

	for i, s := range rows {
		m, _ := re.FindStringMatch(s)
		for m != nil {
			if ValidatePartNum(rows, i, m.Index, m.Length) {
				n, _ := strconv.Atoi(m.String())
				sum += n
			}

			m, _ = re.FindNextMatch(m)
		}
	}

	return sum
}

func ValidatePartNum(rows []string, rowIndex int, numIndex int, length int) bool {
	re := regexp2.MustCompile("[^\\d\\.]", 0)
	check := func(row string, start int, end int) bool {
		cut := row[start:end]
		m, _ := re.FindStringMatch(cut)
		if m != nil {
			return true
		}

		return false
	}

	startIndex := int(math.Max(0, float64(numIndex)-1))
	endIndex := int(math.Min(float64(len(rows[rowIndex])), float64(numIndex+length)+1))
	startRow := int(math.Max(0, float64(rowIndex)-1))
	endRow := int(math.Min(float64(len(rows)-1), float64(rowIndex)+1))

	for i := startRow; i <= endRow; i++ {
		if check(rows[i], startIndex, endIndex) {
			return true
		}
	}

	return false
}
