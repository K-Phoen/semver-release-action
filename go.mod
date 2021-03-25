module github.com/K-Phoen/semver-release-action

go 1.13

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-github/v28 v28.1.1
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
)

replace github.com/K-Phoen/semver-release-action => github.com/krogon/semver-release-action v1.3.2-0.20200201232446-f3f0333ed926
