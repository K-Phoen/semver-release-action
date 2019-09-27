package main

import (
	"os"

	"github.com/K-Phoen/semver-release-action/bumper/semver"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "bumper",
}

func main() {
	rootCmd.AddCommand(semver.Command())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
