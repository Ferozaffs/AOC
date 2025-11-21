package day20

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

	SetCutoff(60)
	ans1, _ := Solve(sampleData)
	if ans1 != 1 {
		t.Fatalf("Mismatch! Expected 2 got %d", ans1)
	}
	//if ans2 != "6,1" {
	//	t.Fatalf("Mismatch! Expected 6,1 got %s", ans2)
	//}
}
