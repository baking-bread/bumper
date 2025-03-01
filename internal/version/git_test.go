package version

import (
	"os"
	"os/exec"
	"testing"

	"github.com/baking-bread/bumper/internal/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGitReader(t *testing.T) {
	// Skip if git is not installed
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not available")
	}

	tests := []struct {
		name      string
		setup     func(t *testing.T, dir string)
		wantVer   string
		wantError bool
	}{
		{
			name: "with valid tag",
			setup: func(t *testing.T, dir string) {
				setupGitRepo(t, dir, true)
			},
			wantVer:   "v1.0.0",
			wantError: false,
		},
		{
			name: "without tags",
			setup: func(t *testing.T, dir string) {
				setupGitRepo(t, dir, false)
			},
			wantVer:   "",
			wantError: true,
		},
		{
			name:      "not a git repository",
			setup:     func(t *testing.T, dir string) {},
			wantVer:   "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := testutil.CreateTempDir(t)
			prevWd, _ := os.Getwd()
			t.Cleanup(func() { os.Chdir(prevWd) })
			os.Chdir(dir)

			tt.setup(t, dir)

			reader := &GitReader{log: logrus.New()}
			version, err := reader.ReadVersion("")

			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, version)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVer, version)
			}
		})
	}
}

func setupGitRepo(t *testing.T, dir string, withTag bool) {
	cmds := [][]string{
		{"git", "init"},
		{"git", "config", "user.email", "test@example.com"},
		{"git", "config", "user.name", "Test User"},
		{"git", "commit", "--allow-empty", "-m", "Initial commit"},
	}

	if withTag {
		cmds = append(cmds, []string{"git", "tag", "v1.0.0"})
	}

	for _, cmd := range cmds {
		err := exec.Command(cmd[0], cmd[1:]...).Run()
		assert.NoError(t, err)
	}
}
