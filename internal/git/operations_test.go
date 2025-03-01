package git

import (
	"os"
	"os/exec"
	"testing"

	"github.com/baking-bread/bumper/internal/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGitOperations(t *testing.T) {
	// Skip if git is not installed
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not available")
	}

	tests := []struct {
		name      string
		version   string
		skipGit   bool
		setup     func(t *testing.T, dir string)
		wantError bool
	}{
		{
			name:      "successful commit and tag",
			version:   "v1.2.3",
			skipGit:   false,
			setup:     setupTestRepo,
			wantError: false,
		},
		{
			name:      "skip git operations",
			version:   "v1.2.3",
			skipGit:   true,
			setup:     nil,
			wantError: false,
		},
		{
			name:      "not a git repository",
			version:   "v1.2.3",
			skipGit:   false,
			setup:     nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := testutil.CreateTempDir(t)
			prevWd, _ := os.Getwd()
			t.Cleanup(func() { os.Chdir(prevWd) })
			os.Chdir(dir)

			if tt.setup != nil {
				tt.setup(t, dir)
			}

			ops := NewOperations(logrus.New(), tt.skipGit)
			err := ops.CommitAndTag(tt.version)

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if !tt.skipGit {
					assertTagExists(t, tt.version)
				}
			}
		})
	}
}

func setupTestRepo(t *testing.T, dir string) {
	cmds := [][]string{
		{"git", "init"},
		{"git", "config", "user.email", "test@example.com"},
		{"git", "config", "user.name", "Test User"},
		{"git", "config", "commit.gpgsign", "false"},
	}

	for _, cmd := range cmds {
		err := exec.Command(cmd[0], cmd[1:]...).Run()
		assert.NoError(t, err)
	}

	testutil.CreateTempFile(t, "testfile", "test content")

	cmds = [][]string{
		{"git", "add", "testfile"},
		{"git", "commit", "-m", "Initial commit"},
	}

	for _, cmd := range cmds {
		err := exec.Command(cmd[0], cmd[1:]...).Run()
		assert.NoError(t, err)
	}
}

func assertTagExists(t *testing.T, version string) {
	cmd := exec.Command("git", "tag", "-l", version)
	output, err := cmd.Output()
	assert.NoError(t, err)
	assert.Contains(t, string(output), version, "Tag should exist")
}
