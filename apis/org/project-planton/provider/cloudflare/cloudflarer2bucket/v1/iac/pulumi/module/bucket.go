package module

import (
	cloudflarer2bucketv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/cloudflare/cloudflarer2bucket/v1"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// bucket provisions the R2 bucket (and optional managed domain) and exports outputs.
func bucket(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudflare.Provider,
) (*cloudflare.R2Bucket, error) {

	// 1. Translate the enum into Cloudflare's location string.
	// Cloudflare expects UPPERCASE location values.
	var bucketLocation string
	switch locals.CloudflareR2Bucket.Spec.Location {
	case cloudflarer2bucketv1.CloudflareR2Location_WNAM:
		bucketLocation = "WNAM"
	case cloudflarer2bucketv1.CloudflareR2Location_ENAM:
		bucketLocation = "ENAM"
	case cloudflarer2bucketv1.CloudflareR2Location_WEUR:
		bucketLocation = "WEUR"
	case cloudflarer2bucketv1.CloudflareR2Location_EEUR:
		bucketLocation = "EEUR"
	case cloudflarer2bucketv1.CloudflareR2Location_APAC:
		bucketLocation = "APAC"
	case cloudflarer2bucketv1.CloudflareR2Location_OC:
		bucketLocation = "OC"
	default:
		bucketLocation = "auto"
	}

	// 2. Create the bucket.
	createdBucket, err := cloudflare.NewR2Bucket(
		ctx,
		"bucket",
		&cloudflare.R2BucketArgs{
			AccountId: pulumi.String(locals.CloudflareR2Bucket.Spec.AccountId),
			Name:      pulumi.String(locals.CloudflareR2Bucket.Spec.BucketName),
			Location:  pulumi.String(bucketLocation),
		},
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Cloudflare R2 bucket")
	}

	// 4. Warn about unsupported versioning.
	if locals.CloudflareR2Bucket.Spec.VersioningEnabled {
		ctx.Log.Warn("Cloudflare provider does not yet support R2 bucket versioning – field will be ignored.", nil)
	}

	// 5. Export stack outputs.
	ctx.Export(OpBucketName, createdBucket.Name)

	return createdBucket, nil
}
