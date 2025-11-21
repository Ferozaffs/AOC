package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "aoc2024",
	Short: "A collection of advent of code challanges",
	Long: `Advent of code 2024 made as a CLI tool
	to make it easier to manage and run each of the challanges for the day`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		f, err := os.Create("cpu.prof")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating CPU profile: %v\n", err)
			os.Exit(1)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting CPU profile: %v\n", err)
			os.Exit(1)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			f.Close()
			os.Exit(0)
		}()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		pprof.StopCPUProfile()
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
