package day5

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type DataRange struct {
	s int
	r int
}

var day5_2Cmd = &cobra.Command{
	Use:   "day5_2",
	Short: "Day 5, challenge 2",
	Long:  `Day 5, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day5_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day5_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	entries := []DataRange{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Scan()
	s := strings.Split(scanner.Text(), ":")
	for _, m := range regexp.MustCompile(`\d+ \d+`).FindAllString(s[1], -1) {
		s = strings.Split(m, " ")
		n, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])
		entries = append(entries, DataRange{s: n, r: r})
	}

	//Prime scanner
	scanner.Scan()
	scanner.Scan()

	mapDatas := [][]MapData{}
	mapData := []MapData{}
	for scanner.Scan() {
		if len(scanner.Text()) <= 1 {
			mapDatas = append(mapDatas, mapData)
			mapData = []MapData{}

			//Prime scanner
			scanner.Scan()
		} else {
			m := regexp.MustCompile(`\d+`).FindAllString(scanner.Text(), -1)
			d, _ := strconv.Atoi(m[0])
			s, _ := strconv.Atoi(m[1])
			r, _ := strconv.Atoi(m[2])
			mapData = append(mapData, MapData{dst: d, src: s, rng: r})
		}
	}

	//Each seed one at a time due to size, poor poor CPU
	min := math.MaxInt
	for _, n := range entries {
		for s := n.s; s < n.s+n.r; s++ {
			f := Recurse2(s, 0, mapDatas)
			if f < min {
				min = f
			}
		}

	}

	return min
}

func Recurse2(seed int, mapIndex int, mapDatas [][]MapData) int {
	if mapIndex == len(mapDatas) {
		return seed
	}

	seed = MapEntry(seed, &mapDatas[mapIndex])
	return Recurse2(seed, mapIndex+1, mapDatas)
}

func MapEntry(seed int, mapDatas *[]MapData) int {
	for _, md := range *mapDatas {
		if seed >= md.src && seed < md.src+md.rng {
			return md.dst + (seed - md.src)
		}
	}

	return seed
}
