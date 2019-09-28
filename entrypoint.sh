#!/bin/sh -l

set -e
set -u

if [ -z "$GITHUB_TOKEN" ]
then
    echo "The GITHUB_TOKEN environment variable is not defined."
    exit 1
fi

RELEASE_BRANCH="$1"

/bumper guard "$RELEASE_BRANCH" "$GITHUB_EVENT_PATH"

LATEST_TAG=$(/bumper latest-tag "$GITHUB_REPOSITORY" "$GITHUB_TOKEN")

INCREMENT=$(/bumper increment "$GITHUB_EVENT_PATH")
NEXT_TAG=$(/bumper semver "$LATEST_TAG" $INCREMENT)

/bumper release "$GITHUB_REPOSITORY" "$GITHUB_SHA" "$NEXT_TAG" "$GITHUB_TOKEN"

echo ::set-output name=tag::$NEXT_TAG