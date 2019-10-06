package action

import (
	"os"

	"github.com/spf13/cobra"
)

func Skip(cmd *cobra.Command, format string, args ...interface{}) {
	cmd.PrintErrf(format, args...)
	os.Exit(78)
}
