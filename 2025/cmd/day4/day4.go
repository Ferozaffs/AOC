package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"2025/cmd"

	"github.com/spf13/cobra"
)

type Coord struct {
	x int
	y int
}

var DayCmd = &cobra.Command{
	Use:   "day4",
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
	dat, _ := os.ReadFile("inputs/day4_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	rows := [][]rune{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		rows = append(rows, []rune(string(scanner.Bytes())))
	}

	ans1 := 0
	ans2 := 0

	found := true
	firstPass := true
	for found {
		found = false

		foundRollCoords := []Coord{}
		for currentRow, row := range rows {
			for currentColumn, roll := range row {
				if roll == '@' {
					adjacent := 0

					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							if x != 0 || y != 0 {
								deltaRow := currentRow + x
								deltaColumn := currentColumn + y

								if deltaRow >= 0 && deltaColumn >= 0 && deltaRow < len(rows) && deltaColumn < len(row) {
									if rows[deltaRow][deltaColumn] == '@' {
										adjacent++
									}
								}
							}
						}
					}

					if adjacent < 4 {
						if firstPass {
							ans1++
						}
						ans2++

						foundRollCoords = append(foundRollCoords, Coord{x: currentRow, y: currentColumn})
						found = true
					}
				}
			}
		}
		for _, coord := range foundRollCoords {
			rows[coord.x][coord.y] = '.'
		}

		firstPass = false
	}

	return ans1, ans2
}
