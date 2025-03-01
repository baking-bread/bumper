package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "version-tool",
		Short: "A CLI tool for managing project versions",
	}

	// Add subcommands
	rootCmd.AddCommand(
		newReadCmd(),
		newBumpCmd(),
		newTagCmd(),
	)

	return rootCmd
}
