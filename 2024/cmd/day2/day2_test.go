package day2

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
9 8 9 7 6`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 2 {
		t.Fatalf("Mismatch! Expected 2 got %d", ans1)
	}
	if ans2 != 5 {
		t.Fatalf("Mismatch! Expected 5 got %d", ans2)
	}
}
