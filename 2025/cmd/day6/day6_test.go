package day6

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	ans1, ans2 := Solve(sampleData)
	if ans1 != 4277556 {
		t.Fatalf("Mismatch! Expected 4277556 got %d", ans1)
	}
	if ans2 != 3263827 {
		t.Fatalf("Mismatch! Expected 3263827 got %d", ans2)
	}
}
