package day13

import (
	"testing"
)

func TestSample1(t *testing.T) {
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

	ans := Solve1(sampleData)
	if ans != 405 {
		t.Fatalf("Mismatch! Expected 405 got %d", ans)
	}
}
