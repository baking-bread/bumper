package cmd

import (
	"github.com/baking-bread/bumper/internal/logger"
	"github.com/baking-bread/bumper/internal/version"
	"github.com/spf13/cobra"
)

func newBumpCmd() *cobra.Command {
	var (
		bumpType string
		dryRun   bool
		force    bool
	)

	cmd := &cobra.Command{
		Use:   "bump",
		Short: "Bump the project version (patch, minor, major)",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.Init()
			log.Info("Executing 'bump' command")

			reader := version.NewFileReader(log)
			currentVersion, err := reader.DetectAndRead()
			if err != nil {
				return err
			}

			sv := version.NewSemVer(log, dryRun, force)
			if err := sv.Parse(currentVersion); err != nil {
				return err
			}

			if err := sv.Bump(bumpType); err != nil {
				return err
			}

			if dryRun {
				cmd.Println("Dry run: version would be bumped to:", sv.String())
			} else {
				cmd.Println("Bumped version to:", sv.String())
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&bumpType, "type", "t", "patch", "Version bump type (patch|minor|major)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would happen without making changes")
	cmd.Flags().BoolVar(&force, "force", false, "Force version change even if it would result in a downgrade")
	return cmd
}
