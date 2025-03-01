package version

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type GitReader struct {
	log *logrus.Logger
}

func (r *GitReader) ReadVersion(filename string) (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--match", "v*.*.*")
	output, err := cmd.Output()
	if err != nil {
		r.log.Error("No valid Git tag found")
		return "", errors.New("no valid Git tag found")
	}

	version := strings.TrimSpace(string(output))
	r.log.Infof("Extracted version from Git tag: %s", version)
	return version, nil
}
