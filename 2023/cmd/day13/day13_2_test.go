package day13

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `#.##..##.
	..#.##.#.
	##......#
	##......#
	..#.##.#.
	..##..##.
	#.#.##.#.
	
	#...##..#
	#....#..#
	..##..###
	#####.##.
	#####.##.
	..##..###
	#....#..#
	`

	ans := Solve2(sampleData)
	if ans != 400 {
		t.Fatalf("Mismatch! Expected 400 got %d", ans)
	}
}
