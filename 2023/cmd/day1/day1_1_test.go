package day1

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	ans := Solve1(sampleData)
	if ans != 142 {
		t.Fatalf("Mismatch! Expected 142 got %d", ans)
	}
}
