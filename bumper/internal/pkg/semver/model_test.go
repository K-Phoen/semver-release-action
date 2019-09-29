package semver

import (
	"testing"

	"github.com/K-Phoen/semver-release-action/bumper/internal/pkg"

	"github.com/stretchr/testify/require"
)

func TestVersionBump(t *testing.T) {
	cases := []struct {
		name string

		version   pkg.Version
		increment pkg.Increment

		expectedVersion string
	}{
		{
			name:            "patch bump",
			version:         pkg.Version{major: 1, minor: 2, patch: 3},
			increment:       pkg.IncrementPatch,
			expectedVersion: "v1.2.4",
		},
		{
			name:            "minor bump",
			version:         pkg.Version{major: 1, minor: 2, patch: 3},
			increment:       pkg.IncrementMinor,
			expectedVersion: "v1.3.0",
		},
		{
			name:            "major bump",
			version:         pkg.Version{major: 1, minor: 2, patch: 3},
			increment:       pkg.IncrementMajor,
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

		expectedIncrement pkg.Increment
		expectedError     error
	}{
		{
			input:             "patch",
			expectedIncrement: pkg.IncrementPatch,
			expectedError:     nil,
		},
		{
			input:             "PaTcH",
			expectedIncrement: pkg.IncrementPatch,
			expectedError:     nil,
		},
		{
			input:             "minor",
			expectedIncrement: pkg.IncrementMinor,
			expectedError:     nil,
		},
		{
			input:             "mInOr",
			expectedIncrement: pkg.IncrementMinor,
			expectedError:     nil,
		},
		{
			input:             "major",
			expectedIncrement: pkg.IncrementMajor,
			expectedError:     nil,
		},
		{
			input:             "mAjOR",
			expectedIncrement: pkg.IncrementMajor,
			expectedError:     nil,
		},

		{
			input:             "micro",
			expectedIncrement: pkg.IncrementPatch,
			expectedError:     pkg.ErrInvalidIncrement,
		},
	}

	for _, testCase := range cases {
		inc, err := pkg.ParseIncrement(testCase.input)

		require.Equal(t, testCase.expectedIncrement, inc)
		require.Equal(t, testCase.expectedError, err)
	}
}

func TestParseVersion(t *testing.T) {
	cases := []struct {
		input string

		expectedVersion pkg.Version
		expectError     bool
	}{
		{
			input:           "1.2.3",
			expectedVersion: pkg.Version{major: 1, minor: 2, patch: 3},
			expectError:     false,
		},
		{
			input:           "v1.2.3",
			expectedVersion: pkg.Version{major: 1, minor: 2, patch: 3},
			expectError:     false,
		},
		{
			input:           "v1.2",
			expectedVersion: pkg.Version{major: 1, minor: 2, patch: 0},
			expectError:     false,
		},
		{
			input:           "vlala.2.3",
			expectedVersion: pkg.Version{},
			expectError:     true,
		},
		{
			input:           "v1.lala.3",
			expectedVersion: pkg.Version{},
			expectError:     true,
		},
		{
			input:           "v1.2.lala",
			expectedVersion: pkg.Version{},
			expectError:     true,
		},
	}

	for _, testCase := range cases {
		ver, err := pkg.ParseVersion(testCase.input)

		require.Equal(t, testCase.expectedVersion, ver)
		if testCase.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
