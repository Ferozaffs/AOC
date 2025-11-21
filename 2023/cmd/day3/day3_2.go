package day3

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/spf13/cobra"
)

type Gear struct {
	gearRatio       int
	rowConnection   int
	indexConnection int
}

var day3_2Cmd = &cobra.Command{
	Use:   "day3_2",
	Short: "Day 3, challenge 2",
	Long:  `Day 3, challenge 2`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day3_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day3_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sum := 0

	re := regexp2.MustCompile("\\d+", 0)
	rows := []string{}
	gears := []Gear{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\t", "")
		rows = append(rows, line)
	}

	for i, s := range rows {
		m, _ := re.FindStringMatch(s)
		for m != nil {
			gear := FindPotentialGear(rows, i, m.Index, m.Length)
			if gear != nil {
				(*gear).gearRatio, _ = strconv.Atoi(m.String())
				gears = append(gears, *gear)
			}

			m, _ = re.FindNextMatch(m)
		}
	}

	for i := 0; i < len(gears); i++ {
		for j := i + 1; j < len(gears); j++ {
			g1 := gears[i]
			g2 := gears[j]
			if g1.indexConnection == g2.indexConnection && g1.rowConnection == g2.rowConnection {
				sum += g1.gearRatio * g2.gearRatio
			}
		}
	}

	return sum
}

func FindPotentialGear(rows []string, rowIndex int, numIndex int, length int) *Gear {
	re := regexp2.MustCompile("\\*", 0)
	check := func(row string, start int, end int) int {
		cut := row[start:end]
		m, _ := re.FindStringMatch(cut)
		if m != nil {
			return start + m.Index
		}

		return -1
	}

	startIndex := int(math.Max(0, float64(numIndex)-1))
	endIndex := int(math.Min(float64(len(rows[rowIndex])), float64(numIndex+length)+1))
	startRow := int(math.Max(0, float64(rowIndex)-1))
	endRow := int(math.Min(float64(len(rows)-1), float64(rowIndex)+1))

	for i := startRow; i <= endRow; i++ {
		c := check(rows[i], startIndex, endIndex)
		if c != -1 {
			g := &Gear{
				gearRatio:       0,
				rowConnection:   i,
				indexConnection: c,
			}
			return g
		}
	}

	return nil
}
