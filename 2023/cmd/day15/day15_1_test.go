package day15

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	ans := Solve1(sampleData)
	if ans != 1320 {
		t.Fatalf("Mismatch! Expected 1320 got %d", ans)
	}
}
