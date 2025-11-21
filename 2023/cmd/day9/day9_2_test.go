package day9

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `0 3 6 9 12 15
	1 3 6 10 15 21
	10 13 16 21 30 45`

	ans := Solve(sampleData, true)
	if ans != 2 {
		t.Fatalf("Mismatch! Expected 2 got %d", ans)
	}
}
