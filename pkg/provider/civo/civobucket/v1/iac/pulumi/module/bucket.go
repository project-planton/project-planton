package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// bucket provisions the Civo Object Store bucket and exports its ID, endpoint & keys.
func bucket(
	ctx *pulumi.Context,
	locals *Locals,
	civoProvider *civo.Provider,
) (*civo.ObjectStore, error) {

	// 1. Create an Object‑Store credential (no explicit keys → Civo generates them).
	createdCredential, err := civo.NewObjectStoreCredential(
		ctx,
		"bucket-creds",
		&civo.ObjectStoreCredentialArgs{},
		pulumi.Provider(civoProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create object‑store credential")
	}

	// 2. Build resource arguments directly from proto fields.
	bucketArgs := &civo.ObjectStoreArgs{
		Name:        pulumi.String(locals.CivoBucket.Spec.BucketName),
		Region:      pulumi.String(locals.CivoBucket.Spec.Region.String()),
		AccessKeyId: createdCredential.AccessKeyId,
	}

	// 3. Create the bucket.
	createdBucket, err := civo.NewObjectStore(
		ctx,
		"bucket",
		bucketArgs,
		pulumi.Provider(civoProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create civo object‑store bucket")
	}

	// 4. Export stack outputs.
	ctx.Export(OpBucketId, createdBucket.ID())
	ctx.Export(OpEndpointUrl, createdBucket.BucketUrl)
	ctx.Export(OpAccessKeySecretRef, createdCredential.AccessKeyId)
	ctx.Export(OpSecretKeySecretRef, createdCredential.SecretAccessKey)

	return createdBucket, nil
}
