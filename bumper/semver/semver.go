package semver

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:  "semver [VERSION] [INCREMENT]",
		Args: cobra.MinimumNArgs(2),
		Run:  execute,
	}
}

func execute(cmd *cobra.Command, args []string) {
	currentVersion := args[0]
	increment := args[1]

	version, err := parseVersion(currentVersion)
	assertNoError(err, "invalid version: %s\n", currentVersion)

	inc, err := parseIncrement(increment)
	assertNoError(err, "invalid increment: %s\n", increment)

	fmt.Print(version.bump(inc))
}

func assertNoError(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
