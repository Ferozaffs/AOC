package day12

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 1930 {
		t.Fatalf("Mismatch! Expected 1930 got %d", ans1)
	}
	if ans2 != 1206 {
		t.Fatalf("Mismatch! Expected 1206 got %d", ans2)
	}
}

func TestSampleSimple(t *testing.T) {
	sampleData := `AAAA
BBCD
BBCC
EEEC`

	_, ans2 := Solve(sampleData)
	if ans2 != 80 {
		t.Fatalf("Mismatch! Expected 80 got %d", ans2)
	}
}

func TestSampleEasy(t *testing.T) {
	sampleData := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

	_, ans2 := Solve(sampleData)
	if ans2 != 236 {
		t.Fatalf("Mismatch! Expected 236 got %d", ans2)
	}
}

func TestSampleMedium(t *testing.T) {
	sampleData := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

	_, ans2 := Solve(sampleData)
	if ans2 != 368 {
		t.Fatalf("Mismatch! Expected 368 got %d", ans2)
	}
}

func TestSampleOs(t *testing.T) {
	sampleData := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

	_, ans2 := Solve(sampleData)
	if ans2 != 436 {
		t.Fatalf("Mismatch! Expected 436 got %d", ans2)
	}
}
