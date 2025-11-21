package day3

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	ans1, _ := Solve(sampleData)
	if ans1 != 161 {
		t.Fatalf("Mismatch! Expected 161 got %d", ans1)
	}

	sampleData2 := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	_, ans2 := Solve(sampleData2)
	if ans2 != 48 {
		t.Fatalf("Mismatch! Expected 48 got %d", ans2)
	}
}
