package errors

import (
	"fmt"
	"os"
)

func AssertNone(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
