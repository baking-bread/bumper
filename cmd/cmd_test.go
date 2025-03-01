package cmd

import (
	"bytes"
	"io"
	"testing"
)

func executeCommand(t *testing.T, cmd string) (string, error) {
	rootCmd := NewRootCmd()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{cmd})
	err := rootCmd.Execute()
	out, _ := io.ReadAll(b)
	return string(out), err
}
