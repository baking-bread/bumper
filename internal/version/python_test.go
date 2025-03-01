package version

import (
	"path/filepath"
	"testing"

	"github.com/baking-bread/bumper/internal/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPythonReader(t *testing.T) {
	dir := testutil.CreateTempDir(t)
	filename := filepath.Join(dir, "setup.py")
	content := `setup(name='test', version='3.4.5')`

	testutil.CreateTempFile(t, filename, content)

	reader := &PythonReader{log: logrus.New()}
	version, err := reader.ReadVersion(filename)

	assert.NoError(t, err)
	assert.Equal(t, "3.4.5", version)
}
