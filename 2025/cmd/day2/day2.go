package day2

import (
	"fmt"
	"strconv"
	"strings"

	"2025/cmd"

	"github.com/spf13/cobra"
)

var DayCmd = &cobra.Command{
	Use:   "day2",
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
	ans1, ans2 := Solve("3335355312-3335478020,62597156-62638027,94888325-95016472,4653-6357,54-79,1-19,314-423,472-650,217886-298699,58843645-58909745,2799-3721,150748-178674,9084373-9176707,1744-2691,17039821-17193560,2140045-2264792,743-1030,6666577818-6666739950,22946-32222,58933-81008,714665437-714803123,9972438-10023331,120068-142180,101-120,726684-913526,7575737649-7575766026,8200-11903,81-96,540949-687222,35704-54213,991404-1009392,335082-425865,196-268,3278941-3383621,915593-991111,32-47,431725-452205")

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	ans1 := 0
	ans2 := 0

	for _, productID := range strings.Split(data, ",") {
		ids := strings.Split(productID, "-")
		id0, _ := strconv.Atoi(ids[0])
		id1, _ := strconv.Atoi(ids[1])

		for id := id0; id <= id1; id++ {
			idStr := strconv.Itoa(id)

			for section := len(idStr) / 2; section > 0; section-- {

				if len(idStr)%section != 0 {
					continue
				}

				strs := []string{}

				prev := 0
				for i := section; i <= len(idStr); i += section {
					strs = append(strs, idStr[prev:i])
					prev = i
				}

				prevStr := ""
				match := true
				for _, str := range strs {
					if len(prevStr) == 0 {
						prevStr = str
						continue
					}

					if prevStr != str {
						match = false
						break
					}
				}

				if match {
					if len(strs) == 2 {
						ans1 += id
					}
					ans2 += id
					break
				}
			}
		}
	}

	return ans1, ans2
}
