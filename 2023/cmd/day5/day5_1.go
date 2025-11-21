package day5

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type MapData struct {
	dst int
	src int
	rng int
}

var day5_1Cmd = &cobra.Command{
	Use:   "day5_1",
	Short: "Day 5, challenge 1",
	Long:  `Day 5, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day5_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day5_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	entries := []int{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Scan()
	s := strings.Split(scanner.Text(), ":")
	for _, m := range regexp.MustCompile(`\d+`).FindAllString(s[1], -1) {
		n, _ := strconv.Atoi(m)
		entries = append(entries, n)
	}

	//Prime scanner
	scanner.Scan()
	scanner.Scan()

	Recurse(&entries, scanner)

	min := entries[0]
	for _, n := range entries {
		if n < min {
			min = n
		}
	}

	return min
}

func Recurse(entries *[]int, scanner *bufio.Scanner) {
	mapDatas := []MapData{}

	fmt.Print("Recurse")

	for scanner.Scan() {
		if len(scanner.Text()) <= 1 {

			MapEntries(entries, &mapDatas)

			//Prime scanner
			scanner.Scan()

			Recurse(entries, scanner)
		} else {
			m := regexp.MustCompile(`\d+`).FindAllString(scanner.Text(), -1)
			d, _ := strconv.Atoi(m[0])
			s, _ := strconv.Atoi(m[1])
			r, _ := strconv.Atoi(m[2])
			mapDatas = append(mapDatas, MapData{dst: d, src: s, rng: r})
		}
	}
}

func MapEntries(entries *[]int, mapDatas *[]MapData) {
	for i, e := range *entries {
		for _, md := range *mapDatas {
			if e >= md.src && e < md.src+md.rng {
				(*entries)[i] = md.dst + (e - md.src)
			}
		}
	}
}
