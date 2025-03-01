package version

import (
	"path/filepath"
	"testing"

	"github.com/baking-bread/bumper/internal/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNPMReader(t *testing.T) {
	dir := testutil.CreateTempDir(t)
	filename := filepath.Join(dir, "package.json")
	content := `{"name": "test", "version": "2.3.4"}`

	testutil.CreateTempFile(t, filename, content)

	reader := &NPMReader{log: logrus.New()}
	version, err := reader.ReadVersion(filename)

	assert.NoError(t, err)
	assert.Equal(t, "2.3.4", version)
}
