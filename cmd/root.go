package cmd

import (
	"github.com/spf13/cobra"
)

var (
	filepath        string
	package_manager string
	regexpPattern   string
)

var (
	maven_project_version = `(?s)(?:<project.*?<version>)(?P<version>.*?)(?:<\/version>.*?<\/project>)`
	npm_project_version   = `(?:["']?version["']?\s*:.*?["'])(?P<version>.*?)(?:["'])`
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "version-tool",
		Short: "A CLI tool for managing project versions",
	}

	// Add subcommands
	rootCmd.AddCommand(
		newGetCmd(),
		newSetCmd(),
	)

	return rootCmd
}
