package day5

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 3 {
		t.Fatalf("Mismatch! Expected 3 got %d", ans1)
	}
	if ans2 != 14 {
		t.Fatalf("Mismatch! Expected 14 got %d", ans2)
	}
}
