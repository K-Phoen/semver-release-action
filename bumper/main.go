package main

import (
	"os"

	"github.com/K-Phoen/semver-release-action/bumper/event"
	"github.com/K-Phoen/semver-release-action/bumper/release"
	"github.com/K-Phoen/semver-release-action/bumper/semver"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "bumper",
}

func main() {
	rootCmd.AddCommand(semver.Command())
	rootCmd.AddCommand(release.Command())
	rootCmd.AddCommand(event.GuardCommand())
	rootCmd.AddCommand(event.IncrementCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
