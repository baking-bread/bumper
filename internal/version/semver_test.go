package version

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSemVer(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		bumpType  string
		dryRun    bool
		force     bool
		want      string
		wantError bool
	}{
		{"patch bump", "v1.2.3", "patch", false, false, "v1.2.4", false},
		{"minor bump", "v1.2.3", "minor", false, false, "v1.3.0", false},
		{"major bump", "v1.2.3", "major", false, false, "v2.0.0", false},
		{"patch bump dry-run", "v1.2.3", "patch", true, false, "v1.2.3", false},
		{"minor bump dry-run", "v1.2.3", "minor", true, false, "v1.2.3", false},
		{"major bump dry-run", "v1.2.3", "major", true, false, "v1.2.3", false},
		{"prevent downgrade", "v2.0.0", "patch", false, false, "", true},
		{"force downgrade", "v2.0.0", "patch", false, true, "v2.0.1", false},
		{"invalid version", "v1.2", "", false, false, "", true},
		{"invalid bump type", "v1.2.3", "invalid", false, false, "", true},
		{"version without v", "1.2.3", "patch", false, false, "v1.2.4", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := NewSemVer(logrus.New(), tt.dryRun, tt.force)
			err := sv.Parse(tt.version)
			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.bumpType != "" {
				err = sv.Bump(tt.bumpType)
				assert.NoError(t, err)
				if tt.dryRun {
					// In dry-run mode, version should remain unchanged
					assert.Equal(t, tt.version, sv.String())
				} else {
					assert.Equal(t, tt.want, sv.String())
				}
			}
		})
	}
}

func TestSemVerCompare(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{"equal versions", "v1.2.3", "v1.2.3", 0},
		{"greater major", "v2.0.0", "v1.2.3", 1},
		{"less major", "v1.2.3", "v2.0.0", -1},
		{"greater minor", "v1.3.0", "v1.2.3", 1},
		{"less minor", "v1.2.3", "v1.3.0", -1},
		{"greater patch", "v1.2.4", "v1.2.3", 1},
		{"less patch", "v1.2.3", "v1.2.4", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv1 := NewSemVer(logrus.New(), false, false)
			sv2 := NewSemVer(logrus.New(), false, false)
			sv1.Parse(tt.version1)
			sv2.Parse(tt.version2)
			assert.Equal(t, tt.want, sv1.Compare(sv2))
		})
	}
}
