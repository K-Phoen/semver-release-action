package release

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/K-Phoen/semver-release-action/internal/pkg/action"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

const releaseTypeNone = "none"
const releaseTypeRelease = "release"
const releaseTypeTag = "tag"

type repository struct {
	owner string
	name  string
	token string
}

type releaseDetails struct {
	version string
	target  string
	autoReleaseNotes bool
}

func Command() *cobra.Command {
	var releaseType string

	cmd := &cobra.Command{
		Use:  "release [REPOSITORY] [TARGET_COMMITISH] [VERSION] [GH_TOKEN] [GENERATE_RELEASE_NOTES]",
		Args: cobra.ExactArgs(5),
		Run: func(cmd *cobra.Command, args []string) {
			execute(cmd, releaseType, args)
		},
	}

	cmd.Flags().StringVarP(&releaseType, "strategy", "s", releaseTypeRelease, "Release strategy")

	return cmd
}

func execute(cmd *cobra.Command, releaseType string, args []string) {
	parts := strings.Split(args[0], "/")
	autoReleaseNotes,_ := strconv.ParseBool(args[4])
	repo := repository{
		owner: parts[0],
		name:  parts[1],
		token: args[3],
	}

	release := releaseDetails{
		version: args[2],
		target:  args[1],
		autoReleaseNotes: autoReleaseNotes,
	}

	ctx := context.Background()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: repo.token})
	client := github.NewClient(oauth2.NewClient(ctx, tokenSource))

	switch releaseType {
	case releaseTypeNone:
		return
	case releaseTypeRelease:
		if err := createGithubRelease(ctx, client, repo, release); err != nil {
			action.AssertNoError(cmd, err, "could not create GitHub release: %s", err)
		}
		return
	case releaseTypeTag:
		if err := createLightweightTag(ctx, client, repo, release); err != nil {
			action.AssertNoError(cmd, err, "could not create lightweight tag: %s", err)
		}
		return
	default:
		action.Fail(cmd, "unknown release strategy: %s", releaseType)
	}
}

func createLightweightTag(ctx context.Context, client *github.Client, repo repository, release releaseDetails) error {
	_, _, err := client.Git.CreateRef(ctx, repo.owner, repo.name, &github.Reference{
		Ref: github.String(fmt.Sprintf("refs/tags/%s", release.version)),
		Object: &github.GitObject{
			SHA: &release.target,
		},
	})

	return err
}

func createGithubRelease(ctx context.Context, client *github.Client, repo repository, release releaseDetails) error {
	_, _, err := client.Repositories.CreateRelease(ctx, repo.owner, repo.name, &github.RepositoryRelease{
		Name:            &release.version,
		TagName:         &release.version,
		TargetCommitish: &release.target,
		Draft:           github.Bool(false),
		Prerelease:      github.Bool(false),
		GenerateReleaseNotes: github.Bool(release.autoReleaseNotes),
	})

	return err
}
