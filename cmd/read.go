package cmd

import (
	"github.com/baking-bread/bumper/internal/logger"
	"github.com/baking-bread/bumper/internal/version"
	"github.com/spf13/cobra"
)

func newReadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "read",
		Short: "Read the current version from project files or Git",
		Run: func(cmd *cobra.Command, args []string) {
			log := logger.Init()
			log.Info("Executing 'read' command")

			reader := version.NewFileReader(log)
			ver, err := reader.DetectAndRead()
			if err != nil {
				log.Error(err)
				return
			}

			cmd.Println("Current version:", ver)
		},
	}
}
