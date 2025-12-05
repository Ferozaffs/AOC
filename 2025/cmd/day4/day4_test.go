package day4

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 13 {
		t.Fatalf("Mismatch! Expected 13 got %d", ans1)
	}
	if ans2 != 43 {
		t.Fatalf("Mismatch! Expected 43 got %d", ans2)
	}
}
