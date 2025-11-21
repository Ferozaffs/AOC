package day9

import (
	"aoc2024/cmd"
	"fmt"
	"math"
	"os"

	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day9",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day1_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day9_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	extractedValues := [][]int{}
	extractedFree := [][]int{}
	index := 0
	for i, r := range data {
		v := int(r - '0')

		var values []int
		for j := 0; j < v; j++ {
			if i%2 == 0 {
				values = append(values, index)
			} else {
				values = append(values, -1)
			}
		}

		if i%2 == 0 {
			extractedValues = append(extractedValues, values)
		} else {
			extractedFree = append(extractedFree, values)
			index++
		}
	}

	combined := MoveAndCombine(DeepCopy(extractedValues), DeepCopy(extractedFree), false)

	ans1 := 0
	counter := 0
out2:
	for _, s := range combined {
		for _, v := range s {
			if v == -1 {
				break out2
			}

			ans1 += counter * v
			counter++
		}
	}

	combined = MoveAndCombine(DeepCopy(extractedValues), DeepCopy(extractedFree), true)

	ans2 := 0
	counter = 0
	for _, s := range combined {
		for _, v := range s {
			if v == -1 {
				counter++
				continue
			}

			ans2 += counter * v
			counter++
		}
	}

	return ans1, ans2
}

func MoveAndCombine(extractedValues [][]int, extractedFree [][]int, smart bool) [][]int {
	var combined [][]int

	if !smart {
		indexSpace := 0
		indexValue := 0
	out:
		for i := len(extractedValues) - 1; i >= 0; i-- {
			for j := len(extractedValues[i]) - 1; j >= 0; j-- {
				if indexSpace >= i {
					break out
				}

				extractedFree[indexSpace][indexValue] = extractedValues[i][j]
				indexValue++
				if indexValue == len(extractedFree[indexSpace]) {
					for {
						indexSpace++
						indexValue = 0
						if indexSpace == len(extractedFree) {
							break out
						}
						if extractedFree[indexSpace] != nil && len(extractedFree[indexSpace]) != 0 {
							break
						}
					}

				}

				extractedValues[i] = append(extractedValues[i][:j], extractedValues[i][j+1:]...)
			}
			if len(extractedValues[i]) == 0 {
				extractedValues = append(extractedValues[:i], extractedValues[i+1:]...)
			}
		}
	} else {
		for i := len(extractedValues) - 1; i >= 0; i-- {
			for j, f := range extractedFree {
				if j < i {
					count := 0
					offset := 0
					for _, v := range f {
						if v == -1 {
							count++
						} else {
							offset++
						}
					}

					if count >= len(extractedValues[i]) {
						numToAppend := 0
						for k := 0; k < len(extractedValues[i]); {
							f[offset] = extractedValues[i][k]
							offset++
							numToAppend++
							extractedValues[i] = append(extractedValues[i][:k], extractedValues[i][k+1:]...)
						}

						for a := 0; a < numToAppend; a++ {
							if i-1 < len(extractedFree) {
								extractedFree[i-1] = append(extractedFree[i-1], -1)
							}
						}
					}
				}
			}
		}
	}

	maxLen := int(math.Max(float64(len(extractedValues)), float64(len(extractedFree))))

	for i := 0; i < maxLen; i++ {
		if i < len(extractedValues) {
			combined = append(combined, extractedValues[i])
		}
		if i < len(extractedFree) && extractedFree[i] != nil {
			combined = append(combined, extractedFree[i])
		}
	}

	return combined
}

func DeepCopy(original [][]int) [][]int {
	result := make([][]int, len(original))

	for i, inner := range original {
		result[i] = make([]int, len(inner))
		copy(result[i], inner)
	}

	return result
}
