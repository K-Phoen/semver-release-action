# Internal Use: Hnry-Semver-Release

## How we can use this internally

Since this is being used via a private registry. This action is being executed manually so that docker authentication can occur prior to image pull. This is how you manually implement this (see [example implementation](https://github.com/HnryNZ/hnry-rails/blob/master/.github/workflows/auto-releaser.yml#L36))

```yaml
generate_release:
  name: Generate Release
  runs-on: ubuntu-latest
  container:
    image: hnrynz/semver-release-action:latest
    credentials:
      username: ${{ secrets.DOCKER_USER }}
      password: ${{ secrets.DOCKER_ACCESS_TOKEN }}
  needs: [suspend_release]
  if: github.event.pull_request.merged

  steps:
    - name: Run Action
      run: |
        # manual execution of the action
        chmod +x /entrypoint.sh
        export GITHUB_TOKEN=${{ secrets.GITHUB_TOKEN }}
        /entrypoint.sh "master" "release" "" "%major%.%minor%.%patch%"
```

This is what you need:

- Docker Credentials: `${{ secrets.DOCKER_USER }}` & `${{ secrets.DOCKER_ACCESS_TOKEN }}`
- The trunk branch name needs to be consistent with script call `/entrypoint.sh "master" ...`

## How to roll in your changes to main

What you need:

1. Merge your changes to `main`
2. Build the image `docker build -t hnrynz/semver-release-action:latest .`
3. Push the image `docker push hnrynz/semver-release-action:latest`
4. Since this is a `:latest` tag, this should be **automatically** rolled in to any workflow pulling this image 👌


<!-- # Semver Release Github Action ![](https://github.com/K-Phoen/semver-release-action/workflows/CI/badge.svg)

Automatically create [SemVer](https://semver.org/) compliant releases based on
PR labels.

Assuming that a PR is tagged with a "*semver-compliant*" label (*patch*, *minor* or *major*),
then this action can create a tag and a GitHub release when it is merged.

**Note:** to determine the base tag for the increment, this action will try to
find the most recent tag complying to [SemVer](https://semver.org/). No
additional setup is required.

## Inputs

### `release_branch`

**Required** Branch to tag. Default `"master"`.

### `release_strategy`

**Required** Release strategy. Default `"release"` (`release`: creates a GitHub
release ; `tag`: creates a lightweight tag ; `none`: computes the next
[SemVer](https://semver.org/) version but does not create a release or tag).

### `tag_format`

**Optional** Format used to create tags. Default `"v%major%.%minor%.%patch%"`.

### `tag`

**Optional** Tag to use. If left undefined, it will be computed using the tags
already present in the repository.

## Outputs

### `tag`

The newly created tag.

## Example usage

```yaml
# .github/workflows/release.yml
name: Release

on:
  pull_request:
    types: [closed]

jobs:
  build:
    runs-on: ubuntu-latest

    if: github.event.pull_request.merged

    steps:
      - name: Tag
        uses: K-Phoen/semver-release-action@master
        with:
          release_branch: master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

```

## License

This library is under the [MIT](LICENSE.md) license. -->
