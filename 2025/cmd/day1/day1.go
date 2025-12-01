package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"2025/cmd"

	"github.com/spf13/cobra"
)

var DayCmd = &cobra.Command{
	Use:   "day1",
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
	dat, _ := os.ReadFile("inputs/day1_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	instructions := []int{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		instruction := string(scanner.Bytes())
		direction, size := utf8.DecodeRuneInString(instruction)
		steps, _ := strconv.Atoi(instruction[size:])

		if direction == 'R' {
			instructions = append(instructions, steps)
		} else {
			instructions = append(instructions, -steps)
		}
	}

	value := 1000000 + 50

	ans1 := 0
	ans2 := 0

	current := value % 100
	for _, instruction := range instructions {

		value += instruction

		newCurrent := current + instruction
		ans2 += int(math.Abs(float64(newCurrent)) / 100.0)
		if newCurrent <= 0 && current != 0 {
			ans2 += 1
		}

		current = value % 100

		if current == 0 {
			ans1 += 1
		}
	}

	return ans1, ans2
}
