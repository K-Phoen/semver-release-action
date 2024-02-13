#!/bin/sh -l

set -u

if [ -z "${GITHUB_TOKEN}" ]
then
    echo "The GITHUB_TOKEN environment variable is not defined."
    exit 1
fi

RELEASE_BRANCH="$1"
RELEASE_STRATEGY="$2"
NEXT_TAG="$3"
TAG_FORMAT="$4"
GENERATE_RELEASE_NOTES="$5"

echo ::Executing bumper guard ::debug release_branch=${RELEASE_BRANCH},github_event_path=${GITHUB_EVENT_PATH}
/bumper guard "${RELEASE_BRANCH}" "${GITHUB_EVENT_PATH}"
if [ $? -eq 78 ]
then
    echo ::debug ::Guard returned a neutral code, stopping the execution.
    exit 0
fi

if [ -z "${NEXT_TAG}" ]
then
    echo ::debug ::Executing bumper latest-tag github_repository=${GITHUB_REPOSITORY}
    LATEST_TAG=$(/bumper latest-tag "${GITHUB_REPOSITORY}" "${GITHUB_TOKEN}")

    echo ::debug ::Executing bumper increment github_event_path=${GITHUB_EVENT_PATH}
    INCREMENT=$(/bumper increment "${GITHUB_EVENT_PATH}")

    echo ::debug ::Executing bumper semver latest_tag=${LATEST_TAG},increment=${INCREMENT}
    NEXT_TAG=$(/bumper semver "${LATEST_TAG}" "${INCREMENT}" "${TAG_FORMAT}")
fi

echo ::debug ::Executing bumper release github_repository=${GITHUB_REPOSITORY},github_sha=${GITHUB_SHA},next_tag=${NEXT_TAG},strategy=${RELEASE_STRATEGY},gen_release_notes=${GENERATE_RELEASE_NOTES}
/bumper release --strategy "${RELEASE_STRATEGY}" "${GITHUB_REPOSITORY}" "${GITHUB_SHA}" "${NEXT_TAG}" "${GITHUB_TOKEN}" "${GENERATE_RELEASE_NOTES}"

echo "tag=${NEXT_TAG}" >> $GITHUB_OUTPUT
