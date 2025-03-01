package cmd

import (
	"github.com/baking-bread/bumper/internal/git"
	"github.com/baking-bread/bumper/internal/logger"
	"github.com/baking-bread/bumper/internal/version"
	"github.com/spf13/cobra"
)

func newTagCmd() *cobra.Command {
	var skipGit bool

	cmd := &cobra.Command{
		Use:   "tag",
		Short: "Create a Git tag for the current version",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.Init()
			log.Info("Executing 'tag' command")

			reader := version.NewFileReader(log)
			ver, err := reader.DetectAndRead()
			if err != nil {
				return err
			}

			gitOps := git.NewOperations(log, skipGit)
			if err := gitOps.CommitAndTag(ver); err != nil {
				return err
			}

			cmd.Printf("Successfully tagged version: %s\n", ver)
			return nil
		},
	}

	cmd.Flags().BoolVar(&skipGit, "skip-git", false, "Skip Git operations")
	return cmd
}
