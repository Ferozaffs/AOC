package day3

import (
	"aoc2024/cmd"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/dlclark/regexp2"
	"github.com/spf13/cobra"
)

var day1_1Cmd = &cobra.Command{
	Use:   "day3",
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
	dat, _ := os.ReadFile("inputs/day3_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	extractMul := `mul\(\d{1,3},\d{1,3}\)`
	extractMulAdv := `((^|do\(\))(?:[^d]|d(?!on't\())*?(?=(don't\(\)|$)))`
	extractDigits := `\d{1,3}`
	reMul := regexp.MustCompile(extractMul)
	reMulAdv := regexp2.MustCompile(extractMulAdv, regexp2.None)
	reDigit := regexp.MustCompile(extractDigits)

	ans1 := 0
	ans2 := 0
	matches := reMul.FindAllString(data, -1)
	for _, m := range matches {
		digits := reDigit.FindAllString(m, -1)
		a, _ := strconv.Atoi(digits[0])
		b, _ := strconv.Atoi(digits[1])
		ans1 += a * b
	}

	for am, _ := reMulAdv.FindStringMatch(data); am != nil; am, _ = reMulAdv.FindNextMatch(am) {
		matches = reMul.FindAllString(am.String(), -1)
		for _, m := range matches {
			digits := reDigit.FindAllString(m, -1)
			a, _ := strconv.Atoi(digits[0])
			b, _ := strconv.Atoi(digits[1])
			ans2 += a * b
		}
	}

	return ans1, ans2
}
