package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"2025/cmd"

	"github.com/spf13/cobra"
)

var DayCmd = &cobra.Command{
	Use:   "day6",
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
	dat, _ := os.ReadFile("inputs/day6_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	numbers := [][]int{}
	runes := [][]rune{}
	operators := []string{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		row := string(scanner.Bytes())
		if strings.Contains(row, "*") {
			operators = strings.Fields(row)
		} else {
			runes = append(runes, []rune(row))

			numberStrs := strings.Fields(row)

			numberRow := []int{}
			for _, numberStr := range numberStrs {
				num, _ := strconv.Atoi(numberStr)
				numberRow = append(numberRow, num)
			}

			numbers = append(numbers, numberRow)
		}
	}

	ans1 := 0
	ans2 := 0

	runeIndex := 0
	for i := 0; i < len(numbers[0]); i++ {
		operator := operators[i]
		value := numbers[0][i]

		for j := 1; j < len(numbers); j++ {
			if operator == "*" {
				value *= numbers[j][i]
			} else {
				value += numbers[j][i]
			}
		}

		ans1 += value

		createdNumbers := []string{}
		empty := false
		for !empty && runeIndex < len(runes[0]) {
			empty = true

			number := ""
			for j := 0; j < len(runes); j++ {
				if runes[j][runeIndex] != ' ' {
					number += string(runes[j][runeIndex])
					empty = false
				}
			}

			runeIndex++
			if runeIndex == len(runes[0]) {
				createdNumbers = append(createdNumbers, number)
			}

			if empty || runeIndex == len(runes[0]) {
				num, _ := strconv.Atoi(createdNumbers[0])
				value = num
				for j := 1; j < len(createdNumbers); j++ {
					num, _ = strconv.Atoi(createdNumbers[j])

					if operator == "*" {
						value *= num
					} else {
						value += num
					}
				}
				ans2 += value
			} else {
				createdNumbers = append(createdNumbers, number)
			}
		}
	}

	return ans1, ans2
}
