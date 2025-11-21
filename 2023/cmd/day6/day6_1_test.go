package day6

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `Time:      7  15   30
	Distance:  9  40  200`

	ans := Solve1(sampleData)
	if ans != 288 {
		t.Fatalf("Mismatch! Expected 288 got %d", ans)
	}
}
