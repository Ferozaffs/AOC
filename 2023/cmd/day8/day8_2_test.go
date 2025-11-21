package day8

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `LR

	11A = (11B, XXX)
	11B = (XXX, 11Z)
	11Z = (11B, XXX)
	22A = (22B, XXX)
	22B = (22C, 22C)
	22C = (22Z, 22Z)
	22Z = (22B, 22B)
	XXX = (XXX, XXX)`

	ans := Solve2(sampleData)
	if ans != 6 {
		t.Fatalf("Mismatch! Expected 6 got %d", ans)
	}
}
