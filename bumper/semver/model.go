package semver

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var errInvalidVersion = errors.New("invalid version")
var ErrInvalidIncrement = errors.New("invalid increment")

var semverPattern = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)$`)

type Increment string

const (
	IncrementPatch Increment = "patch"
	IncrementMinor Increment = "minor"
	IncrementMajor Increment = "major"
)

type version struct {
	major int
	minor int
	patch int
}

func (v version) bump(inc Increment) version {
	switch inc {
	case IncrementPatch:
		return version{
			patch: v.patch + 1,
			minor: v.minor,
			major: v.major,
		}
	case IncrementMinor:
		return version{
			patch: 0,
			minor: v.minor + 1,
			major: v.major,
		}
	case IncrementMajor:
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

func ParseIncrement(inc string) (Increment, error) {
	switch strings.ToLower(inc) {
	case "patch":
		return IncrementPatch, nil
	case "minor":
		return IncrementMinor, nil
	case "major":
		return IncrementMajor, nil
	}

	return IncrementPatch, ErrInvalidIncrement
}
