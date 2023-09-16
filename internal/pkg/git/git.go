package git

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/cloud-crafts/semver-release-action/internal/pkg/action"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func LatestTagCommand() *cobra.Command {
	var githubServerBaseHostname string
	var githubServerUploadBaseHostname string

	cmd := &cobra.Command{
		Use:  "latest-tag [REPOSITORY] [GH_TOKEN]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			executeLatestTag(cmd, githubServerBaseHostname, githubServerUploadBaseHostname, args)
		},
	}

	cmd.PersistentFlags().StringVarP(&githubServerBaseHostname, "baseHost", "b", "api.github.com",
		"GitHub Enterprise Server Base URL.")
	cmd.PersistentFlags().StringVarP(&githubServerUploadBaseHostname, "uploadHost", "u", "uploads.github.com",
		"GitHub Enterprise Server Upload URL.")

	return cmd
}

func executeLatestTag(cmd *cobra.Command, githubServerBaseHostname, githubServerUploadBaseHostname string, args []string) {
	repository := args[0]
	githubToken := args[1]

	ctx := context.Background()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})

	client, err := github.NewEnterpriseClient(fmt.Sprintf("https://%s/", githubServerBaseHostname),
		fmt.Sprintf("https://%s/", githubServerUploadBaseHostname), oauth2.NewClient(ctx, tokenSource))
	if err != nil {
		cmd.Print(fmt.Errorf("github server client could not be created: %w", err))
		return
	}

	parts := strings.Split(repository, "/")
	owner := parts[0]
	repo := parts[1]

	refs, response, err := client.Git.ListMatchingRefs(ctx, owner, repo, &github.ReferenceListOptions{
		Ref: "tags",
	})

	if response != nil && response.StatusCode == http.StatusNotFound {
		cmd.Print("v0.0.0")
		return
	}
	action.AssertNoError(cmd, err, "could not list git refs: %s", err)

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

	cmd.Printf("v%s", latest)
}
