package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ecrRepo creates an AWS ECR repository and optional lifecycle and repository policies
// based on the fields in AwsEcrRepoSpec.
func ecrRepo(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	spec := locals.AwsEcrRepo.Spec

	imageTagMutability := "MUTABLE"

	if spec.ImageImmutable {
		imageTagMutability = "IMMUTABLE"
	}

	repo, err := ecr.NewRepository(ctx, locals.AwsEcrRepo.Metadata.Name, &ecr.RepositoryArgs{
		Name:               pulumi.String(spec.RepositoryName),
		ImageTagMutability: pulumi.String(imageTagMutability),
		ForceDelete:        pulumi.Bool(spec.ForceDelete),
		Tags:               pulumi.ToStringMap(locals.AwsTags),
		EncryptionConfigurations: ecr.RepositoryEncryptionConfigurationArray{
			&ecr.RepositoryEncryptionConfigurationArgs{
				EncryptionType: pulumi.String(spec.EncryptionType),
				KmsKey:         pulumi.String(spec.KmsKeyId),
			},
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create AWS ECR Repository")
	}

	ctx.Export(OpEcrRepoName, repo.Name)
	ctx.Export(OpEcrRepoUrl, repo.RepositoryUrl)
	ctx.Export(OpEcrRepoArn, repo.Arn)
	ctx.Export(OpEcrRepoRegistryId, repo.RegistryId)

	return nil
}
