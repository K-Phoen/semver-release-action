package semver

import (
	"github.com/cloud-crafts/semver-release-action/internal/pkg/action"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:  "semver [VERSION] [INCREMENT] [FORMAT]",
		Args: cobra.ExactArgs(3),
		Run:  execute,
	}
}

func execute(cmd *cobra.Command, args []string) {
	currentVersion := args[0]
	increment := args[1]
	format := args[2]

	version, err := ParseVersion(currentVersion)
	action.AssertNoError(cmd, err, "invalid Version: %s\n", currentVersion)

	inc, err := ParseIncrement(increment)
	action.AssertNoError(cmd, err, "invalid increment: %s\n", increment)

	cmd.Print(version.bump(inc).format(format))
}
