package day19

import (
	"aoc2023/cmd"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Condition struct {
	token  string
	larger bool
	value  int
	result string
}

type Workflow struct {
	conditions []Condition
	fallback   string
}

func (wf Workflow) WorkflowProcess(wfs map[string]Workflow, p Part) string {
	for _, c := range wf.conditions {
		val := -1
		if c.token == "x" {
			val = p.x
		} else if c.token == "m" {
			val = p.m
		} else if c.token == "a" {
			val = p.a
		} else if c.token == "s" {
			val = p.s
		}

		conditionMet := false
		if val != -1 {
			if c.larger == true && val > c.value {
				conditionMet = true
			} else if c.larger == false && val < c.value {
				conditionMet = true
			}
		}

		if conditionMet == true {
			if c.result == "A" {
				return "A"
			} else if c.result == "R" {
				return "R"
			} else {
				return wfs[c.result].WorkflowProcess(wfs, p)
			}
		}
	}

	if wf.fallback == "A" {
		return "A"
	} else if wf.fallback == "R" {
		return "R"
	} else {
		return wfs[wf.fallback].WorkflowProcess(wfs, p)
	}
}

var day19_1Cmd = &cobra.Command{
	Use:   "day19_1",
	Short: "Day 19, challenge 1",
	Long:  `Day 19, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day19_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day19_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(data, "\n")

	splitData := regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)

	wfs := ExtractWorkflows(splitData[0])
	parts := ExtractParts(splitData[1])

	sum := 0
	for _, part := range parts {
		result := wfs["in"].WorkflowProcess(wfs, part)
		if result == "A" {
			sum += part.x + part.m + part.a + part.s
		}
	}

	return sum
}

func ExtractWorkflows(data string) map[string]Workflow {
	reConds := regexp.MustCompile("[xmas](>|<)\\d*:\\w*")
	reFallback := regexp.MustCompile("\\w*")

	rows := strings.Split(strings.ReplaceAll(data, "\t", ""), "\n")

	workflows := map[string]Workflow{}
	for _, r := range rows {

		split := strings.Split(r, "{")
		coStrings := reConds.FindAllString(r, -1)
		fallback := DeleteEmpty(reFallback.FindAllString(r, -1))

		cos := []Condition{}
		for _, s := range coStrings {
			co := Condition{}
			co.token = string(s[0])

			if s[1] == '>' {
				co.larger = true
			} else {
				co.larger = false
			}

			n := DeleteEmpty(regexp.MustCompile("\\d*").FindAllString(s, -1))
			co.value, _ = strconv.Atoi(n[0])
			result := regexp.MustCompile("\\w*").FindAllString(s, -1)
			co.result = result[len(result)-1]

			cos = append(cos, co)
		}

		workflows[split[0]] = Workflow{
			conditions: cos, fallback: fallback[len(fallback)-1],
		}
	}
	return workflows
}

func ExtractParts(data string) []Part {
	rows := strings.Split(strings.ReplaceAll(data, "\t", ""), "\n")

	parts := []Part{}
	for _, r := range rows {
		re := regexp.MustCompile("[0-9]*")
		n := DeleteEmpty(re.FindAllString(r, -1))

		xv, _ := strconv.Atoi(n[0])
		mv, _ := strconv.Atoi(n[1])
		av, _ := strconv.Atoi(n[2])
		sv, _ := strconv.Atoi(n[3])

		parts = append(parts, Part{x: xv, m: mv, a: av, s: sv})
	}

	return parts
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
