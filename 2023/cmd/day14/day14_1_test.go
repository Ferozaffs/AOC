package day14

import (
	"testing"
)

func TestSample1(t *testing.T) {
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

	ans := Solve1(sampleData)
	if ans != 136 {
		t.Fatalf("Mismatch! Expected 136 got %d", ans)
	}
}
