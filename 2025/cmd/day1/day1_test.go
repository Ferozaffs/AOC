package day1

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 3 {
		t.Fatalf("Mismatch! Expected 3 got %d", ans1)
	}
	if ans2 != 6 {
		t.Fatalf("Mismatch! Expected 6 got %d", ans2)
	}
}
