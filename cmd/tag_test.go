package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagCmd(t *testing.T) {
	output, err := executeCommand(t, "tag")
	assert.NoError(t, err)
	assert.True(t, strings.Contains(output, "Tagging version"))
}
