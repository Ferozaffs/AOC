package day11

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var day11_2Cmd = &cobra.Command{
	Use:   "day11_2",
	Short: "Day 11, challenge 2",
	Long:  `Day 11, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day11_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day11_data.txt")
	ans := Solve2(string(dat), 1000000)

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string, expansion int) int {
	sum := 0

	universe := []string{}
	expRows := []int{}
	expCols := []int{}
	scanner := bufio.NewScanner(strings.NewReader(data))

	row := 0
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "\t") {
			s = s[1:]
		}
		galaxyFound := strings.Index(s, "#")

		universe = append(universe, s)
		if galaxyFound == -1 {
			expRows = append(expRows, row)
		}

		row++
	}

	for i := 0; i < len(universe[0]); i++ {
		foundGalaxy := false
		for j := 0; j < len(universe); j++ {
			if universe[j][i] == '#' {
				foundGalaxy = true
				break
			}
		}

		if foundGalaxy == false {
			expCols = append(expCols, i)
		}
	}

	galaxies := []Coordinate{}
	for y, row := range universe {
		for x, column := range row {
			if column == '#' {
				galaxies = append(galaxies, Coordinate{x: x, y: y})
			}
		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += int(math.Abs(float64(galaxies[i].x) - float64(galaxies[j].x)))
			sum += int(math.Abs(float64(galaxies[i].y) - float64(galaxies[j].y)))

			for _, r := range expRows {
				if (galaxies[i].y > r && galaxies[j].y < r) || (galaxies[i].y < r && galaxies[j].y > r) {
					sum += expansion - 1
				}
			}
			for _, c := range expCols {
				if (galaxies[i].x > c && galaxies[j].x < c) || (galaxies[i].x < c && galaxies[j].x > c) {
					sum += expansion - 1
				}
			}
		}
	}

	return sum
}
