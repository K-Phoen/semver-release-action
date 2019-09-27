package semver

import (
	"fmt"

	"github.com/K-Phoen/semver-release-action/bumper/errors"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:  "semver [VERSION] [INCREMENT]",
		Args: cobra.ExactArgs(2),
		Run:  execute,
	}
}

func execute(cmd *cobra.Command, args []string) {
	currentVersion := args[0]
	increment := args[1]

	version, err := parseVersion(currentVersion)
	errors.AssertNone(err, "invalid version: %s\n", currentVersion)

	inc, err := ParseIncrement(increment)
	errors.AssertNone(err, "invalid increment: %s\n", increment)

	fmt.Print(version.bump(inc))
}
