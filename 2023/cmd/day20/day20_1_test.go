package day20

import (
	"testing"
)

func TestSample1(t *testing.T) {
	sampleData := `broadcaster -> a, b, c
	%a -> b
	%b -> c
	%c -> inv
	&inv -> a`

	ans := Solve1(sampleData)
	if ans != 32000000 {
		t.Fatalf("Mismatch! Expected 32000000 got %d", ans)
	}

	sampleData = `broadcaster -> a
	%a -> inv, con
	&inv -> b
	%b -> con
	&con -> output`

	ans = Solve1(sampleData)
	if ans != 11687500 {
		t.Fatalf("Mismatch! Expected 11687500 got %d", ans)
	}
}
