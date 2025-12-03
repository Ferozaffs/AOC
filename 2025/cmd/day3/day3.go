package day3

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

		batteryPair := map[int]int{}
		batteryArray := map[int]int{}

		for compare := 9; compare >= 0; compare-- {

			lowestPairIndex := -1
			seekValue := 10
			for idx, value := range batteryPair {
				if lowestPairIndex == -1 {
					lowestPairIndex = idx
				} else if idx < lowestPairIndex || value < seekValue {
					lowestPairIndex = idx
					seekValue = value
				}
			}

			lowestArrayIndex := -1
			seekValue = 10
			for idx, value := range batteryArray {
				if lowestArrayIndex == -1 {
					lowestArrayIndex = idx
				} else if idx < lowestArrayIndex || value < seekValue {
					lowestArrayIndex = idx
					seekValue = value
				}
			}

			for index := len(bank) - 1; index >= 0; index-- {
				value := bank[index]
				if value == compare {
					if len(batteryPair) < 2 && (lowestPairIndex > len(bank)-3-len(batteryPair) || index > lowestPairIndex) {
						_, ok := batteryPair[index]
						if ok == false {
							batteryPair[index] = value
						}
					}
					if len(batteryArray) < 12 && (lowestArrayIndex > len(bank)-13-len(batteryArray) || index > lowestArrayIndex) {
						_, ok := batteryArray[index]
						if ok == false {
							batteryArray[index] = value
						}
					}
				}
			}
		}

		index0 := -1
		index1 := -1
		for idx := range batteryPair {
			if index0 != -1 {
				index1 = idx
			} else {
				index0 = idx
			}
		}

		if index0 < index1 {
			ans1 += batteryPair[index0]*10 + batteryPair[index1]
		} else {
			ans1 += batteryPair[index1]*10 + batteryPair[index0]
		}

		index := make([]int, 0, len(batteryArray))
		for idx := range batteryArray {
			index = append(index, idx)
		}

		sort.Ints(index)

		value := 0
		count := 1
		for i := len(index) - 1; i >= 0; i-- {
			idx := index[i]

			value += batteryArray[idx] * count

			count *= 10
		}

		ans2 += value
	}

	return ans1, ans2
}
