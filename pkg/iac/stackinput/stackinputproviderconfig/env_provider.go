package stackinputproviderconfig

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// BuildProviderConfigOptionsFromEnv creates provider config files from environment variables
// for all supported cloud providers. Returns the provider config options and a cleanup function
// to remove temporary files. The cleanup function should be called when done.
func BuildProviderConfigOptionsFromEnv() (StackInputProviderConfigOptions, func(), error) {
	opts := StackInputProviderConfigOptions{}
	cleanupFuncs := []func(){}

	// AWS Provider Config
	if awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID"); awsAccessKeyID != "" {
		if awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY"); awsSecretKey != "" {
			awsRegion := os.Getenv("AWS_REGION")
			if awsRegion == "" {
				awsRegion = os.Getenv("AWS_DEFAULT_REGION")
			}
			if awsRegion == "" {
				awsRegion = "us-east-1" // default region
			}

			// Try to get account ID from environment variable first
			awsAccountID := os.Getenv("AWS_ACCOUNT_ID")
			if awsAccountID == "" {
				// Try to get account ID from AWS STS using the credentials
				accountID, err := getAwsAccountIDFromSTS(awsAccessKeyID, awsSecretKey, awsRegion, os.Getenv("AWS_SESSION_TOKEN"))
				if err == nil && accountID != "" {
					awsAccountID = accountID
				} else {
					// If we can't get it, use a placeholder (validation may fail, but it's better than nothing)
					// The user should set AWS_ACCOUNT_ID environment variable for best results
					awsAccountID = "000000000000" // placeholder - user should set AWS_ACCOUNT_ID
				}
			}

			file, cleanup, err := createAwsProviderConfigFile(awsAccessKeyID, awsSecretKey, awsRegion, awsAccountID, os.Getenv("AWS_SESSION_TOKEN"))
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
	}

	// GCP Provider Config
	if gcpCredPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); gcpCredPath != "" {
		file, cleanup, err := createGcpProviderConfigFileFromPath(gcpCredPath)
		if err != nil {
			// Cleanup already created files
			for _, fn := range cleanupFuncs {
				fn()
			}
			return opts, nil, errors.Wrap(err, "failed to create GCP provider config")
		}
		opts.GcpProviderConfig = file
		cleanupFuncs = append(cleanupFuncs, cleanup)
	} else if gcpCreds := os.Getenv("GOOGLE_CREDENTIALS"); gcpCreds != "" {
		// GOOGLE_CREDENTIALS contains the JSON directly
		file, cleanup, err := createGcpProviderConfigFileFromJSON(gcpCreds)
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
	armClientID := os.Getenv("ARM_CLIENT_ID")
	armClientSecret := os.Getenv("ARM_CLIENT_SECRET")
	armTenantID := os.Getenv("ARM_TENANT_ID")
	armSubscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if armClientID != "" && armClientSecret != "" && armTenantID != "" && armSubscriptionID != "" {
		file, cleanup, err := createAzureProviderConfigFile(armClientID, armClientSecret, armTenantID, armSubscriptionID)
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

	cleanup := func() {
		for _, fn := range cleanupFuncs {
			fn()
		}
	}

	return opts, cleanup, nil
}

// createAwsProviderConfigFile creates a temporary AWS provider config file
func createAwsProviderConfigFile(accessKeyID, secretAccessKey, region, accountID, sessionToken string) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "aws-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Write AWS provider config YAML
	configContent := fmt.Sprintf("accountId: %s\naccessKeyId: %s\nsecretAccessKey: %s\nregion: %s\n", accountID, accessKeyID, secretAccessKey, region)
	if sessionToken != "" {
		configContent += fmt.Sprintf("sessionToken: %s\n", sessionToken)
	}

	if _, err := tmpFile.WriteString(configContent); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write AWS provider config: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createGcpProviderConfigFileFromPath creates a GCP provider config file from a JSON key file path
func createGcpProviderConfigFileFromPath(credPath string) (string, func(), error) {
	keyBytes, err := os.ReadFile(credPath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read GCP credentials file: %w", err)
	}

	return createGcpProviderConfigFileFromJSON(string(keyBytes))
}

// createGcpProviderConfigFileFromJSON creates a GCP provider config file from JSON credentials
func createGcpProviderConfigFileFromJSON(jsonCreds string) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "gcp-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Base64 encode the JSON credentials
	serviceAccountKeyBase64 := base64.StdEncoding.EncodeToString([]byte(jsonCreds))

	// Write GCP provider config YAML
	configContent := fmt.Sprintf("serviceAccountKeyBase64: %s\n", serviceAccountKeyBase64)

	if _, err := tmpFile.WriteString(configContent); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write GCP provider config: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// createAzureProviderConfigFile creates a temporary Azure provider config file
func createAzureProviderConfigFile(clientID, clientSecret, tenantID, subscriptionID string) (string, func(), error) {
	tmpFile, err := os.CreateTemp("", "azure-provider-config-*.yaml")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	// Write Azure provider config YAML
	configContent := fmt.Sprintf("clientId: %s\nclientSecret: %s\ntenantId: %s\nsubscriptionId: %s\n", clientID, clientSecret, tenantID, subscriptionID)

	if _, err := tmpFile.WriteString(configContent); err != nil {
		tmpFile.Close()
		cleanup()
		return "", nil, fmt.Errorf("failed to write Azure provider config: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), cleanup, nil
}

// getAwsAccountIDFromSTS attempts to get the AWS account ID using AWS STS
func getAwsAccountIDFromSTS(accessKeyID, secretAccessKey, region, sessionToken string) (string, error) {
	// Set up environment for AWS CLI
	env := os.Environ()
	env = append(env, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", accessKeyID))
	env = append(env, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", secretAccessKey))
	if region != "" {
		env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", region))
	}
	if sessionToken != "" {
		env = append(env, fmt.Sprintf("AWS_SESSION_TOKEN=%s", sessionToken))
	}

	// Run aws sts get-caller-identity
	cmd := exec.Command("aws", "sts", "get-caller-identity", "--output", "json")
	cmd.Env = env
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get AWS account ID from STS: %w", err)
	}

	// Parse JSON output
	var result struct {
		Account string `json:"Account"`
	}
	if err := json.Unmarshal(output, &result); err != nil {
		return "", fmt.Errorf("failed to parse STS output: %w", err)
	}

	if result.Account == "" {
		return "", fmt.Errorf("account ID not found in STS response")
	}

	return result.Account, nil
}
