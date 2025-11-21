package day5

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 143 {
		t.Fatalf("Mismatch! Expected 143 got %d", ans1)
	}
	if ans2 != 123 {
		t.Fatalf("Mismatch! Expected 123 got %d", ans2)
	}
}
