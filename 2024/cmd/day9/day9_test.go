package day9

import (
	"testing"
)

func TestSampleSimple(t *testing.T) {
	sampleData := `12345`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 60 {
		t.Fatalf("Mismatch! Expected 60 got %d", ans1)
	}
	if ans2 != 132 {
		t.Fatalf("Mismatch! Expected 132 got %d", ans2)
	}

}

func TestSampleExtra(t *testing.T) {
	sampleData := `354631466260`

	_, ans2 := Solve(sampleData)
	if ans2 != 1325 {
		t.Fatalf("Mismatch! Expected 1325 got %d", ans2)
	}

}

func TestSample(t *testing.T) {
	sampleData := `2333133121414131402`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 1928 {
		t.Fatalf("Mismatch! Expected 1928 got %d", ans1)
	}
	if ans2 != 2858 {
		t.Fatalf("Mismatch! Expected 2858 got %d", ans2)
	}
}
