package day3

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"2025/cmd"

	"github.com/spf13/cobra"
)

var DayCmd = &cobra.Command{
	Use:   "day3",
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
	dat, _ := os.ReadFile("inputs/day3_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	banks := [][]int{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		bankStr := string(scanner.Bytes())

		values := []int{}

		for _, r := range bankStr {
			values = append(values, int(r-'0'))
		}

		banks = append(banks, values)
	}

	ans1 := 0
	ans2 := 0

	for _, bank := range banks {

		batteryPair := []int{}
		batteryArray := []int{}

		breakPointIndex := -1

		for compare := 9; compare >= 0; compare-- {
			for index, value := range bank {
				if value == compare {
					if len(batteryPair) < 2 && index+1-len(batteryPair) < len(bank) && index > breakPointIndex {
						breakPointIndex = index
						batteryPair = append(batteryPair, value)
						compare = 10
						break
					}
				}
			}
		}

		breakPointIndex = -1
		for compare := 9; compare >= 0; compare-- {
			for index, value := range bank {
				if value == compare {
					if len(batteryArray) < 12 && index+11-len(batteryArray) < len(bank) && index > breakPointIndex {
						breakPointIndex = index
						batteryArray = append(batteryArray, value)
						compare = 10
						break
					}
				}
			}
		}

		slices.Reverse(batteryPair)
		slices.Reverse(batteryArray)

		for i, value := range batteryPair {
			ans1 += value * int(math.Pow(10.0, float64(i)))
		}
		for i, value := range batteryArray {
			ans2 += value * int(math.Pow(10.0, float64(i)))
		}
	}

	return ans1, ans2
}
