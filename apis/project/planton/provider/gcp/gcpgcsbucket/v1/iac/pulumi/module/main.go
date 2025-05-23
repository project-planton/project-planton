package module

import (
	"fmt"
	"github.com/pkg/errors"
	gcpgcsbucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgcsbucket/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gcpgcsbucketv1.GcpGcsBucketStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	createdBucket, err := storage.NewBucket(ctx,
		locals.GcpGcsBucket.Metadata.Name,
		&storage.BucketArgs{
			ForceDestroy:             pulumi.Bool(true),
			Labels:                   pulumi.ToStringMap(locals.GcpLabels),
			Location:                 pulumi.String(locals.GcpGcsBucket.Spec.GcpRegion),
			Name:                     pulumi.String(locals.GcpGcsBucket.Metadata.Name),
			Project:                  pulumi.String(locals.GcpGcsBucket.Spec.GcpProjectId),
			UniformBucketLevelAccess: pulumi.Bool(!locals.GcpGcsBucket.Spec.IsPublic),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create bucket resource")
	}

	ctx.Export(OpBucketId, createdBucket.ID())

	if !locals.GcpGcsBucket.Spec.IsPublic {
		return nil
	}

	//grant bucket-reader role to allUsers
	_, err = storage.NewBucketAccessControl(ctx,
		fmt.Sprintf("%s-public", locals.GcpGcsBucket.Metadata.Name),
		&storage.BucketAccessControlArgs{
			Bucket: createdBucket.Name,
			Role:   pulumi.String("READER"),
			Entity: pulumi.String("allUsers"),
		}, pulumi.Parent(createdBucket))
	if err != nil {
		return errors.Wrap(err, "failed to create public access control rule")
	}

	//grant object-reader role to allUsers
	_, err = storage.NewBucketAccessControl(ctx,
		fmt.Sprintf("%s-public-object-reader", locals.GcpGcsBucket.Metadata.Name),
		&storage.BucketAccessControlArgs{
			Bucket: createdBucket.Name,
			Role:   pulumi.String("READER"),
			Entity: pulumi.String("allUsers"),
		},
		pulumi.Parent(createdBucket))
	if err != nil {
		return errors.Wrap(err, "failed to create public access control rule for object reader role")
	}
	return nil
}
