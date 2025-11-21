package day22

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `1
10
100
2024`

	ans1, _ := Solve(sampleData)
	if ans1 != 37327623 {
		t.Fatalf("Mismatch! Expected 37327623 got %d", ans1)
	}
}

func TestSample2(t *testing.T) {
	sampleData := `1
2
3
2024`

	_, ans2 := Solve(sampleData)
	if ans2 != 23 {
		t.Fatalf("Mismatch! Expected 23 got %d", ans2)
	}
}
