package action

import (
	"os"

	"github.com/spf13/cobra"
)

func AssertNoError(cmd *cobra.Command, err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	Fail(cmd, format, args...)
}

func Fail(cmd *cobra.Command, format string, args ...interface{}) {
	cmd.PrintErrf(format, args...)
	os.Exit(1)
}
