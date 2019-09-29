package action

import (
	"fmt"
	"os"
)

func Skip(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(0)
}
