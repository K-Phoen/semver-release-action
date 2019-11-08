package event

import (
	"io/ioutil"
	"os"

	"github.com/K-Phoen/semver-release-action/internal/pkg/action"
	"github.com/K-Phoen/semver-release-action/internal/pkg/semver"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func GuardCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "guard [RELEASE_BRANCH] [GH_EVENT_PATH] [DEFAULT_VERSION]",
		Args: cobra.ExactArgs(3),
		Run:  executeGuard,
	}
}

func IncrementCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "increment [GH_EVENT_PATH] [DEFAULT_VERSION]",
		Args: cobra.ExactArgs(2),
		Run:  executeIncrement,
	}
}

func executeGuard(cmd *cobra.Command, args []string) {
	releaseBranch := args[0]
	event := parseEvent(cmd, args[1])
	defaultVersion := args[2]

	if event.Action == nil || *event.Action != "closed" {
		action.Skip(cmd, "pull request not closed")
	}

	if event.PullRequest.Merged == nil || !*event.PullRequest.Merged {
		action.Skip(cmd, "pull request not merged")
	}

	if event.PullRequest.Base == nil || event.PullRequest.Base.Ref == nil {
		action.Fail(cmd, "could not determine pull request base branch")
	}

	if *event.PullRequest.Base.Ref != releaseBranch {
		action.Skip(
			cmd,
			"pull request not merged into the release branch (expected '%s', got '%s'",
			releaseBranch,
			*event.PullRequest.Base.Ref,
		)
	}

	_, incrementFound := extractIncrement(cmd, event.PullRequest)
	if !incrementFound && defaultVersion == "" {
		action.Skip(cmd, "no valid semver label found or no default version added")
	}
}

func executeIncrement(cmd *cobra.Command, args []string) {
	event := parseEvent(cmd, args[0])
	defaultVersion := args[1]

	increment, found := extractIncrement(cmd, event.PullRequest)
	if !found {
		if defaultVersion == "" {
			action.Fail(cmd, "no valid semver label found")
		} else {
			increment, _ = semver.ParseIncrement(defaultVersion)
		}
	}

	cmd.Print(increment)
}

func extractIncrement(cmd *cobra.Command, pr *github.PullRequest) (semver.Increment, bool) {
	validLabelFound := false
	increment := semver.IncrementPatch
	for _, label := range pr.Labels {
		if label.Name == nil {
			continue
		}

		inc, err := semver.ParseIncrement(*label.Name)
		if err != nil {
			continue
		}

		// we already found one valid label: something is fishy.
		if validLabelFound {
			action.Fail(cmd, "several valid semver label found")
		}

		validLabelFound = true
		increment = inc
	}

	return increment, validLabelFound
}

func parseEvent(cmd *cobra.Command, filePath string) *github.PullRequestEvent {
	parsed, err := github.ParseWebHook("pull_request", readEvent(cmd, filePath))
	action.AssertNoError(cmd, err, "could not parse GitHub event: %s", err)

	event, ok := parsed.(*github.PullRequestEvent)
	if !ok {
		action.Fail(cmd, "could not parse GitHub event into a PullRequestEvent: %s", err)
	}

	return event
}

func readEvent(cmd *cobra.Command, filePath string) []byte {
	file, err := os.Open(filePath)
	action.AssertNoError(cmd, err, "could not open GitHub event file: %s", err)
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	action.AssertNoError(cmd, err, "could not read GitHub event file: %s", err)

	return b
}
