package semver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	version "github.com/blang/semver"
)

var ErrInvalidIncrement = errors.New("invalid increment")

type Increment string

const (
	IncrementSkip  Increment = "skip"
	IncrementPatch Increment = "patch"
	IncrementMinor Increment = "minor"
	IncrementMajor Increment = "major"
)

type Version struct {
	major uint64
	minor uint64
	patch uint64
}

func (v Version) bump(inc Increment) Version {
	switch inc {
	case IncrementPatch:
		return Version{
			patch: v.patch + 1,
			minor: v.minor,
			major: v.major,
		}
	case IncrementMinor:
		return Version{
			patch: 0,
			minor: v.minor + 1,
			major: v.major,
		}
	case IncrementMajor:
		return Version{
			patch: 0,
			minor: 0,
			major: v.major + 1,
		}
	}

	return v
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}

func (v Version) format(format string) string {
	formatted := format

	formatted = strings.ReplaceAll(formatted, "%major%", strconv.FormatUint(v.major, 10))
	formatted = strings.ReplaceAll(formatted, "%minor%", strconv.FormatUint(v.minor, 10))
	formatted = strings.ReplaceAll(formatted, "%patch%", strconv.FormatUint(v.patch, 10))

	return formatted
}

func ParseVersion(input string) (Version, error) {
	v, err := version.ParseTolerant(input)
	if err != nil {
		return Version{}, err
	}

	return Version{
		major: v.Major,
		minor: v.Minor,
		patch: v.Patch,
	}, nil
}

func ParseIncrement(inc string) (Increment, error) {
	switch strings.ToLower(inc) {
	case "skip":
		return IncrementSkip, nil
	case "patch":
		return IncrementPatch, nil
	case "minor":
		return IncrementMinor, nil
	case "major":
		return IncrementMajor, nil
	}

	return IncrementPatch, ErrInvalidIncrement
}
