package main

import (
	"os"

	"github.com/baking-bread/bumper/cmd"
	"github.com/baking-bread/bumper/internal/logger"
)

func main() {
	log := logger.Init()

	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
