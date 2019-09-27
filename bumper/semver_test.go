package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionBump(t *testing.T) {
	cases := []struct {
		name string

		version   version
		increment increment

		expectedVersion string
	}{
		{
			name:            "patch bump",
			version:         version{major: 1, minor: 2, patch: 3},
			increment:       incrementPatch,
			expectedVersion: "v1.2.4",
		},
		{
			name:            "minor bump",
			version:         version{major: 1, minor: 2, patch: 3},
			increment:       incrementMinor,
			expectedVersion: "v1.3.0",
		},
		{
			name:            "major bump",
			version:         version{major: 1, minor: 2, patch: 3},
			increment:       incrementMajor,
			expectedVersion: "v2.0.0",
		},
	}

	for _, testCase := range cases {
		bumped := testCase.version.bump(testCase.increment)

		require.Equal(t, testCase.expectedVersion, bumped.String())
	}
}

func TestParseIncrement(t *testing.T) {
	cases := []struct {
		input string

		expectedIncrement increment
		expectedError     error
	}{
		{
			input:             "patch",
			expectedIncrement: incrementPatch,
			expectedError:     nil,
		},
		{
			input:             "PaTcH",
			expectedIncrement: incrementPatch,
			expectedError:     nil,
		},
		{
			input:             "minor",
			expectedIncrement: incrementMinor,
			expectedError:     nil,
		},
		{
			input:             "mInOr",
			expectedIncrement: incrementMinor,
			expectedError:     nil,
		},
		{
			input:             "major",
			expectedIncrement: incrementMajor,
			expectedError:     nil,
		},
		{
			input:             "mAjOR",
			expectedIncrement: incrementMajor,
			expectedError:     nil,
		},

		{
			input:             "micro",
			expectedIncrement: incrementPatch,
			expectedError:     errInvalidIncrement,
		},
	}

	for _, testCase := range cases {
		inc, err := parseIncrement(testCase.input)

		require.Equal(t, testCase.expectedIncrement, inc)
		require.Equal(t, testCase.expectedError, err)
	}
}

func TestParseVersion(t *testing.T) {
	cases := []struct {
		input string

		expectedVersion version
		expectedError   error
	}{
		{
			input:           "1.2.3",
			expectedVersion: version{major: 1, minor: 2, patch: 3},
			expectedError:   nil,
		},
		{
			input:           "v1.2.3",
			expectedVersion: version{major: 1, minor: 2, patch: 3},
			expectedError:   nil,
		},
		{
			// could be improved... or I should probably just re-use a lib that already properly handles semver.
			input:           "v1.2",
			expectedVersion: version{},
			expectedError:   errInvalidVersion,
		},
		{
			input:           "vlala.2.3",
			expectedVersion: version{},
			expectedError:   errInvalidVersion,
		},
		{
			input:           "v1.lala.3",
			expectedVersion: version{},
			expectedError:   errInvalidVersion,
		},
		{
			input:           "v1.2.lala",
			expectedVersion: version{},
			expectedError:   errInvalidVersion,
		},
	}

	for _, testCase := range cases {
		ver, err := parseVersion(testCase.input)

		require.Equal(t, testCase.expectedVersion, ver)
		require.Equal(t, testCase.expectedError, err)
	}
}
