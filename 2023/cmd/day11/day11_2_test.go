package day11

import (
	"testing"
)

func TestSample2(t *testing.T) {
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

	ans := Solve2(sampleData, 10)
	if ans != 1030 {
		t.Fatalf("Mismatch! Expected 1030 got %d", ans)
	}

	ans = Solve2(sampleData, 100)
	if ans != 8410 {
		t.Fatalf("Mismatch! Expected 8410 got %d", ans)
	}
}
