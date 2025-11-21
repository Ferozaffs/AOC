package day3

import (
	"testing"
)

func TestSample2(t *testing.T) {
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

	ans := Solve2(sampleData)
	if ans != 467835 {
		t.Fatalf("Mismatch! Expected 467835 got %d", ans)
	}
}
