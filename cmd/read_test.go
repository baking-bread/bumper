package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCmd(t *testing.T) {
	output, err := executeCommand(t, "read")
	assert.NoError(t, err)
	assert.True(t, strings.Contains(output, "Reading version"))
}
