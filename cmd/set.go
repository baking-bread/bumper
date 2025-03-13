package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/baking-bread/bumper/internal/version"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the version",
	RunE: func(cmd *cobra.Command, args []string) error {
		var update string = args[0]

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

		// update version field
		reg := regexp.MustCompile(regexpPattern)
		content := version.Update(reg, string(data), update)

		// write content
		err = os.WriteFile(filepath, []byte(content), os.ModePerm)
		if err != nil {
			return err
		}

		return nil
	},
}

func newSetCmd() *cobra.Command {
	setCmd.PersistentFlags().StringVarP(&filepath, "file", "f", "", "Path to file")
	setCmd.PersistentFlags().StringVarP(&package_manager, "package-manager", "p", "", "Package manager of the project")

	return setCmd
}
