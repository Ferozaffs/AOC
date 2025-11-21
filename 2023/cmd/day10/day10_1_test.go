package day10

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `.....
	.S-7.
	.|.|.
	.L-J.
	.....`

	ans := Solve1(sampleData)
	if ans != 4 {
		t.Fatalf("Mismatch! Expected 4 got %d", ans)
	}

	sampleData = `..F7.
	.FJ|.
	SJ.L7
	|F--J
	LJ...`

	ans = Solve1(sampleData)
	if ans != 8 {
		t.Fatalf("Mismatch! Expected 8 got %d", ans)
	}
}
