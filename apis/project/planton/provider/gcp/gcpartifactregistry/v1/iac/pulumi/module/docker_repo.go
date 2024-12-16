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

// dockerRepo creates docker repository and also grants reader role to the reader service account and writer, admin roles to
// writer service account.
func dockerRepo(ctx *pulumi.Context, locals *Locals, gcpProvider *pulumigcp.Provider,
	readerServiceAccount *serviceaccount.Account, writerServiceAccount *serviceaccount.Account) error {
	//create a variable with descriptive name for the api-resource in the input
	gcpArtifactRegistry := locals.GcpArtifactRegistry

	//create a name for the docker repo since the name of this repository should be unique with in the gcp project.
	dockerRepoName := fmt.Sprintf("%s-docker", gcpArtifactRegistry.Metadata.Id)

	//create docker repository
	createdDockerRepo, err := artifactregistry.NewRepository(ctx,
		dockerRepoName,
		&artifactregistry.RepositoryArgs{
			Project:      pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
			Location:     pulumi.String(gcpArtifactRegistry.Spec.Region),
			RepositoryId: pulumi.String(dockerRepoName),
			Format:       pulumi.String("DOCKER"),
			Labels:       pulumi.ToStringMap(locals.GcpLabels),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create docker repo")
	}

	if gcpArtifactRegistry.Spec.IsExternal {
		//grant "reader" role for all users on the repo to make it public
		_, err = artifactregistry.NewRepositoryIamMember(ctx,
			fmt.Sprintf("%s-reader-for-all-users", dockerRepoName),
			&artifactregistry.RepositoryIamMemberArgs{
				Project:    pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
				Location:   pulumi.String(gcpArtifactRegistry.Spec.Region),
				Repository: createdDockerRepo.RepositoryId,
				Role:       pulumi.String("roles/artifactregistry.reader"),
				//"allUsers" is a special identifier on google identity system which is used
				//for granting permissions to for everyone.
				Member: pulumi.Sprintf("allUsers"),
			}, pulumi.Provider(gcpProvider))
		if err != nil {
			return errors.Wrap(err, "failed to grant reader role on docker repo for reader service account")
		}
	} else {
		//grant "reader" role for the writer service account on the repo
		_, err = artifactregistry.NewRepositoryIamMember(ctx,
			fmt.Sprintf("%s-reader", dockerRepoName),
			&artifactregistry.RepositoryIamMemberArgs{
				Project:    pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
				Location:   pulumi.String(gcpArtifactRegistry.Spec.Region),
				Repository: createdDockerRepo.RepositoryId,
				Role:       pulumi.String("roles/artifactregistry.reader"),
				Member:     pulumi.Sprintf("serviceAccount:%s", readerServiceAccount.Email),
			}, pulumi.Provider(gcpProvider))
		if err != nil {
			return errors.Wrap(err, "failed to grant reader role on docker repo for reader service account")
		}
	}

	//grant "writer" role for the writer service account on the repo
	_, err = artifactregistry.NewRepositoryIamMember(ctx, fmt.Sprintf("%s-writer",
		dockerRepoName), &artifactregistry.RepositoryIamMemberArgs{
		Project:    pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
		Location:   pulumi.String(gcpArtifactRegistry.Spec.Region),
		Repository: createdDockerRepo.RepositoryId,
		Role:       pulumi.String("roles/artifactregistry.writer"),
		Member:     pulumi.Sprintf("serviceAccount:%s", writerServiceAccount.Email),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to grant writer role on docker repo for writer service account")
	}

	//grant "admin" role for writer service account on the repo
	_, err = artifactregistry.NewRepositoryIamMember(ctx, fmt.Sprintf("%s-admin",
		dockerRepoName), &artifactregistry.RepositoryIamMemberArgs{
		Project:    pulumi.String(gcpArtifactRegistry.Spec.ProjectId),
		Location:   pulumi.String(gcpArtifactRegistry.Spec.Region),
		Repository: createdDockerRepo.RepositoryId,
		Role:       pulumi.String("roles/artifactregistry.repoAdmin"),
		Member:     pulumi.Sprintf("serviceAccount:%s", writerServiceAccount.Email),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to grant admin role on docker repo for writer service account")
	}

	//export important attributes of the docker repository as outputs
	ctx.Export(outputs.DockerRepoName, createdDockerRepo.RepositoryId)
	//export docker-repo hostname
	ctx.Export(outputs.DockerRepoHostname, pulumi.Sprintf(
		"%s-docker.pkg.dev", createdDockerRepo.Location))
	//export complete docker repo url based on the attributes of the created docker-repo
	// ex: artifactregistry://us-central1-docker.pkg.dev/my-gcp-project-id/my-company-docker-repo
	ctx.Export(outputs.DockerRepoUrl, pulumi.Sprintf(
		"%s-docker.pkg.dev/%s/%s",
		createdDockerRepo.Location, createdDockerRepo.Project, createdDockerRepo.Name))

	return nil
}
