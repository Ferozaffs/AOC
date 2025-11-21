package day11

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `...#......
	.......#..
	#.........
	..........
	......#...
	.#........
	.........#
	..........
	.......#..
	#...#.....`

	ans := Solve1(sampleData)
	if ans != 374 {
		t.Fatalf("Mismatch! Expected 374 got %d", ans)
	}
}
