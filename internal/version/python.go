package version

import (
	"errors"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
)

type PythonReader struct {
	log *logrus.Logger
}

func (r *PythonReader) ReadVersion(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`version\s*=\s*['"]([^'"]+)['"]`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		return "", errors.New("version not found in setup.py")
	}

	r.log.Infof("Extracted version from setup.py: %s", matches[1])
	return matches[1], nil
}
