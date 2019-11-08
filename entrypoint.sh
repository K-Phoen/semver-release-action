#!/bin/sh -l

set -e
set -u

if [ -z "${GITHUB_TOKEN}" ]
then
    echo "The GITHUB_TOKEN environment variable is not defined."
    exit 1
fi

RELEASE_BRANCH="$1"
RELEASE_STRATEGY="$2"
NEXT_TAG="$3"
DEFAULT_VERSION="$4"


echo ::Executing bumper guard ::debug release_branch=${RELEASE_BRANCH},github_event_path=${GITHUB_EVENT_PATH},default_version=${DEFAULT_VERSION}
/bumper guard "${RELEASE_BRANCH}" "${GITHUB_EVENT_PATH}" "${DEFAULT_VERSION}"
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
    INCREMENT=$(/bumper increment "${GITHUB_EVENT_PATH}" "${DEFAULT_VERSION}")

    echo ::debug ::Executing bumper semver latest_tag=${LATEST_TAG},increment=${INCREMENT}
    NEXT_TAG=$(/bumper semver "${LATEST_TAG}" "${INCREMENT}")
fi

echo ::debug ::Executing bumper release github_repository=${GITHUB_REPOSITORY},github_sha=${GITHUB_SHA},next_tag=${NEXT_TAG},strategy=${RELEASE_STRATEGY}
/bumper release --strategy "${RELEASE_STRATEGY}" "${GITHUB_REPOSITORY}" "${GITHUB_SHA}" "${NEXT_TAG}" "${GITHUB_TOKEN}"

echo ::set-output name=tag::${NEXT_TAG}
