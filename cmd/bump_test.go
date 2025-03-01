package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBumpCmd(t *testing.T) {
	output, err := executeCommand(t, "bump")
	assert.NoError(t, err)
	assert.True(t, strings.Contains(output, "Bumping version"))
}
