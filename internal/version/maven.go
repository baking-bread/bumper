package version

import (
	"encoding/xml"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type MavenReader struct {
	log *logrus.Logger
}

type mavenProject struct {
	Version string `xml:"version"`
}

func (r *MavenReader) ReadVersion(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var project mavenProject
	if err := xml.Unmarshal(data, &project); err != nil {
		return "", err
	}

	if project.Version == "" {
		return "", errors.New("version not found in pom.xml")
	}

	r.log.Infof("Extracted version from pom.xml: %s", project.Version)
	return project.Version, nil
}
