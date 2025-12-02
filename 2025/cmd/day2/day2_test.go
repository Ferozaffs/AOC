package day2

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 1227775554 {
		t.Fatalf("Mismatch! Expected 1227775554 got %d", ans1)
	}
	if ans2 != 4174379265 {
		t.Fatalf("Mismatch! Expected 4174379265 got %d", ans2)
	}
}
