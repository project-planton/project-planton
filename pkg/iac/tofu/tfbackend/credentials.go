package tfbackend

import (
	"encoding/base64"
	"github.com/pkg/errors"
	terraformbackendcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/terraformbackendcredential/v1"
)

func GetCredentialEnvVars(backendCredentialSpec *terraformbackendcredentialv1.TerraformBackendCredentialSpec) ([]string, error) {
	credentialEnvVars := map[string]string{}
	switch backendCredentialSpec.Type {
	case terraformbackendcredentialv1.TerraformBackendType_gcs:
		if backendCredentialSpec.Gcs == nil {
			return nil, errors.New("GCS credential spec is nil")
		}
		serviceAccountKey, err := base64.StdEncoding.DecodeString(backendCredentialSpec.Gcs.ServiceAccountKeyBase64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode service account key")
		}
		credentialEnvVars["GOOGLE_CREDENTIALS"] = string(serviceAccountKey)
	case terraformbackendcredentialv1.TerraformBackendType_s3:
		if backendCredentialSpec.S3 == nil {
			return nil, errors.New("S3 credential spec is nil")
		}
		credentialEnvVars["AWS_REGION"] = backendCredentialSpec.S3.Region
		credentialEnvVars["AWS_ACCESS_KEY_ID"] = backendCredentialSpec.S3.AwsSecretAccessKey
		credentialEnvVars["AWS_SECRET_ACCESS_KEY"] = backendCredentialSpec.S3.AwsSecretAccessKey
	case terraformbackendcredentialv1.TerraformBackendType_azurerm:
		if backendCredentialSpec.Azurerm == nil {
			return nil, errors.New("Azurerm credential spec is nil")
		}
		credentialEnvVars["ARM_CLIENT_ID"] = backendCredentialSpec.Azurerm.ClientId
		credentialEnvVars["ARM_CLIENT_SECRET"] = backendCredentialSpec.Azurerm.ClientSecret
		credentialEnvVars["ARM_TENANT_ID"] = backendCredentialSpec.Azurerm.TenantId
		credentialEnvVars["ARM_SUBSCRIPTION_ID"] = backendCredentialSpec.Azurerm.SubscriptionId
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
