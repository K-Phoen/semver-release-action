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

INCREMENT=$(/bumper increment "$GITHUB_EVENT_PATH")

LATEST_TAG_REF=$(git rev-list --tags --max-count=1)

if [ -z "$LATEST_TAG_REF" ]
then
    LATEST_TAG_NAME="v0.0.0"
else
    LATEST_TAG_NAME=$(git describe --tags "$LATEST_TAG_REF")
fi

NEXT_TAG=$(/bumper semver "$LATEST_TAG_NAME" $INCREMENT)

/bumper release "$GITHUB_REPOSITORY" "$GITHUB_SHA" "$NEXT_TAG" "$GITHUB_TOKEN"

echo ::set-output name=tag::$LATEST_TAG_NAME