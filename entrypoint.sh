#!/bin/sh -l

set -e
set -u

if [ -z "$GITHUB_TOKEN" ]
then
    echo "The GITHUB_TOKEN environment variable is not defined."
    exit 1
fi

RELEASE_BRANCH="$1"
BASE_REF=$(jq -j '.pull_request?.base.ref' $GITHUB_EVENT_PATH)

if [ "$RELEASE_BRANCH" != "$BASE_REF" ]
then
    echo "Release branch is '$RELEASE_BRANCH', got '$BASE_REF'. Nothing to do."
    exit 0
fi

if ! jq -e '.pull_request?.merged' $GITHUB_EVENT_PATH >/dev/null
then
    echo "Pull request closed, but not merged. Nothing to do."
    exit 0
fi

if jq -e '.pull_request?.labels | any(.name == "patch")' $GITHUB_EVENT_PATH >/dev/null
then
    INCREMENT="patch"
elif jq -e '.pull_request?.labels | any(.name == "minor")' $GITHUB_EVENT_PATH >/dev/null
then
    INCREMENT="minor"
elif jq -e '.pull_request?.labels | any(.name == "major")' $GITHUB_EVENT_PATH >/dev/null
then
    INCREMENT="major"
else
    echo "No semver label found. Nothing to do."
    exit 0
fi

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