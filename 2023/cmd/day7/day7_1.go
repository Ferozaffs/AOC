package day7

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Hand struct {
	powerType int
	powerHand int
	bid       int
	rank      int
}

var day7_1Cmd = &cobra.Command{
	Use:   "day7_1",
	Short: "Day 7, challenge 1",
	Long:  `Day 7, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day7_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day7_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0

	hands := []Hand{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(s[1])

		hands = append(hands, Hand{powerType: GetTypePower(s[0]), powerHand: GetHandPower(s[0]), bid: n, rank: 1})
	}

	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if hands[i].powerType > hands[j].powerType || (hands[i].powerType == hands[j].powerType && hands[i].powerHand > hands[j].powerHand) {
				hands[i].rank++
			} else {
				hands[j].rank++
			}
		}

		sum += hands[i].rank * hands[i].bid
	}

	return sum
}

func GetTypePower(hand string) int {
	s := strings.Split(hand, "")
	sort.Strings(s)
	sortedHand := strings.Join(s, "")

	subTypes := []int{}
	prevRune := '-'
	count := 0
	for _, c := range sortedHand {
		if prevRune != c {
			if count > 1 {
				subTypes = append(subTypes, count)
			}
			count = 0
		}
		prevRune = c
		count++
	}
	if count > 1 {
		subTypes = append(subTypes, count)
	}

	power := 0
	for _, n := range subTypes {
		power += int(math.Pow(float64(n), float64(n)))
	}

	return power
}

func GetHandPower(hand string) int {
	powerRank := []rune{
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
		'T',
		'J',
		'Q',
		'K',
		'A',
	}

	power := 0
	for pos, char := range hand {
		for i, c := range powerRank {
			if c == char {
				power += i * int(math.Pow(100, float64(5-pos)))
				break
			}
		}
	}

	return power
}
