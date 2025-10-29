package module

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// resolveScriptContent handles both inline content and R2 references.
// Uses AWS S3 provider for IaC-native private bucket access.
func resolveScriptContent(
	ctx *pulumi.Context,
	locals *Locals,
	r2Provider *aws.Provider,
) pulumi.StringOutput {

	scriptSource := locals.CloudflareWorker.Spec.ScriptSource

	// Option 1: Inline script content
	if scriptSource.GetValue() != "" {
		return pulumi.String(scriptSource.GetValue()).ToStringOutput()
	}

	// Option 2: R2 bucket reference (private bucket via S3 API)
	// Format: "r2://<bucket-name>/<key-path>"
	if scriptSource.GetValueFrom() != nil && scriptSource.GetValueFrom().FieldPath != "" {
		ref := scriptSource.GetValueFrom().FieldPath

		// Parse R2 reference
		bucketName, objectKey, err := parseR2Reference(ref)
		if err != nil {
			// Log error and return empty (will fail Cloudflare validation)
			return pulumi.Sprintf("").ToStringOutput()
		}

		// Use AWS S3 LookupBucketObject (IaC-native, Terraform-compatible)
		scriptObject := s3.LookupBucketObjectOutput(ctx, s3.LookupBucketObjectOutputArgs{
			Bucket: pulumi.String(bucketName),
			Key:    pulumi.String(objectKey),
		}, pulumi.Provider(r2Provider))

		// Return script body content
		return scriptObject.Body()
	}

	// Neither provided
	return pulumi.String("").ToStringOutput()
}

// parseR2Reference extracts bucket and key from r2://bucket/key format.
// Example: "r2://my-artifacts/workers/abc123/script.js"
// Returns: ("my-artifacts", "workers/abc123/script.js", nil)
func parseR2Reference(ref string) (string, string, error) {
	if !strings.HasPrefix(ref, "r2://") {
		return "", "", errors.New("R2 reference must start with r2://")
	}

	path := strings.TrimPrefix(ref, "r2://")
	parts := strings.SplitN(path, "/", 2)

	if len(parts) != 2 {
		return "", "", errors.New("R2 reference must be r2://bucket/key")
	}

	return parts[0], parts[1], nil
}
