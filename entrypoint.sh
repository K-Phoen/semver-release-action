#!/bin/sh -l

set -e
set -u

if [ -z "${GITHUB_TOKEN}" ]
then
    echo "The GITHUB_TOKEN environment variable is not defined."
    exit 1
fi

RELEASE_BRANCH="$1"

echo ::Executing bumper guard ::debug release_branch=${RELEASE_BRANCH},github_event_path=${GITHUB_EVENT_PATH}
/bumper guard "${RELEASE_BRANCH}" "${GITHUB_EVENT_PATH}"

echo ::debug ::Executing bumper latest-tag github_repository=${GITHUB_REPOSITORY}
LATEST_TAG=$(/bumper latest-tag "${GITHUB_REPOSITORY}" "${GITHUB_TOKEN}")

echo ::debug ::Executing bumper increment github_event_path=${GITHUB_EVENT_PATH}
INCREMENT=$(/bumper increment "${GITHUB_EVENT_PATH}")

echo ::debug ::Executing bumper semver latest_tag=${LATEST_TAG},increment=${INCREMENT}
NEXT_TAG=$(/bumper semver "${LATEST_TAG}" "${INCREMENT}")

echo ::debug ::Executing bumper release github_repository=${GITHUB_REPOSITORY},github_sha=${GITHUB_SHA},next_tag=${NEXT_TAG}
/bumper release "${GITHUB_REPOSITORY}" "${GITHUB_SHA}" "${NEXT_TAG}" "${GITHUB_TOKEN}"

echo ::set-output name=tag::${NEXT_TAG}