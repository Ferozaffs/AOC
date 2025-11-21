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

type Coordinate struct {
	x int
	y int
}

var day11_1Cmd = &cobra.Command{
	Use:   "day11_1",
	Short: "Day 11, challenge 1",
	Long:  `Day 11, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day11_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day11_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0

	expandedUniverse := []string{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "\t") {
			s = s[1:]
		}
		galaxyFound := strings.Index(s, "#")

		expandedUniverse = append(expandedUniverse, s)
		if galaxyFound == -1 {
			expandedUniverse = append(expandedUniverse, s)
		}
	}

	for i := 0; i < len(expandedUniverse[0]); i++ {
		foundGalaxy := false
		for j := 0; j < len(expandedUniverse); j++ {
			if expandedUniverse[j][i] == '#' {
				foundGalaxy = true
				break
			}
		}

		if foundGalaxy == false {
			for j := 0; j < len(expandedUniverse); j++ {
				expandedUniverse[j] = expandedUniverse[j][:i] + "." + expandedUniverse[j][i:]
			}
			i++
		}
	}

	galaxies := []Coordinate{}
	for y, row := range expandedUniverse {
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
		}
	}

	return sum
}
