package day16

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `.|...\....
	|.-.\.....
	.....|-...
	........|.
	..........
	.........\
	..../.\\..
	.-.-/..|..
	.|....-|.\
	..//.|....`

	ans := Solve1(sampleData)
	if ans != 46 {
		t.Fatalf("Mismatch! Expected 46 got %d", ans)
	}
}
