package day9

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `0 3 6 9 12 15
	1 3 6 10 15 21
	10 13 16 21 30 45`

	ans := Solve(sampleData, false)
	if ans != 114 {
		t.Fatalf("Mismatch! Expected 144 got %d", ans)
	}

}
