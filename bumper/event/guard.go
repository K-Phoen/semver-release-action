package event

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/K-Phoen/semver-release-action/bumper/errors"
	"github.com/K-Phoen/semver-release-action/bumper/semver"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func GuardCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "guard [RELEASE_BRANCH] [GH_EVENT_PATH]",
		Args: cobra.ExactArgs(2),
		Run:  executeGuard,
	}
}

func IncrementCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "increment [GH_EVENT_PATH]",
		Args: cobra.ExactArgs(1),
		Run:  executeIncrement,
	}
}

func executeGuard(cmd *cobra.Command, args []string) {
	releaseBranch := args[0]
	event := parseEvent(args[1])

	if event.Action == nil || *event.Action != "closed" {
		errors.Fail("pull request not closed")
	}

	if event.PullRequest.Merged == nil || !*event.PullRequest.Merged {
		errors.Fail("pull request not merged")
	}

	if event.PullRequest.Base == nil || event.PullRequest.Base.Ref == nil {
		errors.Fail("could not determine pull request base branch")
	}

	if *event.PullRequest.Base.Ref != releaseBranch {
		errors.Fail("pull request not merged into the release branch (expected '%s', got '%s'", releaseBranch, *event.PullRequest.Base.Ref)
	}

	validLabelFound := false
	for _, label := range event.PullRequest.Labels {
		if label.Name == nil {
			continue
		}

		_, err := semver.ParseIncrement(*label.Name)
		if err != nil {
			continue
		}

		// we already found one valid label: something is fishy.
		if validLabelFound {
			errors.Fail("several valid semver label found")
		}

		validLabelFound = true
	}

	if !validLabelFound {
		errors.Fail("no valid semver label found")
	}
}

func executeIncrement(cmd *cobra.Command, args []string) {
	event := parseEvent(args[0])

	validLabelFound := false
	increment := semver.IncrementPatch
	for _, label := range event.PullRequest.Labels {
		if label.Name == nil {
			continue
		}

		inc, err := semver.ParseIncrement(*label.Name)
		if err != nil {
			continue
		}

		// we already found one valid label: something is fishy.
		if validLabelFound {
			errors.Fail("several valid semver label found")
		}

		validLabelFound = true
		increment = inc
	}

	if !validLabelFound {
		errors.Fail("no valid semver label found")
	}

	fmt.Print(increment)
}

func parseEvent(filePath string) *github.PullRequestEvent {
	parsed, err := github.ParseWebHook("pull_request", readEvent(filePath))
	errors.AssertNone(err, "could not parse GitHub event: %s", err)

	event, ok := parsed.(*github.PullRequestEvent)
	if !ok {
		errors.Fail("could not parse GitHub event into a PullRequestEvent: %s", err)
	}

	return event
}

func readEvent(filePath string) []byte {
	file, err := os.Open(filePath)
	errors.AssertNone(err, "could not open GitHub event file: %s", err)
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	errors.AssertNone(err, "could not read GitHub event file: %s", err)

	return b
}
