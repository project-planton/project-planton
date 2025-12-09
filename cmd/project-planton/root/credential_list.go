package root

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"github.com/project-planton/project-planton/app/backend/apis/gen/go/proto/backendv1connect"
	"github.com/spf13/cobra"
)

var CredentialListCmd = &cobra.Command{
	Use:   "credential:list",
	Short: "list all cloud provider credentials",
	Long:  "List all stored cloud provider credentials with optional filtering by provider type (gcp, aws, azure).",
	Run:   credentialListHandler,
}

func init() {
	CredentialListCmd.Flags().StringP("provider", "p", "", "filter by cloud provider: gcp, aws, or azure (optional)")
}

func credentialListHandler(cmd *cobra.Command, args []string) {
	// Get provider filter (optional)
	provider, _ := cmd.Flags().GetString("provider")

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
	req := &backendv1.ListCredentialsRequest{}

	// Add provider filter if specified
	if provider != "" {
		provider = strings.ToLower(provider)
		var providerEnum backendv1.CredentialProvider
		switch provider {
		case "gcp":
			providerEnum = backendv1.CredentialProvider_GCP
		case "aws":
			providerEnum = backendv1.CredentialProvider_AWS
		case "azure":
			providerEnum = backendv1.CredentialProvider_AZURE
		default:
			fmt.Printf("Error: Unsupported provider '%s'. Supported providers: gcp, aws, azure\n", provider)
			os.Exit(1)
		}
		req.Provider = &providerEnum
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Make the API call
	resp, err := client.ListCredentials(ctx, connect.NewRequest(req))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeUnavailable {
			fmt.Printf("Error: Cannot connect to backend service at %s. Please check:\n", backendURL)
			fmt.Printf("  1. The backend service is running\n")
			fmt.Printf("  2. The backend URL is correct\n")
			fmt.Printf("  3. Network connectivity\n")
			os.Exit(1)
		}
		fmt.Printf("Error listing credentials: %v\n", err)
		os.Exit(1)
	}

	credentials := resp.Msg.Credentials

	// Display results
	if len(credentials) == 0 {
		if provider != "" {
			fmt.Printf("No credentials found for provider: %s\n", strings.ToUpper(provider))
		} else {
			fmt.Println("No credentials found")
		}
		return
	}

	// Print header
	fmt.Printf("%-24s %-30s %-10s %-20s\n", "ID", "NAME", "PROVIDER", "CREATED")
	fmt.Println(strings.Repeat("-", 84))

	// Print credentials
	for _, cred := range credentials {
		providerStr := providerEnumToStr(cred.Provider)
		createdAt := ""
		if cred.CreatedAt != nil {
			createdAt = cred.CreatedAt.AsTime().Format("2006-01-02 15:04:05")
		}

		fmt.Printf("%-24s %-30s %-10s %-20s\n",
			cred.Id,
			truncate(cred.Name, 30),
			strings.ToUpper(providerStr),
			createdAt,
		)
	}

	// Print summary
	fmt.Printf("\nTotal: %d credential(s)", len(credentials))
	if provider != "" {
		fmt.Printf(" (filtered by provider: %s)", strings.ToUpper(provider))
	}
	fmt.Println()
}

// providerEnumToStr converts CredentialProvider enum to string.
func providerEnumToStr(provider backendv1.CredentialProvider) string {
	switch provider {
	case backendv1.CredentialProvider_GCP:
		return "gcp"
	case backendv1.CredentialProvider_AWS:
		return "aws"
	case backendv1.CredentialProvider_AZURE:
		return "azure"
	default:
		return "unknown"
	}
}

// truncate truncates a string to the specified length.
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
