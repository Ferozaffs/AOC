package day7

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `32T3K 765
	T55J5 684
	KK677 28
	KTJJT 220
	QQQJA 483`

	ans := Solve2(sampleData)
	if ans != 5905 {
		t.Fatalf("Mismatch! Expected 5905 got %d", ans)
	}
}
