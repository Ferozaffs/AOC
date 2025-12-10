package day8

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"2025/cmd"

	"github.com/spf13/cobra"
)

type Vec3 struct {
	x, y, z float64
}

type Junction struct {
	id      int
	point   Vec3
	circuit int
}

type JunctionPair struct {
	a        *Junction
	b        *Junction
	distance float64
}

var DayCmd = &cobra.Command{
	Use:   "day8",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(DayCmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day8_data.txt")
	ans1, ans2 := Solve(string(dat), 1000)

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string, connectionsCap int) (int, int) {
	var junctions []*Junction
	id := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		row := string(scanner.Bytes())
		values := strings.Split(row, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])
		junctions = append(junctions, &Junction{id, Vec3{float64(x), float64(y), float64(z)}, -1})
		id++
	}

	var pairs []JunctionPair
	for i := 0; i < len(junctions); i++ {
		for j := i + 1; j < len(junctions); j++ {
			pairs = append(pairs, JunctionPair{junctions[i], junctions[j], Distance(junctions[i].point, junctions[j].point)})
		}
	}

	pairs = DedupPairs(pairs)

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	currentCircuitId := 0
	connections := 0

	ans1 := 1
	ans2 := 0.0

	for _, p := range pairs {
		if connections == connectionsCap {
			// Solver for ans1
			circuits := map[int]int{}
			for _, j := range junctions {
				if j.circuit == -1 {
					continue
				}
				circuits[j.circuit]++
			}

			var sizes []int
			for _, v := range circuits {
				sizes = append(sizes, v)
			}

			sort.Slice(sizes, func(i, j int) bool {
				return sizes[i] > sizes[j]
			})

			for _, v := range sizes[:3] {
				ans1 *= v
			}
		}

		if p.a.circuit != -1 {
			if p.b.circuit == -1 {
				p.b.circuit = p.a.circuit
			} else {
				circtuiToChange := p.b.circuit
				for _, jn := range junctions {
					if jn.circuit == circtuiToChange {
						jn.circuit = p.a.circuit
					}
				}
			}
		} else if p.b.circuit != -1 {
			if p.a.circuit == -1 {
				p.a.circuit = p.b.circuit
			}
		} else {
			p.a.circuit = currentCircuitId
			p.b.circuit = currentCircuitId
			currentCircuitId++
		}

		circuitId := junctions[0].circuit
		allConnected := true
		for _, j := range junctions {
			if j.circuit != circuitId {
				allConnected = false
				break
			}
		}

		if allConnected {
			ans2 = p.a.point.x * p.b.point.x
			break
		}

		connections++
	}

	return ans1, int(ans2)
}

func Distance(a, b Vec3) float64 {
	dx := b.x - a.x
	dy := b.y - a.y
	dz := b.z - a.z
	return dx*dx + dy*dy + dz*dz
}

func DedupPairs(pairs []JunctionPair) []JunctionPair {
	seen := make(map[[2]int]bool)
	result := []JunctionPair{}

	for _, p := range pairs {
		idA := p.a.id
		idB := p.b.id

		key := [2]int{idA, idB}
		if idA > idB {
			key = [2]int{idB, idA}
		}

		if !seen[key] {
			seen[key] = true
			result = append(result, p)
		}
	}

	return result
}
