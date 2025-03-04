package module

import (
	"fmt"
	"github.com/pkg/errors"
	gcsbucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcsbucket/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcsbucket/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gcsbucketv1.GcsBucketStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	createdBucket, err := storage.NewBucket(ctx,
		locals.GcsBucket.Metadata.Name,
		&storage.BucketArgs{
			ForceDestroy:             pulumi.Bool(true),
			Labels:                   pulumi.ToStringMap(locals.GcpLabels),
			Location:                 pulumi.String(locals.GcsBucket.Spec.GcpRegion),
			Name:                     pulumi.String(locals.GcsBucket.Metadata.Name),
			Project:                  pulumi.String(locals.GcsBucket.Spec.GcpProjectId),
			UniformBucketLevelAccess: pulumi.Bool(!locals.GcsBucket.Spec.IsPublic),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create bucket resource")
	}

	ctx.Export(outputs.BucketId, createdBucket.ID())

	if !locals.GcsBucket.Spec.IsPublic {
		return nil
	}

	//grant bucket-reader role to allUsers
	_, err = storage.NewBucketAccessControl(ctx,
		fmt.Sprintf("%s-public", locals.GcsBucket.Metadata.Name),
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
		fmt.Sprintf("%s-public-object-reader", locals.GcsBucket.Metadata.Name),
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
