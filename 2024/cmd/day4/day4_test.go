package day4

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 18 {
		t.Fatalf("Mismatch! Expected 18 got %d", ans1)
	}
	if ans2 != 9 {
		t.Fatalf("Mismatch! Expected 9 got %d", ans2)
	}
}
