package day1

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `3   3
2   4
1   3
4   9
3   5
3   3`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 11 {
		t.Fatalf("Mismatch! Expected 11 got %d", ans1)
	}
	if ans2 != 31 {
		t.Fatalf("Mismatch! Expected 31 got %d", ans2)
	}
}
