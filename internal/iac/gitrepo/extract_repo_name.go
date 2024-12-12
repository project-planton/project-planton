package gitrepo

import (
	"github.com/pkg/errors"
	"strings"
)

// ExtractRepoName takes a repository URL and returns the repository name.
func ExtractRepoName(repoUrl string) (string, error) {
	parts := strings.Split(repoUrl, "/")
	if len(parts) < 1 {
		return "", errors.New("invalid repository URL format, expected format <domain>/<user>/<repo>.git")
	}
	repoNameWithGit := parts[len(parts)-1]
	repoName := strings.TrimSuffix(repoNameWithGit, ".git")
	return repoName, nil
}
