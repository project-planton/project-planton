package tfbackend

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/credential/terraformbackendcredential/v1"
	"os"
	"path/filepath"
)

// WriteBackendFile creates a `backend.tf` file in projectDir using the backend type specified by tofuBackendType.
func WriteBackendFile(projectDir string, tofuBackendType terraformbackendcredentialv1.TerraformBackendType) error {
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

func BackendTypeFromString(backendTypeStr string) terraformbackendcredentialv1.TerraformBackendType {
	switch backendTypeStr {
	case "local":
		return terraformbackendcredentialv1.TerraformBackendType_local
	case "s3":
		return terraformbackendcredentialv1.TerraformBackendType_s3
	case "gcs":
		return terraformbackendcredentialv1.TerraformBackendType_gcs
	case "azurerm":
		return terraformbackendcredentialv1.TerraformBackendType_azurerm
	default:
		return terraformbackendcredentialv1.TerraformBackendType_terraform_backend_type_unspecified
	}
}
