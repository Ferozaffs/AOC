package day13

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

	ans1, _ := Solve(sampleData)
	if ans1 != 480 {
		t.Fatalf("Mismatch! Expected 480 got %d", ans1)
	}
	//if ans2 != 1206 {
	//	t.Fatalf("Mismatch! Expected 1206 got %d", ans2)
	//}
}
