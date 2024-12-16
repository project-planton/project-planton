package module

import (
	"fmt"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistry/v1/iac/pulumi/module/outputs"
	pulumigcp "github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/artifactregistry"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// pythonRepo creates python repository and also grants reader role to the reader service account and writer, admin roles to
// writer service account.
func pythonRepo(ctx *pulumi.Context, locals *Locals, gcpProvider *pulumigcp.Provider,
	readerServiceAccount *serviceaccount.Account, writerServiceAccount *serviceaccount.Account) error {
	//create a variable with descriptive name for the api-resource in the input
	gcpArtifactRegistry := locals.GcpArtifactRegistry

	//create a name for the python repo since the name of this repository should be unique with in the gcp project.
	pythonRepoName := fmt.Sprintf("%s-python", gcpArtifactRegistry.Metadata.Id)

	//create python repository
	createdPythonRepo, err := artifactregistry.NewRepository(ctx,
		pythonRepoName,
		&artifactregistry.RepositoryArgs{
			Project:      pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
			Location:     pulumi.String(gcpArtifactRegistry.Spec.Region),
			RepositoryId: pulumi.String(pythonRepoName),
			Format:       pulumi.String("PYTHON"),
			Labels:       pulumi.ToStringMap(locals.GcpLabels),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create python repo")
	}

	//grant "reader" role for the writer service account on the repo
	_, err = artifactregistry.NewRepositoryIamMember(ctx,
		fmt.Sprintf("%s-reader", pythonRepoName),
		&artifactregistry.RepositoryIamMemberArgs{
			Project:    pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
			Location:   pulumi.String(gcpArtifactRegistry.Spec.Region),
			Repository: createdPythonRepo.RepositoryId,
			Role:       pulumi.String("roles/artifactregistry.reader"),
			Member:     pulumi.Sprintf("serviceAccount:%s", readerServiceAccount.Email),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to grant reader role on python repo for reader service account")
	}

	//grant "writer" role for the writer service account on the repo
	_, err = artifactregistry.NewRepositoryIamMember(ctx, fmt.Sprintf("%s-writer",
		pythonRepoName), &artifactregistry.RepositoryIamMemberArgs{
		Project:    pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
		Location:   pulumi.String(gcpArtifactRegistry.Spec.Region),
		Repository: createdPythonRepo.RepositoryId,
		Role:       pulumi.String("roles/artifactregistry.writer"),
		Member:     pulumi.Sprintf("serviceAccount:%s", writerServiceAccount.Email),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to grant writer role on python repo for writer service account")
	}

	//grant "admin" role for writer service account on the repo
	_, err = artifactregistry.NewRepositoryIamMember(ctx, fmt.Sprintf("%s-admin",
		pythonRepoName), &artifactregistry.RepositoryIamMemberArgs{
		Project:    pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
		Location:   pulumi.String(gcpArtifactRegistry.Spec.Region),
		Repository: createdPythonRepo.RepositoryId,
		Role:       pulumi.String("roles/artifactregistry.repoAdmin"),
		Member:     pulumi.Sprintf("serviceAccount:%s", writerServiceAccount.Email),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to grant admin role on python repo for writer service account")
	}

	//export the name of the python repository as output
	ctx.Export(outputs.PYTHON_REPO_NAME, createdPythonRepo.RepositoryId)

	return nil
}
