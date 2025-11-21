package day8

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `RL

	AAA = (BBB, CCC)
	BBB = (DDD, EEE)
	CCC = (ZZZ, GGG)
	DDD = (DDD, DDD)
	EEE = (EEE, EEE)
	GGG = (GGG, GGG)
	ZZZ = (ZZZ, ZZZ)`

	ans := Solve1(sampleData)
	if ans != 2 {
		t.Fatalf("Mismatch! Expected 2 got %d", ans)
	}

	sampleData = `LLR

	AAA = (BBB, BBB)
	BBB = (AAA, ZZZ)
	ZZZ = (ZZZ, ZZZ)`

	ans = Solve1(sampleData)
	if ans != 6 {
		t.Fatalf("Mismatch! Expected 6 got %d", ans)
	}
}
