package cmd

import (
	"davideimola.dev/ddd-onion/cmd/server"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-onion-ddd",
		Short: "Project template for DDD with Onion Architecture in Go",
	}
)

func init() {
	rootCmd.AddCommand(server.Cmd)
}

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
