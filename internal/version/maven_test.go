package version

import (
	"path/filepath"
	"testing"

	"github.com/baking-bread/bumper/internal/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMavenReader(t *testing.T) {
	dir := testutil.CreateTempDir(t)
	filename := filepath.Join(dir, "pom.xml")
	content := `<project><version>1.2.3</version></project>`

	testutil.CreateTempFile(t, filename, content)

	reader := &MavenReader{log: logrus.New()}
	version, err := reader.ReadVersion(filename)

	assert.NoError(t, err)
	assert.Equal(t, "1.2.3", version)
}
