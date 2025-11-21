package day18

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

	SetSize(6, 6, 12)
	ans1, ans2 := Solve(sampleData)
	if ans1 != 22 {
		t.Fatalf("Mismatch! Expected 22 got %d", ans1)
	}
	if ans2 != "6,1" {
		t.Fatalf("Mismatch! Expected 6,1 got %s", ans2)
	}
}
