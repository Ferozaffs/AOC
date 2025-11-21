package day19

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 6 {
		t.Fatalf("Mismatch! Expected 6 got %d", ans1)
	}
	if ans2 != 16 {
		t.Fatalf("Mismatch! Expected 16 got %d", ans2)
	}
}
