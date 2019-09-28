package action

import (
	"fmt"
	"os"
)

func AssertNoError(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	Fail(format, args...)
}

func Fail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
