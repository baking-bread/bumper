package git

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

type Operations struct {
	log     *logrus.Logger
	skipGit bool
}

func NewOperations(log *logrus.Logger, skipGit bool) *Operations {
	return &Operations{
		log:     log,
		skipGit: skipGit,
	}
}

func (g *Operations) CommitAndTag(version string) error {
	if g.skipGit {
		g.log.Info("Skipping Git commit and tag as per user request")
		return nil
	}

	if err := g.commit(version); err != nil {
		return err
	}

	if err := g.tag(version); err != nil {
		return err
	}

	g.log.Infof("Successfully committed and tagged version %s", version)
	return nil
}

func (g *Operations) commit(version string) error {
	cmd := exec.Command("git", "commit", "-am", fmt.Sprintf("Bump version to %s", version))
	if err := cmd.Run(); err != nil {
		g.log.Error("Failed to commit version change")
		return errors.New("failed to commit version change")
	}
	return nil
}

func (g *Operations) tag(version string) error {
	cmd := exec.Command("git", "tag", version)
	if err := cmd.Run(); err != nil {
		g.log.Error("Failed to create Git tag")
		return errors.New("failed to create Git tag")
	}
	return nil
}
