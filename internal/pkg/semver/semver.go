package semver

import (
	"github.com/K-Phoen/semver-release-action/internal/pkg/action"
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

	version, err := ParseVersion(currentVersion)
	action.AssertNoError(err, "invalid Version: %s\n", currentVersion)

	inc, err := ParseIncrement(increment)
	action.AssertNoError(err, "invalid increment: %s\n", increment)

	cmd.Print(version.bump(inc))
}
