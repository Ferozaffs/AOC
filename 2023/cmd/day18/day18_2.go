package day18

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day18_2Cmd = &cobra.Command{
	Use:   "day18_2",
	Short: "Day 18, challenge 2",
	Long:  `Day 18, challenge 2 | Shoelace formula`,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day18_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day18_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	data = strings.ReplaceAll(data, "(", "")
	data = strings.ReplaceAll(data, ")", "")
	data = strings.ReplaceAll(data, "#", "")
	data = strings.ReplaceAll(data, "\t", "")
	rows := strings.Split(data, "\n")

	c := Coordinate{x: 0, y: 0}
	vertices := []Coordinate{c}

	area := 0
	for _, row := range rows {
		s := strings.Fields(row)
		hex := s[2]
		dir := hex[len(hex)-1]
		hex = hex[:len(hex)-1]

		hexVal, _ := strconv.ParseInt(hex, 16, 64)

		area += int(hexVal)

		if dir == '0' {
			c.x += int(hexVal)
		} else if dir == '2' {
			c.x -= int(hexVal)
		} else if dir == '1' {
			c.y -= int(hexVal)
		} else if dir == '3' {
			c.y += int(hexVal)
		}

		vertices = append(vertices, c)
	}

	for i := len(vertices) - 2; i >= 0; i-- {
		area += vertices[i+1].x*vertices[i].y - vertices[i+1].y*vertices[i].x
	}

	area /= 2
	area += 1

	return area
}
