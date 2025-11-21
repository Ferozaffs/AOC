package day1

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day1_2Cmd = &cobra.Command{
	Use:   "day1_2",
	Short: "Day 1, challenge 2",
	Long:  `Day 1, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day1_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day1_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sum := 0
	//re := regexp.MustCompile("([0-9]|zero|one|two|three|four|five|six|seven|eight|nine)")

	words := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		//numbers := re.FindAllString(scanner.Text(), -1)

		numbers := []string{}
		line := scanner.Text()
		for i := 0; i < len(line); {
			for _, val := range words {
				if strings.HasPrefix(line[i:], val) {
					numbers = append(numbers, val)
					break
				}
			}
			i++
		}

		first, err := strconv.Atoi(numbers[0])
		if err != nil {
			first = ConvertWordToNum(numbers[0])
		}
		first *= 10

		last, err := strconv.Atoi(numbers[len(numbers)-1])
		if err != nil {
			last = ConvertWordToNum(numbers[len(numbers)-1])
		}
		sum += first + last
	}

	return sum
}

func ConvertWordToNum(word string) int {
	switch w := word; w {
	case "zero":
		return 0
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		panic(word)
	}
}
