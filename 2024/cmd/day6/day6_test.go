package day6

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 41 {
		t.Fatalf("Mismatch! Expected 41 got %d", ans1)
	}
	if ans2 != 6 {
		t.Fatalf("Mismatch! Expected 6 got %d", ans2)
	}
}
