package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		description string
		version     string
		major       int
		minor       int
		patch       int
		preRelease  string
		build       string
		wantError   bool
	}{
		{"simple", "v1.2.3", 1, 2, 3, "", "", false},
		{"no prefix", "1.2.3", 1, 2, 3, "", "", false},
		{"zero major version", "v0.1.0", 0, 1, 0, "", "", false},
		{"with pre-release", "v1.2.3-beta", 1, 2, 3, "beta", "", false},
		{"with build", "v1.2.3+20230314", 1, 2, 3, "", "20230314", false},
		{"with pre-release and build", "v1.2.3-beta+20230314", 1, 2, 3, "beta", "20230314", false},
		{"non-numeric version", "vX.Y.Z", 0, 0, 0, "", "", true},
		{"negative version numbers", "v1.-2.3", 0, 0, 0, "", "", true},
		{"extra dots", "v1.2.3.4", 1, 2, 3, "", "", false},
		{"missing numbers", "v1..3", 0, 0, 0, "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			sv, err := Parse(tt.version)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.major, sv.Major)
			assert.Equal(t, tt.minor, sv.Minor)
			assert.Equal(t, tt.patch, sv.Patch)
			assert.Equal(t, tt.preRelease, sv.PreRelease)
			assert.Equal(t, tt.build, sv.Build)
		})
	}
}
