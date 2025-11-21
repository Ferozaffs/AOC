package day21

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `029A
980A
179A
456A
379A`

	ans1, _ := Solve(sampleData)
	if ans1 != 126384 {
		t.Fatalf("Mismatch! Expected 126384 got %d", ans1)
	}
	//if ans2 != 16 {
	//	t.Fatalf("Mismatch! Expected 16 got %d", ans2)
	//}
}
