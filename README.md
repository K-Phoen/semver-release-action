# Semver Release Github Action

Automatically create Semver compliant releases based on PR labels.

## Inputs

### `release_branch`

**Required** Branch to tag. Default `"master"`.

## Outputs

### `tag`

The newly created tag.

## Example usage

```yaml
# .github/workflows/release.yml
name: Release

on:
  pull_request:
    types: closed

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Tag
        uses: K-Phoen/semver-release-action@master
        with:
          release_branch: master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

```