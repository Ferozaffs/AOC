package day1

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	ans := Solve2(sampleData)
	if ans != 281 {
		t.Fatalf("Mismatch! Expected 281 got %d", ans)
	}
}
