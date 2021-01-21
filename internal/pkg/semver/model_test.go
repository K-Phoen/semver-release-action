package semver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionBump(t *testing.T) {
	cases := []struct {
		name string

		version   Version
		increment Increment

		expectedVersion string
	}{
		{
			name:            "skip bump",
			version:         Version{major: 1, minor: 2, patch: 3},
			increment:       IncrementSkip,
			expectedVersion: "v1.2.3",
		},
		{
			name:            "patch bump",
			version:         Version{major: 1, minor: 2, patch: 3},
			increment:       IncrementPatch,
			expectedVersion: "v1.2.4",
		},
		{
			name:            "minor bump",
			version:         Version{major: 1, minor: 2, patch: 3},
			increment:       IncrementMinor,
			expectedVersion: "v1.3.0",
		},
		{
			name:            "major bump",
			version:         Version{major: 1, minor: 2, patch: 3},
			increment:       IncrementMajor,
			expectedVersion: "v2.0.0",
		},
	}

	for _, testCase := range cases {
		bumped := testCase.version.bump(testCase.increment)

		require.Equal(t, testCase.expectedVersion, bumped.String())
	}
}

func TestFormat(t *testing.T) {
	cases := []struct {
		version Version
		format  string

		expectedVersion string
	}{
		{
			version:         Version{major: 1, minor: 2, patch: 3},
			format:          "v%major%.%minor%.%patch%-RC",
			expectedVersion: "v1.2.3-RC",
		},
	}

	for _, testCase := range cases {
		require.Equal(t, testCase.expectedVersion, testCase.version.format(testCase.format))
	}
}

func TestParseIncrement(t *testing.T) {
	cases := []struct {
		input string

		expectedIncrement Increment
		expectedError     error
	}{
		{
			input:             "skip",
			expectedIncrement: IncrementSkip,
			expectedError:     nil,
		},
		{
			input:             "patch",
			expectedIncrement: IncrementPatch,
			expectedError:     nil,
		},
		{
			input:             "PaTcH",
			expectedIncrement: IncrementPatch,
			expectedError:     nil,
		},
		{
			input:             "minor",
			expectedIncrement: IncrementMinor,
			expectedError:     nil,
		},
		{
			input:             "mInOr",
			expectedIncrement: IncrementMinor,
			expectedError:     nil,
		},
		{
			input:             "major",
			expectedIncrement: IncrementMajor,
			expectedError:     nil,
		},
		{
			input:             "mAjOR",
			expectedIncrement: IncrementMajor,
			expectedError:     nil,
		},

		{
			input:             "micro",
			expectedIncrement: IncrementPatch,
			expectedError:     ErrInvalidIncrement,
		},
	}

	for _, testCase := range cases {
		inc, err := ParseIncrement(testCase.input)

		require.Equal(t, testCase.expectedIncrement, inc)
		require.Equal(t, testCase.expectedError, err)
	}
}

func TestParseVersion(t *testing.T) {
	cases := []struct {
		input string

		expectedVersion Version
		expectError     bool
	}{
		{
			input:           "1.2.3",
			expectedVersion: Version{major: 1, minor: 2, patch: 3},
			expectError:     false,
		},
		{
			input:           "v1.2.3",
			expectedVersion: Version{major: 1, minor: 2, patch: 3},
			expectError:     false,
		},
		{
			input:           "v1.2",
			expectedVersion: Version{major: 1, minor: 2, patch: 0},
			expectError:     false,
		},
		{
			input:           "vlala.2.3",
			expectedVersion: Version{},
			expectError:     true,
		},
		{
			input:           "v1.lala.3",
			expectedVersion: Version{},
			expectError:     true,
		},
		{
			input:           "v1.2.lala",
			expectedVersion: Version{},
			expectError:     true,
		},
	}

	for _, testCase := range cases {
		ver, err := ParseVersion(testCase.input)

		require.Equal(t, testCase.expectedVersion, ver)
		if testCase.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
