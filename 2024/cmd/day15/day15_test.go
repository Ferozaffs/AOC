package day15

import (
	"testing"
)

func TestSample(t *testing.T) {
	sampleData := `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

	ans1, _ := Solve(sampleData)
	if ans1 != 2028 {
		t.Fatalf("Mismatch! Expected 2028 got %d", ans1)
	}
	//if ans2 != 1206 {
	//	t.Fatalf("Mismatch! Expected 1206 got %d", ans2)
	//}
}

func TestSample2(t *testing.T) {
	sampleData := `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`

	ans1, ans2 := Solve(sampleData)
	if ans1 != 10092 {
		t.Fatalf("Mismatch! Expected 10092 got %d", ans1)
	}
	if ans2 != 9021 {
		t.Fatalf("Mismatch! Expected 9021 got %d", ans2)
	}
}

func TestSample3(t *testing.T) {
	sampleData := `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

	ans1, _ := Solve(sampleData)
	if ans1 != -1 {
		t.Fatal("Should fail")
	}
	//if ans2 != 1206 {
	//	t.Fatalf("Mismatch! Expected 1206 got %d", ans2)
	//}
}
