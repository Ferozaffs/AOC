package day6

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `Time:      7  15   30
	Distance:  9  40  200`

	ans := Solve2(sampleData)
	if ans != 71503 {
		t.Fatalf("Mismatch! Expected 71503 got %d", ans)
	}
}
