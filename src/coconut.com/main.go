package main

import (
	"github.com/spf13/cobra"
	"coconut.com/agent"
	"fmt"
)

var name = "deploygate"

var rootCmd = &cobra.Command{Use: name}

func main() {
	rootCmd.AddCommand(agent.Cmd)
	if err := rootCmd.Execute(); err != nil {
		panic(fmt.Sprintf("error executing rootCmd: %s", err))
	}
}
