package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var errInvalidVersion = errors.New("invalid version")
var errInvalidIncrement = errors.New("invalid increment")

var semverPattern = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)$`)

type increment int

const (
	incrementPatch increment = iota
	incrementMinor
	incrementMajor
)

type version struct {
	major int
	minor int
	patch int
}

func (v version) bump(inc increment) version {
	switch inc {
	case incrementPatch:
		return version{
			patch: v.patch + 1,
			minor: v.minor,
			major: v.major,
		}
	case incrementMinor:
		return version{
			patch: 0,
			minor: v.minor + 1,
			major: v.major,
		}
	case incrementMajor:
		return version{
			patch: 0,
			minor: 0,
			major: v.major + 1,
		}
	}

	return v
}

func (v version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}

func parseVersion(v string) (version, error) {
	parts := semverPattern.FindStringSubmatch(v)
	if len(parts) != 4 {
		return version{}, errInvalidVersion
	}

	major, err := strconv.Atoi(parts[1])
	if err != nil {
		return version{}, errInvalidVersion
	}

	minor, err := strconv.Atoi(parts[2])
	if err != nil {
		return version{}, errInvalidVersion
	}

	patch, err := strconv.Atoi(parts[3])
	if err != nil {
		return version{}, errInvalidVersion
	}

	return version{
		major: major,
		minor: minor,
		patch: patch,
	}, nil
}

func parseIncrement(inc string) (increment, error) {
	switch strings.ToLower(inc) {
	case "patch":
		return incrementPatch, nil
	case "minor":
		return incrementMinor, nil
	case "major":
		return incrementMajor, nil
	}

	return incrementPatch, errInvalidIncrement
}
