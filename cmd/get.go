package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/baking-bread/bumper/internal/version"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Prints the current version",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch package_manager {
		case "maven":
			regexpPattern = maven_project_version
		case "npm":
			regexpPattern = npm_project_version
		}

		// check if file exists
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", filepath)
		}

		// read file
		data, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}

		// match version field
		reg := regexp.MustCompile(regexpPattern)
		match, err := version.Match(reg, string(data))
		if err != nil {
			return err
		}

		// parse version
		semVer, err := version.Parse(match)
		if err != nil {
			return err
		}

		fmt.Print(semVer)

		return nil
	},
}

func newGetCmd() *cobra.Command {
	getCmd.PersistentFlags().StringVarP(&filepath, "file", "f", "", "Path to file")
	getCmd.PersistentFlags().StringVarP(&package_manager, "package-manager", "p", "", "Package manager of the project")

	return getCmd
}
