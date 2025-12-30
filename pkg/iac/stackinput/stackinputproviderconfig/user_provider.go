package stackinputproviderconfig

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	atlasv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/atlas"
	auth0v1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/auth0"
	awsv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/aws"
	azurev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/azure"
	cloudflarev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/cloudflare"
	confluentv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/confluent"
	gcpv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp"
	kubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	snowflakev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/snowflake"
	"gopkg.in/yaml.v3"
)

// BuildProviderConfigOptionsFromUserCredentials converts user-provided credentials
// from API request to temporary credential files, matching the CLI pattern.
// This function creates temporary YAML files in the same format as CLI credential files.
func BuildProviderConfigOptionsFromUserCredentials(
	awsConfig *awsv1.AwsProviderConfig,
	gcpConfig *gcpv1.GcpProviderConfig,
	azureConfig *azurev1.AzureProviderConfig,
	atlasConfig *atlasv1.AtlasProviderConfig,
	auth0Config *auth0v1.Auth0ProviderConfig,
	cloudflareConfig *cloudflarev1.CloudflareProviderConfig,
	confluentConfig *confluentv1.ConfluentProviderConfig,
	snowflakeConfig *snowflakev1.SnowflakeProviderConfig,
	kubernetesConfig *kubernetesv1.KubernetesProviderConfig,
) (StackInputProviderConfigOptions, func(), error) {
	opts := StackInputProviderConfigOptions{}
	cleanupFuncs := []func(){}

	// AWS Provider Config
	if awsConfig != nil {
		file, cleanup, err := createAwsProviderConfigFileFromProto(awsConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create AWS provider config")
		}
		opts.AwsProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// GCP Provider Config
	if gcpConfig != nil {
		file, cleanup, err := createGcpProviderConfigFileFromProto(gcpConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create GCP provider config")
		}
		opts.GcpProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// Azure Provider Config
	if azureConfig != nil {
		file, cleanup, err := createAzureProviderConfigFileFromProto(azureConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create Azure provider config")
		}
		opts.AzureProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// Atlas Provider Config
	if atlasConfig != nil {
		file, cleanup, err := createAtlasProviderConfigFileFromProto(atlasConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create Atlas provider config")
		}
		opts.AtlasProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// Auth0 Provider Config
	if auth0Config != nil {
		file, cleanup, err := createAuth0ProviderConfigFileFromProto(auth0Config)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create Auth0 provider config")
		}
		opts.Auth0ProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// Cloudflare Provider Config
	if cloudflareConfig != nil {
		file, cleanup, err := createCloudflareProviderConfigFileFromProto(cloudflareConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create Cloudflare provider config")
		}
		opts.CloudflareProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// Confluent Provider Config
	if confluentConfig != nil {
		file, cleanup, err := createConfluentProviderConfigFileFromProto(confluentConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create Confluent provider config")
		}
		opts.ConfluentProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// Snowflake Provider Config
	if snowflakeConfig != nil {
		file, cleanup, err := createSnowflakeProviderConfigFileFromProto(snowflakeConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create Snowflake provider config")
		}
		opts.SnowflakeProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	// Kubernetes Provider Config
	if kubernetesConfig != nil {
		file, cleanup, err := createKubernetesProviderConfigFileFromProto(kubernetesConfig)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create Kubernetes provider config")
		}
		opts.KubernetesProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	cleanup := func() {
		for _, fn := range cleanupFuncs {
			fn()
		}
	}

	return opts, cleanup, nil
}

// createAwsProviderConfigFileFromProto creates a temporary AWS provider config file from proto message
func createAwsProviderConfigFileFromProto(awsConfig *awsv1.AwsProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "aws-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching proto JSON field names (snake_case)
	// protojson expects snake_case keys based on JSON tags in the proto
	awsCredMap := map[string]interface{}{
		"account_id":        awsConfig.AccountId,
		"access_key_id":     awsConfig.AccessKeyId,
		"secret_access_key": awsConfig.SecretAccessKey,
	}

	// Region is required for Pulumi AWS provider - default to us-east-1 if not provided
	region := "us-east-1" // default region
	if awsConfig.Region != "" {
		region = awsConfig.Region
	}
	awsCredMap["region"] = region

	if awsConfig.SessionToken != "" {
		awsCredMap["session_token"] = awsConfig.SessionToken
	}

	yamlBytes, err := yaml.Marshal(awsCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal AWS credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write AWS credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createGcpProviderConfigFileFromProto creates a GCP provider config file from proto message
func createGcpProviderConfigFileFromProto(gcpConfig *gcpv1.GcpProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "gcp-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	gcpCredMap := map[string]interface{}{
		"serviceAccountKeyBase64": gcpConfig.ServiceAccountKeyBase64,
	}

	yamlBytes, err := yaml.Marshal(gcpCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal GCP credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write GCP credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createAzureProviderConfigFileFromProto creates an Azure provider config file from proto message
func createAzureProviderConfigFileFromProto(azureConfig *azurev1.AzureProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "azure-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	azureCredMap := map[string]interface{}{
		"clientId":       azureConfig.ClientId,
		"clientSecret":   azureConfig.ClientSecret,
		"tenantId":       azureConfig.TenantId,
		"subscriptionId": azureConfig.SubscriptionId,
	}

	yamlBytes, err := yaml.Marshal(azureCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal Azure credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Azure credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createAtlasProviderConfigFileFromProto creates an Atlas provider config file from proto message
func createAtlasProviderConfigFileFromProto(atlasConfig *atlasv1.AtlasProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "atlas-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	atlasCredMap := map[string]interface{}{
		"publicKey":  atlasConfig.PublicKey,
		"privateKey": atlasConfig.PrivateKey,
	}

	yamlBytes, err := yaml.Marshal(atlasCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal Atlas credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Atlas credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createAuth0ProviderConfigFileFromProto creates an Auth0 provider config file from proto message
func createAuth0ProviderConfigFileFromProto(auth0Config *auth0v1.Auth0ProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "auth0-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	auth0CredMap := map[string]interface{}{
		"domain":       auth0Config.Domain,
		"clientId":     auth0Config.ClientId,
		"clientSecret": auth0Config.ClientSecret,
	}

	yamlBytes, err := yaml.Marshal(auth0CredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal Auth0 credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Auth0 credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createCloudflareProviderConfigFileFromProto creates a Cloudflare provider config file from proto message
func createCloudflareProviderConfigFileFromProto(cloudflareConfig *cloudflarev1.CloudflareProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "cloudflare-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	cloudflareCredMap := map[string]interface{}{
		"authScheme": int(cloudflareConfig.AuthScheme),
	}

	if cloudflareConfig.ApiToken != "" {
		cloudflareCredMap["apiToken"] = cloudflareConfig.ApiToken
	}
	if cloudflareConfig.ApiKey != "" {
		cloudflareCredMap["apiKey"] = cloudflareConfig.ApiKey
	}
	if cloudflareConfig.Email != "" {
		cloudflareCredMap["email"] = cloudflareConfig.Email
	}

	yamlBytes, err := yaml.Marshal(cloudflareCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal Cloudflare credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Cloudflare credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createConfluentProviderConfigFileFromProto creates a Confluent provider config file from proto message
func createConfluentProviderConfigFileFromProto(confluentConfig *confluentv1.ConfluentProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "confluent-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	confluentCredMap := map[string]interface{}{
		"apiKey":    confluentConfig.ApiKey,
		"apiSecret": confluentConfig.ApiSecret,
	}

	yamlBytes, err := yaml.Marshal(confluentCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal Confluent credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Confluent credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createSnowflakeProviderConfigFileFromProto creates a Snowflake provider config file from proto message
func createSnowflakeProviderConfigFileFromProto(snowflakeConfig *snowflakev1.SnowflakeProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "snowflake-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	snowflakeCredMap := map[string]interface{}{
		"account":  snowflakeConfig.Account,
		"region":   snowflakeConfig.Region,
		"username": snowflakeConfig.Username,
		"password": snowflakeConfig.Password,
	}

	yamlBytes, err := yaml.Marshal(snowflakeCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal Snowflake credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Snowflake credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createKubernetesProviderConfigFileFromProto creates a Kubernetes provider config file from proto message
func createKubernetesProviderConfigFileFromProto(kubernetesConfig *kubernetesv1.KubernetesProviderConfig) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "kubernetes-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Build YAML content matching CLI credential file format
	kubernetesCredMap := map[string]interface{}{
		"provider": int(kubernetesConfig.Provider),
	}

	if kubernetesConfig.GcpGke != nil {
		kubernetesCredMap["gcpGke"] = map[string]interface{}{
			"clusterEndpoint":         kubernetesConfig.GcpGke.ClusterEndpoint,
			"clusterCaData":           kubernetesConfig.GcpGke.ClusterCaData,
			"serviceAccountKeyBase64": kubernetesConfig.GcpGke.ServiceAccountKeyBase64,
		}
	}

	if kubernetesConfig.DigitalOceanDoks != nil {
		kubernetesCredMap["digitalOceanDoks"] = map[string]interface{}{
			"kubeConfig": kubernetesConfig.DigitalOceanDoks.KubeConfig,
		}
	}

	yamlBytes, err := yaml.Marshal(kubernetesCredMap)
	if err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to marshal Kubernetes credentials to YAML: %w", err)
	}

	if _, err := tmpFile.Write(yamlBytes); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Kubernetes credentials: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}
