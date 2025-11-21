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

var day7_2Cmd = &cobra.Command{
	Use:   "day7_2",
	Short: "Day 7, challenge 2",
	Long:  `Day 7, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day7_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day7_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sum := 0

	hands := []Hand{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(s[1])

		hands = append(hands, Hand{powerType: GetTypePowerJoker(s[0]), powerHand: GetHandPowerJoker(s[0]), bid: n, rank: 1})
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

func GetTypePowerJoker(hand string) int {
	s := strings.Split(hand, "")
	sort.Strings(s)
	sortedHand := strings.Join(s, "")

	subTypes := []int{}
	prevRune := '-'
	count := 0
	jokerCount := 0
	for _, c := range sortedHand {
		if prevRune != c {
			if prevRune != '-' {
				subTypes = append(subTypes, count)
			}
			if c == 'J' {
				jokerCount++
				prevRune = '-'
				count = 0
				continue
			}
			count = 0
		}

		prevRune = c
		count++
	}
	if count > 0 {
		subTypes = append(subTypes, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(subTypes)))
	if len(subTypes) == 0 {
		subTypes = append(subTypes, 5)
	} else {
		subTypes[0] += jokerCount
	}

	power := 0
	for _, n := range subTypes {
		power += int(math.Pow(float64(n), float64(n)))
	}

	return power
}

func GetHandPowerJoker(hand string) int {
	powerRank := []rune{
		'J',
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
