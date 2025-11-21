package day14

import (
	"testing"
)

func TestSingle(t *testing.T) {
	sampleData := `p=2,4 v=2,-3`

	SetSize(11, 7)
	ans1, _ := Solve(sampleData)
	if ans1 != 1 {
		t.Fatalf("Mismatch! Expected 1 got %d", ans1)
	}
}

func TestSample(t *testing.T) {
	sampleData := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	SetSize(11, 7)
	ans1, _ := Solve(sampleData)
	if ans1 != 12 {
		t.Fatalf("Mismatch! Expected 12 got %d", ans1)
	}
	//if ans2 != 1206 {
	//	t.Fatalf("Mismatch! Expected 1206 got %d", ans2)
	//}
}
