package day11

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `125 17`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 55312 {
		t.Fatalf("Mismatch! Expected 55312 got %d", ans1)
	}
	if ans2 != 65601038650482 {
		t.Fatalf("Mismatch! Expected 65601038650482 got %d", ans2)
	}
}
