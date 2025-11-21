package day7

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `32T3K 765
	T55J5 684
	KK677 28
	KTJJT 220
	QQQJA 483`

	ans := Solve1(sampleData)
	if ans != 6440 {
		t.Fatalf("Mismatch! Expected 6440 got %d", ans)
	}
}
