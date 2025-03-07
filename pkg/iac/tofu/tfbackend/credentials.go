package tfbackend

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
)

func GetCredentialEnvVars(terraformBackend *terraform.TerraformBackend) ([]string, error) {
	credentialEnvVars := map[string]string{}
	switch terraformBackend.Type {
	case terraform.TerraformBackendType_gcs:
		if terraformBackend.Gcs == nil {
			return nil, errors.New("GCS credential spec is nil")
		}
		serviceAccountKey, err := base64.StdEncoding.DecodeString(terraformBackend.Gcs.ServiceAccountKeyBase64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode service account key")
		}
		credentialEnvVars["GOOGLE_CREDENTIALS"] = string(serviceAccountKey)
	case terraform.TerraformBackendType_s3:
		if terraformBackend.S3 == nil {
			return nil, errors.New("S3 credential spec is nil")
		}
		credentialEnvVars["AWS_REGION"] = terraformBackend.S3.Region
		credentialEnvVars["AWS_ACCESS_KEY_ID"] = terraformBackend.S3.AwsAccessKeyId
		credentialEnvVars["AWS_SECRET_ACCESS_KEY"] = terraformBackend.S3.AwsSecretAccessKey
	case terraform.TerraformBackendType_azurerm:
		if terraformBackend.Azurerm == nil {
			return nil, errors.New("Azurerm credential spec is nil")
		}
		credentialEnvVars["ARM_CLIENT_ID"] = terraformBackend.Azurerm.ClientId
		credentialEnvVars["ARM_CLIENT_SECRET"] = terraformBackend.Azurerm.ClientSecret
		credentialEnvVars["ARM_TENANT_ID"] = terraformBackend.Azurerm.TenantId
		credentialEnvVars["ARM_SUBSCRIPTION_ID"] = terraformBackend.Azurerm.SubscriptionId
	}
	return mapToSlice(credentialEnvVars), nil
}

// mapToSlice converts a map of string to string into a slice of string slices by joining key-value pairs with an equals sign.
func mapToSlice(inputMap map[string]string) []string {
	var result []string
	for key, value := range inputMap {
		result = append(result, key+"="+value)
	}
	return result
}
