package day21

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `...........
	.....###.#.
	.###.##..#.
	..#.#...#..
	....#.#....
	.##..S####.
	.##..#...#.
	.......##..
	.##.#.####.
	.##..##.##.
	...........`

	ans := Solve1(sampleData, 6, false)
	if ans != 16 {
		t.Fatalf("Mismatch! Expected 16 got %d", ans)
	}
}
