package day10

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `.F----7F7F7F7F-7....
	.|F--7||||||||FJ....
	.||.FJ||||||||L7....
	FJL7L7LJLJ||LJ.L-7..
	L--J.L7...LJS7F-7L7.
	....F-J..F7FJ|L7L7L7
	....L7.F7||L7|.L7L7|
	.....|FJLJ|FJ|F7|.LJ
	....FJL-7.||.||||...
	....L---J.LJ.LJLJ...`

	ans := Solve2(sampleData)
	if ans != 8 {
		t.Fatalf("Mismatch! Expected 8 got %d", ans)
	}

	sampleData = `FF7FSF7F7F7F7F7F---7
	L|LJ||||||||||||F--J
	FL-7LJLJ||||||LJL-77
	F--JF--7||LJLJ7F7FJ-
	L---JF-JLJ.||-FJLJJ7
	|F|F-JF---7F7-L7L|7|
	|FFJF7L7F-JF7|JL---7
	7-L-JL7||F7|L7F-7F7|
	L.L7LFJ|||||FJL7||LJ
	L7JLJL-JLJLJL--JLJ.L`

	ans = Solve2(sampleData)
	if ans != 10 {
		t.Fatalf("Mismatch! Expected 10 got %d", ans)
	}
}
