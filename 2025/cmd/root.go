package cmd

import (
	"2025/tui"
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "A collection of advent of code challanges",
	Long: `Advent of code 2025 made as a CLI tool
	to make it easier to manage and run each of the challanges for the day`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

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
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return tui.RunTUI(cmd)
		}

		return nil
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
