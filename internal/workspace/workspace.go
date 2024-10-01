package workspace

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/ulidgen"
	"os"
	"path/filepath"
	"strings"
)

const (
	ProjectPlantonDir = ".project-planton"
)

// GetWorkspaceDir returns the path of the workspace directory to be used while initializing stack using automation api.
func GetWorkspaceDir(stackFqdn string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get home directory")
	}
	generator := ulidgen.NewGenerator()
	//base directory will always be ${HOME}/.planton-cloud/pulumi
	stackWorkspaceDir := filepath.Join(homeDir, ProjectPlantonDir, "pulumi", stackFqdn, strings.ToLower(generator.Generate().String()))
	if !isDirExists(stackWorkspaceDir) {
		if err := os.MkdirAll(stackWorkspaceDir, os.ModePerm); err != nil {
			return "", errors.Wrapf(err, "failed to ensure %s dir", stackWorkspaceDir)
		}
	}
	return stackWorkspaceDir, nil
}

// isDirExists check if a directory exists
func isDirExists(d string) bool {
	if d == "" {
		return false
	}
	info, err := os.Stat(d)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		//todo: should return an error instead
		return false
	}
	return info.IsDir()
}
