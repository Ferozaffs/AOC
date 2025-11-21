package day10

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 36 {
		t.Fatalf("Mismatch! Expected 36 got %d", ans1)
	}
	if ans2 != 81 {
		t.Fatalf("Mismatch! Expected 81 got %d", ans2)
	}
}
