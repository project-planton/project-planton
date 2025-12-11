package root

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"github.com/project-planton/project-planton/app/backend/apis/gen/go/proto/backendv1connect"
	"github.com/spf13/cobra"
)

var CredentialGetCmd = &cobra.Command{
	Use:   "credential:get",
	Short: "get a credential by ID",
	Long:  "Retrieve detailed information about a credential by providing its unique ID.",
	Run:   credentialGetHandler,
}

func init() {
	CredentialGetCmd.Flags().StringP("id", "i", "", "unique identifier of the credential (required)")
	CredentialGetCmd.MarkFlagRequired("id")
}

func credentialGetHandler(cmd *cobra.Command, args []string) {
	// Get credential ID
	id, _ := cmd.Flags().GetString("id")
	if id == "" {
		fmt.Println("Error: --id flag is required. Provide the credential ID")
		fmt.Println("Usage: project-planton credential:get --id=<credential-id>")
		os.Exit(1)
	}

	// Get backend URL from configuration
	backendURL, err := GetBackendURL()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Create Connect-RPC client
	client := backendv1connect.NewCredentialServiceClient(
		http.DefaultClient,
		backendURL,
	)

	// Prepare request
	req := &backendv1.GetCredentialRequest{
		Id: id,
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Make the API call
	resp, err := client.GetCredential(ctx, connect.NewRequest(req))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeUnavailable {
			fmt.Printf("Error: Cannot connect to backend service at %s. Please check:\n", backendURL)
			fmt.Printf("  1. The backend service is running\n")
			fmt.Printf("  2. The backend URL is correct\n")
			fmt.Printf("  3. Network connectivity\n")
			os.Exit(1)
		}
		if connect.CodeOf(err) == connect.CodeNotFound {
			fmt.Printf("Error: Credential with ID '%s' not found\n", id)
			os.Exit(1)
		}
		fmt.Printf("Error getting credential: %v\n", err)
		os.Exit(1)
	}

	cred := resp.Msg.Credential

	// Display credential details
	fmt.Println("Credential Details:")
	fmt.Println("===================")
	fmt.Printf("ID:         %s\n", cred.Id)
	fmt.Printf("Name:       %s\n", cred.Name)

	providerName := "UNKNOWN"
	switch cred.Provider {
	case backendv1.CredentialProvider_GCP:
		providerName = "GCP"
	case backendv1.CredentialProvider_AWS:
		providerName = "AWS"
	case backendv1.CredentialProvider_AZURE:
		providerName = "Azure"
	}
	fmt.Printf("Provider:   %s\n", providerName)

	if cred.CreatedAt != nil {
		fmt.Printf("Created At: %s\n", cred.CreatedAt.AsTime().Format("2006-01-02 15:04:05"))
	}
	if cred.UpdatedAt != nil {
		fmt.Printf("Updated At: %s\n", cred.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\nCredential Data:")
	fmt.Println("----------------")
	switch cred.Provider {
	case backendv1.CredentialProvider_GCP:
		if gcp := cred.GetGcp(); gcp != nil {
			fmt.Printf("Service Account Key (Base64): %s\n", maskSensitive(gcp.ServiceAccountKeyBase64))
		}
	case backendv1.CredentialProvider_AWS:
		if aws := cred.GetAws(); aws != nil {
			fmt.Printf("Account ID:       %s\n", aws.AccountId)
			fmt.Printf("Access Key ID:   %s\n", maskSensitive(aws.AccessKeyId))
			fmt.Printf("Secret Key:       %s\n", maskSensitive(aws.SecretAccessKey))
			if aws.Region != nil {
				fmt.Printf("Region:          %s\n", *aws.Region)
			}
			if aws.SessionToken != nil {
				fmt.Printf("Session Token:   %s\n", maskSensitive(*aws.SessionToken))
			}
		}
	case backendv1.CredentialProvider_AZURE:
		if azure := cred.GetAzure(); azure != nil {
			fmt.Printf("Client ID:        %s\n", azure.ClientId)
			fmt.Printf("Client Secret:    %s\n", maskSensitive(azure.ClientSecret))
			fmt.Printf("Tenant ID:        %s\n", azure.TenantId)
			fmt.Printf("Subscription ID:  %s\n", azure.SubscriptionId)
		}
	}
}

func maskSensitive(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) <= 8 {
		return "***"
	}
	return s[:4] + "..." + s[len(s)-4:]
}
