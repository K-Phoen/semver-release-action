package release

import (
	"context"
	"strings"

	"github.com/K-Phoen/semver-release-action/internal/pkg/action"
	"github.com/google/go-github/v28/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:  "release [REPOSITORY] [TARGET_COMMITISH] [VERSION] [GH_TOKEN]",
		Args: cobra.ExactArgs(4),
		Run:  execute,
	}
}

func execute(cmd *cobra.Command, args []string) {
	repository := args[0]
	targetCommitish := args[1]
	version := args[2]
	githubToken := args[3]

	ctx := context.Background()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	client := github.NewClient(oauth2.NewClient(ctx, tokenSource))

	parts := strings.Split(repository, "/")
	owner := parts[0]
	repo := parts[1]

	_, _, err := client.Repositories.CreateRelease(ctx, owner, repo, &github.RepositoryRelease{
		Name:            &version,
		TagName:         &version,
		TargetCommitish: &targetCommitish,
		Draft:           github.Bool(false),
		Prerelease:      github.Bool(false),
	})
	action.AssertNoError(cmd, err, "could not create GitHub release: %s", err)
}
