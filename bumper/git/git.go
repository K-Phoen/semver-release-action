package git

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/K-Phoen/semver-release-action/bumper/action"
	"github.com/blang/semver"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func LatestTagCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "latest-tag [REPOSITORY] [GH_TOKEN]",
		Args: cobra.ExactArgs(2),
		Run:  executeLatestTag,
	}
}

func executeLatestTag(cmd *cobra.Command, args []string) {
	repository := args[0]
	githubToken := args[1]

	ctx := context.Background()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	client := github.NewClient(oauth2.NewClient(ctx, tokenSource))

	parts := strings.Split(repository, "/")
	owner := parts[0]
	repo := parts[1]

	refs, response, err := client.Git.ListRefs(ctx, owner, repo, &github.ReferenceListOptions{
		Type: "tag",
	})
	if response.StatusCode == http.StatusNotFound {
		fmt.Print("v0.0.0")
		return
	}
	action.AssertNoError(err, "could not list git refs: %s", err)

	latest := semver.MustParse("0.0.0")
	for _, ref := range refs {
		version, err := semver.ParseTolerant(strings.Replace(*ref.Ref, "refs/tags/", "", 1))
		if err != nil {
			continue
		}

		if version.GT(latest) {
			latest = version
		}
	}

	fmt.Printf("v%s", latest)
}
