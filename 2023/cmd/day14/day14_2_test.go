package day14

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `O....#....
	O.OO#....#
	.....##...
	OO.#O....O
	.O.....O#.
	O.#..O.#.#
	..O..#O..O
	.......O..
	#....###..
	#OO..#....`

	ans := Solve2(sampleData)
	if ans != 64 {
		t.Fatalf("Mismatch! Expected 64 got %d", ans)
	}
}
