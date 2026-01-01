package stackinputproviderconfig

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
	"testing"

	gcpv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp"
	"sigs.k8s.io/yaml"
)

// TestCreateGcpProviderConfigFileFromProto tests that long base64 strings are not folded
func TestCreateGcpProviderConfigFileFromProto(t *testing.T) {
	// Create a realistic GCP service account key JSON
	serviceAccountKey := map[string]interface{}{
		"type":                    "service_account",
		"project_id":              "project-planton-testing",
		"private_key_id":          "abc123def456",
		"private_key":             "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7VJTUt9Us8cKj\nMzEfYyjiWA4R4/M2bS1+fWIcPm15j9fKy4gZFGz5aEg6qPgH6cWFpGzT4F7MSpPO\n-----END PRIVATE KEY-----\n",
		"client_email":            "test@project-planton-testing.iam.gserviceaccount.com",
		"client_id":               "123456789",
		"auth_uri":                "https://accounts.google.com/o/oauth2/auth",
		"token_uri":               "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":    "https://www.googleapis.com/robot/v1/metadata/x509/test%40project-planton-testing.iam.gserviceaccount.com",
	}

	// Marshal to JSON and then base64 encode
	jsonBytes, err := json.Marshal(serviceAccountKey)
	if err != nil {
		t.Fatalf("Failed to marshal service account key: %v", err)
	}
	base64Key := base64.StdEncoding.EncodeToString(jsonBytes)

	// Create config
	gcpConfig := &gcpv1.GcpProviderConfig{
		ServiceAccountKeyBase64: base64Key,
	}

	// Create temp file
	filePath, cleanup, err := createGcpProviderConfigFileFromProto(gcpConfig)
	if err != nil {
		t.Fatalf("Failed to create GCP provider config file: %v", err)
	}
	defer cleanup()

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	// Verify it's valid YAML
	var yamlMap map[string]interface{}
	if err := yaml.Unmarshal(content, &yamlMap); err != nil {
		t.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Verify the base64 key is present and not folded
	extractedKey, ok := yamlMap["serviceAccountKeyBase64"].(string)
	if !ok {
		t.Fatal("serviceAccountKeyBase64 not found in YAML")
	}

	// Remove any whitespace (in case there are newlines)
	extractedKey = strings.ReplaceAll(extractedKey, "\n", "")
	extractedKey = strings.ReplaceAll(extractedKey, " ", "")

	if extractedKey != base64Key {
		t.Errorf("Base64 key mismatch.\nExpected: %s\nGot: %s", base64Key, extractedKey)
	}

	// Convert YAML to JSON to simulate the actual deployment flow
	jsonBytes2, err := yaml.YAMLToJSON(content)
	if err != nil {
		t.Fatalf("Failed to convert YAML to JSON: %v", err)
	}

	// Verify JSON can be parsed
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(jsonBytes2, &jsonMap); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v\nJSON: %s", err, string(jsonBytes2))
	}

	// Verify the key is in the JSON
	jsonKey, ok := jsonMap["serviceAccountKeyBase64"].(string)
	if !ok {
		t.Fatal("serviceAccountKeyBase64 not found in JSON")
	}

	if jsonKey != base64Key {
		t.Errorf("Base64 key mismatch in JSON.\nExpected: %s\nGot: %s", base64Key, jsonKey)
	}

	t.Logf("Test passed! Base64 key length: %d characters", len(base64Key))
}

// TestCreateGcpProviderConfigWithVeryLongKey tests with a very long base64 string (4000+ chars)
func TestCreateGcpProviderConfigWithVeryLongKey(t *testing.T) {
	// Create a very long private key to simulate real-world scenario
	longPrivateKey := "-----BEGIN PRIVATE KEY-----\n"
	for i := 0; i < 50; i++ {
		longPrivateKey += "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7VJTUt9Us8cKj\n"
	}
	longPrivateKey += "-----END PRIVATE KEY-----\n"

	serviceAccountKey := map[string]interface{}{
		"type":                    "service_account",
		"project_id":              "project-planton-testing",
		"private_key_id":          "abc123def456ghi789jkl012mno345pqr678stu901vwx234yz",
		"private_key":             longPrivateKey,
		"client_email":            "test@project-planton-testing.iam.gserviceaccount.com",
		"client_id":               "123456789012345678901",
		"auth_uri":                "https://accounts.google.com/o/oauth2/auth",
		"token_uri":               "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":    "https://www.googleapis.com/robot/v1/metadata/x509/test%40project-planton-testing.iam.gserviceaccount.com",
	}

	// Marshal to JSON and then base64 encode
	jsonBytes, err := json.Marshal(serviceAccountKey)
	if err != nil {
		t.Fatalf("Failed to marshal service account key: %v", err)
	}
	base64Key := base64.StdEncoding.EncodeToString(jsonBytes)

	t.Logf("Testing with base64 key of length: %d characters", len(base64Key))

	// Create config
	gcpConfig := &gcpv1.GcpProviderConfig{
		ServiceAccountKeyBase64: base64Key,
	}

	// Create temp file
	filePath, cleanup, err := createGcpProviderConfigFileFromProto(gcpConfig)
	if err != nil {
		t.Fatalf("Failed to create GCP provider config file: %v", err)
	}
	defer cleanup()

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	t.Logf("YAML content:\n%s", string(content))

	// Convert YAML to JSON - this is where the error was happening
	jsonBytes2, err := yaml.YAMLToJSON(content)
	if err != nil {
		t.Fatalf("Failed to convert YAML to JSON: %v\nYAML content:\n%s", err, string(content))
	}

	// Verify JSON can be parsed - this is where "unexpected token" errors would occur
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(jsonBytes2, &jsonMap); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v\nJSON: %s", err, string(jsonBytes2))
	}

	// Verify the key is correct
	jsonKey, ok := jsonMap["serviceAccountKeyBase64"].(string)
	if !ok {
		t.Fatal("serviceAccountKeyBase64 not found in JSON")
	}

	if jsonKey != base64Key {
		t.Errorf("Base64 key mismatch.\nExpected length: %d\nGot length: %d", len(base64Key), len(jsonKey))
	}

	t.Logf("Test passed! Successfully handled base64 key with %d characters", len(base64Key))
}

