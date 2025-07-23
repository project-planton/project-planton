package module

import (
	"github.com/pkg/errors"
	digitaloceancertificatev1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceancertificate/v1"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// bucket provisions the Spaces bucket and exports its ID & endpoint.
func bucket(
	ctx *pulumi.Context,
	locals *Locals,
	digitalOceanProvider *digitalocean.Provider,
) (*digitalocean.SpacesBucket, error) {

	// 1. Map proto enum → ACL string.
	var acl pulumi.StringPtrInput
	if locals.DigitalOceanCertificate.Spec.AccessControl == digitaloceancertificatev1.DigitalOceanCertificateAccessControl_PUBLIC_READ {
		acl = pulumi.String("public-read")
	} else {
		acl = pulumi.String("private")
	}

	// 2. Build resource arguments directly from proto fields.
	bucketArgs := &digitalocean.SpacesBucketArgs{
		Name:   pulumi.String(locals.DigitalOceanCertificate.Spec.BucketName),
		Region: pulumi.String(locals.DigitalOceanCertificate.Spec.Region.String()),
		Acl:    acl,
	}

	// 3. Optional versioning.
	if locals.DigitalOceanCertificate.Spec.VersioningEnabled {
		bucketArgs.Versioning = &digitalocean.SpacesBucketVersioningArgs{
			Enabled: pulumi.Bool(true),
		}
	}

	// 4. Create the bucket.
	createdBucket, err := digitalocean.NewSpacesBucket(
		ctx,
		"bucket",
		bucketArgs,
		pulumi.Provider(digitalOceanProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean spaces bucket")
	}

	// 5. Export stack outputs.
	ctx.Export(OpBucketId, createdBucket.ID())
	ctx.Export(OpEndpoint, createdBucket.Endpoint)

	return createdBucket, nil
}
