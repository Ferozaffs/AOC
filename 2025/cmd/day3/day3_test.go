package day3

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `987654321111111
811111111111119
234234234234278
818181911112111`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 357 {
		t.Fatalf("Mismatch! Expected 357 got %d", ans1)
	}
	if ans2 != 3121910778619 {
		t.Fatalf("Mismatch! Expected 3121910778619 got %d", ans2)
	}
}
