package day3

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	ans := Solve1(sampleData)
	if ans != 4361 {
		t.Fatalf("Mismatch! Expected 4361 got %d", ans)
	}
}
