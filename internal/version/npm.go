package version

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type NPMReader struct {
	log *logrus.Logger
}

func (r *NPMReader) ReadVersion(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", err
	}

	if version, ok := result["version"].(string); ok {
		r.log.Infof("Extracted version from package.json: %s", version)
		return version, nil
	}

	return "", errors.New("version not found in package.json")
}
