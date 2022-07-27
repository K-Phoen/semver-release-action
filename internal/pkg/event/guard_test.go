package event

import (
	"testing"

	"github.com/K-Phoen/semver-release-action/internal/pkg/semver"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestExtractLabelNoLabelsFound(t *testing.T) {
	cases := []struct {
		cmd cobra.Command
		pullRequest github.PullRequest
	}{
		{
			cmd:         cobra.Command{},
			pullRequest: github.PullRequest{},
		},
	}

	for _, testCase := range cases {
		increment, success := extractIncrement(&testCase.cmd, &testCase.pullRequest)

		require.Equal(t, increment, semver.Increment("patch"))
		require.Equal(t, success, true)
	}
}
