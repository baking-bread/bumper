package version

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type Reader interface {
	ReadVersion(filename string) (string, error)
}

type FileReader struct {
	log *logrus.Logger
}

func NewFileReader(log *logrus.Logger) *FileReader {
	return &FileReader{log: log}
}

func (r *FileReader) DetectAndRead() (string, error) {
	// Try Git first
	gitReader := &GitReader{log: r.log}
	if version, err := gitReader.ReadVersion(""); err == nil {
		return version, nil
	}

	// Try file-based readers
	files := []struct {
		filename string
		reader   Reader
	}{
		{"pom.xml", &MavenReader{log: r.log}},
		{"package.json", &NPMReader{log: r.log}},
		{"setup.py", &PythonReader{log: r.log}},
	}

	for _, f := range files {
		if _, err := os.Stat(f.filename); err == nil {
			r.log.Infof("Detected version file: %s", f.filename)
			return f.reader.ReadVersion(f.filename)
		}
	}

	r.log.Error("No known version file found")
	return "", errors.New("no version file found")
}
