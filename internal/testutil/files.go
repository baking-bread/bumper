package testutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CreateTempFile creates a temporary file with given content
func CreateTempFile(t *testing.T, filename, content string) {
	t.Helper()
	err := os.WriteFile(filename, []byte(content), 0644)
	assert.NoError(t, err)
	t.Cleanup(func() {
		os.Remove(filename)
	})
}

// CreateTempDir creates a temporary directory and returns its path
func CreateTempDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "version-test-*")
	assert.NoError(t, err)
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})
	return dir
}
