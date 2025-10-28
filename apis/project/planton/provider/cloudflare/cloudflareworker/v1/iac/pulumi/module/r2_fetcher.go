package module

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// resolveScriptContent handles both inline content and URL references.
// Returns pulumi.StringOutput for WorkerScript.Content field.
//
// Note on R2 URL accessibility:
// - R2 URLs must be publicly accessible for Pulumi's HTTP fetch
// - This is acceptable because worker scripts are client-side code (not sensitive)
// - Content-hash based paths provide obscurity
// - No directory listing enabled
// - Alternative: use signed URLs with expiration generated during build
func resolveScriptContent(
	ctx *pulumi.Context,
	locals *Locals,
) pulumi.StringOutput {

	scriptSource := locals.CloudflareWorker.Spec.ScriptSource

	// Option 1: Inline script content
	if scriptSource.GetValue() != "" {
		return pulumi.String(scriptSource.GetValue()).ToStringOutput()
	}

	// Option 2: URL reference (from R2 via pipeline)
	// Download during Pulumi apply (IaC-native approach)
	if scriptSource.GetValueFrom() != nil && scriptSource.GetValueFrom().FieldPath != "" {
		url := scriptSource.GetValueFrom().FieldPath

		// Use Pulumi Apply to download during deployment
		return pulumi.String(url).ToStringOutput().ApplyT(func(u string) (string, error) {
			content, err := fetchScriptFromURL(u)
			if err != nil {
				return "", errors.Wrapf(err, "failed to fetch script from %s", u)
			}
			return content, nil
		}).(pulumi.StringOutput)
	}

	// Neither provided - return empty (will fail Cloudflare validation)
	return pulumi.String("").ToStringOutput()
}

// fetchScriptFromURL downloads script content from a public R2 URL.
// Simple HTTP GET - works in Pulumi and translates to Terraform (http data source).
func fetchScriptFromURL(url string) (string, error) {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to create HTTP request")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", errors.Wrapf(err, "failed to fetch from URL")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("HTTP %d from R2 URL", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}

	return string(body), nil
}

