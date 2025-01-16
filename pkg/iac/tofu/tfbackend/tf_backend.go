package tfbackend

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/tofu"
	"os"
	"path/filepath"
)

// WriteBackendFile creates a `backend.tf` file in projectDir using the backend type specified by tofuBackendType.
func WriteBackendFile(projectDir string, tofuBackendType tofu.TofuBackendType) error {
	backendName := tofuBackendType.String()

	// Construct a minimal backend configuration.
	// The user can further configure the backend by editing this file or using -backend-config flags.
	backendContent := fmt.Sprintf(`terraform {
  backend "%s" {}
}
`, backendName)

	backendFilePath := filepath.Join(projectDir, "backend.tf")
	if err := os.WriteFile(backendFilePath, []byte(backendContent), 0644); err != nil {
		return errors.Wrap(err, "failed to write backend file")
	}

	return nil
}

func BackendTypeFromString(backendTypeStr string) tofu.TofuBackendType {
	switch backendTypeStr {
	case "local":
		return tofu.TofuBackendType_local
	case "s3":
		return tofu.TofuBackendType_s3
	case "remote":
		return tofu.TofuBackendType_remote
	case "gcs":
		return tofu.TofuBackendType_gcs
	case "azurerm":
		return tofu.TofuBackendType_azurerm
	case "consul":
		return tofu.TofuBackendType_consul
	case "http":
		return tofu.TofuBackendType_http
	case "etcdv3":
		return tofu.TofuBackendType_etcdv3
	case "manta":
		return tofu.TofuBackendType_manta
	case "swift":
		return tofu.TofuBackendType_swift
	case "artifactory":
		return tofu.TofuBackendType_artifactory
	case "oss":
		return tofu.TofuBackendType_oss
	default:
		return tofu.TofuBackendType_tofu_backend_type_unspecified
	}
}
