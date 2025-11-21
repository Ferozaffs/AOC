package day20

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Module struct {
	name             string
	moduleType       int
	destinationNames []string
	destinations     []*Module
	value            bool
	connectionValues map[*Module]bool
}

type Signal struct {
	value       bool
	sender      *Module
	destination *Module
}

var day20_1Cmd = &cobra.Command{
	Use:   "day20_1",
	Short: "Day 20, challenge 1",
	Long:  `Day 20, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day20_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day20_data.txt")
	ans := Solve1(string(dat), false)

	fmt.Printf("ANSWER: %d\n", ans)
}

var highSignals int = 0
var lowSignals int = 0
var pressToRX = 0
var rxSignal bool = true

func Solve1(data string, checkRX bool) int {
	highSignals = 0
	lowSignals = 0

	modules := ParseModules(data)
	ConfigureModules(&modules)

	var broadcaster *Module
	for _, m := range modules {
		if m.name == "broadcaster" {
			broadcaster = &m
			break
		}
	}

	for i := 0; i < 1000; i++ {
		//Quick RX hack
		if checkRX == true {
			i = 0
		}

		if rxSignal == false {
			return pressToRX
		}
		pressToRX++
		//

		queue := []Signal{Signal{false, nil, broadcaster}}
		for j := 0; j < len(queue); j++ {
			s := SendSignal(queue[j])
			queue = append(queue, s...)
		}
	}

	return highSignals * lowSignals
}

func ParseModules(data string) []Module {
	modules := []Module{}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "->")

		s[0] = strings.TrimPrefix(s[0], "\t")

		mt := -1
		if s[0][0] == '%' {
			mt = 0
			s[0] = s[0][1:]
		} else if s[0][0] == '&' {
			mt = 1
			s[0] = s[0][1:]
		}

		d := strings.Split(s[1], ",")

		modules = append(modules, Module{name: s[0][:len(s[0])-1], moduleType: mt, destinationNames: d, value: false})
	}

	return modules
}

func ConfigureModules(modules *[]Module) {
	for i, m := range *modules {
	Names:
		for _, d := range m.destinationNames {
			for j, n := range *modules {
				if n.name == d[1:] {
					(*modules)[i].destinations = append((*modules)[i].destinations, &(*modules)[j])
					continue Names
				}
			}

			(*modules)[i].destinations = append((*modules)[i].destinations, nil)
		}
	}

	for i, m := range *modules {
		if m.moduleType == 1 {
			(*modules)[i].connectionValues = make(map[*Module]bool)

			for j, n := range *modules {
				for _, d := range n.destinations {
					if d == &(*modules)[i] {
						(*modules)[i].connectionValues[&(*modules)[j]] = false
					}
				}
			}
		}

	}
}

func SendSignal(s Signal) []Signal {
	if s.value == false {
		lowSignals++
	} else {
		highSignals++
	}

	signals := []Signal{}

	if s.destination == nil {
		return signals
	}

	m := s.destination

	transmitSignal := false
	if m.moduleType == 0 {
		if s.value == true {
			return signals
		}

		m.value = !m.value
		transmitSignal = m.value
	} else if m.moduleType == 1 {
		m.connectionValues[s.sender] = s.value

		highs := true
		for _, c := range m.connectionValues {
			if c == false {
				highs = false
				break
			}
		}

		transmitSignal = !highs
	}

	for _, d := range m.destinations {
		signals = append(signals, Signal{transmitSignal, m, d})
	}

	for _, d := range m.destinationNames {
		if d == "rx" {
			rxSignal = transmitSignal
		}
	}

	return signals
}
