package utils

import "fmt"

const (
	defaultBaseHostName  = "api.github.com"
	defaultUploadBaseURL = "uploads.github.com"
)

// GetOrDefaultGitHubHostnames validates and returns the GitHub Server hostname values. If there is no hostname specified
// for baseHostName and for uploadHostName, the default hostnames for `github.com` will be used.
func GetOrDefaultGitHubHostnames(baseHostName, uploadHostName string) (string, string, error) {
	if baseHostName != "" && uploadHostName != "" {
		return baseHostName, uploadHostName, nil
	}
	if baseHostName == "" && uploadHostName == "" {
		return defaultBaseHostName, defaultUploadBaseURL, nil
	}
	if baseHostName != "" && uploadHostName == "" {
		return baseHostName, baseHostName, nil
	}
	return "", "", fmt.Errorf("invalid input for base hostname (%s) and upload hostname (%s)", baseHostName, uploadHostName)
}
