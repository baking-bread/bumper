package version

import (
	"os"
	"testing"

	"github.com/baking-bread/bumper/internal/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFileReaderNoFiles(t *testing.T) {
	dir := testutil.CreateTempDir(t)
	prevWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(prevWd)

	reader := NewFileReader(logrus.New())
	version, err := reader.DetectAndRead()

	assert.Error(t, err)
	assert.Equal(t, "", version)
}

func TestFileReaderWithFiles(t *testing.T) {
	dir := testutil.CreateTempDir(t)
	prevWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(prevWd)

	testutil.CreateTempFile(t, "package.json", `{"version": "1.0.0"}`)

	reader := NewFileReader(logrus.New())
	version, err := reader.DetectAndRead()

	assert.NoError(t, err)
	assert.Equal(t, "1.0.0", version)
}
