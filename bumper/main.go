package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var currentVersion string
	var increment string

	flag.StringVar(&currentVersion, "current", "", "Current version")
	flag.StringVar(&increment, "increment", "patch", "Semver increment (patch, minor or major")

	flag.Parse()

	version, err := parseVersion(currentVersion)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid version: %s\n", currentVersion)
		os.Exit(1)
	}

	inc, err := parseIncrement(increment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid increment: %s\n", increment)
		os.Exit(1)
	}

	fmt.Print(version.bump(inc))
}
